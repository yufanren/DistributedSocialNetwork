package etcdclient

import (
	"context"
	. "cs9223-final-project/etcdclient/etcdclientpb"
	"cs9223-final-project/utils"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

var blogLock sync.Mutex

func Listblog_(pageNum int, pageSize int) (int, []*Blog) {
	var blogs []*Blog 
	pageNum--
	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	kv := clientv3.NewKV(cli)
	blogLock.Lock()
	resp, err := cli.Get(ctx, "blognumber")
	blogLock.Unlock()
	if err != nil {
		cancel()
		return 0, []*Blog{}
	}
	total, err := strconv.Atoi(string(resp.Kvs[0].Value))
	if err != nil {
		cancel()
		return 0, []*Blog{}
	}
	if pageNum * pageSize > total {
		pageNum = 0
	}
	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))

	opts := []clientv3.OpOption {
		clientv3.WithRange("blog/" + FixedLengthItoa(int(m), 10)),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
	}
	Blogs, err := kv.Get(ctx, "blog/" + FixedLengthItoa(pageNum * pageSize, 10), opts...)
	cancel()
	if err != nil {
		return 0, []*Blog{}
	}

	for _, item := range Blogs.Kvs {
		blog := new(Blog)
		err = utils.Decode([]byte(item.Value), blog)
		if err != nil {
			return 0, []*Blog{}
		}
		blogs = append(blogs, blog)
	}
	return total, blogs
}

func Addblog_(blog *Blog) (int, error) {
	author := (*blog).Author
	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	blogLock.Lock()
	defer blogLock.Unlock()

	resp, _ := cli.Get(ctx, "blognumber")
	blogNumber, _ := strconv.Atoi(string(resp.Kvs[0].Value))
	resp, err = cli.Get(ctx, fmt.Sprintf("user/%s/blogs", author))
	if err != nil {
		return blogNumber, errors.New("can't find user blog record")
	}
	bloglist := make([]int, 0)
	err = utils.Decode(resp.Kvs[0].Value, &bloglist)
	if err != nil {
		return blogNumber, err
	}
	bloglist = append(bloglist, blogNumber)
	(*blog).BlogId = int32(blogNumber)
	encodedBlogs, err := utils.Encode(bloglist)
	if err != nil {
		return blogNumber, err
	}
	encodedBlog, err := utils.Encode(blog)
	if err != nil {
		return blogNumber, err
	}
	_, err = cli.KV.Txn(ctx).Then(clientv3.OpPut(fmt.Sprintf("user/%s/blogs", author), string(encodedBlogs)),
								clientv3.OpPut("blognumber", strconv.Itoa(blogNumber+1)),
								clientv3.OpPut(fmt.Sprintf("blog/%s", FixedLengthItoa(blogNumber, 10)), string(encodedBlog))).Commit()
	if err != nil {
		return blogNumber, errors.New("blog post transaction failed")
	}
	return blogNumber,  nil
}

func Getblog_(blogId int) (*Blog, error) {
	encodedBlog, err := get(fmt.Sprintf("blog/%s", FixedLengthItoa(blogId, 10)))
	if err != nil {
		return nil, errors.New("blog not found")
	}
	blog := new(Blog)
	err = utils.Decode([]byte(encodedBlog), blog)
	if err != nil {
		return nil, errors.New("unable to decode blog")
	}
	return blog, nil
}

func Getuserblog_(username string, pageNum int, pageSize int) (int, []*Blog) {
	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	blogLock.Lock()
	resp, err := cli.Get(ctx, fmt.Sprintf("user/%s/blogs", username))
	blogLock.Unlock()
	if err != nil {
		return 0, []*Blog{}
	}
	bloglist := new([]int)
	err = utils.Decode(resp.Kvs[0].Value, bloglist)
	sort.Sort(sort.Reverse(sort.IntSlice(*bloglist)))
	if err != nil {
		return 0, []*Blog{}
	}
	pageNum--
	total := len(*bloglist)
	if pageNum * pageSize > total {
		pageNum = 0
	}
	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
	blogpagelist := (*bloglist)[pageNum*pageSize:int(m):len(*bloglist)]
	var blogs []*Blog
	for _, blogId := range blogpagelist {
		resp, err := cli.Get(ctx, fmt.Sprintf("blog/%s", FixedLengthItoa(blogId, 10)))
		if err != nil {
			return len(blogs), blogs
		}
		blog := new(Blog)
		err = utils.Decode([]byte(resp.Kvs[0].Value), blog)
		blogs = append(blogs, blog)
	}
	return len(blogs), blogs
}

func Getfollowblog_(username string, pageNum int, pageSize int) (int, []*Blog) {
	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	blogLock.Lock()
	resp, err := cli.Get(ctx, fmt.Sprintf("user/%s/followed", username))
	if err != nil || len(resp.Kvs) == 0 {
		return 0, []*Blog{}
	}
	followedMap := new(map[string]bool)
	err = utils.Decode(resp.Kvs[0].Value, followedMap)
	if err != nil {
		return 0, []*Blog{}
	}
	bloglist := make([]int, 0)
	for user, isfollowed := range *followedMap {
		if isfollowed {
			resp, err = cli.Get(ctx, fmt.Sprintf("user/%s/blogs", user))
			if err != nil {
				continue
			}
			tmplist := new([]int)
			err := utils.Decode(resp.Kvs[0].Value, tmplist)
			if err != nil {
				continue
			}
			bloglist = append(bloglist, (*tmplist)...)
		}
	}
	blogLock.Unlock()
	pageNum--
	total := len(bloglist)
	if pageNum * pageSize > total {
		pageNum = 0
	}
	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
	blogpagelist := bloglist[pageNum*pageSize:int(m):len(bloglist)]
	// sort.Ints(blogpagelist)
	sort.Sort(sort.Reverse(sort.IntSlice(blogpagelist)))
	var blogs []*Blog
	for _, blogId := range blogpagelist {
		resp, err := cli.Get(ctx, fmt.Sprintf("blog/%s", FixedLengthItoa(blogId, 10)))
		if err != nil {
			return len(blogs), blogs
		}
		blog := new(Blog)
		err = utils.Decode([]byte(resp.Kvs[0].Value), blog)
		blogs = append(blogs, blog)
	}
	return len(blogs), blogs
}

//---------------------------------------------------------------------------------------------------------------------------------------------
func StartDB() {
	time.Sleep(time.Second)
	cli, _ := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.KV.Txn(ctx).If(clientv3.Compare(clientv3.Version("blognumber"), ">",
		0)).Else(clientv3.OpPut("blognumber", "0")).Commit()
	cancel()	
	if err != nil {
		fmt.Printf("Exception while preping DB")
	}
	if resp.Succeeded {
		fmt.Printf("DB already exists")
	} else {
		fmt.Printf("DB setup complete")
	}
}
// func GetBlog() ([]*Blog, error) {
// 	encoded, err := get("blog")
// 	if err != nil {
// 		return *new([]*Blog), err
// 	}
// 	blogs := new ([]*Blog)
// 	err = utils.Decode([]byte(encoded), blogs)
// 	if err != nil {
// 		return *new([]*Blog), err
// 	}
// 	return *blogs, nil
// }

// func PutBlog(blog *Blog) error {
// 	blogs, _ := GetBlog()
// 	blogs = append(blogs, blog)
// 	encoded, err := utils.Encode(blogs)
// 	if err != nil {
// 		return err
// 	}
// 	err = put("blog", string(encoded))
// 	return err
// }

// func AddBlog(blog *Blog) int {
// 	PutBlog(blog)
// 	author := (*blog).Author
// 	if userInfo, err := GetUser(author); err != nil {
// 		userInfo.Blogs = append(userInfo.Blogs, blog)
// 		PutUser(author, userInfo)
// 	}
// 	blogs, err := GetBlog()
// 	if err != nil {
// 		return 0
// 	}
// 	return len(blogs) - 1
// }

// func ListBlogs(pageNum int, pageSize int) (int, []*Blog) {
// 	var blogs []*Blog
// 	pageNum--
// 	dbBlogs, err := GetBlog()
// 	if err != nil {
// 		return 0, nil
// 	}
// 	total := len(dbBlogs)
// 	if pageNum * pageSize > total {
// 		pageNum = 0
// 	}
// 	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
// 	for i := pageNum * pageSize; i < int(m); i++ {
// 		blogs = append(blogs, dbBlogs[i])
// 	}
// 	return total, blogs
// }

// func GetBlogById(blogId int) *Blog {
// 	blogs, err := GetBlog()
// 	if err != nil {
// 		return nil
// 	}
// 	if blogId >= len(blogs) {
// 		return nil
// 	}
// 	return blogs[blogId]
// }

// func ListBlogsFromUser(user string, pageNum int, pageSize int) (int, []*Blog) {
// 	var blogs []*Blog
// 	pageNum--
// 	if userInfo, err := GetUser(user); err != nil {
// 		userBlogs := userInfo.Blogs
// 		total := len(userBlogs)
// 		if pageNum * pageSize > total {
// 			pageNum = 0
// 		}
// 		m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
// 		for i := pageNum * pageSize; i < int(m); i++ {
// 			blogs = append(blogs, userBlogs[i])
// 		}
// 		return total, blogs
// 	} else {
// 		return 0, []*Blog{}
// 	}
// }

// //Get everyone who the user follows
// func ListFollowedFromUser(user string) (int, []string) {
// 	var followed []string
// 	if userInfo, err := GetUser(user); err != nil {
// 		for user, isFollowed := range userInfo.Followed {
// 			if isFollowed {
// 				followed = append(followed, user)
// 			}
// 		}
// 		return len(followed), followed
// 	}
// 	return 0, []string{}
// }

// //Get everyone who follows the user
// func ListFollowerFromUser(user string) (int, []string) {
// 	var follower []string
// 	if userInfo, err := GetUser(user); err != nil {
// 		for user, isFollower := range userInfo.Follower {
// 			if isFollower {
// 				follower = append(follower, user)
// 			}
// 		}
// 		return len(follower), follower
// 	}
// 	return 0, []string{}
// }

// func ListFollowBlog(user string, pageNum int, pageSize int) (int, []*Blog) {
// 	var blogs []*Blog
// 	_, followed := ListFollowerFromUser(user)
// 	for _, fuser := range followed {
// 		_, fblogs := ListBlogsFromUser(fuser, 1, math.MaxInt)
// 		blogs = append(blogs, fblogs...)
// 	}
// 	total := len(blogs)
// 	pageNum--
// 	if pageNum * pageSize > total {
// 		pageNum = 0
// 	}
// 	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
// 	var rblogs []*Blog
// 	for i := pageNum * pageSize; i < int(m); i++ {
// 		rblogs = append(rblogs, blogs[i])
// 	}
// 	return total, rblogs
// }
package etcdclient

import (
	"cs9223-final-project/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

var followLock sync.Mutex

func Register_(username string, password string) error {

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return errors.New("error hashing password")
	}
	encodedBlogList, err := utils.Encode(make([]int, 0))  
	if err != nil {
		return errors.New("error encoding bloglist")
	}
	encodedFollowedList, err := utils.Encode(make(map[string]bool))
	if err != nil {
		return errors.New("error encoding followed list")
	}
	encodedFollowerList, err := utils.Encode(make(map[string]bool))
	if err != nil {
		return errors.New("error encoding follower list")
	}	
	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	response, err := cli.KV.Txn(ctx).If(clientv3.Compare(clientv3.Version(fmt.Sprintf("user/%s/pw", username)), ">",
		0)).Else(clientv3.OpPut(fmt.Sprintf("user/%s/pw", username), string(hash)),
		clientv3.OpPut(fmt.Sprintf("user/%s/blogs", username), string(encodedBlogList)),
		clientv3.OpPut(fmt.Sprintf("user/%s/followed", username), string(encodedFollowedList)),
		clientv3.OpPut(fmt.Sprintf("user/%s/follower", username), string(encodedFollowerList))).Commit()
	cancel()	
	if response.Succeeded {
		return errors.New("user name already taken")
	}
	return err
}

func Login_(username string, password string) error {

	hash, err := get(fmt.Sprintf("user/%s/pw", username))
	if err != nil {
		return errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	}
	return err
}

func Follow_(follower string, parent string) error {
	followLock.Lock()
	defer followLock.Unlock()

	encodedParentFollower, err := get(fmt.Sprintf("user/%s/follower", parent)) 
	if err != nil {
		return errors.New("User not found")
	}
	parentFollower := make(map[string]bool)
	err = utils.Decode([]byte(encodedParentFollower), &parentFollower)
	if err != nil {
		return err
	}
	encodedFollowerFollowed, err := get(fmt.Sprintf("user/%s/followed", follower)) 
	if err != nil {
		return errors.New("User not found")
	}
	followerFollowed := make(map[string]bool)
	err = utils.Decode([]byte(encodedFollowerFollowed), &followerFollowed)
	if err != nil {
		return err
	}

	if followerFollowed[parent] {
		followerFollowed[parent] = false
		if parentFollower[follower] {
			parentFollower[follower] = false
		}
	} else {
		followerFollowed[parent] = true
		parentFollower[follower] = true
	}

	finalParentFollower, err := utils.Encode(parentFollower)
	if err != nil {
		return err
	}
	finalFollowerFollowed, err := utils.Encode(followerFollowed)
	if err != nil {
		return err
	}

	cli, err := makeClient()
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.KV.Txn(ctx).Then(clientv3.OpPut(fmt.Sprintf("user/%s/follower", parent), string(finalParentFollower)),
					clientv3.OpPut(fmt.Sprintf("user/%s/followed", follower), string(finalFollowerFollowed))).Commit()
	cancel()
	return err
} 

func Isfollow_(follower string, parent string) bool {

	encodedFollowerFollowed, err := get(fmt.Sprintf("user/%s/followed", follower)) 
	if err != nil {
		return false
	}
	followerFollowed := make(map[string]bool)
	err = utils.Decode([]byte(encodedFollowerFollowed), &followerFollowed)
	if err != nil {
		return false
	}
	return followerFollowed[parent]
}

func Getuserfollowed_(user string)  (int, []string) {
	var followed []string
	encodedUserFollowed, err := get(fmt.Sprintf("user/%s/followed", user)) 
	if err != nil {
		return 0, []string{}
	}
	followerFollowed := make(map[string]bool)
	err = utils.Decode([]byte(encodedUserFollowed), &followerFollowed)
	if err != nil {
		return 0, []string{}
	}
	for user, isfollowed := range followerFollowed {
		if isfollowed {
			followed = append(followed, user)
		}
	}
	return len(followed), followed
}

func Getuserfollower_(user string)  (int, []string) {
	var follower []string
	encodedUserFollower, err := get(fmt.Sprintf("user/%s/follower", user)) 
	if err != nil {
		return 0, []string{}
	}
	userFollower := make(map[string]bool)
	err = utils.Decode([]byte(encodedUserFollower), &userFollower)
	if err != nil {
		return 0, []string{}
	}
	for user, isfollower := range userFollower {
		if isfollower {
			follower = append(follower, user)
		}
	}
	return len(follower), follower
}


//------------------------------------------------------------------------------------------------------------------------------------------

// type UserInfo struct {
// 	PasswordHash []byte
// 	Blogs []*Blog
// 	Followed map[string]bool
// 	Follower map[string]bool
// }

// func PutUser(username string, userInfo *UserInfo) error {
// 	encoded, err := utils.Encode(userInfo)
// 	if err != nil {
// 		return err
// 	}
// 	err = put(fmt.Sprintf("user/%s/", username), string(encoded))
// 	return err
// }

// func GetUser(username string) (*UserInfo, error) {
// 	encoded, err := get(fmt.Sprintf("user/%s/", username))
// 	if err != nil {
// 		return nil, err
// 	}
// 	if encoded == "" {
// 		return nil, errors.New("no such user")
// 	}
// 	userInfo := new (UserInfo)
// 	err = utils.Decode([]byte(encoded), userInfo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return userInfo, nil
// }

// func AddUser(username string, password string) error {
// 	cost := bcrypt.DefaultCost
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
// 	if err != nil {
// 		return errors.New("error hashing password")
// 	}
// 	userInfo := &UserInfo {
// 		PasswordHash: hash,
// 		Blogs: make([]*Blog, 0),
// 		Followed: make(map[string]bool),
// 		Follower: make(map[string]bool),
// 	}
// 	err = PutUser(username, userInfo)
// 	return err
// }

// func HasUser(username string) bool {
// 	_, err := GetUser(username)
// 	return err == nil
// }

// func ValidUser(username string, password string) bool {
// 	userInfo, err := GetUser(username)
// 	if err != nil {
// 		return false
// 	}
// 	err = bcrypt.CompareHashAndPassword(userInfo.PasswordHash, []byte(password))
// 	return err != bcrypt.ErrMismatchedHashAndPassword
// }

// func FollowUser(parent string, follower string) {
// 	parentInfo, err := GetUser(parent)
// 	if err != nil {
// 		return
// 	}
// 	followerInfo, err := GetUser(follower)
// 	if err == nil {
// 		parentMap := followerInfo.Followed
// 		if parentMap[parent] {
// 			parentMap[parent] = false
// 			if _, ok := parentInfo.Follower[follower]; ok {
// 				parentInfo.Follower[follower] = false
// 			}
// 		} else {
// 			parentMap[parent] = true
// 			parentInfo.Follower[follower] = true
// 		}
// 		PutUser(parent, parentInfo)
// 		PutUser(follower, followerInfo)
// 	}
// }

// func IsFollow(parent string, follower string) bool {
// 	userinfo, err := GetUser(follower)
// 	if err == nil {
// 		return userinfo.Followed[parent]
// 	}
// 	return false
// }
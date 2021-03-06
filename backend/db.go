package backend

import (

"errors"
"final-project-0b5a2e16-ly2062-yufanren/backend/backendpb"
"golang.org/x/crypto/bcrypt"
"math"
"sync"
)

type DB struct {
	user map[string]*UserInfo
	blogs []*backendpb.Blog
	// follow map[string]bool
	DBLock sync.RWMutex
}

type UserInfo struct {
	passwordHash []byte
	blogs []*backendpb.Blog
	followed map[string]bool
	follower map[string]bool
}

func NewDB() *DB {
	db := DB{}
	db.user = make(map[string]*UserInfo)
	db.blogs = make([]*backendpb.Blog, 0)
	// db.follow = make(map[string]bool)
	return &db
}

func (db *DB) AddUser(username string, password string) error {
	db.DBLock.Lock()
	defer db.DBLock.Unlock()

	if _, ok := db.user[username]; ok {
		return errors.New("user already exists")
	}
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return errors.New("error hashing password")
	}
	db.user[username] = &UserInfo{
		passwordHash: hash,
		blogs: make([]*backendpb.Blog, 0),
		followed: make(map[string]bool),
		follower: make(map[string]bool),
	}
	return nil
}

func (db *DB) HasUser(username string) bool {
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	_, ok := db.user[username]
	return ok
}

func (db *DB) ValidUser(username string, password string) bool {
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	err := bcrypt.CompareHashAndPassword(db.user[username].passwordHash, []byte(password))
	return err != bcrypt.ErrMismatchedHashAndPassword
}

func (db *DB) FollowUser(parent string, follower string) {
	db.DBLock.Lock()
	defer db.DBLock.Unlock()
	// if db.follow[parent + " " + follower] {
	// 	db.follow[parent + " " + follower] = false
	// } else {
	// 	db.follow[parent + " " + follower] = true
	// }
	parentInfo, ok := db.user[parent]
	if !ok {
		return
	}
	followerInfo, ok := db.user[follower]
	if ok {
		parentMap := followerInfo.followed
		if parentMap[parent] {
			parentMap[parent] = false
			if _, ok = parentInfo.follower[follower]; ok {
				parentInfo.follower[follower] = false
			}
		} else {
			parentMap[parent] = true
			parentInfo.follower[follower] = true
		}
	}
}

func (db *DB) isFollow(parent string, follower string) bool {
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	// return db.follow[parent + " " + follower]
	userinfo, ok := db.user[follower]
	if ok {
		return userinfo.followed[parent]
	}
	return false
}

func (db *DB) AddBlog(blog *backendpb.Blog) int {
	db.DBLock.Lock()
	defer db.DBLock.Unlock()
	db.blogs = append(db.blogs, blog)
	author := (*blog).Author
	if userinfo, ok := db.user[author]; ok {
		userinfo.blogs = append(userinfo.blogs, blog)
	}
	return len(db.blogs) - 1
}

func (db *DB) ListBlogs(pageNum int, pageSize int) (int, []*backendpb.Blog) {
	var blogs []*backendpb.Blog
	pageNum--
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	total := len(db.blogs)
	if pageNum * pageSize > total {
		pageNum = 0
	}
	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
	for i := pageNum * pageSize; i < int(m); i++ {
		blogs = append(blogs, db.blogs[i])
	}
	return total, blogs
}

func (db *DB) GetBlog(blogId int) *backendpb.Blog {
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	if blogId >= len(db.blogs) {
		return nil
	}
	return db.blogs[blogId]
}

// ---------------------------------------------------------------------

func (db *DB) ListBlogsFromUser(user string, pageNum int, pageSize int) (int, []*backendpb.Blog) {
	var blogs []*backendpb.Blog
	pageNum--
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	if userinfo, ok := db.user[user]; ok {
		userBlogs := userinfo.blogs
		total := len(userBlogs)
		if pageNum * pageSize > total {
			pageNum = 0
		}
		m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
		for i := pageNum * pageSize; i < int(m); i++ {
			blogs = append(blogs, userBlogs[i])
		}
		return total, blogs
	} else {
		return 0, []*backendpb.Blog{}
	}
}

//Get everyone who the user follows
func (db *DB) ListFollowedFromUser(user string) (int, []string) {
	var followed []string
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	if userinfo, ok := db.user[user]; ok {
		for user, isFollowed := range userinfo.followed {
			if isFollowed {
				followed = append(followed, user)
			}
		}
		return len(followed), followed
	}
	return 0, []string{}
}

//Get everyone who follows the user
func (db *DB) ListFollowerFromUser(user string) (int, []string) {
	var follower []string
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	if userinfo, ok := db.user[user]; ok {
		for user, isFollower := range userinfo.follower {
			if isFollower {
				follower = append(follower, user)
			}
		}
		return len(follower), follower
	}
	return 0, []string{}
}

func (db *DB) ListFollowBlog(user string, pageNum int, pageSize int) (int, []*backendpb.Blog) {
	var blogs []*backendpb.Blog
	db.DBLock.RLock()
	defer db.DBLock.RUnlock()
	_, followed := db.ListFollowerFromUser(user)
	for _, fuser := range followed {
		_, fblogs := db.ListBlogsFromUser(fuser, 1, math.MaxInt)
		blogs = append(blogs, fblogs...)
	}
	total := len(blogs)
	pageNum--
	if pageNum * pageSize > total {
		pageNum = 0
	}
	m := math.Min(float64(total), float64((pageNum + 1) * pageSize))
	var rblogs []*backendpb.Blog
	for i := pageNum * pageSize; i < int(m); i++ {
		rblogs = append(rblogs, blogs[i])
	}
	return total, rblogs
}
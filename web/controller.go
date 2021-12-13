package web

import (
	"cs9223-final-project/auth"
	"cs9223-final-project/etcdclient/etcdclientpb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res, err := etcdClient.Register(ctx, &etcdclientpb.AuthRequest{Username: username, Password: password})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"token": res.Token,
		}))
	}
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res, err := etcdClient.Login(ctx, &etcdclientpb.AuthRequest{Username: username, Password: password})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"token": res.Token,
		}))
	}
}

func listBlog(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
	res, err := etcdClient.ListBlog(ctx, &etcdclientpb.ListBlogRequest{PageNum: int32(pageNum), PageSize: int32(pageSize)})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"total": res.Total,
			"blogs": res.Blogs,
		}))
	}
}

func addBlog(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	token := ctx.GetHeader("Authorization")
	author, err := auth.GetUsername(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		return
	}
	blog := &etcdclientpb.Blog{Title: title, Content: content, Author: author}
	res, err := etcdClient.AddBlog(ctx, &etcdclientpb.AddBlogRequest{Blog: blog})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"blogId": res.BlogId,
		}))
	}
}

func getBlog(ctx *gin.Context) {
	blogId, _ := strconv.Atoi(ctx.Param("id"))
	res, err := etcdClient.GetBlog(ctx, &etcdclientpb.GetBlogRequest{BlogId: int32(blogId)})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"blog": res.Blog,
		}))
	}
}

func follow(ctx *gin.Context) {
	follower := ctx.Param("follower")
	token := ctx.GetHeader("Authorization")
	parent, err := auth.GetUsername(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		return
	}
	_, err = etcdClient.Follow(ctx, &etcdclientpb.FollowRequest{Parent: parent, Follower: follower})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{}))
	}
}

func isFollow(ctx *gin.Context) {
	follower := ctx.Param("follower")
	token := ctx.GetHeader("Authorization")
	parent, err := auth.GetUsername(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		return
	}
	res, err := etcdClient.IsFollow(ctx, &etcdclientpb.FollowRequest{Parent: parent, Follower: follower})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"res": res.Res,
		}))
	}
}

func listBlogByUser(ctx *gin.Context) {
	user := ctx.Param("user")
	pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
	res, err := etcdClient.GetUserBlog(ctx, &etcdclientpb.UserBlogRequest{Username: user, PageNum: int32(pageNum), PageSize: int32(pageSize)})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"total": res.Total,
			"blogs": res.Blogs,
		}))
	}
}

func getFollowListByFollower(ctx *gin.Context) {
	follower := ctx.Param("follower")
	res, err := etcdClient.GetUserFollowed(ctx, &etcdclientpb.UserRequest{Username: follower})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"total": res.Total,
			"parents": res.Parents,
		}))
	}
}

func getFollowListByParent(ctx *gin.Context) {
	follower := ctx.Param("parent")
	res, err := etcdClient.GetUserFollower(ctx, &etcdclientpb.UserRequest{Username: follower})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"total": res.Total,
			"followers": res.Parents,
		}))
	}
}

func listFollowBlogByUser(ctx *gin.Context) {
	user := ctx.Param("user")
	pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
	res, err := etcdClient.GetFollowBlog(ctx, &etcdclientpb.UserBlogRequest{Username: user, PageNum: int32(pageNum), PageSize: int32(pageSize)})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fail(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, success(gin.H{
			"total": res.Total,
			"blogs": res.Blogs,
		}))
	}
}
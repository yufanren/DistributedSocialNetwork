package main

import (
	"dist-demo/auth"
	"dist-demo/proto"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

const (
	port = ":8091"
	serverPort = ":8090"
)

func success(data gin.H) gin.H {
	return gin.H {
		"code": "200",
		"data": data,
	}
}

func fail(msg string) gin.H {
	return gin.H {
		"code": "400",
		"msg": msg,
	}
}

func main() {
	conn, err := grpc.Dial(serverPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := proto.NewDemoClient(conn)

	g := gin.Default()

	g.Use(cors.AllowAll())

	g.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		res, err := client.Login(ctx, &proto.AuthRequest{Username: username, Password: password})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"token": res.Token,
			}))
		}
	})

	g.POST("/register", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		res, err := client.Register(ctx, &proto.AuthRequest{Username: username, Password: password})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"token": res.Token,
			}))
		}
	})

	g.POST("/listBlog", func(ctx *gin.Context) {
		pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
		pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
		res, err := client.ListBlog(ctx, &proto.ListBlogRequest{PageNum: int32(pageNum), PageSize: int32(pageSize)})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"total": res.Total,
				"blogs": res.Blogs,
			}))
		}
	})

	g.POST("/addBlog", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		content := ctx.PostForm("content")
		token := ctx.GetHeader("Authorization")
		author, err := auth.GetUsername(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
			return
		}
		blog := &proto.Blog{Title: title, Content: content, Author: author}
		res, err := client.AddBlog(ctx, &proto.AddBlogRequest{Blog: blog})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"blogId": res.BlogId,
			}))
		}
	})

	g.GET("/blog/:id", func(ctx *gin.Context) {
		blogId, _ := strconv.Atoi(ctx.Param("id"))
		res, err := client.GetBlog(ctx, &proto.GetBlogRequest{BlogId: int32(blogId)})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"blog": res.Blog,
			}))
		}
	})

	g.GET("/follow/:follower", func(ctx *gin.Context) {
		follower := ctx.Param("follower")
		token := ctx.GetHeader("Authorization")
		parent, err := auth.GetUsername(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
			return
		}
		_, err = client.Follow(ctx, &proto.FollowRequest{Parent: parent, Follower: follower})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{}))
		}
	})

	g.GET("/isFollow/:follower", func(ctx *gin.Context) {
		follower := ctx.Param("follower")
		token := ctx.GetHeader("Authorization")
		parent, err := auth.GetUsername(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
			return
		}
		res, err := client.IsFollow(ctx, &proto.FollowRequest{Parent: parent, Follower: follower})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"res": res.Res,
			}))
		}
	})

	//----------------------------------------------------------------------------
	g.POST("/listBlog/:user", func(ctx *gin.Context) {
		user := ctx.Param("user")
		pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
		pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
		res, err := client.GetUserBlog(ctx, &proto.UserBlogRequest{Username: user, PageNum: int32(pageNum), PageSize: int32(pageSize)})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"total": res.Total,
				"blogs": res.Blogs,
			}))
		}
	})

	//Maybe users can only see their own follow list?
	//how to get user name from jwt/token?
	g.GET("/follow_list/:follower", func(ctx *gin.Context) {
		follower := ctx.Param("follower")
		res, err := client.GetUserFollowed(ctx, &proto.UserRequest{Username: follower})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"total": res.Total,
				"parents": res.Parents,
			}))
		}
	})

	g.GET("/follower_list/:parent", func(ctx *gin.Context) {
		follower := ctx.Param("parent")
		res, err := client.GetUserFollower(ctx, &proto.UserRequest{Username: follower})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"total": res.Total,
				"followers": res.Parents,
			}))
		}
	})

	g.POST("/listFollowBlog/:user", func(ctx *gin.Context) {
		user := ctx.Param("user")
		pageNum, _ := strconv.Atoi(ctx.PostForm("pageNum"))
		pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))
		res, err := client.GetFollowBlog(ctx, &proto.UserBlogRequest{Username: user, PageNum: int32(pageNum), PageSize: int32(pageSize)})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fail(err.Error()))
		} else {
			ctx.JSON(http.StatusOK, success(gin.H{
				"total": res.Total,
				"blogs": res.Blogs,
			}))
		}
	})

	//-----------------------------------------------------------------------------
	if err := g.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
package web

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

func Start() {
	// create client connections to all RPC servers
	dialAll()
	g := gin.Default()
	g.Use(cors.AllowAll())

	// expose REST APIs for frontend
	g.POST("/register", register)
	g.POST("/login", login)
	g.POST("/listBlog", listBlog)
	g.POST("/addBlog", addBlog)
	g.GET("/blog/:id", getBlog)
	g.GET("/follow/:follower", follow)
	g.GET("/isFollow/:follower", isFollow)
	g.POST("/listBlog/:user", listBlogByUser)
	g.GET("/follow_list/:follower", getFollowListByFollower)
	g.GET("/follower_list/:parent", getFollowListByParent)
	g.POST("/listFollowBlog/:user", listFollowBlogByUser)

	if err := g.Run(port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

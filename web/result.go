package web

import "github.com/gin-gonic/gin"

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
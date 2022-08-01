package web

import (
	"fmt"
	"kreutzer/pkg/cluster"
	"kreutzer/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTION")
		c.Next()
	}
}

func StartWebServer() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/namespace", func(ctx *gin.Context) {
		cluster.ListNamespaces()
		ctx.JSON(http.StatusOK, gin.H{
			"namespace": cluster.ListNamespaces(),
		})
	})

	v1 := router.Group("v1")
	{
		v1.Use(Cors())
		v1.GET("/user/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})
	}

	_ = router.Run(fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetUint16("server.port")))
}

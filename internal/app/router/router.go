package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterRouter(r *gin.Engine) {

	//解决跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Any("/echo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello world!"})
	})
	////路由
	//api := r.Group("/api")
	//
	////获取用户信息
	//user := api.Group("/v1/user")
	//user.POST("/get_user_info")
}

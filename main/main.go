package main

import (
	"github.com/gin-gonic/gin"
	"org.beijing.com/Scaffolding-web-han/Scaffolding-web-han/control"
)

func main() {
	r := gin.Default()
	//静态资源处理
	r.Static("/static", "../static")
	//静态页面处理
	r.LoadHTMLGlob("../templates/*")
	//根路径访问登录页面
	r.GET("/", control.Login)
	//用户控制器
	userRoute := r.Group("/user")
	{
		userRoute.POST("/login", control.LoginUser)
	}

	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		//flog.Error(err.Error())
	}

}

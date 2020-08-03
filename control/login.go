package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"org.beijing.com/Scaffolding-web-han/Scaffolding-web-han/until"
)

//登录页面访问
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
	flog, err := until.Log_init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	flog.Info("访问登录页面成功")
}

func LoginUser(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

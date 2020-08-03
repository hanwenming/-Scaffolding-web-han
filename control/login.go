package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录页面访问
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
	//flog ,err := nowLog.Log_init()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//flog.Info("访问登录页面成功")
}

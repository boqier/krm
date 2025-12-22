package auth

import (
	"net/http"

	"github.com/boqier/krm/config"
	"github.com/boqier/krm/utils/jwtutil"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	//1.获取前端的用户名与密码：
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		r.JSON(http.StatusOK, gin.H{
			"message": "登录失败",
			"status":  401,
		})
		return
	}
	logs.Info(map[string]interface{}{"module": "auth"}, "用户登录："+userInfo.Username)
	//2.验证用户名与密码是否正确
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		ss, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			r.JSON(http.StatusOK, gin.H{
				"message": "生成token失败",
				"status":  401,
			})
			return
		}
		//3.返回token
		logs.Info(map[string]interface{}{"module": "auth"}, "用户登录成功，生成token："+ss)
		data := make(map[string]interface{})
		data["token"] = ss
		r.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"status":  200,
			"data":    data,
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "登录失败,用户名或密码错误",
		"status":  401,
	})
}
func Logout(r *gin.Context) {
	r.JSON(http.StatusOK, gin.H{
		"message": "退出成功",
		"status":  200,
	})
}

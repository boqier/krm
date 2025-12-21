package middlerwares

import (
	"github.com/boqier/krm/utils/jwtutil"
	"github.com/boqier/krm/utils/logs"
	"github.com/gin-gonic/gin"
)

func JWTAuth(r *gin.Context) {
	//除login与logout之外都要验证token
	requestUrl := r.FullPath()
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Info(map[string]interface{}{"module": "middlerwares"}, "登录与退出接口，不验证token")
		return
	}
	//验证token
	token := r.Request.Header.Get("Authorization")
	if token == "" {
		logs.Error(map[string]interface{}{"module": "middlerwares"}, "token为空")
		r.JSON(200, gin.H{"status": 401, "msg": "token为空"})
		r.Abort()
		return
	}
	//验证token是否有效
	claims, err := jwtutil.ParseToken(token)
	if err != nil {
		logs.Error(map[string]interface{}{"module": "middlerwares"}, "token解析失败")
		r.JSON(200, gin.H{"status": 401, "msg": "token解析失败"})
		r.Abort()
		return
	}
	//将claims放入上下文
	r.Set("claims", claims)
	r.Next()

}

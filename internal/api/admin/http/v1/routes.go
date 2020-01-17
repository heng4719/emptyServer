/**
 * Created by Wangwei on 2019-06-05 11:26.
 */

package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router := r.Group("/v1/admin_api")
	SetupNoneAuthorized(router)
}

// 不需要token认证的接口
func SetupNoneAuthorized(router gin.IRouter) {
	//authController := AuthController{}
	//router.POST("/login", authController.Login)
}

// 需要token认证的接口
func SetupAuthorized(router gin.IRouter) {

}

package main

import (
	"net/http"
	c "vipback/controller"
	_ "vipback/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router = gin.Default()
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /v1
func main() {
	gin.SetMode(gin.ReleaseMode)
	initRouter()
	url := ginSwagger.URL("http://localhost:8777/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":8777")
	select {}
}
func initRouter() {
	router.MaxMultipartMemory = 8 << 20 //8 MiB
	router.StaticFS("/v1/static", http.Dir("./static"))
	userGroup := router.Group("/v1/user")
	c.AddUserRouter(userGroup)
	userGroupShop := router.Group("/v1/shop")
	c.AddShopRouter(userGroupShop)
}

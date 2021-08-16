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

// @host petstore.swagger.io
// @BasePath /
func main() {
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong1111",
	// 	})
	// })
	gin.SetMode(gin.ReleaseMode)
	// dsn := "vip:8DPRDTGZK5BphnX7@tcp(103.153.139.80:3306)/vip?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// db.AutoMigrate(&model.User{})
	initRouter()
	url := ginSwagger.URL("http://localhost:8777/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":8777")
	// }
}
func initRouter() {
	router.MaxMultipartMemory = 8 << 20 //8 MiB
	router.StaticFS("/static", http.Dir("./static"))
	userGroup := router.Group("/user")
	c.AddUserRouter(userGroup)
}

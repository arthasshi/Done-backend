package controller

import (
	"fmt"
	"log"
	"strconv"
	"vipback/model"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("user controller")
}
func AddUserRouter(rg *gin.RouterGroup) {

	rg.POST("/regist", handlerRegist)
	rg.GET("/", handlerUserList)
	rg.POST("/:id/file", handlerUploadFile)
}

// @summary regist
// @Description user regist
// @Produce  json
// @Param user body string true "reg user data"
// @Resource obj
// @Router /regist/ [post]
// @Success 200 {object} string
func handlerRegist(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(200, model.ResJsonType{
			Code: model.CECODE,
			Msg:  err.Error(),
		})
		return
	}
	id, err := model.RegistUser(&user)
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "regist success",
			Data: id,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary get user list
// @Description this api will return user list by page
// @Produce  json
// @Param page query int false "page"
// @Param size query int false "the req size,if null ,will get all users"
// @Resource obj
// @Router / [get]
// @Success 200 {object} string
func handlerUserList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "-1")
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	list, total, err := model.UserList(pageInt, sizeInt)
	if err == nil {
		res := make(map[string]interface{})
		res["total"] = total
		res["list"] = list
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "success",
			Data: res,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary upload file
// @Description upload file
// @Produce  json
// @Param photo body string true "the file data"
// @Param uid path string true "the user id"
// @Resource obj
// @Router /:id/file [post]
// @Success 200 {object} string
func handlerUploadFile(c *gin.Context) {
	id := c.Param("id")
	if file, err := c.FormFile("file"); err != nil {
		log.Println(err)
		return
	} else {
		dst := fmt.Sprintf(`./static/uploads/` + "id_" + id)
		err := c.SaveUploadedFile(file, dst)
		if err == nil {
			c.JSON(200, model.ResJsonType{
				Code: 200,
				Msg:  "upload success",
			})
		} else {
			c.JSON(200, model.ResJsonType{
				Code: model.SECODE,
				Msg:  err.Error(),
			})
		}
	}

}

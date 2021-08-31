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
	rg.PUT("/:id", handlerChangeUserStatus)
	rg.PUT("/inte/:id", handlerChangeUserInte)
	rg.POST("/login", handlerLogin)
}

// @summary regist
// @Description 新增门店管理员，新增会员都是这个接口，post的参数记得带role，1门店管理员，2会员，其中shopid为门店分组,门店管理员新增会员的时候，会员的shop_id默认为门店管理员的shop_id
// @Produce  json
// @Param user body string true "reg user data"
// @Param shopmid query string false "门店管理员的id，只有新增会员需要传递这个"
// @Resource obj
// @Router /user/regist/ [post]
// @Success 200 {object} string
func handlerRegist(c *gin.Context) {
	var user model.User
	var smIdInt uint64
	shopMaId := c.Query("shopmid")
	if shopMaId != "" {
		smIdInt, _ = strconv.ParseUint(shopMaId, 10, 0)
	}
	if err := c.Bind(&user); err != nil {
		c.JSON(200, model.ResJsonType{
			Code: model.CECODE,
			Msg:  err.Error(),
		})
		return
	}
	id, err := model.RegistUser(&user, uint(smIdInt))
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

// @summary 登录
// @Description 登录必传的参数为用户名和密码，登录成功会返回用户的全部信息
// @Produce  json
// @Param user body string true "reg user data"
// @Resource obj
// @Router /user/login [post]
// @Success 200 {object} string
func handlerLogin(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(200, model.ResJsonType{
			Code: model.CECODE,
			Msg:  err.Error(),
		})
		return
	}
	res, err := model.Login(&user)
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "登录成功",
			Data: res,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary PUT改变会员/管理员的状态，冻结什么的操作都是这个
// @Description 改变会员/管理员的状态，冻结什么的操作都是这个，调用成功后会将此用户返回，用于更新
// @Produce  json
// @Param id query int false "用户的ID"
// @Param status query int false "要改变的值"
// @Resource obj
// @Router /user/:id [put]
// @Success 200 {object} string
func handlerChangeUserStatus(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseUint(id, 10, 0)
	status := c.Query("status")
	statusInt, _ := strconv.Atoi(status)
	user, err := model.ChangeStatus(statusInt, uint(idInt))
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "修改成功",
			Data: user,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary PUT改变会员的积分
// @Description 改变会员的积分，传递过来的是钱，不是换算后的积分
// @Produce  json
// @Param id query int false "用户的ID"
// @Param money query int false "消费金额"
// @Resource obj
// @Router /user/inte/:id [put]
// @Success 200 {object} string
func handlerChangeUserInte(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseUint(id, 10, 0)
	money := c.Query("money")
	moneyInt, _ := strconv.ParseFloat(money, 10)
	user, err := model.ChangeUserInte(moneyInt, uint(idInt))
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "修改成功",
			Data: user,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary GET获取用户列表
// @Description 获取用户列表，分页，可以根据角色进行区分
// @Produce  json
// @Param page query int false "page"
// @Param size query int false "the req size,if null ,will get all users"
// @Param role query int false "角色 admin：0，门店管理员：1，会员2:"
// @Resource obj
// @Router /user/ [get]
// @Success 200 {object} string
func handlerUserList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "-1")
	role := c.DefaultQuery("role", "2") //当role不传时，默认获取会员
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	roleInt, _ := strconv.Atoi(role)
	list, total, err := model.UserList(pageInt, sizeInt, roleInt)
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

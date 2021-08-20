package controller

import (
	"strconv"
	"vipback/model"

	"github.com/gin-gonic/gin"
)

func AddShopRouter(rg *gin.RouterGroup) {
	rg.POST("/", handlerAddShop)
	rg.GET("/", handlerShops)
	rg.PUT("/:id", handlerChangeShopStatus)
	rg.PUT("/:id/scale", handlerChangeShopScale)
}

// @summary 新增店铺
// @Description 新增店铺，必须包含name,address 字段
// @Produce  json
// @Param shop body string true "店铺信息"
// @Resource obj
// @Router /shop/ [post]
// @Success 200 {object} string
func handlerAddShop(c *gin.Context) {
	var shop model.Shop
	if err := c.Bind(&shop); err != nil {
		c.JSON(200, model.ResJsonType{
			Code: model.CECODE,
			Msg:  err.Error(),
		})
		return
	}
	id, err := model.AddShop(&shop)
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "添加成功",
			Data: id,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

// @summary 获取店铺列表
// @Description 获取店铺列表，店铺页面，创建管理员，创建会员都会用到，可分页，
// @Produce  json
// @Param page query int false "page"
// @Param size query int false "size，如果不传，默认获取所有"
// @Resource obj
// @Router /shop/ [get]
// @Success 200 {object} string
func handlerShops(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "-1")
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	list, total, err := model.GetShopList(pageInt, sizeInt)
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

// @summary PUT改变店铺的状态
// @Description 改变店铺的状态，0关闭，1开张
// @Produce  json
// @Param id query int false "店铺的id"
// @Param status query int false "要改变的状态值"
// @Resource obj
// @Router /shop/:id [put]
// @Success 200 {object} string
func handlerChangeShopStatus(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseUint(id, 10, 0)
	status := c.Query("status")
	statusInt, _ := strconv.Atoi(status)
	err := model.ChangeShopStatus(uint(idInt), statusInt)
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "修改成功",
			Data: nil,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}

}

// @summary PUT改变店铺的积分比例，如果是管理员则直接修改所有店铺的积分比例
// @Description PUT改变店铺的积分比例，如果是管理员则传递店铺id的值为0，直接修改所有店铺的积分比例，
// @Produce  json
// @Param id query int false "店铺的id，管理员账户传0，慎用"
// @Param scale query int false "要改变的比例值，浮点"
// @Resource obj
// @Router /shop/:id/scale [put]
// @Success 200 {object} string
func handlerChangeShopScale(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseUint(id, 10, 0)
	scale := c.Query("scale")
	scaleFloat, _ := strconv.ParseFloat(scale, 10)
	err := model.ChangeShopScale(uint(idInt), scaleFloat)
	if err == nil {
		c.JSON(200, model.ResJsonType{
			Code: 200,
			Msg:  "修改成功",
			Data: nil,
		})
	} else {
		c.JSON(200, model.ResJsonType{
			Code: model.SECODE,
			Msg:  err.Error(),
		})
	}
}

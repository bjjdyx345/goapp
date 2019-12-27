package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/app/manageweb/controllers/card"
)

func RegisterRouterCard(app *gin.RouterGroup) {
	door:= card.DoorcardCon{}
	app.POST("/card/create",door.Create)
	app.POST("/card/delete",door.DeleteById)
	app.POST("/card/update",door.Update)
	app.GET("/card/getall",door.Getall)
	app.GET("/card/list",door.List)
	app.GET("/card/detail",door.Detail)
	app.POST("/card/bindvillage",door.BindApartmentId)
	//app.GET("/card/query",door.Query)

	village:= card.VillageCon{}
	app.POST("/village/create",village.Create)
	app.POST("/village/delete",village.DeleteById)
	app.POST("/village/update",village.Update)
	app.GET("/village/getall",village.Getall)
	app.GET("/village/list",village.List)
	app.GET("/village/detail",village.Detail)
	app.GET("/village/getalladdress",village.GetAllAddress)

	order:= card.OrderCon{}
	app.POST("/order/create",order.Create)
	app.GET("/order/getall",order.Getall)
	app.POST("/order/delete",order.DeleteById)
	app.POST("/order/update",order.Update)
	app.GET("/order/list",order.List)
	app.GET("/order/detail",order.Detail)

	city:= card.CityInfoCon{}
	app.GET("/city/getfirstcity",city.GetFirstCityAll)
	app.GET("/city/getsecondcity",city.GetSecondCity)
	app.GET("/city/getthirdcity",city.GetThirdCity)
}

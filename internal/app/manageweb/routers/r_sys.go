package routers

import (
	"github.com/it234/goapp/internal/app/manageweb/controllers/sys"
	"github.com/gin-gonic/gin"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := sys.Menu{}
	app.GET("/menu/list", menu.List)
	app.GET("/menu/detail", menu.Detail)
	app.GET("/menu/allmenu", menu.AllMenu)
	app.GET("/menu/menubuttonlist", menu.MenuButtonList)
	app.POST("/menu/delete", menu.Delete)
	app.POST("/menu/update", menu.Update)
	app.POST("/menu/create", menu.Create)
	user := sys.User{}
	app.GET("/user/info", user.Info)
	app.POST("/user/login", user.Login)
	app.POST("/user/logout", user.Logout)
	app.POST("/user/editpwd", user.EditPwd)
	admins := sys.Admins{}
	app.GET("/admins/list", admins.List)
	app.GET("/admins/detail", admins.Detail)
	app.GET("/admins/adminsroleidlist", admins.AdminsRoleIDList)
	app.POST("/admins/delete", admins.Delete)
	app.POST("/admins/update", admins.Update)
	app.POST("/admins/create", admins.Create)
	app.POST("/admins/setrole", admins.SetRole)
	role := sys.Role{}
	app.GET("/role/list", role.List)
	app.GET("/role/detail", role.Detail)
	app.GET("/role/rolemenuidlist", role.RoleMenuIDList)
	app.GET("/role/allrole", role.AllRole)
	app.POST("/role/delete", role.Delete)
	app.POST("/role/update", role.Update)
	app.POST("/role/create", role.Create)
	app.POST("/role/setrole", role.SetRole)

	door:=sys.DoorcardCon{}
	app.POST("/card/create",door.Create)
	app.POST("/card/delete",door.DeleteById)
	app.POST("/card/update",door.Update)
	app.GET("/card/getall",door.Getall)
	app.GET("/card/list",door.List)
	app.GET("/card/detail",door.Detail)
	app.POST("/card/bindvillage",door.BindApartmentId)
	//app.GET("/card/query",door.Query)

	village:=sys.VillageCon{}
	app.POST("/village/create",village.Create)
	app.POST("/village/delete",village.DeleteById)
	app.POST("/village/update",village.Update)
	app.GET("/village/getall",village.Getall)
	app.GET("/village/list",village.List)
	app.GET("/village/detail",village.Detail)
	app.GET("/village/getalladdress",village.GetAllAddress)

	order:=sys.OrderCon{}
	app.POST("/order/create",order.Create)
	app.GET("/order/getall",order.Getall)
	app.POST("/order/delete",order.DeleteById)
	app.POST("/order/update",order.Update)
	app.GET("/order/list",order.List)
	app.GET("/order/detail",order.Detail)

	city:=sys.CityInfoCon{}
	app.GET("/city/getfirstcity",city.GetFirstCityAll)
	app.GET("/city/getsecondcity",city.GetSecondCity)
	app.GET("/city/getthirdcity",city.GetThirdCity)
/*
	order:=sys.Order{}
	app.Get("/order/create",order.Add)
*/

}

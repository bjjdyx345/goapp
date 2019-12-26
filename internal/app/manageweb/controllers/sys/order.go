package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	"github.com/it234/goapp/internal/pkg/models/sys"
	models "github.com/it234/goapp/internal/pkg/models/common"
	"net/http"
)

type OrderCon struct {}

/*
兼容作者写法，剩余的函数可自行调用
*/

func (OrderCon) Detail(c *gin.Context) {
	id := common.GetQueryToUint64(c, "id")
	var order sys.Order
	where := sys.Order{}
	where.ID = id
	_, err := models.First(&where, &order)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &order)
}

// 分页数据
func (OrderCon) List(c *gin.Context) {
	page := common.GetPageIndex(c)
	limit := common.GetPageLimit(c)
	sort := common.GetPageSort(c)
	key := common.GetPageKey(c)
	var whereOrder []models.PageWhereOrder
	order := "ID DESC"
	if len(sort) >= 2 {
		orderType := sort[0:1]
		order = sort[1:len(sort)]
		if orderType == "+" {
			order += " ASC"
		} else {
			order += " DESC"
		}
	}
	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
	if key != "" {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		// 此处需要修改 code name？？？
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "name like ? or code like ?", Value: arr})
	}
	var total uint64
	list:= []sys.Order{}
	err := models.GetPage(&sys.Order{}, &sys.Order{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessPage(c, total, &list)
}


func (OrderCon) DeleteById(c *gin.Context) {
	var ids []uint64
	err := c.Bind(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(c, err)
		return
	}
	// 这需要改
	//err=sys.DeleteOneOrder(ids)
	if err != nil{
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessMsg(c)
}


func (OrderCon) Getall(c *gin.Context) {
	// 所有菜单
	var orders []sys.Order
	err := models.Find(&sys.Order{}, &orders, "id asc", "order_status asc")
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &orders)
}


// 更新
func (OrderCon) Update(c *gin.Context) {
	model := sys.Order{}
	err := c.Bind(&model)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	err = models.Save(&model)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccessMsg(c)
}

//新增
func (OrderCon) Create(c *gin.Context) {
	order := sys.Order{}
	err := c.Bind(&order)

	if err != nil {
		fmt.Println("bind err,",err)
		common.ResErrSrv(c, err)
		return
	}
	err=sys.IsPhoneNum(order.OrderPhoneNumber)
	if err != nil {
		fmt.Println("phone number err,",err)
		common.ResFail(c, "操作失败")
		return
	}
	err = models.Create(&order)
	if err != nil {
		fmt.Println("create err,",err)
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccess(c, gin.H{"id": order.ID})
}









/*
parameters
input:
	username
	phonenumber
	address
	price


OrderStatus
	1 	已下单待审核
	2	已审核待配送
	3	正在配送
	4	配送完成
*/
func (OrderCon) Add(c *gin.Context){
	var orderjson=sys.Order{}
	//addressid 等参数并未做实际,下单时cardid并没有
	if err:=c.BindJSON(&orderjson);err!=nil&&(sys.IsnotNullOrder(orderjson)!=nil){
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var queryvillage sys.Village
	queryvillage.ID=sys.StringToUint64(orderjson.OrderAddress)
	_,res_count:=sys.QueryVillage(queryvillage)
	if(res_count<1){
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"village ID is unusable,please check it",
			"data":"",
			"trace_id":"",
		})
	}else{
		orderjson.OrderStatus="0"
		if(sys.IsPhoneNum(orderjson.OrderPhoneNumber)==nil){
			err:=sys.AddNewOrder(orderjson)
			if(err==nil){
				c.JSON(200,gin.H{
					"errno":0,
					"errmsg":"Create the new order successfully",
					"data":orderjson,
					"trace_id":"",
				})
			}else{
				c.JSON(200,gin.H{
					"errno":-2,
					"errmsg":err,
					"data":orderjson,
					"trace_id":"",
				})
			}

		}else{
			c.JSON(200,gin.H{
				"errno":-3,
				"errmsg":"phone number is not correct",
				"data":"",
				"trace_id":"",
			})
		}
	}
}
/*
参数
	id &
	useranme &
	phonenumber
*/
func (OrderCon) Delete_back(c *gin.Context){
	var orderjson=sys.Order{}
	if err:=c.BindJSON(&orderjson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if(orderjson.ID>0){
		var queryorder=sys.Order{}
		queryorder.ID=orderjson.ID
		_,res_count:=sys.QueryOrder(orderjson)
		if(res_count>0){
			err:=sys.DeleteOrders(orderjson)
			if(err==nil){
				c.JSON(200,gin.H{
					"errno":0,
					"errmsg":"delete successful",
					"data":"",
					"data_count":res_count,
					"trace_id":"",
				})
			}else {
				c.JSON(200,gin.H{
					"errno":-1,
					"errmsg":"delete failed",
					"data":"",
					"data_count":0,
					"trace_id":"",
				})
			}
		}else {
			c.JSON(200,gin.H{
				"errno":-2,
				"errmsg":"ID is invalid",
				"data":"",
				"data_count":0,
				"trace_id":"",
			})
		}
	}else{
		c.JSON(200,gin.H{
			"errno":-3,
			"errmsg":" please input parameters id",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
	}
}

/*
parameters：
Input
	username & phonenumber

*/
func (OrderCon) Query_back(c *gin.Context) {
	var orderjson = sys.Order{}
	if err := c.BindJSON(&orderjson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//if ID /username/phonenum are not null
	if(sys.IsPhoneNum(orderjson.OrderPhoneNumber)==nil&&orderjson.OrderUsername!=""){
		ordermap,res_count:=sys.QueryOrder(orderjson)
		if(res_count>0){
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"query successfully,find records",
				"data":ordermap,
				"data_count":res_count,
				"trace_id":"",
			})
		} else{
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":"nothing can find",
				"data":ordermap,
				"data_count":res_count,
				"trace_id":"",
			})
		}
	}else{
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"please input correct username and phone number",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
	}
}

/*
parameters
Input
	id

*/

func (OrderCon) Update_back(c *gin.Context) {
	var orderjson = sys.Order{}

	if err := c.BindJSON(&orderjson); (err != nil && sys.IsnotNullOrder(orderjson) != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if(orderjson.ID>0){
		if err:=sys.UpdateOrder(orderjson);err==nil{
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"Update successfully",
				"data":orderjson,
				"trace_id":"",
			})
		}else{
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":err,
				"data":orderjson,
				"trace_id":"",
			})

		}

	}else{
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"please input correct orderID",
			"data":"",
			"trace_id":"",
		})
	}
}


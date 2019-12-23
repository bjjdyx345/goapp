package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/pkg/models/sys"
	"net/http"
)

type OrderCon struct {}


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
func (OrderCon) Delete(c *gin.Context){
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
func (OrderCon) Query(c *gin.Context) {
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

func (OrderCon) Update(c *gin.Context) {
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


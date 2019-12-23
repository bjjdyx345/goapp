package sys

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	models "github.com/it234/goapp/internal/pkg/models/common"
	"github.com/it234/goapp/internal/pkg/models/sys"
	"io/ioutil"
	"net/http"
	_ "strconv"
)

type DoorcardCon struct {}


/*
兼容作者写法，剩余的函数可自行调用
*/

func (DoorcardCon) Detail(c *gin.Context) {
	id := common.GetQueryToUint64(c, "id")
	var card sys.Door_card
	where := sys.Door_card{}
	where.ID = id
	_, err := models.First(&where, &card)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &card)
}

// 分页数据
func (DoorcardCon) List(c *gin.Context) {
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
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "name like ? or code like ?", Value: arr})
	}
	var total uint64
	list:= []sys.Door_card{}
	err := models.GetPage(&sys.Door_card{}, &sys.Door_card{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessPage(c, total, &list)
}


func (DoorcardCon) DeleteById(c *gin.Context) {
	var ids []uint64
	err := c.Bind(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(c, err)
		return
	}
	err=sys.DeleteOneCardById(ids)
	if err != nil{
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessMsg(c)
}


func (DoorcardCon) Getall(c *gin.Context) {
	// 所有菜单
		var doors []sys.Door_card
		err := models.Find(&sys.Door_card{}, &doors, "id asc", "card_id asc")
		if err != nil {
			common.ResErrSrv(c, err)
			return
		}
		common.ResSuccess(c, &doors)
}


// 更新
func (DoorcardCon) Update(c *gin.Context) {
	model := sys.Door_card{}
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
func (DoorcardCon) Create(c *gin.Context) {
	card := sys.Door_card{}
	err := c.Bind(&card)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	err = models.Create(&card)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccess(c, gin.H{"id": card.ID})
}

/*---------------------------分割线，预留函数-----------------------------------------------*/
func(DoorcardCon) Create_back(c *gin.Context){
	var dcjson=sys.Door_card{}
	if err:=c.BindJSON(&dcjson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if((sys.IsCardId(dcjson.CardId)==nil)&&(sys.IsCardType(dcjson.CardType)==0)&& (sys.IsCardFilename(dcjson.CardContentFile)==nil)){
		fmt.Println("----------------------------")
		fmt.Println(dcjson)
		fmt.Println("----------------------------")

		_,res_count:=sys.QueryCardId(dcjson)
		//has same card id please update card
		if(res_count>0){
			//if find the same raws
			c.JSON(200,gin.H{
				"errno":-2,
				"errmsg":"has same cardid ,please update the data",
				"data":"",
				"trace_id":"",
			})

		}else{
			//add new card to table
			err:=sys.AddNewCard(dcjson)
			if err!=nil{
				//if add new rows wrong
				c.JSON(200,gin.H{
					"errno":-1,
					"errmsg":"add rows error",
					"data":dcjson,
					"trace_id":"",
				})

			} else{
				//add successful
				c.JSON(200,gin.H{
					"errno":0,
					"errmsg":"successful",
					"data":dcjson,
					"trace_id":"",
				})
			}
		}
	}else {
		//the para has wrong format
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"parameters has wrong format",
			"data":"",
			"trace_id":"",
		})
	}

}

/*
errorcode
-2	same cardid
-1	add error
0	successful
*/
func (DoorcardCon) Addlist(c *gin.Context) {
	var Cdc = sys.Door_card{}
	//判断参数是否正确
	cardid :=c.Query("cardid")
	cardtype :=c.Query("cardType")
	cardcontentfile :=c.Query("cardContentFile")
	belong2apartment :=c.Query("belongToApartmentId")
	fmt.Println(cardtype,cardid,cardcontentfile)
	if((sys.IsCardId(cardid)==nil)&&(sys.IsCardFilename(cardcontentfile)==nil)&&(sys.IsCardType(cardtype)==0)){
		//此处需要判断是否belong2apartment为可用。后期再补充
		Cdc.CardId =cardid
		_,res_count:=sys.QueryCardId(Cdc)
		//has same card id please update card
		if(res_count>0){
			//if find the same raws
			c.JSON(200,gin.H{
				"errno":-2,
				"errmsg":"has same cardid ,please update the data",
				"data":"",
				"trace_id":"",
			})

		}else{
			//add new card to table
			Cdc.CardType =cardtype
			Cdc.CardContentFile =cardcontentfile
			Cdc.BelongToApartmentId =belong2apartment
			err:=sys.AddNewCard(Cdc)
			if err!=nil{
				//if add new rows wrong
				c.JSON(200,gin.H{
					"errno":-1,
					"errmsg":"add rows error",
					"data":Cdc,
					"trace_id":"",
				})

			} else{
				//add successful
				c.JSON(200,gin.H{
					"errno":0,
					"errmsg":"successful",
					"data":Cdc,
					"trace_id":"",
				})
			}
		}
	}else {
		//the para has wrong format
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"parameters has wrong format",
			"data":"",
			"trace_id":"",
		})
	}
}
/*
门禁卡修改操作
查询是否有对应的数据
errno:
1		more records
0		successful
-1		find nothing
-2		data has wrong format

*/
func (DoorcardCon) Update_back(c *gin.Context) {
	var dcjson = sys.Door_card{}
	if err:=c.BindJSON(&dcjson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if(dcjson.ID<1||sys.IsCardId(c.Query(dcjson.CardId))!=nil){
		c.JSON(200,gin.H{
			"errno":-3,
			"errmsg":"please upload id  ",
			"data":"",
			"trace_id":"",
		})
		return
	}
	//此处可以不查询卡id，如果查询卡id就不能更改id了，查询只按唯一key查
	var (
		_, res_count = sys.QueryCardId(sys.Door_card{
			Model: basemodel.Model{
				ID: dcjson.ID,
			},
			//CardId: Cdc.CardId,
		})
	)

	if(res_count>1){
		//has find more than 1 raws
		c.JSON(200,gin.H{
			"errno":1,
			"errmsg":"error,has find more than one records ",
			"data":"",
			"trace_id":"",
		})
	}else if (res_count<1){
		//find no result
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"has find nothing ",
			"data":"",
			"trace_id":"",
		})
	}else{
		err:=sys.UpdateCard(dcjson)
		if(err!=0){
			c.JSON(200,gin.H{
				"errno":-2,
				"errmsg":"has update failed,please check your data",
				"data":"",
				"trace_id":err,
			})
		}else{
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"has update successful ",
				"data":"",
				"trace_id":"",
			})
		}
	}

}

func (DoorcardCon) Query(c *gin.Context) {
	var dcjson = sys.Door_card{}

	if err:=c.BindJSON(&dcjson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	//保证其中一个有值
	if(sys.IsCardId(dcjson.CardId)!=nil&&sys.IsCardType(dcjson.CardType)!=0){
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"the cardid or cardtype must has value",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
		return
	}
	/*此处需要应答一个cdc结构体，不需要返回值*/
	var doorcdc []sys.Door_card
	var res_count=0
	doorcdc,res_count=sys.QueryCardId(dcjson)
	if(res_count>0){
		c.JSON(200,gin.H{
			"errno":0,
			"errmsg":"has find data rows",
			"data":doorcdc,
			"data_count":res_count,
			"trace_id":"",
		})
	}else {
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"has find nothing ",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
	}

}

/*
传入id，cardid来删除数据
其中id必须，
cardid不是必须的

*/
func (DoorcardCon) Delete(c *gin.Context) {
	var dcjson = sys.Door_card{}

	if err:=c.BindJSON(&dcjson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	//err1:=sys.IsCardId(cid)

	if(dcjson.ID<1){
		c.JSON(200,gin.H{
			"errno":-1,
			"errmsg":"delete error ,is not a valid id",
			"data":"",
			"trace_id":"",
		})
		return
	}else {
		err:=sys.DeleteCards(dcjson)
		if(err==nil){
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":"delete successful",
				"data":dcjson,
				"trace_id":"",
			})
		}else {
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":"delete error",
				"data":dcjson,
				"trace_id":"",
			})
		}
		return
	}
}
//中间件，负责打印请求
func HelloMiddleware() gin.HandlerFunc {
	fmt.Println("USB middleware")
	return func(ctx *gin.Context) {
		data,err := ctx.GetRawData()
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Printf("data: %v\n",string(data))

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
		ctx.Next()
	}

}

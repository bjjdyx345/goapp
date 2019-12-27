package card

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"github.com/it234/goapp/internal/pkg/models/card"
	models "github.com/it234/goapp/internal/pkg/models/common"
	"net/http"
)

type VillageCon struct {}
type Address struct{
	Address_id 		uint64 	`json:"address_id" form:"address_id" binding:"-"`
	Address_content	string 	`json:"address_content" form:"address_content" binding:"-"`
	Address_city	string 	`json:"address_city" form:"address_city" binding:"-"`
	Adress_village	string 	`json:"address_village" form:"address_village" binding:"-"`
	Id_and_address	string 	`json:"id_and_address" form:"id_and_address" binding:"-"`
}
/*
兼容作者写法，剩余的函数可自行调用
*/

func (VillageCon) Detail(c *gin.Context) {
	id := common.GetQueryToUint64(c, "id")
	var village card.Village
	where := card.Village{}
	where.ID = id
	_, err := models.First(&where, &village)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &village)
}

// 分页数据
func (VillageCon) List(c *gin.Context) {
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
	list:= []card.Village{}
	err := models.GetPage(&card.Village{}, &card.Village{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessPage(c, total, &list)
}


func (VillageCon) DeleteById(c *gin.Context) {
	var ids []uint64
	err := c.Bind(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(c, err)
		return
	}
	err= card.DeleteOneVillageById(ids)
	if err != nil{
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessMsg(c)
}


func (VillageCon) Getall(c *gin.Context) {
	// 所有菜单
	var villages []card.Village
	err := models.Find(&card.Village{}, &villages, "id asc", "village_name asc")
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &villages)
}


// 更新
func (VillageCon) Update(c *gin.Context) {
	model := card.Village{}
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
func (VillageCon) Create(c *gin.Context) {
	card := card.Village{}
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

//新增
func (VillageCon) GetAllAddress(c *gin.Context) {
	// 所有菜单
	var queryvil card.Village
	var address []Address
	var addressapd Address
	var villages []card.Village
	var resCount int
	villages,resCount = card.QueryVillage(queryvil)
	if resCount < 1 {
		common.ResErrSrv(c, errors.New("has  not find the record"))
		return
	}
	for i:=0;i<resCount;i++ {
		//fmt.Println(villages[i])
		addressapd.Address_id = villages[i].ID
		addressapd.Address_city= fmt.Sprintf("%s-%s-%s",villages[i].VillageName,villages[i].VillageId,villages[i].VillageApartmentId)
		addressapd.Address_content = fmt.Sprintf("%s-%s-%s-%s-%s",villages[i].VillageAtProvince,villages[i].VillageAtCity,villages[i].VillageAtDistrict,villages[i].VillageAtAddress,addressapd.Address_city)
		addressapd.Id_and_address = fmt.Sprintf("%v|%s|%s|",villages[i].ID,villages[i].VillageAtCity,addressapd.Address_city)
		fmt.Println(addressapd.Id_and_address)
		address = append(address, addressapd)
	}
	common.ResSuccess(c,&address)
}




/*----------------------分割线----------------*/
func (VillageCon) Add_back(c *gin.Context){
	var villagejson= card.Village{}
	if err:=c.BindJSON(&villagejson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	_,res_count:= card.QueryVillage(villagejson);
	if(res_count>0){
		c.JSON(200,gin.H{
			"errno":-3,
			"errmsg":"has the same data ,please directly use the data",
			"data":"",
			"trace_id":"",
		})
		return
	}
	//if has exist the city
	if(card.IsExistCity(villagejson.VillageAtDistrict)>0){
		//entireaddress:=sys.GetCityInfo(villagejson.VillageAtDistrict)[0].MergerName
		err:= card.AddNewVillage(villagejson)
		if(err==nil){
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"successful and return the records",
				"data":villagejson,
				"trace_id":"",
			})
		}else{
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":err,
				"data":"",
				"trace_id":"",
			})
		}
	} else{
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"has not find the entire address",
			"data":"",
			"trace_id":"",
		})
	}
}

func (VillageCon) Query_back(c *gin.Context){
	var villagejson= card.Village{}
	if err:=c.BindJSON(&villagejson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	//if can query village address,we can return the json
	if(villagejson.VillageAtProvince!=""||villagejson.VillageAtDistrict!=""||villagejson.VillageAtCity!=""){
		//entireaddress:=sys.GetCityInfo(villagejson.VillageAtDistrict)[0].MergerName

		res,res_count:= card.QueryVillage(villagejson)
		if(res_count>0){
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"find raws and return the records",
				"data":res,
				"data_count":res_count,
				"trace_id":"",
			})
		}else{
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":"has not find about records",
				"data":0,
				"trace_id":"",
			})
		}
	} else{
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"wrong data name,please check the data you uploaded",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
	}
}


func (VillageCon) Delete_back(c *gin.Context){
	var villagejson= card.Village{}
	var res_count int
	var err error
	if err:=c.BindJSON(&villagejson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	//if can query village address,we can return the json
	if(card.IsNotNull(villagejson)==nil){
		//entireaddress:=sys.GetCityInfo(villagejson.VillageAtDistrict)[0].MergerName
		_,res_count= card.QueryVillage(villagejson)
		if(res_count<1){
			c.JSON(200,gin.H{
				"errno":-3,
				"errmsg":"no such records",
				"data":"",
				"data_count":0,
				"trace_id":"",
			})
			return
		}
		res_count,err= card.DeleteVillages(villagejson)
		if(err==nil){
			c.JSON(200,gin.H{
				"errno":0,
				"errmsg":"Delete successful",
				"data":"",
				"data_count":res_count,
				"trace_id":"",
			})
		}else{
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":err,
				"data":"",
				"data_count":0,
				"trace_id":"",
			})
		}
	} else{
		c.JSON(200,gin.H{
			"errno":-2,
			"errmsg":"please write the complete address",
			"data":"",
			"data_count":0,
			"trace_id":"",
		})
	}
}

/*
error code
0	successful
-1	find more than one records
-2	can't find records
-3 	data struct is not available
*/

func (VillageCon)Update_back(c *gin.Context){
	var villagejson= card.Village{}
	if err:=c.BindJSON(&villagejson);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	//villagejson.ID=sys.StringToUint64(c.Query("id"))
	//all cow are not null
	if(card.IsAllNotNull(villagejson)==nil&&(villagejson.ID>0)){
		var (
			_,res_count= card.QueryVillage(card.Village{
				Model: basemodel.Model{
					ID: villagejson.ID,
				},
			})
		)
		fmt.Println(res_count)
		fmt.Println(villagejson.ID)
		if(res_count>1){
			c.JSON(200,gin.H{
				"errno":-1,
				"errmsg":"error,has find more than one records,please delete",
				"data":"",
				"trace_id":"",
			})
		}else if(res_count==1){
			card.UpdateVillage(villagejson)
				c.JSON(200,gin.H{
					"errno":0,
					"errmsg":",successful update one records",
					"data":"",
					"trace_id":"",
				})
		}else{// rescount<1
			c.JSON(200,gin.H{
				"errno":-2,
				"errmsg":"can't find the record you want to update ",
				"data":"",
				"trace_id":"",
			})
		}
	}else{
		c.JSON(200,gin.H{
			"errno":-3,
			"errmsg":"please check the struct is complete and the ID is not null and must be int type",
			"data":"",
			"trace_id":"",
		})
	}
}
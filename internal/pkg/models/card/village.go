package card
import (
	"errors"
	"fmt"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"github.com/it234/goapp/internal/pkg/models/db"
	"github.com/it234/goapp/internal/pkg/models/sys"
	"time"
)
//订单表结构
type Village struct {
	basemodel.Model
	VillageId			string	`gorm:"column:village_id;type:varchar(30);" json:"village_id" form:"village_id" binding:"-"`
	VillageName 		string 	`gorm:"column:village_name;type:varchar(30);" json:"village_name" form:"village_name" binding:"required"`
	VillageApartmentId 	string	`gorm:"column:village_apartment_id;type:varchar(30);" json:"village_apartment_id" form:"village_apartment_id" binding:"-"`
	VillageAtProvince	string	`gorm:"column:village_at_province;type:varchar(30);" json:"village_at_province" form:"village_at_province" binding:"required"`
	VillageAtCity		string 	`gorm:"column:village_at_city;type:varchar(30);" json:"village_at_city" form:"village_at_city" binding:"required"`
	VillageAtDistrict	string	`gorm:"column:village_at_district;type:varchar(30);" json:"village_at_district" form:"village_at_district" binding:"required"`
	VillageAtAddress	string	`gorm:"column:village_at_address;type:varchar(255);" json:"village_at_address" form:"village_at_address" binding:"-"`
}
func (Village) TableName() string {
	return sys.TableName("village")
}



//删除卡
func DeleteOneVillageById(ids []uint64) error {
	var villagejson = Village{}
	for _, v := range ids {
		villagejson.ID = v
		if err := db.DB.Delete(villagejson).Error; err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

//添加订单信息
func AddNewVillage(village Village)error{
	village.CreatedAt=time.Now()
	err :=db.DB.Create(&village).Error    //must give pointer
	if err!=nil{
		fmt.Println("插入失败", err)
	}
	return err
}

//更新订单信息
func UpdateVillage(village Village) {
	db.DB.Model(&village).Updates(village)
	fmt.Println(village.ID)
}

//查询订单信息
func QueryVillage(village Village)([]Village,int) {
	var value []Village
	res_count:=1
	res:=db.DB.Where(&village).Find(&value).Count(&res_count)
	if(res.Error==nil) {//has find value
		fmt.Println("get the value")
	}else {
		res_count=-1
		return value,-1
	}
	return value ,res_count
}



//删除订单（未知订单数量）
func DeleteVillages(village Village)(int,error){
	var value,res_count=QueryVillage(village)
	var res_count1 int
	var err error
	if (res_count>0) {
		res_count1,err=DeleteMultiVillages(value)
	}
	return res_count1,err
}
//删除一张订单
func DeleteOneVillage(village Village) error{
	err:=db.DB.Delete(village).Error
	if err!=nil{
		fmt.Println("删除失败", err)
	}else{
		fmt.Println("删除成功",err)
	}
	return err
}
//删除多个订单
func DeleteMultiVillages(village []Village) (int, error){
	var err error
	var i int
	for i=0;i<len(village);i++{
		err=DeleteOneVillage(village[i])
	}
	return i,err
}

func IsAllNotNull(villagejson Village)error{
	if(villagejson.VillageAtProvince!=""&&villagejson.VillageAtDistrict!=""&&villagejson.VillageAtCity!=""&&villagejson.VillageName!=""&&villagejson.VillageId!=""&&villagejson.VillageApartmentId!=""){
		return nil
	}else {
		return errors.New("has null value")
	}
}


func IsNotNull(villagejson Village)error{
	if(villagejson.VillageAtProvince!=""&&villagejson.VillageAtDistrict!=""&&villagejson.VillageAtCity!=""&&villagejson.VillageName!=""&&villagejson.VillageId!=""&&villagejson.VillageApartmentId!=""){
		return nil
	}else {
		return errors.New("has null value")
	}
}
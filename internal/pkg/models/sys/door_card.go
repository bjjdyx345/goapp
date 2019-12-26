package sys

import (
	"fmt"
	"github.com/it234/goapp/internal/pkg/models/db"
	"github.com/jinzhu/gorm"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Door_card struct {
	basemodel.Model
	CardType 				string `gorm:"column:card_type;type:varchar(30);not null;" json:"card_type" binding:"-" form:"card_type"`
	CardId					string `gorm:"column:card_id;type:varchar(30);not null;" json:"card_id" binding:"-" form:"card_id"`
	CardContentFile 		string `gorm:"column:card_contentfile;type:varchar(30);not null;" json:"card_contentfile" binding:"-" form:"card_contentfile"`
	BelongToApartmentId 	string `gorm:"column:belong_to_apartment_id;type:varchar(30);" json:"belong_to_apartment_id" binding:"-" form:"belong_to_apartment_id"`
}

func (Door_card) TableName() string {
	return TableName("doorcard")
}

// 添加前
func (m *Door_card) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

//添加新卡，传入门禁卡信息
func AddNewCard(dc Door_card) error {
	dc.UpdatedAt=time.Now()
	err :=db.DB.Create(&dc).Error    //must give pointer
	if err!=nil{
		fmt.Println("插入失败", err)
	}
	return err
}

//更新卡信息
func UpdateCard(dc Door_card) int{
//	var dc1 Door_card
//	db.DB.First(&dc1)
	db.DB.Model(&dc).Updates(dc)
	fmt.Println(dc.ID)
		return 0
}

//查询卡信息，提供卡id
func  QueryCardId(cardInfo Door_card)([]Door_card,int) {

	var value []Door_card
	res_count:=1
	res:=db.DB.Where(&cardInfo).Find(&value).Count(&res_count)
	//fmt.Println("result has finded,result is")
	if(res.Error==nil) {//has find value
		fmt.Println("get the value")
	}else {
		res_count=-1
		return value,-1
	}
	//fmt.Println(value)
	return value ,res_count
}

func DeleteCards(dc Door_card)error{
	var value,dccount=QueryCardId(dc)
	var err error
	if (dccount>0) {
		err =DeleteMultiCards(value)
	}
	return err
}
//删除卡
func DeleteOneCard(dc Door_card) error{
	err:=db.DB.Delete(dc).Error
	fmt.Println("--------------start-----------")
	fmt.Println(dc)
	fmt.Println("--------------end-------------")
	if err!=nil{
		fmt.Println("删除失败", err)
	}else{
		fmt.Println("删除成功",err)
	}
	return err
}
//删除卡
func DeleteOneCardById(ids []uint64) error {

	for _, v := range ids {
		fmt.Println(v)
		var dcjson = Door_card{}
		dcjson.ID = v
		if err := DeleteOneCard(dcjson); err != nil {
			return err
		}
	}
	return nil
}
//删除卡
func DeleteMultiCards(dc []Door_card) error{
	var err error
	for i:=0;i<len(dc);i++{
		err=DeleteOneCard(dc[i])
	}
	return err
}


func IsCardId(cid string)error{
	var _,err =strconv.Atoi(cid)
	if err !=nil{
		return err
	}
	return nil
}

func IsCardFilename(cfn string)error{
	_,err :=regexp.MatchString("p([a-z]+)ch",cfn)
	if err !=nil{
		return  err
	}
	return nil
}

func IsCardType(ctp string)int{
	err :=strings.Compare(ctp,"ISO14443A")+strings.Compare(ctp,"ISO14443B")+strings.Compare(ctp,"LF125K")
	if(err < -2){
		return err
	}else {
		return 0
	}
}

func IsDepartmentId(cdid string)error{
	_,err :=regexp.MatchString("p([a-z]+)ch",cdid)
	if err !=nil{
		return  err
	}
	return nil
}

func StringToUint64(s string) uint64{
	res,_:=strconv.ParseInt(s,10,64)
	return uint64(res)
}
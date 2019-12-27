package card

import (
	"errors"
	"fmt"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"github.com/it234/goapp/internal/pkg/models/db"
	"github.com/it234/goapp/internal/pkg/models/sys"
	"github.com/jinzhu/gorm"
	"regexp"
	"strconv"
	"time"
)

//订单表结构
type Order struct {
	basemodel.Model
	OrderCardId			string	`gorm:"column:order_card_id;type:varchar(30);not null;" json:"order_card_id" form:"order_card_id binding:"-"`
	OrderUsername 		string 	`gorm:"column:order_username;type:varchar(30);" json:"order_username" form:"order_username" binding:"required"`
	OrderPhoneNumber 	string	`gorm:"column:order_phone_number;type:varchar(15);" json:"order_phone_number" form:"order_phone_number" binding:"required"`
	OrderAddress		string	`gorm:"column:order_address;type:varchar(30);not null;" json:"order_address" form:"order_address" binding:"-"`
	OrderPrice			int 	`gorm:"column:order_price;type:int;"json:"order_price" form:"order_price" binding:"-"`
	OrderStatus			string  `gorm:"column:order_status;varchar(30);not null;default:start" json:"order_status" form:"order_status" binding:"-"`
}

func (Order) TableName() string {
	return sys.TableName("order")
}
// 添加前
func (m *Order) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	m.ID= m.GetOrderID()
	return nil
}

func (Order)GetOrderID()uint64{
	tm1:=time.Now().Format("20060102030405")
	res,_:=strconv.ParseInt(tm1,10,64)
	return uint64(res)
}

//删除卡
func DeleteOneOrderById(ids []uint64) error {
	var orderjson = Order{}
	for _, v := range ids {
		orderjson.ID = v
		if err := db.DB.Delete(orderjson).Error; err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

//添加订单信息
func AddNewOrder(order Order)error{
	order.UpdatedAt=time.Now()
	err :=db.DB.Create(&order).Error    //must give pointer
	if err!=nil{
		fmt.Println("插入失败", err)
	}
	return err
}

//更新订单信息
func UpdateOrder(order Order)error {
	db.DB.Model(&order).Updates(order)
	fmt.Println(order.ID)
	return nil
}

//查询订单信息
func QueryOrder(order Order)([]Order,int) {
	var value []Order
	res_count:=1
	res:=db.DB.Where(&order).Find(&value).Count(&res_count)
	if(res.Error==nil) {//has find value
		fmt.Println("get the value")
	}else {
		res_count=-1
		return value,-1
	}

	return value ,res_count
}

//删除订单（未知订单数量）
func DeleteOrders(order Order)error{
	var value,res_count=QueryOrder(order)
	var err error
	if (res_count>0) {
		err =DeleteMultiOrders(value)
	}
	return err
}
//删除一张订单
func DeleteOneOrder(order Order) error{
	err:=db.DB.Delete(order).Error
	if err!=nil{
		fmt.Println("删除失败", err)
	}else{
		fmt.Println("删除成功",err)
	}
	return err
}
//删除多个订单
func DeleteMultiOrders(order []Order) error{
	var err error
	for i:=0;i<len(order);i++{
		err=DeleteOneOrder(order[i])
	}
	return err
}

func IsPhoneNum(mobile string)error{
	isorno, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if isorno {
		if(len(mobile)==11) {
			return nil
		}else {
			return errors.New("length not match")
		}
	} else {
		return errors.New("is not a correct phone number")
	}
}

func IsnotNullOrder(orderjson Order)error{

	if (orderjson.OrderAddress!=""&&orderjson.OrderStatus!=""&&orderjson.OrderPrice!=0) {
		return nil
	}else{
		return errors.New("have not null value")
	}

}

func IsOrderStatus(status string)error{

	return nil
}
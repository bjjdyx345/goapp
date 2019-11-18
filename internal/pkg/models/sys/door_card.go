package sys

import (
	"fmt"
	"github.com/it234/goapp/internal/pkg/models/basemodel"
	"github.com/it234/goapp/internal/pkg/models/db"
	"github.com/jinzhu/gorm"
	"time"
)

type Door_card struct {
	basemodel.Model
	cardType string `gorm:"column:cardType;size:30" json:"card_type" form:"card_type"`
	cardContentFile string `gorm:"column:cardContentFile;size:30" json:"cardContentFile" form:"cardContentFile"`
	belongToApartmentId string `gorm:"column:belongToApartmentId;size:30" json:"belongToApartmentId" form:"belongToApartmentId"`
}


// 添加前
func (m *Door_card) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

//添加新卡，传入门禁卡信息
func (*Door_card) addNewCard(dc Door_card) error{
	return(db.DB.Create(dc).Error)
}

//删除卡
func (*Door_card) deleteCard(dc Door_card) error{
	return(db.DB.Delete(dc).Error)
}

//更新卡信息
func (*Door_card) updateCard(dc Door_card)error {
	return(db.DB.Update(dc).Error)
}

//查询卡信息，提供卡id
func (m *Door_card) queryCardId(cardInfo Door_card)(Door_card) {
	res = db.DB.Where(cardInfo).First(&cardInfo)
	fmt.Sprint(res)
	return cardInfo
}


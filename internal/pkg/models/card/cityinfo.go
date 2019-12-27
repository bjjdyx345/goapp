package card

import (
	"fmt"
	"github.com/it234/goapp/internal/pkg/models/db"
	"github.com/it234/goapp/internal/pkg/models/sys"
)

type CityInfo struct {
	ID        			uint64  `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	Regionid			string	`gorm:"column:regionid;type:varchar(10);" json:"region_id" form:"region_id"`
	Name 				string 	`gorm:"column:name;type:varchar(10);" json:"name" form:"name"`
	ParentId 			string	`gorm:"column:parent_id;type:varchar(10);" json:"parent_id" form:"parent_id"`
	ShortName			string	`gorm:"column:short_name;type:varchar(10);" json:"short_name" form:"short_name"`
	LevelType			int 	`gorm:"column:level_type;type:varchar(10);" json:"level_type" form:"level_type"`
	CityCode			string	`gorm:"column:city_code;type:varchar(10);" json:"city_code" form:"city_code"`
	ZipCode				string	`gorm:"column:zip_code;type:varchar(10);" json:"zip_code" form:"zip_code"`
	MergerName			string 	`gorm:"column:merger_name;type:varchar(50);" json:"merger_name" form:"merger_name"`
	Longitude			string	`gorm:"column:longitude;type:varchar(30);" json:"longitude" form:"longitude"`
	Latitude			string	`gorm:"column:latitude;type:varchar(30);" json:"latitude" form:"latitude"`
	Pinyin				string	`gorm:"column:pinyin;type:varchar(50);" json:"pinyin" form:"pinyin"`
}

func (CityInfo) TableName() string {
	return sys.TableName("cityinfo")
}

//添加订单信息
func IsExistCity(cityname string) int{
	var cinfo CityInfo
	var value []CityInfo
	var res_count=1
	cinfo.Name=cityname
	db.DB.Where(&cinfo).Find(&value).Count(&res_count)
	fmt.Println(cityname)
	fmt.Println(value)
	fmt.Println(res_count)
	return res_count
}

func GetCitys(cityname string,citylevel int) []string{
	var cinfo CityInfo
	var value []CityInfo
	var res_count=1
	nextlevelname:=[]string{}
	cinfo.LevelType = citylevel
	fmt.Println("level type is",citylevel)
	if(citylevel==1){
		cinfo.Name=cityname
		db.DB.Where(&cinfo).Find(&value).Count(&res_count)
	} else {
		cinfo.Name = fmt.Sprintf("%%%s%%",cityname)
		fmt.Println(cinfo)
		db.DB.Where("level_type = ? AND merger_name LIKE ?", cinfo.LevelType, cinfo.Name).Find(&value).Count(&res_count)
	}

	for i:=0;i<res_count;i++ {
		nextlevelname = append(nextlevelname, value[i].Name)
		//must use *firstlever[i]取出里边的内容
	}
	return nextlevelname
}

//about all kind of level city to return the complete rows
func GetCityInfo(cityname string) []CityInfo{
	var cinfo CityInfo
	var value []CityInfo
	var res_count=1
	cinfo.Name=cityname
	db.DB.Where(&cinfo).Find(&value).Count(&res_count)
	return value
}
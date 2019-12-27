package card

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	"github.com/it234/goapp/internal/pkg/models/card"
)

type CityInfoCon struct {}

func (CityInfoCon)GetFirstCityAll(c *gin.Context)  {
	citytree:= card.GetCitys("",1)
	fmt.Println(citytree)
	common.ResSuccess(c,citytree)
	//common.ResSuccessMsg(c)
}

func (CityInfoCon)GetSecondCity(c *gin.Context){

	city_name:=c.Query("cityname")
	fmt.Println("has get second city")
	citytree:= card.GetCitys(city_name,2)
	common.ResSuccess(c,citytree)

}
func (CityInfoCon)GetThirdCity(c *gin.Context){
	fmt.Println("get third city")
	city_name:=c.Query("cityname")
	citytree:= card.GetCitys(city_name,3)
	common.ResSuccess(c,citytree)

}


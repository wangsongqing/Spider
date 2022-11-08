package controllers

import (
	"Spider/app/models/house"
	"Spider/helpers"
	publicHelpers "Spider/pkg/helpers"
	"github.com/gocolly/colly"
	"github.com/gookit/color"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"sync"
	"time"
)

type House struct {
}

var w sync.WaitGroup

func (h *House) Start(city string) {
	w.Add(1)

	go h.initSpider(city)

	w.Wait()

}

func (h *House) initSpider(cityName string) {
	var i int64

	if cityName == "all" {
		for key, _ := range helpers.CityName {
			area, err := helpers.CityArea[key]
			if err == false {
				continue
			}

			for _, areaName := range area {
				_, pageNum := h.totalAll(key, areaName)
				for i = 1; i <= pageNum; i++ {
					h.getHouseInfo(key, i, areaName)
					sleepTime := publicHelpers.MakeRandInt()
					time.Sleep(time.Second * time.Duration(sleepTime)) // 随机sleep，防止请求频繁
				}
			}
		}

	} else {
		area, err := helpers.CityArea[cityName]
		if err == false {
			color.Redln("城市的区域错误")
			return
		}
		for _, areaName := range area {
			_, pageNum := h.totalAll(cityName, areaName)
			for i = 1; i <= pageNum; i++ {
				h.getHouseInfo(cityName, i, areaName)
				sleepTime := publicHelpers.MakeRandInt()
				time.Sleep(time.Second * time.Duration(sleepTime)) // 随机sleep，防止请求频繁
			}
		}
	}

	w.Done()
}

// TotalAll PageAll 获取房源总数
func (h *House) totalAll(city string, area string) (int64, int64) {
	var TotalNum int64
	var TotalPage int64
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// 获取房源总数
	c.OnHTML("div.content > .leftContent", func(e *colly.HTMLElement) {
		T := e.ChildText("div.resultDes > h2 > span")
		TotalNum, _ = strconv.ParseInt(T, 10, 64)
	})

	// 获取页数
	c.OnHTML(".contentBottom .house-lst-page-box", func(e *colly.HTMLElement) {
		TotalPage = gjson.Get(e.Attr("page-data"), "totalPage").Int()
	})
	c.Visit("https://" + city + ".lianjia.com/ershoufang/" + area)
	c.Wait()
	return TotalNum, TotalPage
}

func (h *House) getHouseInfo(city string, page int64, area string) {

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// 获取房源总数
	c.OnHTML(".sellListContent > li", func(e *colly.HTMLElement) {
		listInfoName := e.ChildText("div.info > div.flood > div.positionInfo > a")
		info := strings.Split(listInfoName, " ")
		if len(info) != 2 {
			return
		}

		url := e.ChildAttr("div.info > div.title > a", "href")
		//fmt.Println(url)
		houseIcon := e.ChildText("div.info > div.address > div.houseInfo") // 详情

		priceTotal := e.ChildText("div.info > div.priceInfo > div.totalPrice span") // 总价
		price, _ := strconv.ParseFloat(priceTotal, 64)

		priceEverything := e.ChildText("div.info > div.priceInfo > div.unitPrice span") // 每平米价格
		priceEverythingRepOne := strings.Replace(priceEverything, "元/平", "", 1)
		priceEverythingRep := strings.Replace(priceEverythingRepOne, ",", "", 1)

		houseModel := house.House{
			City:       helpers.CityName[city],
			Name:       info[0],
			Address:    info[1],
			TotalPrice: price,
			Info:       houseIcon,
			Url:        url,
			DanPrice:   priceEverythingRep,
			Area:       area,
		}
		houseModel.Create()
		if houseModel.ID < 0 {
			panic("写入异常")
		}
	})
	c.Visit("https://" + city + ".lianjia.com/ershoufang/pg" + strconv.FormatInt(page, 10) + "/" + area)
	c.Wait()
}

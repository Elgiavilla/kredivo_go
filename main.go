package main

import (
	"net/url"
	"time"

	"github.com/elgiavilla/kredivo/middleware"
	"github.com/elgiavilla/kredivo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	_contactHttp "github.com/elgiavilla/kredivo/contact/http"
	_contactRepo "github.com/elgiavilla/kredivo/contact/repository"
	_contactSvc "github.com/elgiavilla/kredivo/contact/service"
)

func main() {
	db, err := gorm.Open("mysql", "root:pass@/kredivo_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Contact{})
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	e := echo.New()
	middl := middleware.InitMiddleware()
	timeContext := time.Duration(2) * time.Second

	contactRepo := _contactRepo.NewContactRepo(db)
	contactSvc := _contactSvc.NewContactSvc(contactRepo, timeContext)

	_contactHttp.NewContactHandler(e, contactSvc)
	e.Use(middl.CORS)
	e.Start(":8090")
}

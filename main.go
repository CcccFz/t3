package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"main/myTest/common/jwt"
	"main/myTest/common/store"
	"main/myTest/interface/api"
)

func init() {
	println("main")
}

func main() {
	fmt.Println("码运昌隆！")
	store.MyAuto()
	myHttp := gin.Default()

	myHttp.POST("/api/t3/user/user/register", api.UserRegister)     // 乘客/注册
	myHttp.POST("/api/t3/user/driver/register", api.DriverRegister) // 司机/注册
	myHttp.POST("/api/t3/user/user/login", api.UserLogin)           // 乘客/登录
	myHttp.POST("/api/t3/user/driver/login", api.DriverLogin)       // 司机/登录

	// t3go URL
	// 乘客，司机，车，行程，订单，支付
	myHttp.Use(jwt.JWT(store.CACHE))
	myHttp.POST("/api/t3/user/user/logout", api.Userlogout) // 乘客/退出
	myHttp.POST("/api/t3/user/user/update", api.UserUpdate) // 乘客/修改
	myHttp.POST("/api/t3/user/user/detail", api.UserDetail) // 乘客/详情

	myHttp.POST("/api/t3/user/driver/logout", api.Driverlogout) // 司机/退出
	myHttp.POST("/api/t3/user/driver/update", api.DriverUpdate) // 司机/修改
	myHttp.POST("/api/t3/user/driver/detail", api.DriverDetail) // 司机/详情

	myHttp.POST("/api/t3/car/detail", api.CarDetail) // 车/详情

	myHttp.POST("/api/t3/track/create", api.TrackCreate)        // 行程/创建
	myHttp.POST("/api/t3/track/cancel", api.TrackCancel)        // 行程/取消
	myHttp.POST("/api/t3/track/status/set", api.TrackStatusSet) // 行程/状态/修改
	myHttp.POST("/api/t3/track/update", api.TrackUpdate)        // 行程/修改
	myHttp.POST("/api/t3/track/current", api.TrackCurrent)      // 行程/当前
	myHttp.POST("/api/t3/track/list", api.TrackList)            // 行程/列表
	myHttp.POST("/api/t3/track/detail", api.TrackDetail)        // 行程/详情

	myHttp.POST("/api/t3/order/detail", api.OrderDetail) // 订单/详情

	myHttp.POST("/api/t3/payment/notify/pay", api.PaymentNotifyPay) // 支付/创建
	myHttp.POST("/api/t3/payment/detail", api.PaymentDetail)        // 支付/详情

	err := myHttp.Run(":8010")
	if err != nil {
		log.Info().Err(err)
		return
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	myHttp.Use(jwt.JWT(store.CACHE))

	// t3go URL
	// 乘客，司机，车，行程，订单，支付
	myHttp.POST("/t3go/api/user/passenger/login", api.UserLogin)   //乘客注册登录
	myHttp.POST("/t3go/api/user/passenger/delete", api.UserDelete) //乘客注销账号
	myHttp.POST("/t3go/api/user/passenger/update", api.UserUpdate) //更新乘客基本信息
	myHttp.POST("/t3go/api/user/passenger/detail", api.UserDetail) //查看乘客详细信息

	myHttp.POST("/t3go/api/user/driver/login", api.DriverLogin)   //司机注册登录
	myHttp.POST("/t3go/api/user/driver/delete", api.DriverDelete) //司机注销账号
	myHttp.POST("/t3go/api/user/driver/update", api.DriverUpdate) //更新司机基本信息
	myHttp.POST("/t3go/api/user/driver/detail", api.DriverDetail) //查看司机详细信息

	myHttp.POST("/t3go/api/car/create", api.CarCreate) //新增网约车
	myHttp.POST("/t3go/api/car/update", api.CarUpdate) //更新网约车基本信息
	myHttp.POST("/t3go/api/car/delete", api.CarDelete) //删除网约车
	myHttp.POST("/t3go/api/car/detail", api.CarDetail) //查看网约车详细信息

	myHttp.POST("/t3go/api/route/create", api.RouteCreate) //新增行程
	myHttp.POST("/t3go/api/route/set", api.RouteSet)       //更新行程状态（司机已到达目的地，行程开始，取消行程）
	myHttp.POST("/t3go/api/route/update", api.RouteUpdate) //修改行程(例如修改目的地)
	myHttp.POST("/t3go/api/route/delete", api.RouteDelete) //删除行程
	myHttp.POST("/t3go/api/order/list", api.RouteList)     //查询行程列表
	myHttp.POST("/t3go/api/order/detail", api.RouteDetail) //查看行程详细信息

	myHttp.POST("/t3go/api/order/create", api.OrderCreate) //新增订单(行程状态为司机正在赶来)
	myHttp.POST("/t3go/api/order/set", api.OrderSet)       //完成订单
	myHttp.POST("/t3go/api/order/delete", api.OrderDelete) //删除订单

	myHttp.POST("/t3go/api/payment/create", api.PaymentCreate) //新增支付(行程结束)
	myHttp.POST("/t3go/api/payment/set", api.PaymentSet)       //支付成功
	myHttp.POST("/t3go/api/payment/detail", api.PaymentDetail) //查看支付详细信息

	err := myHttp.Run(":8010")
	if err != nil {
		fmt.Println(err)
		return
	}
}

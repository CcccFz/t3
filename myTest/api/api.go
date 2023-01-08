package api

import (
	"github.com/gin-gonic/gin"
	"main/myTest/dto"
)

func UserLogin(c *gin.Context) {
	req := new(dto.UserLoginReq)
}

func UserDelete(c *gin.Context) {
	req := new(dto.UserReq)
}

func UserUpdate(c *gin.Context) {
	req := new(dto.UserUpdateReq)
}

func UserDetail(c *gin.Context) {
	req := new(dto.UserReq)
	rsp := new(dto.UserDetailRsp)
}

func DriverLogin(c *gin.Context) {
	req := new(dto.DriverLoginReq)
}

func DriverDelete(c *gin.Context) {
	req := new(dto.DriverReq)
}

func DriverUpdate(c *gin.Context) {
	req := new(dto.DriverUpdateReq)
}

func DriverDetail(c *gin.Context) {
	req := new(dto.DriverReq)
	rsp := new(dto.DriverDetailRsp)
}

func CarCreate(c *gin.Context) {
	req := new(dto.CarCreateReq)
}

func CarUpdate(c *gin.Context) {
	req := new(dto.CarUpdateReq)
}

func CarDelete(c *gin.Context) {
	req := new(dto.CarReq)
}

func CarDetail(c *gin.Context) {
	req := new(dto.CarReq)
	rsp := new(dto.CarDetailRsp)
}

func RouteCreate(c *gin.Context) {
	req := new(dto.RouteCreateReq)
	rsp := new(dto.RouteCreateRsp)
}

func RouteMatch(c *gin.Context) {
	req := new(dto.RouteMatchReq)
}

func RouteSet(c *gin.Context) {
	req := new(dto.RouteSetReq)
}

func RouteUpdate(c *gin.Context) {
	req := new(dto.RouteUpdateReq)
}

func RouteDelete(c *gin.Context) {
	req := new(dto.RouteReq)
}

func RouteList(c *gin.Context) {
	req := new(dto.UserReq)
	rsp := new(dto.RouteListRsp)
}

func RouteDetail(c *gin.Context) {
	req := new(dto.RouteReq)
	rsp := new(dto.RouteDetailRsp)
}

func OrderCreate(c *gin.Context) {
	req := new(dto.OrderCreateReq)
	rsp := new(dto.OrderCreateRsp)
}

func OrderSet(c *gin.Context) {
	req := new(dto.OrderSetReq)
}

func OrderDelete(c *gin.Context) {
	req := new(dto.OrderReq)
}

func PaymentCreate(c *gin.Context) {
	req := new(dto.PaymentCreateReq)
	rsq := new(dto.PaymentCreateRsp)
}

func PaymentSet(c *gin.Context) {
	req := new(dto.PaymentSetReq)
}

func PaymentDetail(c *gin.Context) {
	req := new(dto.PaymentReq)
	rsp := new(dto.RouteDetailRsp)
}

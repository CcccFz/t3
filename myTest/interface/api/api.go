package api

import (
	"github.com/gin-gonic/gin"
	"main/myTest/interface/dto"
)

func UserRegister(c *gin.Context) {
	req := new(dto.UserRegisterReq)
}

func DriverRegister(c *gin.Context) {
	req := new(dto.DriverRegisterReq)
}

func UserLogin(c *gin.Context) {
	req := new(dto.UserLoginReq)
}

func DriverLogin(c *gin.Context) {
	req := new(dto.DriverLoginReq)
}

func Userlogout(c *gin.Context) {
}

func UserUpdate(c *gin.Context) {
	req := new(dto.UserUpdateReq)
}

func UserDetail(c *gin.Context) {
	rsp := new(dto.UserDetailRsp)
}

func Driverlogout(c *gin.Context) {}

func DriverUpdate(c *gin.Context) {
	req := new(dto.DriverUpdateReq)
}

func DriverDetail(c *gin.Context) {
	rsp := new(dto.DriverDetailRsp)
}

func CarDetail(c *gin.Context) {
	req := new(dto.CarReq)
	rsp := new(dto.CarDetailRsp)
}

func TrackCreate(c *gin.Context) {
	req := new(dto.TrackCreateReq)
	rsp := new(dto.TrackCreateRsp)
}

func TrackCancel(c *gin.Context) {
}

func TrackStatusSet(c *gin.Context) {
	req := new(dto.TrackSetReq)
	rsp := new(dto.TrackSetRsp)
}

func TrackUpdate(c *gin.Context) {
	req := new(dto.TrackUpdateReq)
}

func TrackCurrent(c *gin.Context) {
	rsp := new(dto.TrackDetailRsp)
}

func TrackList(c *gin.Context) {
	rsp := new(dto.TrackListRsp)
}

func TrackDetail(c *gin.Context) {
	req := new(dto.TrackReq)
	rsp := new(dto.TrackDetailRsp)
}

func OrderDetail(c *gin.Context) {
	rsp := new(dto.OrderDetailRsp)
}

func PaymentNotifyPay(c *gin.Context) {
	req := new(dto.PaymentNotifyPayReq)
}

func PaymentDetail(c *gin.Context) {
	rsp := new(dto.PaymentDetailRsp)
}

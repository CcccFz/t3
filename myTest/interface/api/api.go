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
	req := new(dto.UserReq)
}

func UserUpdate(c *gin.Context) {
	req := new(dto.UserUpdateReq)
}

func UserDetail(c *gin.Context) {
	req := new(dto.UserReq)
	rsp := new(dto.UserDetailRsp)
}

func Driverlogout(c *gin.Context) {}

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

func TrackCreate(c *gin.Context) {
	req := new(dto.TrackCreateReq)
	rsp := new(dto.TrackCreateRsp)
}

func TrackCancel(c *gin.Context) {
	req := new(dto.TrackCancelReq)
}

func TrackStatusSet(c *gin.Context) {
	req := new(dto.TrackSetReq)
}

func TrackUpdate(c *gin.Context) {
	req := new(dto.TrackUpdateReq)
}

func TrackDelete(c *gin.Context) {
	req := new(dto.TrackReq)
}

func TrackList(c *gin.Context) {
	req := new(dto.UserReq)
	rsp := new(dto.TrackListRsp)
}

func TrackDetail(c *gin.Context) {
	req := new(dto.TrackReq)
	rsp := new(dto.TrackDetailRsp)
}

func OrderDetail(c *gin.Context) {
	rsq := new(dto.OrderDetailReq)
}

func PaymentNotifyPay(c *gin.Context) {
	req := new(dto.PaymentNotifyPayReq)
}
}

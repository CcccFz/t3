package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"main/myTest/common/http"
	"main/myTest/common/store"
	"main/myTest/domain/entity"
	"main/myTest/interface/dto"
)

func UserRegister(c *gin.Context) {
	req := new(dto.UserRegisterReq)
	err := c.Bind(req)
	if err != nil {
		http.ErrJson(c, err)
	}

	checkUser := &entity.TUser{Phone: req.Phone}
	err = store.DB.First(checkUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/注册")
		http.ErrJson(c, err)
		return
	} else if err == nil {
		err = errors.New("用户已注册，请直接登录")
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/注册")
		http.ErrJson(c, err)
		return
	}

	user := new(entity.TUser)
	if err = copier.Copy(user, req); err != nil {
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/注册")
		http.ErrJson(c, err)
		return
	}
	if err = store.DB.Create(&user).Error; err != nil {
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/注册")
		http.ErrJson(c, err)
		return
	}

	log.Info().Str("phone", user.Phone).Uint("user_id", user.UserId).Msg("乘客/注册")
	http.OkJson(c, nil)
	return
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
	req := new(dto.DriverReq)
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
	req := new(dto.TrackListReq)
	rsp := new(dto.TrackListRsp)
}

func TrackDetail(c *gin.Context) {
	req := new(dto.TrackReq)
	rsp := new(dto.TrackDetailRsp)
}

func OrderDetail(c *gin.Context) {
	req := new(dto.OrderReq)
	rsp := new(dto.OrderDetailRsp)
}

func PaymentNotifyPay(c *gin.Context) {
	req := new(dto.PaymentNotifyPayReq)
}

func PaymentDetail(c *gin.Context) {
	req := new(dto.PaymentReq)
	rsp := new(dto.PaymentDetailRsp)
}

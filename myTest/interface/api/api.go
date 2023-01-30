package api

import (
	"errors"
	"main/myTest/common/http"
	"main/myTest/common/jwt"
	"main/myTest/common/store"
	"main/myTest/domain/entity"
	"main/myTest/interface/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	req := new(dto.UserRegisterReq)
	err := c.Bind(req)
	if err != nil {
		http.ErrJson(c, err)
		return
	}

	checkUser := new(entity.TUser)
	err = store.DB.Where("phone = ?", req.Phone).First(checkUser).Error
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
	if err = store.DB.Create(user).Error; err != nil {
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/注册")
		http.ErrJson(c, err)
		return
	}

	log.Info().Str("phone", user.Phone).Uint("user_id", user.UserId).Msg("乘客/注册")
	http.OkJson(c, nil)
	return
}

func UserLogin(c *gin.Context) {
	req := new(dto.UserLoginReq)
	err := c.Bind(req)
	if err != nil {
		http.ErrJson(c, err)
		return
	}

	checkUser := new(entity.TUser)
	err = store.DB.Where("phone = ?", req.Phone).First(checkUser).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户未注册，请注册后登录")
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/登录")
		http.ErrJson(c, err)
		return
	} else if err != nil {
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/登录")
		http.ErrJson(c, err)
		return
	} else if checkUser.Password != req.Password {
		err = errors.New("密码错误，请重新登录")
		log.Error().Err(err).Str("phone", req.Phone).Msg("乘客/登录")
		http.ErrJson(c, err)
		return
	}
	var token string
	token, err = jwt.GenToken(checkUser.UserId, entity.UserTypeUser, store.CACHE)
	c.Header("Authorization", token)

	log.Info().Str("phone", checkUser.Phone).Uint("user_id", checkUser.UserId).Msg("乘客/登录")
	http.OkJson(c, nil)
	return
}

func DriverRegister(c *gin.Context) {
	req := new(dto.DriverRegisterReq)
	_ = req

}

func DriverLogin(c *gin.Context) {
	req := new(dto.DriverLoginReq)
	_ = req
}

func Userlogout(c *gin.Context) {
}

func UserUpdate(c *gin.Context) {
	req := new(dto.UserUpdateReq)
	_ = req
}

func UserDetail(c *gin.Context) {
	rsp := new(dto.UserDetailRsp)
	_ = rsp
}

func Driverlogout(c *gin.Context) {}

func DriverUpdate(c *gin.Context) {
	req := new(dto.DriverUpdateReq)
	_ = req
}

func DriverDetail(c *gin.Context) {
	req := new(dto.DriverReq)
	rsp := new(dto.DriverDetailRsp)
	_ = rsp
	_ = req
}

func CarDetail(c *gin.Context) {
	req := new(dto.CarReq)
	rsp := new(dto.CarDetailRsp)
	_ = req
	_ = rsp
}

func TrackCreate(c *gin.Context) {
	req := new(dto.TrackCreateReq)
	rsp := new(dto.TrackCreateRsp)
	_ = req
	_ = rsp
}

func TrackCancel(c *gin.Context) {
}

func TrackStatusSet(c *gin.Context) {
	req := new(dto.TrackSetReq)
	// rsp := new(dto.TrackSetRsp)
	_ = req
	// _ = rsp
}

func TrackUpdate(c *gin.Context) {
	req := new(dto.TrackUpdateReq)
	_ = req
}

func TrackCurrent(c *gin.Context) {
	rsp := new(dto.TrackDetailRsp)
	_ = rsp
}

func TrackList(c *gin.Context) {
	req := new(dto.TrackListReq)
	rsp := new(dto.TrackListRsp)
	_ = req
	_ = rsp
}

func TrackDetail(c *gin.Context) {
	req := new(dto.TrackReq)
	rsp := new(dto.TrackDetailRsp)
	_ = req
	_ = rsp
}

func OrderDetail(c *gin.Context) {
	req := new(dto.OrderReq)
	rsp := new(dto.OrderDetailRsp)
	_ = req
	_ = rsp
}

func PaymentNotifyPay(c *gin.Context) {
	req := new(dto.PaymentNotifyPayReq)
	_ = req
}

func PaymentDetail(c *gin.Context) {
	req := new(dto.PaymentReq)
	rsp := new(dto.PaymentDetailRsp)
	_ = req
	_ = rsp
}

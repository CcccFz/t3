package dto

import (
	"main/myTest/common/http"
	"main/myTest/domain/entity"
)

type UserReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"10"`
}

type UserRegisterReq struct {
	// 用户名
	UserName string `json:"user_name" form:"user_name" binding:"required" example:"duome"`
	// 密码
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"required,len=11"`
}

type UserLoginReq struct {
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"required,len=11" example:"13981536645"`
	// 密码
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
}

type UserUpdateReq struct {
	// 用户名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 密码
	Password string `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"len=11"`
	// 旧密码
	Old_Password string `json:"old_Password" form:"old_password" example:"123456"`
}

type UserDetailRsp struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"10"`
	// 用户名
	UserName string `json:"user_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone"`
}

type DriverReq struct {
	// 司机ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"4"`
}

type DriverRegisterReq struct {
	// 司机名
	UserName string `json:"user_name" form:"user_name" binding:"required" example:"ccccfz"`
	// 密码
	Password uint `json:"password" form:"password" binding:"required" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" binding:"required" example:"13981556666"`
	// 车牌号/
	CarNo string `json:"car_no" form:"car_no" binding:"required" example:"川AA85409"`
	// 车类型（1-特享，2-惠享）
	CarType entity.CarType `json:"car_type" form:"car_type" binding:"required" enums:"1,2" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" binding:"required" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" binding:"required" example:"10"`
}

type DriverLoginReq struct {
	// 手机号
	Phone uint `json:"phone" form:"phone" binding:"required" example:"13981556666"`
	// 密码
	Password uint `json:"password" form:"password" binding:"required" example:"123456"`
}

type DriverUpdateReq struct {
	// 司机名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 密码
	Password uint `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" binding:"len=11" example:"13981556666"`
	// 老密码
	OldPassword string `json:"old_password" form:"old_password" example:"123456"`
}

type DriverDetailRsp struct {
	// 司机ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"4"`
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 司机名
	UserName string `json:"user_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone" example:"13981556666"`
	// 总里程数
	Courses uint `json:"courses" example:"100"`
}

type CarReq struct {
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
}

type CarDetailRsp struct {
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 车牌号
	CarNo string `json:"car_no" example:"川AA85409"`
	// 车类型（1-特享，2-惠享）
	CarType entity.CarType `json:"car_type" enums:"1,2" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" example:"10"`
}

type TrackReq struct {
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
}

type TrackCreateReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 出发地点
	SrcPoint string `json:"src_point" form:"src_point" binding:"required" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" form:"dest_point" binding:"required" example:"华为技术有限公司三号门"`
	// 车类型（1-特享，2-惠享）
	CarType []entity.CarType `json:"car_type" form:"car_type" binding:"required" enums:"[1,2]" example:"2"`
	// 里程数（公里）
	Distance uint `json:"distance" example:"100"`
	// 时长(分钟)
	Duration uint `json:"duration" example:"9"`
}

type TrackCreateRsp struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 出发地点
	SrcPoint string `json:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" example:"华为技术有限公司三号门"`
	// 车类型（1-特享，2-惠享）
	CarType []entity.CarType `json:"car_type" enums:"[1,2]" example:"2"`
	// 里程数（公里）
	Distance uint `json:"distance" example:"100"`
	// 时长(分钟)
	Duration uint `json:"duration" example:"9"`
}

type TrackSetReq struct {
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 行程进度(2-待赶来,3-待上车,4-进行中,5-待支付)
	TrackStatus entity.TrackStatus `json:"track_status" form:"track_status" binding:"required" enums:"2,3,4,5" example:"3"`
}

type TrackUpdateReq struct {
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 目的地
	DestPoint string `json:"dest_point" form:"dest_point" example:"华为技术有限公司三号门"`
}

type TrackListReq struct {
	*http.Page
	// 行程状态(1-待接单,2-待赶来,3-待上车,4-进行中,5-待支付,6-已完成,10-已取消)
	TrackStatus entity.TrackStatus `json:"track_status" form:"track_status" enums:"1,2,3,4,5,6,10" example:"2"`
	// 创建时间（时间戳）
	CreatedAt string `json:"created_at" form:"created_at" example:"2022-04-05 15:30:12"`
}

type TrackListRsp struct {
	Tracks []*TrackItem `json:"tracks"`
}

type TrackItem struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 车类型（1-特享，2-惠享）
	CarType entity.CarType `json:"car_type" enums:"1,2" example:"2"`
	// 出发地点
	SrcPoint string `json:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" example:"华为技术有限公司三号门"`
	// 行程状态(1-待接单,2-待赶来,3-待上车,4-进行中,5-待支付,6-已完成,10-已取消)
	TrackStatus entity.TrackStatus `json:"track_status" enums:"1,2,3,4,5,6,10" example:"2"`
	// 创建时间（时间戳）
	CreatedAt string `json:"created_at" example:"2022-04-05 15:30:12"`
}

type TrackDetailRsp struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 司机ID
	ServicerId uint `json:"servicer_id" example:"4"`
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 出发地点
	SrcPoint string `json:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" example:"华为技术有限公司三号门"`
	// 司机名
	UserName string `json:"user_name" example:"duome"`
	// 车牌号
	CarNo string `json:"car_no" example:"川AA85409"`
	// 车类型（1-特享，2-惠享）
	CarType entity.CarType `json:"car_type" enums:"1,2" example:"2"`
	// 行程状态(1-待接单,2-待赶来,3-待上车,4-进行中,5-待支付,6-已完成,10-已取消)
	TrackStatus entity.TrackStatus `json:"status" enums:"1,2,3,4,5,6,10" example:"2"`
	// 金额（元）
	Cost uint `json:"cost" example:"9.60"`
	// 创建时间（时间戳）
	CreatedAt string `json:"created_at" example:"2022-04-05 15:30:12"`
}

type OrderReq struct {
	// 订单ID
	OrderId uint `json:"order_id" example:"8"`
}

type OrderDetailRsp struct {
	// 订单ID
	OrderId uint `json:"order_id" example:"8"`
	// 用户ID
	UserId uint `json:"user_id" example:"10"`
	// 司机ID
	ServicerId uint `json:"servicer_id" example:"4"`
	// 支付ID
	PaymentId uint `json:"payment_id" example:"9"`
	// 起步价（元）
	StartCost uint `json:"start_cost" example:"9"`
	// 单价(每分钟)
	UnitCost uint `json:"unit_cost" example:"0.2"`
	// 创建时间
	CreatedAt string `json:"created_at" example:"2022-04-05 15:30:12"`
}

type PaymentNotifyPayReq struct {
	// 支付ID
	PaymentId uint `json:"payment_id" example:"9"`
	// 支付方式(1-支付宝支付，2-微信支付，3-银行卡支付，4-云闪付支付)
	PaymentMethod entity.PaymentMethod `json:"payment_method" form:"payment_method" binding:"required" enums:"1,2,3,4" example:"1"`
}

type PaymentReq struct {
	// 支付ID
	PaymentId uint `json:"payment_id" example:"9"`
}

type PaymentDetailRsp struct {
	// 金额（元）
	Cost uint `json:"cost" example:"9.60"`
	// 支付状态(1-待支付，2-已支付)
	PaymentStatus entity.PaymentStatus `json:"payment_status" enums:"1,2" example:"1"`
	// 支付方式(1-支付宝支付，2-微信支付，3-银行卡支付，4-云闪付支付)
	PaymentMethod entity.PaymentMethod `json:"payment_method" form:"payment_method" enums:"1,2,3,4" example:"1"`
}

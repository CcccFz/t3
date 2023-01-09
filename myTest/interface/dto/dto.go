package dto

import (
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
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" binding:"required" enums:"1,2" example:"1"`
}

type UserLoginReq struct {
	// 密码
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"required,len=11" example:"13981536645"`
}

type UserUpdateReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"10"`
	// 用户名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 密码
	Password string `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone string `json:"phone" form:"phone" binding:"len=11"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 旧密码
	Old_Password string `json:"old_Password" form:"old_password" example:"123456"`
}

type UserDetailRsp struct {
	// 用户ID
	UserId uint `json:"user_id" example:"10"`
	// 用户名
	UserName string `json:"user_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" enums:"1,2" example:"1"`
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
	// 性别（1-男， 2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"2"`
	// 车牌号// 这里len是多少
	CarNo string `json:"car_no" form:"car_no" binding:"required" example:"川AA85409"`
	// 车类型（1-网约车，2-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" binding:"required" enums:"1,2" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" binding:"required" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" binding:"required" example:"10"`
}

type DriverLoginReq struct {
	// 密码
	Password uint `json:"password" form:"password" binding:"required" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" binding:"required" example:"13981556666"`
}

type DriverUpdateReq struct {
	// 司机ID
	UserId uint `json:"user_id" form:"user_id" binding:"required" example:"4"`
	// 司机名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 密码
	Password uint `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" binding:"len=11" example:"13981556666"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 老密码
	OldPassword string `json:"old_password" form:"old_password" example:"123456"`
}

type DriverDetailRsp struct {
	// 司机ID
	UserId uint `json:"user_id" example:"10"`
	// 司机名
	UserName string `json:"user_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone" example:"13981556666"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" enums:"1,2" example:"1"`
	// 等级（（1-大众，2-白银，3-黄金，4-铂金，5-黑金））
	UserLevel entity.UserLevel `json:"driver_level" enums:"1,2,3,4,5" example:"3"`
	// 成长值
	GrowthNumber uint `json:"growth_number" example:"10"`
	// 总里程数
	Courses uint `json:"courses" example:"100"`
}

type CarReq struct {
	// 车Id
	CarId uint `json:"car_id" form:"car_id" binding:"required" example:"1"`
	// 司机ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
}

type CarCreateReq struct {
	// 司机ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
}

type CarUpdateReq struct {
	// 车牌号// 这里len是多少
	CarNo string `json:"car_no" form:"car_no" example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" example:"10"`
}

type CarDetailRsp struct {
	// 司机ID
	UserId uint `json:"user_id" example:"10"`
	// 车ID
	CarId string `json:"car_id" example:"川AA85409"`
	// 车牌号// 这里len是多少
	CarNo string `json:"car_no" example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享）
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
	// 路线
	Path string `json:"path" form:"path" binding:"required" example:"1"`
	// 车类型（1-网约车特享，2-网约车惠享）
	CarType entity.CarType `json:"car_type" form:"car_type" binding:"required" enums:"1,2" example:"2"`
}

type TrackCreateRsp struct {
	// 出发地点
	SrcPoint string `json:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" example:"华为技术有限公司三号门"`
	// 车类型（1-网约车特享，2-网约车惠享）
	CarTypes []*entity.CarType `json:"car_type" enums:"[1,2]" example:"2"`
	// 里程数（公里）
	TrackCourses uint `json:"track_courses" example:"100"`
	// 时长(分钟)
	TrackTime uint `json:"track_time" example:"9"`
}

type TrackCancelReq struct {
	// 行程进度(10-取消)
	TrackStatus entity.TrackStatus `json:"Schedule" form:"Schedule" binding:"required" enums:"10" example:"10"`
}

type TrackSetReq struct {
	// 行程进度(2-待赶来,3-待上车,4-进行中,5-待支付)
	TrackStatus entity.TrackStatus `json:"Schedule" form:"Schedule" binding:"required" enums:"2,3,4,5" example:"3"`
}

type TrackUpdateReq struct {
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 目的地
	DestPoint string `json:"dest_point" form:"dest_point" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" example:"1"`
}

type TrackListRsp struct {
	Tracks []*TrackItem `json:"tracks"`
}

type TrackItem struct {
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" example:"1"`
	// 车类型（1-网约车特享，2-网约车惠享）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2" example:"2"`
	// 出发地点
	SrcPoint string `json:"src_point" form:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" form:"dest_point"  example:"华为技术有限公司三号门"`
	// 行程状态(1-待接单,2-待赶来,3-待上车,4-进行中,5-待支付,6-已完成,10-已取消)
	TrackStatus entity.TrackStatus `json:"status" form:"status" enums:"1,2,3,4,5,6,10" example:"2"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type TrackDetailRsp struct {
	// 出发地点
	SrcPoint string `json:"src_point" example:"亚洲湾小区南门"`
	// 目的地
	DestPoint string `json:"dest_point" example:"华为技术有限公司三号门"`
	// 司机名
	UserName string `json:"user_name" example:"duome"`
	// 车牌号// 这里len是多少
	CarNo string `json:"car_no" example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享）
	CarType entity.CarType `json:"car_type" enums:"1,2" example:"2"`
	// 行程状态(1-待接单,2-待赶来,3-待上车,4-进行中,5-待支付,6-已完成,10-已取消)
	TrackStatus entity.TrackStatus `json:"status" enums:"1,2,3,4,5,6,10" example:"2"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type OrderDetailReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 司机ID
	ServicerId uint `json:"servicer_id" form:"servicer_id" binding:"required" example:"4"`
	// 车ID
	CarId string `json:"car_id" form:"car_id"  example:"川AA85409"`
	// 行程ID
	TrackId uint `json:"track_id" form:"track_id" binding:"required" example:"1"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type PaymentNotifyPayReq struct {
	// 金额
	Cost uint `json:"cost" example:"9.60"`
	// 支付方式(1-支付宝支付，2-微信支付，3-银行卡支付，4-云闪付支付)
	PaymentMethod entity.PaymentMethod `json:"payment_method" form:"payment_method" binding:"required" enums:"1,2,3,4" example:"1"`
}

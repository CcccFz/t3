package dto

import (
	"main/myTest/domain/entity"
)

type UserReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" bonding:"required" example:"10"`
}

type UserLoginReq struct {
	// 用户名
	UserName string `json:"user_name" form:"user_name" bonding:"required" example:"duome"`
	// 密码
	Password string `json:"password" form:"password" bonding:"required" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" bonding:"required,len=11"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" bonding:"required" enums:"1,2" example:"1"`
}

type UserUpdateReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" bonding:"required" example:"10"`
	// 用户名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 密码
	Password string `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" bonding:"len=11"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 旧密码
	Old_Password string `json:"old_Password" form:"old_password" example:"123456"`
}

type UserDetailRsp struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 用户名
	UserName string `json:"user_name" form:"user_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone" form:"phone"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 用户等级（1-大众，2-白银，3-黄金，4-铂金，5-黑金）
	UserLevel entity.UserLevel `json:"user_level" form:"user_level" enums:"1,2,3,4,5" example:"3"`
	// 成长值
	GrowthNumber uint `json:"growth_number" form:"growth_number" example:"10"`
}

type DriverReq struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" bonding:"required" example:"4"`
}

type DriverLoginReq struct {
	// 司机名
	DriverName string `json:"driver_name" form:"driver_name" bonding:"required" example:"ccccfz"`
	// 密码
	Password uint `json:"password" form:"password" bonding:"required" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" bonding:"required" example:"13981556666"`
	// 性别（1-男， 2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"2"`
}

type DriverUpdateReq struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" bonding:"required" example:"4"`
	// 司机名
	DriverName string `json:"driver_name" form:"driver_name" example:"duome"`
	// 密码
	Password uint `json:"password" form:"password" example:"123456"`
	// 手机号
	Phone uint `json:"phone" form:"phone" bonding:"len=11" example:"13981556666"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 老密码
	OldPassword string `json:"old_password" form:"old_password" example:"123456"`
}

type DriverDetailRsp struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" example:"10"`
	// 司机名
	DriverName string `json:"driver_name" form:"driver_name" example:"duome"`
	// 手机号
	Phone uint `json:"phone" form:"phone" example:"13981556666"`
	// 性别（1-男，2-女）
	Sex entity.Sex `json:"sex" form:"sex" enums:"1,2" example:"1"`
	// 等级（（1-大众，2-白银，3-黄金，4-铂金，5-黑金））
	UserLevel entity.UserLevel `json:"driver_level" form:"driver_level" enums:"1,2,3,4,5" example:"3"`
	// 成长值
	GrowthNumber uint `json:"growth_number" form:"growth_number" example:"10"`
	// 总里程数
	Courses uint `json:"courses" form:"courses" example:"100"`
}

type CarReq struct {
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id" bonding:"required" example:"川AA85409"`
}

type CarCreateReq struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" example:"10"`
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id" bonding:"required" example:"川AA85409"`
	// 车类型（1-网约车，2-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" bonding:"required" enums:"1,2" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" bonding:"required" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" bonding:"required" example:"10"`
}

type CarUpdateReq struct {
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id" example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2，3，4，5，6，7" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" example:"10"`
}

type CarDetailRsp struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" example:"10"`
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id"  example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2，3，4，5，6，7" example:"2"`
	// 车品牌
	CarBrand string `json:"car_brand" form:"car_brand" example:"东风"`
	// 车年限
	CarAge uint `json:"car_age" form:"car_age" example:"10"`
}

type RouteReq struct {
	// 行程ID
	RouteId uint `json:"route_id" form:"route_id" bonding:"required" example:"1"`
}

type RouteCreateReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 出发地点
	Outset string `json:"outset" form:"outset" bonding:"required" example:"亚洲湾小区南门"`
	// 目的地
	Destination string `json:"destination" form:"destination" bonding:"required" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" bonding:"required" example:"1"`
	// 车类型
	CarType entity.CarType `json:"car_type" form:"car_type" bonding:"required" enums:"1,2,3,4,5,6,7" example:"2"`
}

type RouteCreateRsp struct {
	// 出发地点
	Outset string `json:"outset" form:"outset" example:"亚洲湾小区南门"`
	// 目的地
	Destination string `json:"destination" form:"destination" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" example:"1"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2,3,4,5,6,7" example:"2"`
	// 里程数（公里）
	RouteCourses uint `json:"route_courses" form:"route_courses" example:"100"`
	// 时长(分钟)
	RouteTime uint `json:"route_time" form:"route_time" example:"9"`
}

type RouteMatchReq struct {
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" bonding:"required" example:"4"`
}

type RouteSetReq struct {
	// 行程进度(1-司机正在赶过来，2-司机已到达，3-开始行程，4-结束行程，5-取消行程)
	RouteSchedule entity.RouteSchedule `json:"Schedule" form:"Schedule" bonding:"required" enums:"1,2,3,4,5" example:"3"`
}

type RouteUpdateReq struct {
	// 行程ID
	RouteId uint `json:"route_id" form:"route_id" bonding:"required" example:"1"`
	// 目的地
	Destination string `json:"destination" form:"destination" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" example:"1"`
}

type RouteListRsp struct {
	Routes []*RouteItem `json:"routes"`
}

type RouteItem struct {
	// 行程ID
	RouteId uint `json:"route_id" form:"route_id" example:"1"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2，3，4，5，6，7" example:"2"`
	// 出发地点
	Outset string `json:"outset" form:"outset" example:"亚洲湾小区南门"`
	// 目的地
	Destination string `json:"destination" form:"destination"  example:"华为技术有限公司三号门"`
	// 行程状态(1-进行中，2-已完成，3-已取消)
	RouteStatus entity.RouteStatus `json:"status" form:"status" enums:"1,2,3" example:"2"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type RouteDetailRsp struct {
	// 出发地点
	Outset string `json:"outset" form:"outset" example:"亚洲湾小区南门"`
	// 目的地
	Destination string `json:"destination" form:"destination" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" example:"1"`
	// 金额
	Cost uint `json:"cost" form:"cost" example:"8.64"`
	// 成长值
	GrowthNumber uint `json:"growth_number" form:"growth_number" example:"10"`
	// 司机名
	DriverName string `json:"driver_name" form:"driver_name" example:"duome"`
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id"  example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2，3，4，5，6，7" example:"2"`
	// 行程状态(1-进行中，2-已完成，3-已取消)
	RouteStatus entity.RouteStatus `json:"status" form:"status" enums:"1,2,3" example:"2"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type OrderReq struct {
	// 订单ID
	OrderId uint `json:"order_id" form:"order_id" bonding:"required" example:"5"`
}

type OrderCreateReq struct {
	// 用户ID
	UserId uint `json:"user_id" form:"user_id" example:"10"`
	// 司机ID
	DriverId uint `json:"driver_id" form:"driver_id" bonding:"required" example:"4"`
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id"  example:"川AA85409"`
	// 行程ID
	RouteId uint `json:"route_id" form:"route_id" bonding:"required" example:"1"`
}

type OrderCreateRsp struct {
	// 出发地点
	Outset string `json:"outset" form:"outset" example:"亚洲湾小区南门"`
	// 目的地
	Destination string `json:"destination" form:"destination" example:"华为技术有限公司三号门"`
	// 路线
	Path string `json:"path" form:"path" example:"1"`
	// 金额
	Cost uint `json:"cost" form:"cost" example:"8.64"`
	// 司机名
	DriverName string `json:"driver_name" form:"driver_name" example:"duome"`
	// 车牌号// 这里len是多少
	CarId string `json:"car_id" form:"car_id"  example:"川AA85409"`
	// 车类型（1-网约车特享，2-网约车惠享，3-网约车快享，4-网约车速享，5-拼车，6-自营出租车，7-出租车）
	CarType entity.CarType `json:"car_type" form:"car_type" enums:"1,2，3，4，5，6，7" example:"2"`
	// 行程状态(1-进行中，2-已完成，3-已取消)
	RouteStatus entity.RouteStatus `json:"status" form:"status" enums:"1,2,3" example:"2"`
	// 创建时间（时间戳）
	CreatedAt int64 `json:"created_at" example:"1" example:"1630724893"`
}

type OrderSetReq struct {
	// 订单状态(1-开启，2-关闭)
	OrderStatus entity.OrderStatus `json:"order_status" form:"order_status" bonding:"required" enums:"1,2" example:"1"`
}

type PaymentReq struct {
	// 支付ID
	PaymentId uint `json:"payment_id" form:"payment_id" bonding:"required" example:"6"`
}

type PaymentCreateReq struct {
	// 里程数（公里）
	RouteCourses uint `json:"route_courses" form:"route_courses" bonding:"required" example:"100"`
	// 时长(分钟)
	RouteTime uint `json:"route_time" form:"route_time" bonding:"required" example:"9"`
}

type PaymentCreateRsp struct {
	// 金额
	Cost uint `json:"cost" form:"cost" bonding:"required" example:"9.60"`
	// 支付状态（1-待支付，2-已支付）
	PaymentStatus entity.PaymentStatus `json:"payment_status" form:"payment_status" bonding:"required" enums:"1,2" example:"1"`
	// 行程进度(1-司机正在赶过来，2-司机已到达，3-开始行程，4-结束行程，5-取消行程)
	RouteSchedule entity.RouteSchedule `json:"Schedule" form:"Schedule" bonding:"required" enums:"1,2,3,4,5" example:"3"`
}

type PaymentSetReq struct {
	// 支付方式(1-支付宝支付，2-微信支付，3-银行卡支付，4-云闪付支付)
	PaymentMethod entity.PaymentMethod `json:"payment_method" form:"payment_method" bonding:"required" enums:"1,2,3,4" example:"1"`
	// 支付状态（1-待支付，2-已支付）
	PaymentStatus entity.PaymentStatus `json:"payment_status" form:"payment_status" bonding:"required" enums:"1,2" example:"1"`
}

type PaymentDetailRsp struct {
	// 支付ID
	PaymentId uint `json:"payment_id" form:"payment_id" bonding:"required" example:"6"`
	// 金额
	Cost uint `json:"cost" form:"cost" bonding:"required" example:"9.60"`
	// 里程数（公里）
	RouteCourses uint `json:"route_courses" form:"route_courses" bonding:"required" example:"100"`
	// 时长(分钟)
	RouteTime uint `json:"route_time" form:"route_time" bonding:"required" example:"9"`
	// 支付方式(1-支付宝支付，2-微信支付，3-银行卡支付，4-云闪付支付)
	PaymentMethod entity.PaymentMethod `json:"payment_method" form:"payment_method" bonding:"required" enums:"1,2,3,4" example:"1"`
	// 支付状态（1-待支付，2-已支付）
	PaymentStatus entity.PaymentStatus `json:"payment_status" form:"payment_status" bonding:"required" enums:"1,2" example:"1"`
}

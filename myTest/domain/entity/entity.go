package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserType uint8

const (
	UserTypeUser   UserType = 1 // 乘客
	UserTypeDriver _        = 2 // 司机
)

type CarType uint8

const (
	CarTypeCosy      CarType = 1 // 优享
	CarTypeFavorable _       = 2 // 惠享
)

type TrackStatus uint8

const (
	TrackStatusOrdering  TrackStatus = 1
	TrackStatusComing    _           = 2
	TrackStatusBoarding  _           = 3
	TrackStatusDriving   _           = 4
	TrackStatusPaying    _           = 5
	TrackStatusCompleted _           = 6
	TrackStatusCancel    _           = 10
)

type PaymentStatus uint8

const (
	PaymentStatusWaitPaid PaymentStatus = 1
	PaymentStatusPaid     _             = 2
)

type PaymentMethod uint8

const (
	PaymentMethodAlipay         PaymentMethod = 1
	PaymentMethodWeChat         _             = 2
	PaymentMethodBankcard       _             = 3
	PaymentMethodCloudQuickPass _             = 4
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TUser struct {
	UserId   uint   `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"colume:user_name;type:varchar(256);not null"`
	Password string `gorm:"colume:password;type:varchar(256);not null"`
	Phone    string `gorm:"colume:phone;type:varchar(20);not null"`
	Model
}

type TDriver struct {
	UserId   uint   `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"colume:user_name;type:varchar(256);not null"`
	Password string `gorm:"colume:password;type:varchar(256);not null"`
	Phone    string `gorm:"colume:phone;type:varchar(256);not null"`
	Distance uint   `gorm:"colume:distance;type:integer;default:0;not null"`
	CarId    uint   `gorm:"colume:car_id;type:integer;not null"`
	*Model
}

type TCar struct {
	CarId    uint    `gorm:"colume:car_id;primary_key;AUTO_INCREMENT"`
	CarNo    string  `gorm:"colume:car_no;type:varchar(256);not null"`
	CarType  CarType `gorm:"colume:car_type;type:tinyint;not null"`
	CarBrand string  `gorm:"colume:car_brand;type:varchar(256);not null"`
	CarAge   uint    `gorm:"colume:car_age:courses;type:integer;not null"`
	*Model
}

type TTrack struct {
	TrackId     uint        `gorm:"colume:track_id;primary_key;AUTO_INCREMENT"`
	UserId      uint        `gorm:"colume:user_id;type:integer;not null"`
	OrderId     uint        `gorm:"colume:order_id;type:integer"`
	ServicerId  uint        `gorm:"colume:servicer_id;type:integer"`
	SrcPoint    string      `gorm:"colume:scr_point;type:varchar(256);not null"`
	DestPoint   string      `gorm:"colume:dest_point;type:varchar(256);not null"`
	TrackStatus TrackStatus `gorm:"colume:track_status;type:tinyint;not null"`
	Distance    uint        `gorm:"colume:distance;type:integer;not null;comment:公里数"`
	Duration    uint        `gorm:"colume:duration;type:integer;not null;comment:时长"`
	StartAt     *time.Time  `gorm:"colume:start_at"`
	EndAt       *time.Time  `gorm:"colume:end_at"`
	*Model
}

type TOrder struct {
	OrderId    uint `gorm:"colume:order_id;primary_key;AUTO_INCREMENT"`
	UserId     uint `gorm:"colume:user_id;type:integer;not null"`
	ServicerId uint `gorm:"colume:servicer_id;type:integer;not null"`
	PaymentId  uint `gorm:"colume:payment_id;primary_key;AUTO_INCREMENT"`
	StartCost  uint `gorm:"colume:start_cost;type:integer;not null"`
	UnitCost   uint `gorm:"colume:unit_cost;type:integer;not null"`
	Cost       uint `gorm:"colume:cost;type:integer"`
	*Model
}

type TPayment struct {
	PaymentId     uint          `gorm:"colume:payment_id;primary_key;AUTO_INCREMENT"`
	Cost          uint          `gorm:"colume:cost;type:integer;not null"`
	PaymentMethod PaymentMethod `gorm:"colume:payment_method;type:tinyint"`
	PaymentStatus PaymentStatus `gorm:"colume:payment_status;type:tinyint;not null"`
	*Model
}

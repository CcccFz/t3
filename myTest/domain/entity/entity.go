package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserType uint8

type CarType uint8

type TrackStatus uint8

type PaymentMethod uint8

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	UserId   uint   `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"colume:user_name;type:varchar(256);not null"`
	Password string `gorm:"colume:password;type:varchar(256);not null"`
	Phone    string `gorm:"colume:phone;type:varchar(20);not null"`
	*Model
}

type Driver struct {
	DriverId   uint   `gorm:"colume:driver_id;primary_key;AUTO_INCREMENT"`
	DriverName string `gorm:"colume:driver_name;type:varchar(256);not null"`
	Password   string `gorm:"colume:password;type:varchar(256);not null"`
	Phone      string `gorm:"colume:phone;type:varchar(256);not null"`
	Courses    uint   `gorm:"colume:courses;type:integer;not null"`
	CarId      uint   `gorm:"colume:car_id;type:integer;not null"`
	*Model
}

type Car struct {
	CarId    uint    `gorm:"colume:car_id;primary_key;AUTO_INCREMENT"`
	CarNo    string  `gorm:"colume:car_no;type:varchar(256);not null"`
	CarType  CarType `gorm:"car_type;type:tinyint;not null"`
	CarBrand string  `gorm:"colume:car_brand;type:varchar(256);not null"`
	CarAge   uint    `gorm:"car_age:courses;type:integer;not null"`
	*Model
}

type Track struct {
	TrackId     uint        `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserId      uint        `gorm:"colume:user_id;type:integer;not null"`
	SrcPoint    string      `gorm:"colume:scr_point;type:varchar(256);not null"`
	DestPoint   string      `gorm:"colume:dest_point;type:varchar(256);not null"`
	TrackStatus TrackStatus `gorm:"colume:track_status;type:tinyint;not null"`
	TrackCourse uint        `gorm:"track_course;type:integer;not null"`
	StartTime   *time.Time  `gorm:"colume:start_time"`
	EndTime     *time.Time  `gorm:"colume:end_time"`
	*Model
}

type Order struct {
	OrderId   uint `gorm:"colume:order_id;primary_key;AUTO_INCREMENT"`
	UserId    uint `gorm:"colume:user_id;type:integer;not null"`
	ServiceId uint `gorm:"colume:service_id;type:integer;not null"`
	StartCost uint `gorm:"colume:start_cost;type:integer;not null"`
	UnitCost  uint `gorm:"colume:unit_cost;type:integet;not null"`
	*Model
}

type Payment struct {
	PaymentId     uint          `gorm:"colume:payment_id;primary_key;AUTO_INCREMENT"`
	OrderId       uint          `gorm:"colume:order_id;type:integer;not null"`
	Cost          uint          `gorm:"colume:cost;type:integer;not null"`
	PaymentMethod PaymentMethod `gorm:"colume:payment_method;type:tinyint"` //这里不传是为空么？
	*Model
}

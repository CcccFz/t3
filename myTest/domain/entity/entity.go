package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserLevel uint8

type CarLevel uint8

type CarType uint8

type Sex uint8

type RouteStatus uint8

type RouteSchedule uint8

type OrderStatus uint8

type PaymentMethod uint8

type PaymentStatus uint8

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	UserId       uint      `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserName     string    `gorm:"colume:user_name;type:varchar(256);not null"`
	Password     string    `gorm:"colume:password;type:varchar(256);not null"`
	Phone        string    `gorm:"colume:phone;type:varchar(20);not null"`
	Sex          Sex       `gorm:"colume:sex;type:tinyint;not null"`
	UserLevel    UserLevel `gorm:"user_level;type:tinyint"`
	GrowthNumber uint      `gorm:"colume:growth_number;type:integer"`
	*Model
}

type Driver struct {
	DriverId     uint      `gorm:"colume:driver_id;primary_key;AUTO_INCREMENT"`
	DriverName   string    `gorm:"colume:driver_name;type:varchar(256);not null"`
	Password     string    `gorm:"colume:password;type:varchar(256);not null"`
	Phone        string    `gorm:"colume:phone;type:varchar(256);not null"`
	Sex          Sex       `gorm:"colume:sex;type:tinyint;not null"`
	UserLevel    UserLevel `gorm:"colume:user_level;type:tinyint"`
	GrowthNumber uint      `gorm:"colume:growth_number;type:integer"`
	Courses      uint      `gorm:"colume:courses;type:integer"`
	*Model
}

type Car struct {
	CarId    uint    `gorm:"colume:car_id;primary_key;AUTO_INCREMENT"`
	DriverId uint    `gorm:"colume:driver_id;type:integer"`
	CarType  CarType `gorm:"car_type;type:tinyint"`
	CarBrand string  `gorm:"colume:car_brand;type:varchar(256);not null"`
	CarAge   uint    `gorm:"car_age:courses;type:integer;not null"`
	*Model
}

type Route struct {
	UserId      uint   `gorm:"colume:user_id;primary_key;type:integer"`
	Outset      string `gorm:"colume:outset;type:varchar(256);not null"`
	Destination string `gorm:"colume:destination;type:varchar(256);not null"`
	Path        string `gorm:"colume:path;type:varchar(256);not null"`
	Cost        uint   `gorm:"car_age:courses;type:integer;not null"`
}

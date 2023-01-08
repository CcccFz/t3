package entity

import (
	"gorm.io/gorm"
	"time"
)

type Level uint8

type CarLevel uint8

type Cartype uint8

type Sex uint8

type RouteStatus uint8

type RouteSchedule uint8

type OrderStatus uint8

type PaymentMethod uint8

type PaymentStatus uint8

type Order struct {
	OrderId   uint   `gorm:"colume:order_id;primary_key;AUTO_INCREMENT"`
	OrderName string `gorm:"colume:order_name;type:varchar(256);not null"`
	UserId    uint   `gorm:"colume:user_id;type:integer;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	UserId    uint   `gorm:"colume:user_id;primary_key;AUTO_INCREMENT"`
	UserName  string `gorm:"colume:user_name;type:varchar(256);not null"`
	Password  string `gorm:"password;type:varchar(256);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

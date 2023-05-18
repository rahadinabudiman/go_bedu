package models

import "github.com/jinzhu/gorm"

type Administrator struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// For Response Get All Admin
type AdminsResponse struct {
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
}

// For Response Create Admin
type AdminCreateRES struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Response for token JWT
type AdminsJWTRES struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	AdministratorID uint   `json:"administrator_id" form:"administrator_id"`
	Title           string `json:"title" form:"title" validate:"required"`
	Content         string `json:"content" form:"content" validate:"required"`
	ImageLink       string `json:"image_link" form:"image_link" validate:"required"`
}

// For Response Get Article
type ArticlesResponse struct {
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	ImageLink string `json:"image_link" form:"image_link"`
}

// For Params Get Article
type ArticleParams struct {
	ID uint `json:"id" form:"id"`
}

package models

import "time"

// Post struct
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type:varchar(180)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

func GetPostsByUserID(id int) (*[]Post, error) {
	var posts []Post
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

func GetPostsByUserIDPageAndLimit(id, page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post

	offset := (page - 1) * limit

	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}

	db.Model(&Post{}).Where("user_id=?", id).Count(&total)
	return &posts, total, nil
}

func GetPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post

	offset := (page - 1) * limit

	if err := db.Preload("User").Order("timestamp desc").Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}

	db.Model(&Post{}).Count(&total)
	return &posts, total, nil
}

package models

import (
	"fmt"
	"github.com/elfgzp/go_blog/config"
	"log"
	"time"
)

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	LastSeen     *time.Time
	AboutMe      string `gorm:"type:varchar(140)"`
	Avatar       string `gorm:"type:varchar(200)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("%s/%s?d=identicon", config.GravatarURL, Md5(email))
}

func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return user.FollowSelf()
}

func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(user).Updates(contents).Error
}

func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}

func (u *User) Follow(username string) error {
	otherUser, err := GetUserByUsername(username)

	if err != nil {
		return err
	}

	return db.Model(otherUser).Association("Followers").Append(u).Error
}

func (u *User) UnFollow(username string) error {
	otherUser, err := GetUserByUsername(username)

	if err != nil {
		return err
	}

	return db.Model(otherUser).Association("Followers").Delete(u).Error
}

func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count() - 1
}

func (u *User) FollowingIDs() []int {
	var ids []int

	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id, follower_id").Rows()

	if err != nil {
		log.Println("Counting Following error: ", err)
	}

	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		_ = rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}

	return ids
}

func (u *User) FollowingCount() int {
	ids := u.FollowingIDs()
	return len(ids) - 1
}

func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post

	ids := u.FollowingIDs()

	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)

	ids := user.FollowingIDs()
	for _, id := range ids {
		if u.ID == id {
			return true
		}
	}
	return false
}

func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

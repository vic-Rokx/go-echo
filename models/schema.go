package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var Tables = []string{"user", "category", "article"}

type User struct {
	gorm.Model
	ID             string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string     `json:"name"`
	Email          string     `gorm:"unique" json:"email"`
	Password       string     `json:"password"`
	UserCategories []Category `gorm:"many2many:user_categories;"`
}

type Category struct {
	gorm.Model
	Name string
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}

type Article struct {
	gorm.Model
	ID                string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title             string     `json:"title"`
	Author            string     `json:"author"`
	Source            string     `json:"source"`
	Image             string     `json:"image"`
	Summary           string     `json:"summary"`
	ArticleCategories []Category `gorm:"many2many:article_categories;"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *Article) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}

type Post struct {
	gorm.Model
	ID      string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Source  string   `json:"source"`
	Image   string   `json:"image"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

type NewsResponse struct {
	gorm.Model
	Author string      `json:"author"`
	Data   interface{} `json:"data"`
}

func ValidateArticle(article Article) error {
	validate := validator.New()
	return validate.Struct(article)
}

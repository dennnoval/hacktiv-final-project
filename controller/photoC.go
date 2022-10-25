package controller

import (
	e "hacktiv-final-project/entity"

	g "github.com/gin-gonic/gin"

	// h "net/http"
	"gorm.io/gorm"
	// s "strconv"
	// b "golang.org/x/crypto/bcrypt"
)

func GetAllPhotos(ctx *g.Context) {}

func AddPhoto(ctx *g.Context) {
	p := ctx.MustGet("photo").(e.Photo)
	sub := ctx.MustGet("sub").(map[string]interface{})
	p.UserID = int(sub["user_id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&p).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add new photo!"})
	} else {
		ctx.JSON(201, g.H{
			"id": p.UserID, "title": p.Title, "caption": p.Caption,
			"photo_url": p.PhotoUrl, "created_at": p.CreatedAt,
		})
	}
}

func UpdatePhoto(ctx *g.Context) {}

func DeletePhoto(ctx *g.Context) {}

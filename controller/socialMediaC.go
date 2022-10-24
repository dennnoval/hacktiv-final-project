package controller

import (
	e "hacktiv-final-project/entity"
	g "github.com/gin-gonic/gin"
	// h "net/http"
	"gorm.io/gorm"
	s "strconv"
	// b "golang.org/x/crypto/bcrypt"
	// "log"
)

func GetAllSocialMedias(ctx *g.Context) {
	sm := []e.SocialMedia{}
	db := ctx.MustGet("db").(*gorm.DB)
	f := db.Preload("User").Find(&sm)
	c := f.RowsAffected
	if err := f.Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": err.Error()})
	} else {
		amm := make([]map[string]interface{}, c)
		var p e.Photo
		for i, _ := range sm {
			smm := make(map[string]interface{})
			smm["id"] = sm[i].SocialMediaID
			smm["name"] = sm[i].Name
			smm["social_media_url"] = sm[i].SocialMediaUrl
			smm["UserId"] = sm[i].UserID
			smm["createdAt"] = sm[i].CreatedAt
			smm["updatedAt"] = sm[i].UpdatedAt
			p = e.Photo{UserID: sm[i].UserID}
			db.Take(&p)
			u := make(map[string]interface{})
			u["id"] = sm[i].UserID
			u["username"] = sm[i].User.Username
			u["profile_image_url"] = p.PhotoUrl
			smm["User"] = u
			amm[i] = smm
		}
		ctx.JSON(200, g.H{"social_medias": amm})
	}
}

func AddSocialMedia(ctx *g.Context) {
	sm := ctx.MustGet("social_media").(e.SocialMedia)
	sub := ctx.MustGet("sub").(map[string]interface{})
	sm.UserID = int(sub["user_id"].(float64))
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&sm).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add social media!"})
	} else {
		ctx.JSON(201, g.H{
			"id": sm.SocialMediaID, "name": sm.Name, "social_media_url": sm.SocialMediaUrl, 
			"user_id": sm.UserID, "created_at": sm.CreatedAt,
		})
	}
}

func UpdateSocialMedia(ctx *g.Context) {
	sm := ctx.MustGet("social_media").(e.SocialMedia)
	smId, _ := s.Atoi(ctx.Param("socialMediaId"))
	sm.SocialMediaID = smId
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Model(&sm).Where("id = ?", sm.SocialMediaID).Updates(&sm).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message": err.Error()})
	} else {
		db.Take(&sm)
		ctx.JSON(201, g.H{
			"id": sm.SocialMediaID, "name": sm.Name, "social_media_url": sm.SocialMediaUrl, 
			"user_id": sm.UserID, "updated_at": sm.CreatedAt,
		})
	}
}

func DeleteSocialMedia(ctx *g.Context) {
	smId, _ := s.Atoi(ctx.Param("socialMediaId"))
	sm := e.SocialMedia{SocialMediaID: smId}
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Delete(&sm).Error; err != nil {
		ctx.AbortWithStatusJSON(200, g.H{"message":"Social media doesn't exists"})
		return
	}
	ctx.JSON(200, g.H{"message":"Your social media has been deleted successfully"})
}

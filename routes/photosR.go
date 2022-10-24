package router

import (
	g "github.com/gin-gonic/gin"
	c "hacktiv-final-project/controller"
)

func PhotosRoutes(r *g.Engine) {
	rg := r.Group("/photos")
	rg.GET("", c.GetAllPhotos)
	rg.POST("", c.AddPhoto)
	rg.PUT("/:photoId", c.UpdatePhoto)
	rg.DELETE("/:photoId", c.DeletePhoto)
}

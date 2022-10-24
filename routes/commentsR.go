package router

import (
	g "github.com/gin-gonic/gin"
	c "hacktiv-final-project/controller"
)

func CommentsRoutes(r *g.Engine) {
	rg := r.Group("/comments")
	rg.GET("", c.GetAllComments)
	rg.POST("", c.AddComment)
	rg.PUT("/:commentId", c.UpdateComment)
	rg.DELETE("/:commentId", c.DeleteComment)
}

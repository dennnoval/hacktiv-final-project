package main

import (
	g "github.com/gin-gonic/gin"
	conf "hacktiv-final-project/config"
	r "hacktiv-final-project/routes"
	m "hacktiv-final-project/middleware"
)

func main() {
	// g.SetMode(g.ReleaseMode)
	router := g.Default()
	router.Use(func(ctx *g.Context) {
		ctx.Set("db", conf.DB)
	})
	r.UsersRoutes(router, m.Authentication(), m.Authorization(), m.RegisterUserValidation(), m.UpdateUserValidation())
	r.PhotosRoutes(router)
	r.CommentsRoutes(router)
	r.SocialMediasRoutes(router, m.Authorization(), m.SocialMediaValidation())
	router.Run()
}

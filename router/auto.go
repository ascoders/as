package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	"newWoku/lib/captcha"
	
)

func AutoRoute(r martini.Router) {
    
    user := user.New()
    r.Get("/api/users/authentication", user.Authentication)
    r.Post("/api/users/authentication", captcha.Check, user.CreateAuthentication)
    r.Get("/api/users", user.Gets)
    r.Get("/api/users/:id", user.Get)
    r.Post("/api/users", user.Add)
    r.Patch("/api/users/:id", user.Update)
    r.Delete("/api/users/:id", user.Delete)
	
    article := article.New()
    r.Get("/api/aa", article.Other)
    r.Get("/api/articles", article.Gets)
    r.Get("/api/articles/:id", article.Get)
    r.Post("/api/articles", article.Add)
    r.Patch("/api/articles/:id", article.Update)
    r.Delete("/api/articles/:id", article.Delete)
	
}

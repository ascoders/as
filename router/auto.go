package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/article"
	"newWoku/controllers/user"
	"newWoku/lib/csrf"
	"newWoku/lib/captcha"
	
)

func AutoRoute(r martini.Router) {
    
    article := article.New()
    r.Get("/api/aa", article.Other)
    r.Get("/api/articles", article.Gets)
    r.Get("/api/articles/:id", article.Get)
    r.Post("/api/articles", csrf.Validate, article.Add)
    r.Patch("/api/articles/:id", csrf.Validate, article.Update)
    r.Delete("/api/articles/:id", csrf.Validate, article.Delete)
	
    user := user.New()
    r.Get("/api/users/authentication", user.Authentication)
    r.Post("/api/users/authentication", captcha.Check, user.AuthenticationCreate)
    r.Post("/api/users/authentication/email", user.CreateEmailAuthentication)
    r.Get("/api/users/current", user.Current)
    r.Delete("/api/users/authentication", user.AuthenticationDelete)
    r.Get("/api/users", user.Gets)
    r.Get("/api/users/:id", user.Get)
    r.Post("/api/users", csrf.Validate, user.Add)
    r.Patch("/api/users/:id", csrf.Validate, user.Update)
    r.Delete("/api/users/:id", csrf.Validate, user.Delete)
	
}

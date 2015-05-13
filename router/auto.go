package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	"newWoku/controllers/app"
	"newWoku/lib/csrf"
)

func AutoRoute(r martini.Router) {
    
    user := user.New()
    r.Get("/api/users/:id", csrf.Validate, user.Get)
    r.Get("/api/example", user.Other1)
    r.Get("/api/example/:id", user.Other2)
    r.Post("/api/example", user.Other3)
    r.Put("/api/example", user.Other3)
    r.Delete("/api/example", csrf.Validate, user.Other4)
    r.Get("/api/example", csrf.Validate, user.Other5)
    r.Get("/api/users", user.Gets)
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
	
    app := app.New()
    r.Get("/api/app/xx", app.Other)
    r.Get("/api/apps", app.Gets)
    r.Get("/api/apps/:id", app.Get)
    r.Post("/api/apps", app.Add)
    r.Patch("/api/apps/:id", app.Update)
    r.Delete("/api/apps/:id", app.Delete)
	
}

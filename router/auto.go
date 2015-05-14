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
    r.Get("/api/users", csrf.Validate, user.Gets)
    r.Post("/api/users", csrf.Validate, user.Add)
    r.Patch("/api/users/:id", csrf.Validate, user.Update)
    r.Delete("/api/users/:id", csrf.Validate, user.Delete)
	
    article := article.New()
    r.Get("/api/aa", article.Other)
    r.Get("/api/articles", csrf.Validate, article.Gets)
    r.Get("/api/articles/:id", csrf.Validate, article.Get)
    r.Post("/api/articles", csrf.Validate, article.Add)
    r.Patch("/api/articles/:id", csrf.Validate, article.Update)
    r.Delete("/api/articles/:id", csrf.Validate, article.Delete)
	
    app := app.New()
    r.Get("/api/app/xx", app.Other)
    r.Get("/api/apps", csrf.Validate, app.Gets)
    r.Get("/api/apps/:id", csrf.Validate, app.Get)
    r.Post("/api/apps", csrf.Validate, app.Add)
    r.Patch("/api/apps/:id", csrf.Validate, app.Update)
    r.Delete("/api/apps/:id", csrf.Validate, app.Delete)
	
}

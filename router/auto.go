package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/article"
	"newWoku/controllers/app"
	"newWoku/controllers/user"
	"newWoku/lib/csrf"
)

func AutoRoute(r martini.Router) {
    
    article := &article.Article{}
    r.Get("/api/aa", article.Before, article.Other)
    r.Get("/api/articles", article.Before, article.Gets)
    r.Get("/api/articles/:id", article.Before, article.Get)
    r.Post("/api/articles", article.Before, article.Add)
    r.Patch("/api/articles/:id", article.Before, article.Update)
    r.Delete("/api/articles/:id", article.Before, article.Delete)
	
    app := &app.App{}
    r.Get("/api/app/xx", app.Before, app.Other)
    r.Get("/api/apps", app.Before, app.Gets)
    r.Get("/api/apps/:id", app.Before, app.Get)
    r.Post("/api/apps", app.Before, app.Add)
    r.Patch("/api/apps/:id", app.Before, app.Update)
    r.Delete("/api/apps/:id", app.Before, app.Delete)
	
    user := &user.User{}
    r.Get("/api/example", user.Before, user.Other1)
    r.Get("/api/example/:id", user.Before, user.Other2)
    r.Post("/api/example", user.Before, user.Other3)
    r.Put("/api/example", user.Before, user.Other3)
    r.Delete("/api/example", csrf.Validate, user.Before, user.Other4)
    r.Get("/api/example", csrf.Validate, user.Before, user.Other5)
    r.Get("/api/users", user.Before, user.Gets)
    r.Get("/api/users/:id", user.Before, user.Get)
    r.Post("/api/users", user.Before, user.Add)
    r.Patch("/api/users/:id", user.Before, user.Update)
    r.Delete("/api/users/:id", user.Before, user.Delete)
	
}

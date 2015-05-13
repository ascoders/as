package router

import (
    "github.com/go-martini/martini"
	"newWoku/lib/csrf"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	"newWoku/controllers/app"
	
)

func AutoRoute(r martini.Router) {
    
    user := &user.User{}
    r.Get("/api/example", csrf.Validate, user.Before, user.Other1)
    r.Get("/api/example/:id", csrf.Validate, user.Before, user.Other2)
    r.post("/api/example", csrf.Validate, user.Before, user.Other3)
    r.put("/api/example", csrf.Validate, user.Before, user.Other3)
    r.delete("/api/example", csrf.Validate, user.Before, user.Other4)
    r.Get("/api/example", csrf.Validate, user.Before, user.Other5)
    r.Get("/api/users", csrf.Validate, user.Before, user.Gets)
    r.Get("/api/users/:id", csrf.Validate, user.Before, user.Get)
    r.Post("/api/users", csrf.Validate, user.Before, user.Add)
    r.Patch("/api/users/:id", csrf.Validate, user.Before, user.Update)
    r.Delete("/api/users/:id", csrf.Validate, user.Before, user.Delete)
	
    article := &article.Article{}
    r.Get("/api/aa", csrf.Validate, article.Before, article.Other)
    r.Get("/api/articles", csrf.Validate, article.Before, article.Gets)
    r.Get("/api/articles/:id", csrf.Validate, article.Before, article.Get)
    r.Post("/api/articles", csrf.Validate, article.Before, article.Add)
    r.Patch("/api/articles/:id", csrf.Validate, article.Before, article.Update)
    r.Delete("/api/articles/:id", csrf.Validate, article.Before, article.Delete)
	
    app := &app.App{}
    r.Get("/api/app/xx", csrf.Validate, app.Before, app.Other)
    r.Get("/api/apps", csrf.Validate, app.Before, app.Gets)
    r.Get("/api/apps/:id", csrf.Validate, app.Before, app.Get)
    r.Post("/api/apps", csrf.Validate, app.Before, app.Add)
    r.Patch("/api/apps/:id", csrf.Validate, app.Before, app.Update)
    r.Delete("/api/apps/:id", csrf.Validate, app.Before, app.Delete)
	
}

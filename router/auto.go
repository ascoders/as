package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	
)

func AutoRoute(r martini.Router) {
    
    user := &user.User{}
    r.Get("/api/users", user.Before ,user.Gets)
    r.Get("/api/users/:id", user.Before ,user.Get)
    r.Post("/api/users", user.Before ,user.Add)
    r.Patch("/api/users", user.Before ,user.Update)
    r.Delete("/api/users/:id", user.Before ,user.Delete)
    r.Get("/api/bb", user.Before ,user.Other)
	
    article := &article.Article{}
    r.Get("/api/articles", article.Before ,article.Gets)
    r.Get("/api/articles/:id", article.Before ,article.Get)
    r.Post("/api/articles", article.Before ,article.Add)
    r.Patch("/api/articles", article.Before ,article.Update)
    r.Delete("/api/articles/:id", article.Before ,article.Delete)
    r.Get("/api/aa", article.Before ,article.Other)
	
}

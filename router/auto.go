package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	
)

func AutoRoute(r martini.Router) {
    
    user := &user.User{}
    r.Get("/users", user.Before ,user.Gets)
    r.Get("/users/:id", user.Before ,user.Get)
    r.Post("/users", user.Before ,user.Add)
    r.Put("/users", user.Before ,user.Update)
    r.Delete("/users/:id", user.Before ,user.Delete)
    r.Get("/other", user.Before ,user.Other)
    r.Post("/xxx", user.Before ,user.Xxxx)
	
    article := &article.Article{}
    r.Get("/articles", article.Before ,article.Gets)
    r.Get("/articles/:id", article.Before ,article.Get)
    r.Post("/articles", article.Before ,article.Add)
    r.Put("/articles", article.Before ,article.Update)
    r.Delete("/articles/:id", article.Before ,article.Delete)
    r.Get("/aa", article.Before ,article.Other)
	
}

package router

import (
    "github.com/go-martini/martini"
    "newWoku/controllers/user"
	"newWoku/controllers/article"
	
)

func AutoRoute(r martini.Router) {
    
    user := &user.User{}
    r.Get("/other", user.Before ,user.Other)
    r.Post("/xxx", user.Before ,user.Xxxx)
	
    article := &article.Article{}
    r.Get("/article", article.Before ,article.Other)
	
}

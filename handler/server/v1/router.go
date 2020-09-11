package v1

import (
	"github.com/go-chi/chi"
	"net/http"
)

// Routes returns post router with each endpoint.
func (pr *PostRouter) RoutesPost() http.Handler {
	newRouter := chi.NewRouter()

	newRouter.Get("/user/{userId}", pr.GetByUserHandler)
	newRouter.Get("/", pr.GetAllPost)
	newRouter.Post("/", pr.CreateHandler)
	newRouter.Get("/{id}", pr.GetOneHandler)
	newRouter.Put("/{id}", pr.UpdateHandler)
	newRouter.Delete("/{id}", pr.DeleteHandler)

	return newRouter
}

// Routes returns user router with each endpoint.
func (ur *UserRouter) RoutesUser() http.Handler {
	newRouter := chi.NewRouter()

	newRouter.Get("/", ur.GetAllUser)
	newRouter.Post("/", ur.CreateHandler)
	newRouter.Get("/{id}", ur.GetOneHandler)
	newRouter.Put("/{id}", ur.UpdateHandler)
	newRouter.Delete("/{id}", ur.DeleteHandler)

	return newRouter
}
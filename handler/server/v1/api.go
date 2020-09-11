package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	data "microblog/database"
	repo "microblog/repository"
)

// New returns the API V1 Handler with configuration.
func New(conn *data.Data) http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &repo.UserRepository{
			Data: conn,
		},
	}
	r.Mount("/users", ur.RoutesUser())

	pr := &PostRouter{
		Repository: &repo.PostRepository{
			Data: conn,
		},
	}
	r.Mount("/posts", pr.RoutesPost())

	return r
}

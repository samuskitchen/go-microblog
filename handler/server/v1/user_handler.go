package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"microblog/domain/response"
	"microblog/domain/user"
)

// UserRouter is the router of the users.
type UserRouter struct {
	Repository user.Repository
}

// CreateHandler Create a new user.
func (ur *UserRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if err := user.HashPassword(); err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = user.Validate("")
	if err != nil {
		response.HTTPError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &user)
	if err != nil {
		response.HTTPError(w, r, http.StatusConflict, err.Error())
		return
	}

	user.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), user.ID))
	response.JSON(w, r, http.StatusCreated, user)
}

// GetAllUser response all the users.
func (ur *UserRouter) GetAllUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := ur.Repository.GetAllUser(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, users)
}

// GetOneHandler response one user by id.
func (ur *UserRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	userResult, err := ur.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, userResult)
}

// UpdateHandler update a stored user by id.
func (ur *UserRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var userUpdate user.User
	err = json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	err = userUpdate.Validate("")
	if err != nil {
		response.HTTPError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Update(ctx, uint(id), userUpdate)
	if err != nil {
		response.HTTPError(w, r, http.StatusConflict, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a user by ID.
func (ur *UserRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusNoContent, response.Map{})
}

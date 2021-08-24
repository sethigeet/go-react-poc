package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/sethigeet/go-react-poc/packages/server/database"
	"github.com/sethigeet/go-react-poc/packages/server/model"
	"github.com/sethigeet/go-react-poc/packages/server/util"
	"github.com/sethigeet/go-react-poc/packages/server/validator"
)

type UserHandler struct {
	r *mux.Router
}

var db *gorm.DB

// Apply creates a new subrouter on the main router and applies the appropriate
// routes on it
func (userHandler UserHandler) Apply() {
	db = database.DB

	subRouter := userHandler.r.PathPrefix("/user").Subrouter()

	subRouter.HandleFunc("/", userHandler.GetAll).Methods(http.MethodGet)
	subRouter.HandleFunc("/", userHandler.Create).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}", userHandler.Get).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", userHandler.Delete).Methods(http.MethodDelete)
}

// GetAll retreives all the users from the databse and gives back a JSON response
func (UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		util.ErrorResponse(w, "Unable to fetch the users")
		return
	}

	util.OKResponse(w, "users", users)
}

// Get retreives the user from the databse according to the ID specified in the
// url and gives back a JSON response
func (UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Get the variables from the url
	vars := mux.Vars(r)

	var user model.User
	err := db.First(&user, "id = ?", vars["id"]).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		util.NotFoundResponse(w, "user")
		return
	} else if err != nil {
		util.ErrorResponse(w, "Unable to fetch the user")
		return
	}

	util.OKResponse(w, "user", user)
}

// Create creates a user in the database and returns a JSON response
func (UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := util.ParseJSONBody(r)
	if err != nil {
		util.BadRequestResponse(w, "Bad request body")
		return
	}

	user := model.User{
		Email:    body["email"],
		Username: body["username"],
	}

	if err := validator.Validate(user); err != nil {
		util.BadRequestResponse(w, err)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			util.BadRequestResponse(w, map[string]string{
				"email": fmt.Sprintf(validator.ErrorMessages["alreadyExists"], "user", "email"),
			})
			return
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			util.BadRequestResponse(w, map[string]string{
				"email": fmt.Sprintf(validator.ErrorMessages["alreadyExists"], "user", "username"),
			})
			return
		} else {
			util.ErrorResponse(w, "Unable to create the user")
			return
		}
	}

	util.CreatedResponse(w, "user", user)
}

// Delete deletes the user from the databse according to the ID specified in the
// url and gives back a JSON response
func (UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Get the variables from the url
	vars := mux.Vars(r)

	var user model.User
	err := db.Delete(&user, "id = ?", vars["id"]).Error
	if err != nil {
		util.ErrorResponse(w, "Unable to delete the user")
		return
	}

	util.OKResponse(w, "deleted", true)
}

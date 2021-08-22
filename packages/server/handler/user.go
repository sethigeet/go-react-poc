package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/sethigeet/go-react-poc/packages/server/database"
	"github.com/sethigeet/go-react-poc/packages/server/model"
	"github.com/sethigeet/go-react-poc/packages/server/util"
)

type UserHandler struct {
	r *mux.Router
}

var db *gorm.DB

func (userHandler UserHandler) Apply() {
	db = database.DB

	subRouter := userHandler.r.PathPrefix("/user").Subrouter()

	// TODO: Accept other methods here too!
	subRouter.HandleFunc("/", userHandler.GetAll).Methods("GET")
	//                         ------------------------ Only accept valid uuids in the url ------------------------
	subRouter.HandleFunc("/{id:[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}}", userHandler.Get).Methods("GET")
}

func (UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		util.ErrorResponse(w, "Unable to fetch the users")
		return
	}

	util.OKResponse(w, "users", users)
}

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

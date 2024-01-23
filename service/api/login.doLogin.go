package api

import (
	"encoding/json"
	"net/http"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var userName UserName

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the username is correct
	userName.UserName = user.UserName
	if !userName.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create the user in the database
	newUser, err := rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("doLogin/db.CreateUser: executing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.FromDatabase(newUser)

	// give back the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

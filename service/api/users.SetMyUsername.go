package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var newUsername UserName

	// check parameters
	userId, err := strconv.ParseUint(ps.ByName("UserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate the user
	requestUserId, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization requestUserId", http.StatusBadRequest)
		return
	}

	if requestUserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// obtain new username from the body and store it in a variable
	err = json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create structure and check it has the correct length
	user.UserId = userId
	user.UserName = newUsername.UserName

	if !newUsername.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the username with the database function
	err = rt.db.ChangeUsername(user.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("setMyUserName/db.ChangeUsername: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

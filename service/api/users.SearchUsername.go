package api

import (
	"encoding/json"
	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

func (rt *_router) searchUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate correct user
	requestUserId, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization requestUserId", http.StatusBadRequest)
		return
	}

	// fill structure and do the action
	requestUser.UserId = requestUserId

	foundUser, err := rt.db.GetUserId(user.ToDatabase(), requestUser.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("searchUsername/db.GetUserId: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.FromDatabase(foundUser)

	// give back the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

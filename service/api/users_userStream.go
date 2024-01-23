package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/MarinaGRey/WASA/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photos []database.Photo

	// check parameters
	userId, err := strconv.ParseUint(ps.ByName("UserId"), 10, 64)
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

	if requestUserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// fill structure and do the action
	user.UserId = userId

	photos, err = rt.db.GetStream(user.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getMyStream/db.GetStream: executing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the stream
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)
}

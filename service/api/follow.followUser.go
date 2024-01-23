package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follower User
	var followed User

	// check parameters
	userId, err := strconv.ParseUint(ps.ByName("UserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	FollowUserId, err := strconv.ParseUint(ps.ByName("FollowUserId"), 10, 64)
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

	// user cannot follow itself
	if requestUserId == FollowUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// necessary structure
	follower.UserId = userId
	followed.UserId = FollowUserId

	// cannot follow a banned user
	isBanned, err := rt.db.UserBanned(follower.ToDatabase(), followed.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("followUser/db.UserBanned: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// the follow is done
	err = rt.db.FollowUser(follower.ToDatabase(), followed.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("followUser/db.FollowUser: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

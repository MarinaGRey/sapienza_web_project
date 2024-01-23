package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var banner User
	var banned User

	// the parameters are cheked
	userId, err := strconv.ParseUint(ps.ByName("UserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banUserId, err := strconv.ParseUint(ps.ByName("BanUserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// the user is who it should be
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

	// user cannot ban itself
	if userId == banUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banner.UserId = userId
	banned.UserId = banUserId

	// a user cannot be banned twice
	isBanned, err := rt.db.UserBanned(banner.ToDatabase(), banned.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.UserBanned: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// user is banned
	err = rt.db.BanUser(banner.ToDatabase(), banned.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.BanUser: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// the follows are erased (both parts)
	err = rt.db.UnfollowUser(banner.ToDatabase(), banned.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.UnfollowUser: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.UnfollowUser(banned.ToDatabase(), banner.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.UnfollowUser: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// erase likes from the photos from that user
	err = rt.db.RemoveLikes(banned.ToDatabase(), banner.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.RemoveLikes: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// erase comments from the photos from that user
	err = rt.db.RemoveComments(banned.ToDatabase(), banner.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("banUser/db.RemoveComments: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

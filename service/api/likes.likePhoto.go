package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	var photo PhotoId
	var user User

	// check parameters
	userId, err := strconv.ParseUint(ps.ByName("UserId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoId, err := strconv.ParseUint(ps.ByName("PhotoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate correct user
	requestUserId, err := strconv.ParseUint(strings.
		Split(r.Header.Get("Authorization"), " ")[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid authorization requestUserID", http.StatusBadRequest)
		return
	}

	if requestUserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// owner of the photo
	photo.PhotoId = photoId

	userPhoto, err := rt.db.GetUserPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("likePhoto/db.GetUserPhoto: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// cannot comment on a photo of a banned user
	user.UserId = userId

	isBanned, err := rt.db.UserBanned(userPhoto, user.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("likePhoto/db.UserBanned: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// structure for the database and call the action
	like.PhotoId = photoId
	like.UserId = userId

	newLike, err := rt.db.LikePhoto(like.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("likePhoto/db.LikePhoto: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	like.FromDatabase((newLike))

	// give back the like
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(like)
}

package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var commentText CommentText
	var comment Comment
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

	// obtain the text and check the length
	err = json.NewDecoder(r.Body).Decode(&commentText)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(commentText.Comment) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// this is the owner of the photo
	photo.PhotoId = photoId

	userPhoto, err := rt.db.GetUserPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("commentPhoto/db.GetUserPhoto: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// cannot comment if is banned
	user.UserId = userId

	isBanned, err := rt.db.UserBanned(userPhoto, user.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("commentPhoto/db.UserBanned: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// structure for the database
	comment.PhotoId = photoId
	comment.UserId = userId
	comment.Comment = commentText.Comment

	newComment, err := rt.db.CommentPhoto(comment.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("commentPhoto/db.CommentPhoto: executing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(newComment)

	// give back the comment in the body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var photo PhotoId
	var newPhoto Photo

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
		http.Error(w, "Invalid authorization requestUserId", http.StatusBadRequest)
		return
	}

	if requestUserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photo.PhotoId = photoId

	getPhoto, err := rt.db.GetPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getPhoto/db.GetPhoto: executing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPhoto.FromDatabase(getPhoto)

	// give back the photo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(newPhoto)
}

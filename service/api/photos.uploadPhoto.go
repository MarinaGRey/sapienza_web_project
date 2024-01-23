package api

import (
	"encoding/json"
	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo

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

	// structure for the database and do the action
	photo.UserId = userId

	photo.File, err = io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPhoto, err := rt.db.UploadPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("uploadPhoto/db.UploadPhoto: executing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(newPhoto)

	// give back the photo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
}

package api

import (
	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo PhotoId
	var file []byte

	// check parameters
	photoId, err := strconv.ParseUint(ps.ByName("PhotoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// action executed on database
	photo.PhotoId = photoId

	file, err = rt.db.GetPhotoFile(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getFile/db.GetPhotoFile: executing error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the file
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

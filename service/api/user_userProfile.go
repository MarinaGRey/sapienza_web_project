package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarinaGRey/WASA/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var requestUser User
	var profileUser User
	var profile Profile

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

	// structures needed
	requestUser.UserId = requestUserId
	profileUser.UserId = userId

	// cannot see profile if the user is banned
	banned, err := rt.db.UserBanned(requestUser.ToDatabase(), profileUser.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// if it is banned, it gives a partial content answer
	banned, err = rt.db.UserBanned(profileUser.ToDatabase(), requestUser.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	// name of the user we want to see
	profileUserResult, err := rt.db.GetUserName(profileUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// rest of the data of the profile
	followers, err := rt.db.GetFollowers(profileUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	following, err := rt.db.GetFollowing(profileUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photos, err := rt.db.GetPhotos(profileUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fill these data
	profile.UserId = profileUserResult.UserId
	profile.UserName = profileUserResult.UserName
	profile.Followers = followers
	profile.Following = following
	profile.Photos = photos

	// give back the profile
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)
}

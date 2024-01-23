package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
// Every endpoint that will be later used in the frontend is here
func (rt *_router) Handler() http.Handler {

	// Login enpoint
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// User Endpoint
	rt.router.PUT("/users/:UserId", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:UserId", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users", rt.wrap(rt.searchUsername))

	// Ban endpoint
	rt.router.PUT("/users/:UserId/banned/:BanUserId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:UserId/banned/:BanUserId", rt.wrap(rt.unbanUser))

	// Followers endpoint
	rt.router.PUT("/users/:UserId/followers/:FollowUserId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:UserId/followers/:FollowUserId", rt.wrap(rt.unfollowUser))

	// Stream endpoint
	rt.router.GET("/users/:UserId/homescreen", rt.wrap(rt.getMyStream))

	// Photo Endpoint
	rt.router.POST("/users/:UserId/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:UserId/photos/:PhotoId", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:UserId/photos/:PhotoId", rt.wrap(rt.getPhoto))
	rt.router.GET("/users/:UserId/photos/:PhotoId/file", rt.wrap(rt.getFile))

	// Comments endpoint
	rt.router.POST("/users/:UserId/photos/:PhotoId/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:UserId/photos/:PhotoId/comments/:CommentId", rt.wrap(rt.uncommentPhoto))

	// Likes endpoint
	rt.router.POST("/users/:UserId/photos/:PhotoId/likes", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:UserId/photos/:PhotoId/likes", rt.wrap(rt.unlikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

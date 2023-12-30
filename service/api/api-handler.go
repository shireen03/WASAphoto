package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	// rt.router.PUT("/user/:username/username", rt.wrap(rt.setMyUsername))
	// rt.router.POST("/user/:userID/follow/:followUserID", rt.wrap(rt.followUser))
	rt.router.GET("/user/:userID/followers", rt.wrap(rt.getFollowers))
	rt.router.GET("/user/:userID/followings", rt.wrap(rt.getFollowings))

	rt.router.GET("/username/:username/follow/:followUsername", rt.wrap(rt.isFollowing))

	rt.router.GET("/username/:username/profile", rt.wrap(rt.getMyProfile))

	// rt.router.GET("/user/:userID/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/profile/:username", rt.wrap(rt.getMyProfile))

	// rt.router.POST("/user/:userID/photo", rt.wrap(rt.uploadPhoto))
	// rt.router.POST("/user/:userID/photo/:photoID/remove", rt.wrap(rt.deletePhoto))

	// rt.router.POST("/user/:userID/photo/:photoID/like", rt.wrap(rt.likePhoto))

	// rt.router.POST("/user/:userID/photo/:photoID/comment", rt.wrap(rt.commentPhoto))
	// rt.router.DELETE("/user/:userID/photo/:photoID/comment", rt.wrap(rt.uncommentPhoto))

	// rt.router.POST("/user/:userID/ban/:banUserID", rt.wrap(rt.banUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

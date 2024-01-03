package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/user/:userID/setusername/:username", rt.wrap(rt.setMyUsername))

	rt.router.POST("/username/:username/follow/:followUsername", rt.wrap(rt.followUser))
	rt.router.DELETE("/username/:username/follow/:followUsername", rt.wrap(rt.unfollowUser))

	rt.router.GET("/user/:userID/followers", rt.wrap(rt.getFollowers))
	rt.router.GET("/user/:userID/followings", rt.wrap(rt.getFollowings))

	rt.router.GET("/username/:username/follow/:followUsername", rt.wrap(rt.isFollowing))

	rt.router.GET("/username/:username/profile", rt.wrap(rt.getMyProfile))

	// rt.router.GET("/user/:userID/stream", rt.wrap(rt.getMyStream))
	//rt.router.GET("/profile/:username", rt.wrap(rt.getMyProfile))

	rt.router.POST("/photo/upload/:userID", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/photo/upload/:userID", rt.wrap(rt.getPhotos))

	// rt.router.POST("/user/:userID/photo/:photoID/remove", rt.wrap(rt.deletePhoto))

	rt.router.POST("/user/:userID/photo/:photoID/like", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/user/:userID/photo/:photoID/like", rt.wrap(rt.unlikePhoto))

	rt.router.POST("/user/:userID/photo/:photoID/comment", rt.wrap(rt.commentPhoto))
	// rt.router.DELETE("/user/:userID/photo/:photoID/comment", rt.wrap(rt.uncommentPhoto))

	// rt.router.POST("/user/:userID/ban/:banUserID", rt.wrap(rt.banUser))
	rt.router.POST("/username/:username/ban/:banUsername", rt.wrap(rt.banUser))
	rt.router.DELETE("/username/:username/ban/:banUsername", rt.wrap(rt.unbanUser))
	rt.router.GET("/username/:username/ban/:banUsername", rt.wrap(rt.isBan))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

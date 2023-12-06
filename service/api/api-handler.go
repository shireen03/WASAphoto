package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// rt.router.POST("/user/:username/doLogin", rt.wrap(rt.LogUser))
	// rt.router.PUT("/user/:userID/username/:username", rt.wrap(rt.setUsername))
	// rt.router.POST("/user/:userID/follow/:followUserID", rt.wrap(rt.followUser))
	// rt.router.GET("/user/:userID/stream", rt.wrap(rt.Stream))

	// rt.router.POST("/photo/:photoID/upload/:userID", rt.wrap(rt.postPhoto))
	// rt.router.POST("/photo/:photoID/remove", rt.wrap(rt.deletePhoto))

	// rt.router.POST("/photo/:photoID/like/:userID", rt.wrap(rt.like))

	// rt.router.POST("/photo/:photoID/comment/:userID", rt.wrap(rt.doComment))
	// rt.router.DELETE("/photo/:photoID/comment/:userID", rt.wrap(rt.doComment))

	// rt.router.POST("/user/:userID/ban/:banUserID", rt.wrap(rt.banUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

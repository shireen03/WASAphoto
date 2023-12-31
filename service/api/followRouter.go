package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	username := ps.ByName("username")

	followUsername := ps.ByName("followUsername")

	var follow database.Follow
	follow.UserID = username
	follow.FollowedID = followUsername

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}
	if !isFollow {
		_ = rt.db.SetFollow(follow)
		message := "User followed successfully"
		w.Write([]byte(message))
	}

}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	username := ps.ByName("username")

	followUsername := ps.ByName("followUsername")

	var follow database.Follow
	follow.UserID = username
	follow.FollowedID = followUsername

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}

	if isFollow {
		_ = rt.db.RemoveFollow(follow)
		message := "User unfollowed successfully"
		w.Write([]byte(message))
	}

}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userID := ps.ByName("userID")

	followers, err := rt.db.GetFollowers(userID)
	if err != nil {
		http.Error(w, "cant get list of followers", http.StatusBadRequest)
		return
	}

	message := "User followers retrieved"
	w.Write([]byte(message))
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(followers)

}

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userID := ps.ByName("userID")

	followers, err := rt.db.GetFollowings(userID)
	if err != nil {
		return
	}

	message := "User followings retrieved"
	w.Write([]byte(message))
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(followers)

}

func (rt *_router) isFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	username := ps.ByName("username")

	followUsername := ps.ByName("followUsername")

	var follow database.Follow
	follow.UserID = username
	follow.FollowedID = followUsername

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(isFollow)

}

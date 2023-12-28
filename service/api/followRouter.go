package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userID := ps.ByName("userID")

	followUserID := ps.ByName("followUserID")

	var follow database.Follow
	follow.UserID = userID
	follow.FollowedID = followUserID

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}
	if isFollow {
		err = rt.db.RemoveFollow(follow) //if already following then we unfollow
		message := "User unfollowed successfully"
		w.Write([]byte(message))

	} else {
		err = rt.db.SetFollow(follow)
		message := "User followed successfully"
		w.Write([]byte(message))
	}

}

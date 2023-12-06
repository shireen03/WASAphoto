package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64) //converting string to integer (uint64)
	if err != nil {
		return
	}
	followUserID, err := strconv.ParseUint(ps.ByName("followUserID"), 10, 64)

	var follow database.Follow
	follow.UserID = userID
	follow.FollowedID = followUserID

	isFollow, err := rt.db.IsFollowing(follow)

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

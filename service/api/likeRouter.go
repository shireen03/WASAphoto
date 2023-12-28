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

func (rt *_router) like(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")
	userID := ps.ByName("userID") //converting string to integer (uint64)

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		return
	}
	var like database.Like
	like.UserID = userID
	like.PhotoID = photoID

	isLiked, err := rt.db.IsLiked(like)

	if isLiked {
		err = rt.db.RemoveLike(like) //if already following then we unfollow
		message := "User unliked the photo"
		w.Write([]byte(message))

	} else {
		err = rt.db.SetLike(like)
		message := "User followed liked the photo"
		w.Write([]byte(message))
	}

}

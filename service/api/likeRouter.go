package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	if err != nil {
		http.Error(w, "couldnt check if liked", http.StatusBadRequest)
		return
	}

	if !isLiked {
		liker, err := rt.db.SetLike(like) //if already following then we unfollow
		if err != nil {
			http.Error(w, "couldnt check if its liked", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(liker)

	}

}
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")
	var like database.Like
	userID := ps.ByName("userID") //converting string to integer (uint64)

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		return
	}

	like.UserID = userID
	like.PhotoID = photoID

	isLiked, err := rt.db.IsLiked(like)
	if err != nil {
		http.Error(w, "couldnt check if its liked", http.StatusBadRequest)
		return
	}

	if isLiked {
		err = rt.db.RemoveLike(like) //if already following then we unfollow
		if err != nil {
			http.Error(w, "remove like failed", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)

	}

}

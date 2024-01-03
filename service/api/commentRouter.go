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

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	userID := ps.ByName("userID")

	var comment database.Comment
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Error decoding", http.StatusInternalServerError)
		return
	}

	comment.UserID = userID

	comment.PhotoID = photoID

	username, err := rt.db.GetUsernameWithUserID(comment.UserID)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)
		return
	}
	comment.Username = username

	commenter, err := rt.db.SetComment(comment)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(commenter)

}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	userID := ps.ByName("userID")

	var comment database.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	comment.UserID = userID
	comment.PhotoID = photoID

	err = rt.db.RemoveComment(comment)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{Message: "Comment removed successfully"})
}

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

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid account creation", http.StatusBadRequest)
		return
	}

	_, err = rt.db.LogUser(user)

	if err != nil {
		http.Error(w, "Invalid account creation", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (rt *_router) getMyProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var user database.User
	user.UserID = ps.ByName("userID")
	user.Username = ps.ByName("username")

	_, err := rt.db.GetProfile(user)

	if err != nil {
		http.Error(w, "Invalid account creation", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	w.Header().Set("content-type", "application/json")
// 	var user database.User

// 	UserID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid userID", http.StatusBadRequest)
// 		return
// 	}
// 	user.UserID = UserID

// 	username := ps.ByName("userID")

// 	newUser, err := rt.db.SetUsername(username, user)

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(newUser)
// }

// func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	w.Header().Set("content-type", "application/json")
// 	var user database.User
// 	var stream database.Streamer
// 	UserID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid userID", http.StatusBadRequest)
// 		return
// 	}
// 	stream.UserID = UserID
// 	user.UserID = UserID

// 	streamPics, err := rt.db.GetStream(user)

// 	stream.StreamedPhotos = streamPics

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(stream)
// }

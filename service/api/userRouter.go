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

	newuser, err := rt.db.LogtheUser(user)

	if err != nil {
		http.Error(w, "Invalid account creationssdkslfjòlskadjfljsadòlfss", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newuser)
	if err != nil {
		http.Error(w, "Invalid account creatioijouiuounssss", http.StatusConflict)
		return
	}

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var profile database.Profile

	profile.Username = ps.ByName("username")
	checkuser, err := rt.db.CheckUserExist(profile.Username)
	if err != nil {
		http.Error(w, "error checking whether username is in db", http.StatusBadRequest)
		return
	}
	if !checkuser {
		http.Error(w, "user with username doesnt exist", http.StatusBadRequest)
		return

	}

	profileuserID, err := rt.db.GetUserIDWithUsername(profile.Username)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)
		return
	}
	profile.UserID = profileuserID

	photocount, err := rt.db.GetPhotoCount(profile.UserID)
	if err != nil {
		http.Error(w, "cant get photo count", http.StatusBadRequest)
		return
	}
	profile.PicNumb = photocount

	followCount, err := rt.db.GetFollowCount(profile.UserID)
	if err != nil {
		http.Error(w, "cant get follow count", http.StatusBadRequest)
		return
	}
	profile.FollowedCount = followCount

	followingCount, err := rt.db.GetFollowingCount(profile.UserID)
	if err != nil {
		http.Error(w, "cant get following count", http.StatusBadRequest)
		return
	}
	profile.FollowingCount = followingCount

	bancount, err := rt.db.GetBanCount(profile.UserID)
	if err != nil {
		http.Error(w, "cant get ban count", http.StatusBadRequest)
		return
	}
	profile.BanCount = bancount

	bans, err := rt.db.GetBans(profile.UserID)
	if err != nil {
		http.Error(w, "cant get ban count", http.StatusBadRequest)
		return
	}
	profile.Bans = bans

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(profile)
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var user database.User

	userID := ps.ByName("userID")

	user.UserID = userID

	username := ps.ByName("username")

	newUser, err := rt.db.SetUsername(username, user)
	if err != nil {
		http.Error(w, "Invalid username update", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var user database.User
	var stream database.Streamer
	userID := ps.ByName("userID")

	stream.UserID = userID
	user.UserID = userID

	streamPics, err := rt.db.GetStream(user)
	if err != nil {
		http.Error(w, "Invalid stream", http.StatusBadRequest)
		return
	}

	stream.StreamedPhotos = streamPics

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(streamPics)
}

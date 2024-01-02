package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	username := ps.ByName("username")
	userID, err := rt.db.GetUserIDWithUsername(username)
	if err != nil {
		return
	}
	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		return
	}

	if !isbanned {
		err = rt.db.SetBan(ban) //if already banned then we unban
		if err != nil {
			http.Error(w, "mffff helpmeeee", http.StatusBadRequest)
			return
		}

		message := "User banned successfully"
		w.Write([]byte(message))

	}
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	username := ps.ByName("username")
	userID, err := rt.db.GetUserIDWithUsername(username)
	if err != nil {
		return
	}
	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		return
	}

	if isbanned {
		err = rt.db.SetUnBan(ban) //if already banned then we unban
		if err != nil {
			http.Error(w, "mffff whyyy", http.StatusBadRequest)
			return
		}

		message := "User unbanned successfully"
		w.Write([]byte(message))

	}
}

func (rt *_router) isBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	username := ps.ByName("username")
	userID, err := rt.db.GetUserIDWithUsername(username)
	if err != nil {
		return
	}
	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		http.Error(w, "mffff helpmeeee", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(isbanned)

}

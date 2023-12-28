package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	userID := ps.ByName("userID")

	banUserID, err := strconv.ParseUint(ps.ByName("banUserID"), 10, 64)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)

	if isbanned {
		err = rt.db.SetUnBan(ban) //if already banned then we unban
		message := "User unbanned successfully"
		w.Write([]byte(message))

	} else {
		err = rt.db.SetBan(ban)
		message := "User banned successfully"
		w.Write([]byte(message))
	}

}

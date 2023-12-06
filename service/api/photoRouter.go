package api

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/shireen03/WASAphoto/service/api/reqcontext"
	"github.com/shireen03/WASAphoto/service/database"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).

func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var pic database.Photo

	userID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	pic.Picture, err = io.ReadAll(r.Body) //reads image file from requestbody
	if err != nil {
		http.Error(w, "Invalid photo file", http.StatusBadRequest)
		return

	}
	pic.UserID = userID
	pic.UploadTime = time.Now()
	pic.Date = pic.UploadTime.Format("2006-01-02 15:04:05")

	err = rt.db.SetPic(pic)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var pic database.Photo

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(ps.ByName("userID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	pic.UserID = userID
	pic.PhotoID = photoID

	err = rt.db.RemovePic(pic)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package httphandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	"github.com/juhonamnam/wedding-invitation-server/types"
)

type GuestbookHandler struct {
	http.Handler
}

func (h *GuestbookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		offsetQ := r.URL.Query().Get("offset")
		limitQ := r.URL.Query().Get("limit")

		offset, err := strconv.Atoi(offsetQ)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		limit, err := strconv.Atoi(limitQ)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		guestbook, err := sqldb.GetGuestbook(offset, limit)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		pbytes, err := json.Marshal(guestbook)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pbytes)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var post types.GuestbookPostForCreate
		err := decoder.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		err = sqldb.CreateGuestbookPost(post.Name, post.Content, post.Password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var post types.GuestbookPostForDelete
		err := decoder.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		err = sqldb.DeleteGuestbookPost(post.Id, post.Password)

		if err != nil {
			if err.Error() == "INCORRECT_PASSWORD" {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Forbidden"))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("InternalServerError"))
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

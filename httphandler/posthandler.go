package httphandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	"github.com/juhonamnam/wedding-invitation-server/types"
)

type PostHandler struct {
	http.Handler
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

		posts, err := sqldb.GetPosts(offset, limit)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		pbytes, err := json.Marshal(posts)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pbytes)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var post types.PostCreate
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

		err = sqldb.CreatePost(post.Name, post.Content, post.Password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var post types.PostDelete
		err := decoder.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		err = sqldb.DeletePost(post.Id, post.Password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

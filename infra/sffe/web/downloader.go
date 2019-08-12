package web

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"

	"geegle.org/infra/sffe/context"
	"github.com/syndtr/goleveldb/leveldb"
)

func ServeFile(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[len("/s/"):]
	f, err := ctx.Get(p)
	if err == leveldb.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parts := strings.Split(p, "/")
	for i := len(parts) - 2; i > 0; i-- {
		j := strings.Index(parts[i], "=")
		if j > -1 {
			if parts[i][:j] == "i" && parts[i][j+1:] == "1" { // service internal
				tknStr := r.Header.Get("X-Geegle-JWT")
				service, err := getJwtServiceName(tknStr, []byte("superSecretJWTKEY"))
				if err != nil || parts[0] != fmt.Sprintf("%x", md5.Sum([]byte(service))) {
					writeJSONError(w, "Invalid JWT", http.StatusUnauthorized)
					return
				}
			}
			// add more option requirements here
			if parts[i][:j] == "s" && parts[i][j+1:] == "1" { // signature required
				// TODO (adamyi@): implement this
			}
		}
	}

	w.Write(f.Content)
}

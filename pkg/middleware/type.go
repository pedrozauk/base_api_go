package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func MiddleContentType(next func(http.ResponseWriter, *http.Request), acceptCt []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ct := r.Header.Get("Content-Type")
		if ct != "" {
			for _, v := range acceptCt {
				fmt.Println(v)
				mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
				if mediaType == v {
					next(w, r)
					return
				}
			}
			msg := fmt.Sprintf("Content-Type header is not %s", acceptCt)
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return

		}
		msg := fmt.Sprintf("Content-Type header is not %s", acceptCt)
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}
}

package handler

import (
	"allfather/pkg/httputil"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	httputil.Success(w, http.StatusAccepted, struct{ ServerStatus string }{ServerStatus: "healthy"}, "success")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}

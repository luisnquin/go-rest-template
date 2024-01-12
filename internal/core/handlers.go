package core

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/api"
	"github.com/luisnquin/server-example/internal/server"
)

func (m Manager) PingHandler() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		api.Response(w, http.StatusOK, api.StdResponse{
			Message: "pong",
		})
	}
}

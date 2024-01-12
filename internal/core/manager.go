package core

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/server"
)

type Manager struct{}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) RegisterHandlers(srv server.Registerer) {
	srv.RegisterHandler("/ping", http.MethodGet, m.PingHandler(), false, false)
}

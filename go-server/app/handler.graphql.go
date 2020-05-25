package app

import (
	"net/http"

	"go-server/controller"
)

func (a *App) ExecQueryHandler(w http.ResponseWriter, r *http.Request) {
	controller.ExecQuery(a.DB, w, r)
}

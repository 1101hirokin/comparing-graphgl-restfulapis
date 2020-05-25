package app

import (
	"net/http"

	"go-server/controller"
)

func (a *App) GetUsersIndexHandler(w http.ResponseWriter, r *http.Request) {
	controller.GetUsersIndex(a.DB, w, r)
}

func (a *App) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	controller.GetUser(a.DB, w, r)
}
func (a *App) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	controller.PostUser(a.DB, w, r)
}
func (a *App) PutUserHandler(w http.ResponseWriter, r *http.Request) {
	controller.PutUser(a.DB, w, r)
}
func (a *App) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	controller.DeleteUser(a.DB, w, r)
}

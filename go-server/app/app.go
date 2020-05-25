package app

import (
	"fmt"
	"log"
	"net/http"

	"go-server/config"
	"go-server/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Init(config *config.Config) {
	connectUri := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Protocol,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, connectUri)
	if err != nil {
		panic(err.Error())
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.SetRouters()
	a.SetInitialData()
}

func (a *App) SetRouters() {
	// Restful API Endpoint
	a.Get("/api/users", a.GetUsersIndexHandler)
	a.Get("/api/users/{id}", a.GetUserHandler)
	a.Post("/api/users", a.PostUserHandler)
	a.Put("/api/users/{id}", a.PutUserHandler)
	a.Delete("/api/users/{id}", a.DeleteUserHandler)

	// GraphQL Endpoint
	a.Post("/graphql", a.ExecQueryHandler)
}

// Wrapper Method of "POST, PUT, GET, DELETE"
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) SetInitialData() {
	initUser := model.User{
		Name:      "中谷 仁貴",
		Email:     "email@example.com",
		Bio:       "こんにちは！中谷 仁貴（ナカタニヒロキ）です！",
		UrlAvatar: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
	}

	a.DB.Create(&initUser)
}

// Run the Server
func (a *App) Run(host string) {
	log.Println("Serving Start! Access http://localhost" + host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

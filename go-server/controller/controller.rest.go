package controller

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"go-server/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetUsersIndex(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	status := http.StatusOK

	users := []model.User{}

	payload := model.UsersResponse{}
	payload.Success = true
	payload.Message = "Succeeded! you have got users data"

	offset := 0
	if r.FormValue("offset") != "" {
		offsetInt, err := strconv.Atoi(r.FormValue("offset"))
		if err == nil {
			offset = offsetInt
		}
	}
	limit := 10
	if r.FormValue("offset") != "" {
		limitInt, err := strconv.Atoi(r.FormValue("offset"))
		if err == nil {
			limit = limitInt
		}
	}

	db.Limit(limit).Offset(offset).Find(&users)
	payload.Users = &users
	if len(users) <= 0 {
		payload.Message = "Succeeded! But, There was no user data ..."
	}

	pagination := model.Pagination{
		Offset: offset,
		Limit:  limit,
	}
	payload.Pagination = &pagination

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	status := http.StatusOK

	vars := mux.Vars(r)
	user := model.User{}

	payload := model.UserResponse{}
	payload.Success = true
	payload.Message = "Succeeded! you have got a user data"

	user_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		payload.Success = false
		payload.Message = "Failed! the ID of Request URI was strange type!"
		status = http.StatusBadRequest
	} else if db.First(&user, user_id).RecordNotFound() {
		payload.Success = false
		payload.Message = "Failed! the user was Not Found !"
		status = http.StatusNotFound
	} else {
		payload.User = &user
	}

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}

func PostUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	status := http.StatusOK

	payload := model.UserResponse{}
	payload.Success = true
	payload.Message = "Succeeded! you have created a user data"

	err := r.ParseForm()
	if err != nil {
		payload.Success = false
		payload.Message = "Failed! some error has occured!"
		status = http.StatusBadRequest
	}

	newUser := model.User{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		Bio:       r.FormValue("bio"),
		UrlAvatar: r.FormValue("url_avatar"),
	}

	db.Create(&newUser)
	payload.User = &newUser

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}

func PutUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	status := http.StatusOK

	vars := mux.Vars(r)
	user := model.User{}

	payload := model.UserResponse{}
	payload.Success = true
	payload.Message = "Succeeded! you have created a user data"

	err := r.ParseForm()
	if err != nil {
		payload.Success = false
		payload.Message = "Failed! some error has occured!"
		status = http.StatusBadRequest
	}

	user_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		payload.Success = false
		payload.Message = "Failed! the ID of Request URI was strange type!"
		status = http.StatusBadRequest
	} else if db.First(&user, user_id).RecordNotFound() {
		payload.Success = false
		payload.Message = "Failed! the user was Not Found !"
		status = http.StatusNotFound
	} else {
		user.Name = r.FormValue("name")
		user.Email = r.FormValue("email")
		user.Bio = r.FormValue("bio")
		user.UrlAvatar = r.FormValue("url_avatar")

		db.Save(&user)
		payload.User = &user
	}

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	status := http.StatusOK

	vars := mux.Vars(r)
	user := model.User{}

	payload := model.UserResponse{}
	payload.Success = true
	payload.Message = "Succeeded! you have created a user data"

	user_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		payload.Success = false
		payload.Message = "Failed! the ID of Request URI was strange type!"
		status = http.StatusBadRequest
	} else if db.First(&user, user_id).RecordNotFound() {
		payload.Success = false
		payload.Message = "Failed! the user was Not Found !"
		status = http.StatusNotFound
	} else {
		db.Delete(&user)
		payload.User = &user
	}

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emmaperez2197/api/db"
	"github.com/emmaperez2197/api/routes/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user no fount"))
		return
	}

	json.NewEncoder(w).Encode(&user)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	existingUser := models.User{}

	var error = db.DB.Where("email = ?", user.Email).First(&existingUser).Error

	if error == nil {
		// Si encontramos un usuario con el mismo email, devolvemos un error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The email already exists"))
		return
	}

	createUser := db.DB.Create(&user)

	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("UPDATE"))

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user no fount"))
		return
	}

	db.DB.Unscoped().Delete(&user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))

}

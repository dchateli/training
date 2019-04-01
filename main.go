package main

import (
	"github.com/dchateli/training/davidDb"
	"github.com/dchateli/training/davidDb/filesystem"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	myDb db.DbContract
)

func main() {
	// COUCOU
	myDb = &filesystem.InMemoryDb{}
	// modif
	router := mux.NewRouter()
	router.HandleFunc("/users", listUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", createUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPatch)

	if err := http.ListenAndServe("0.0.0.0:8888", router); err != nil {
		log.Fatal(err.Error())
	}

}

func listUsers(w http.ResponseWriter, r *http.Request) {

	listUser, err := myDb.ListUser()
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pagesJson, _ := json.MarshalIndent(listUser, "", "\t")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(pagesJson)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.User

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err = myDb.AddUser(newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		_, e := w.Write([]byte(err.Error())) //cast
		if e != nil {
			fmt.Println(err)
		}

		return
	}

	listByte, err := json.MarshalIndent(newUser, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(listByte)
}

/*func UserExist(userId string) (bool, db.User, int) {

	for i := range db.UserDB {
		if userId == db.UserDB[i].Id {
			return true, db.UserDB[i], i
		}

	}
	return false, db.User{}, -1
}*/

func getUser(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	userId := requestVars["id"]


	newUser, err := myDb.GetUser(userId)

	/*isExist, myUser, _ := UserExist(userId)
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

	b, err := json.MarshalIndent(newUser, "", "\t")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		// handle error
		fmt.Println(err)

	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	/*isExist, _, userListNb := UserExist(userId)
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.UserDB = append(db.UserDB[:userListNb], db.UserDB[userListNb+1:]...)
	*/
	err := myDb.DeleteUser(userId)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	return
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.User

	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	UpdatedUser,err:= myDb.UpdateUser(userId, newUser )
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listByte, err := json.MarshalIndent(UpdatedUser, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(listByte)

}

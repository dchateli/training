package training

import (

	"davidDb"
	"davidDb/filesystem"

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
	myDb = filesystem.InMemoryDb{}
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

func listUsers(w http.ResponseWriter, r *http.Request ) {

	listeUser, err := myDb.ListUser()
	if err != nil{
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pagesJson, _ := json.MarshalIndent(listeUser, "", "\t")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(pagesJson)
	if err != nil {
		fmt.Println(err.Error())

	}
}

func createUser(w http.ResponseWriter , r *http.Request){
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
		if e!= nil {
			fmt.Println(err)
		}

		return
	}

	listByte, err:= json.MarshalIndent(newUser, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(listByte)
}


func UserExist(userId string )(bool,db.User,int){

	for i := range db.UserDB {
		if 	userId == db.UserDB[i].Id{
			return true,db.UserDB[i],i
		}

	}
	return false,db.User{},-1
}

func getUser(w http.ResponseWriter , r *http.Request){
	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	isExist,myUser,_ := UserExist(userId)
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}



	b, err:= json.MarshalIndent(myUser, "", "\t")
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	w.WriteHeader(http.StatusOK)

	_ ,err= w.Write(b)
	if err != nil {
		// handle error
		fmt.Println(err)

	}
}

func deleteUser(w http.ResponseWriter , r *http.Request){
	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	isExist,_,userListNb := UserExist(userId)
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.UserDB = append( db.UserDB[:userListNb], db.UserDB[userListNb+1:]...)

	w.WriteHeader(http.StatusNoContent)

	return
}

func updateUser(w http.ResponseWriter , r *http.Request){
	var newUser db.User


	requestVars := mux.Vars(r)
	userId := requestVars["id"]

	isExist,myUser,userListNb := UserExist(userId)
	if !isExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}


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

	if newUser.Id != ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newUser.Name != ""{
		myUser.Name = newUser.Name
	}

	if newUser.Description != ""{
		myUser.Description = newUser.Description
	}


	db.UserDB[userListNb] = myUser


	listByte, err:= json.MarshalIndent(db.UserDB, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(listByte)

}
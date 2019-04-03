package inmemory

import (
	"errors"
	"github.com/dchateli/training/davidDb"
	"time"

	"github.com/satori/go.uuid"
)

type InMemoryDb struct {
	UserDB     []db.User
	lastUpdate time.Time

}

func (d *InMemoryDb) verification(user string) (bool, db.User, int) {

	for i := range d.UserDB {
		if user == d.UserDB[i].Id {
			return true, d.UserDB[i], i
		}
	}
	return false, db.User{}, -1
}

func (d *InMemoryDb) AddUser(u db.User) (db.User, error) {
	storage, err := uuid.NewV4()
	if err != nil {
		return db.User{}, err
	}
	u.Id = storage.String()

	d.UserDB = append(d.UserDB, u)
	return u, nil
}

func (d *InMemoryDb) DeleteUser(userId string) error {

	isExist, _, userListNb := d.verification(userId)
	if !isExist {
		return errors.New("User not found")
	}

	d.UserDB = append(d.UserDB[:userListNb], d.UserDB[userListNb+1:]...)
	return nil
}

func (d *InMemoryDb) UpdateUser(userId string,u db.User) (db.User, error) {
	isExist, myUser, userListNb := d.verification(userId)
	if  !isExist{
		return myUser,errors.New("User not found (UpdateUser request)")
	}

	if u.Name != "" {
		myUser.Name = u.Name
	}

	if u.Description != "" {
		myUser.Description = u.Description
	}

	 d.UserDB[userListNb] = myUser

	return myUser, nil
}

func (d *InMemoryDb) GetUser(userId string) (db.User, error) {

	isExist, myUser, _ := d.verification(userId)
	if !isExist {
		return myUser,errors.New("User not found (getUser Request)")
	}
	return myUser, nil
}

func (d *InMemoryDb) ListUser()([]db.User, error){

	return d.UserDB,nil
}

// Implementer les 3 autres méthodes du contrat

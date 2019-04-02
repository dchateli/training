package filesystem

import (
	"errors"
	"fmt"
	"github.com/dchateli/training/davidDb"
	"time"

	"github.com/satori/go.uuid"
)

type InMemoryDb struct {
	userDB     []db.User
	lastUpdate time.Time
}

func (d *InMemoryDb) verification(user string) (bool, db.User, int) {

	for i := range d.userDB {
		if user == d.userDB[i].Id {
			return true, d.userDB[i], i
		}
	}
	return false, db.User{}, -1
}

func (d *InMemoryDb) AddUser(u db.User) (db.User, error) {
	storage, err := uuid.NewV4()
	fmt.Println(d.userDB)

	if err != nil {
		return db.User{}, err
	}
	u.Id = storage.String()

	d.userDB = append(d.userDB, u)
	return u, nil
}

func (d *InMemoryDb) DeleteUser(userId string) error {

	isExist, _, userListNb := d.verification(userId)
	if !isExist {
		return errors.New("User not found")
	}

	d.userDB = append(d.userDB[:userListNb], d.userDB[userListNb+1:]...)
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

	 d.userDB[userListNb] = myUser

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


	return d.userDB,nil
}

func (d *InMemoryDb) DoubleUser(u db.User) (db.User,db.User, error) {
	storage, err := uuid.NewV4()
	fmt.Println(d.userDB)

	if err != nil {
		return db.User{}, db.User{}, err
	}
	u.Id = storage.String()

	d.userDB = append(d.userDB, u)
	return u,u, nil
}

// Implementer les 3 autres méthodes du contrat

package filesystem

import (
	"time"

	"airbus/datastore/common/uuid"
)

type InMemoryDb struct {
	userDB []db.User
	lastUpdate time.Time
}

func (d InMemoryDb) AddUser(u db.User) ( db.User,error) {
	u.Id = uuid.Generate()
	d.userDB = append(d.userDB, u)
	return u,nil
}

func (d InMemoryDb) DeleteUser(userId string) error {
	return nil
}

// Implementer les 3 autres méthodes du contrat
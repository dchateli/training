package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dchateli/training/davidDb"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

type MysqlDb struct {
	Con *sql.DB
}

func (m *MysqlDb) AddUser(u db.User) (db.User, error) {

	repId, err := m.Con.Exec(`INSERT INTO users(name,description) VALUES( ?, ? )`, u.Name, u.Description)
	if err != nil {
		return db.User{}, errors.Wrap(err, "failed to insert new user")
	}

	var id int64
	id,err = repId.LastInsertId()
	if err != nil {
		return db.User{}, err
	}

	row := m.Con.QueryRow("SELECT userid,name,description FROM users WHERE userid = ?",id )
	var returnedUser db.User
	if err := row.Scan(&returnedUser.Id, &returnedUser.Name, &returnedUser.Description); err != nil {
		return db.User{}, errors.Wrap(err, "failed to select user")
	}
	return returnedUser,err
}

func (m *MysqlDb) ListUser() ([]db.User, error) {
	// queryRows

	rows,err := m.Con.Query("SELECT * FROM users ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listUser []db.User

	for rows.Next() {
		var currentUser db.User

		err := rows.Scan(&currentUser.Id,&currentUser.Name,&currentUser.Description)
		if err != nil {
			log.Fatal(err)
		}
		listUser = append(listUser, currentUser)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return listUser, nil
}
func (m *MysqlDb) DeleteUser(userId string) error {
	deletRetrun, err := m.Con.Exec("DELETE FROM users WHERE userid = ?", userId)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	rowsAffected,err := deletRetrun.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to delete row user")
	}
	if rowsAffected > 0 {
		return nil
	} else {
		return errors.New("no user")
	}
}
func (m *MysqlDb) GetUser(userId string) (db.User, error) {

	// convert string to int64
	n, err := strconv.ParseInt(userId, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", n, n)
	}

	row := m.Con.QueryRow("SELECT userid,name,description FROM users WHERE userid = ?",userId )

	var returnedUser db.User
	if err := row.Scan(&returnedUser.Id, &returnedUser.Name, &returnedUser.Description); err != nil {
		return db.User{}, errors.Wrap(err, "failed to select user")
	}
	return returnedUser,err
}

func (m *MysqlDb) UpdateUser(userId string, u db.User) (db.User, error) {
	_, err := m.Con.Exec("UPDATE users SET name = ?, description = ? WHERE userid = ?",u.Name, u.Description,userId)
	if err != nil {
		return db.User{}, err
	}
	return m.GetUser(userId)
}

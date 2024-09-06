package mysqlrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/salehborhani/todo-list/entity"
)

func (d *MYSQlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8

	query := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := query.Scan(&user.ID, &user.UserName, &user.Password, &user.PhoneNumber, &createdAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, fmt.Errorf("can't scan the query: %w", err)
	}

	return false, nil
}

func (d *MYSQlDB) RepoRegister(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users (name, phone_number, password) values (?, ?, ?)`, u.UserName, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf(`unexpected error: %w`, err)

	}
	id, _ := res.LastInsertId()
	u.ID = uint8(id)

	return u, nil
}

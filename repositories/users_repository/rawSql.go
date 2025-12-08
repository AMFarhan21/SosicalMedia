package users_repository

import (
	"context"
	"database/sql"
	"errors"
	"socialmedia/domain"
)

type RawUsersRepository struct {
	DB  *sql.DB
	ctx context.Context
}

func NewRawUsersRepository(db *sql.DB) *RawUsersRepository {
	return &RawUsersRepository{
		DB:  db,
		ctx: context.Background(),
	}
}

func (r *RawUsersRepository) CreateUser(data domain.Users) (domain.Users, error) {
	var user domain.Users
	query := `
		insert into users (id, first_name, last_name, address, email, username, password, age, role) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning *;
	`

	err := r.DB.QueryRow(query, data.ID, data.FirstName, data.LastName, data.Address, data.Email, data.Username, data.Password, data.Age, data.Role).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Address,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Age,
		&user.Role,
	)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (r *RawUsersRepository) GetUserByID(id string) (domain.Users, error) {
	query := `
		select * from users where id = $1;
	`
	var user domain.Users
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Address, &user.Email, &user.Username, &user.Password, &user.Age, &user.Role)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (r *RawUsersRepository) GetAllUsers(page, limit int) ([]domain.Users, error) {
	query := `
		select * from users order by username offset ($1 - 1) * $2 limit $3;
	`

	row, err := r.DB.Query(query, page, limit, limit)
	if err != nil {
		return nil, err
	}

	var users []domain.Users
	for row.Next() {
		var user domain.Users

		err := row.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Address,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.Age,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *RawUsersRepository) UpdateUser(data domain.Users) error {
	query := `
		update users set id=$1 first_name=$2 last_name=$3 address=$4 email=$5 username=$6 password=$7 age=$8 role=$9 where id = $10;
	`

	row, err := r.DB.Exec(query, data.ID, data.FirstName, data.LastName, data.Address, data.Email, data.Username, data.Password, data.Age, data.Role, data.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *RawUsersRepository) DeleteUser(id string) error {
	query := `
		delete users where id = $1;
	`

	row, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, user doesnt exists")
	}

	return nil
}

func (r *RawUsersRepository) FindByEmail(email string) (domain.Users, error) {
	var user domain.Users

	query := `
		select * from users where email=$1;
	`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Address,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Age,
		&user.Role,
	)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

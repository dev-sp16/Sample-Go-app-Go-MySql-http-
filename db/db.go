package db

import (
	"database/sql"
	"fmt"
	"log"
	"screening/utils"
	"screening/user"
	_ "github.com/go-sql-driver/mysql"
)

const fileName = "db_credentials.txt"

func createConnection() (*sql.DB, error) {
	err := utils.ReadDBCredentials(fileName)
	if err != nil {
		log.Fatal( "Failed to read credentials. ERROR: ", err )
		return nil, err
	}

	username := utils.GetUserName()
    password := utils.GetPassword()
    dbName   := utils.GetDBName()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbName))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255) UNIQUE
	)`

	_, err := db.Exec( query )

	if err != nil {
		log.Fatal( "Failed to create 'users' table. ERROR: ", err )
		return err
	}

	return nil
}

func InitDB() (*sql.DB, error) {
	db, err := createConnection()

	if err != nil {
		return nil, err
	}

	err = createUsersTable( db )
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	_, err := db.Exec(query, name, email)

	if err != nil {
		return err
	}

	return nil
}

func ModifyUser(db *sql.DB, name, email, id string) error {
	query := "Update users set name=?, email=? where id=?"

	_, err := db.Exec(query, name, email, id)

	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers(db *sql.DB) ([]user.User, error) {
	query := "SELECT * FROM users"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
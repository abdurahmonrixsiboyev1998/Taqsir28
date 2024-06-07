package db

import (
	"database/sql"
	"log"
	"taqsir/model"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres dbname=demo password=14022014 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}

func CreateUser(user model.User) error {
	_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", user.ID, user.Name, user.Age)
	return err
}

func GetUser(id string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Age)
	return user, err
}

func UpdateUser(user model.User) error {
	_, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
	return err
}

func DeleteUser(id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func ListUsers() ([]model.User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

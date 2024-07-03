package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	var (
		dbUsername = "root"
		dbPassword = "asdf123"
		dbHost     = "localhost:3306"
		dbName     = "test_db"
	)
	connstStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		dbUsername,
		dbPassword,
		dbHost,
		dbName,
	)
	connDB, err := sql.Open("mysql", connstStr)

	if err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	connDB.SetConnMaxLifetime(time.Minute * 3)
	connDB.SetConnMaxIdleTime(time.Minute * 3)
	connDB.SetMaxOpenConns(10)
	connDB.SetMaxIdleConns(10)

	if err := connDB.Ping(); err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	fmt.Println("mantap dh konek")

	r := repo{db: connDB}

	// users, err := r.FindAllUsers()

	// if err != nil {
	// 	log.Fatalf("failed find all user, error: %s", err.Error())
	// }
	// fmt.Println(users)

	// user, err := r.FindUserByID(1)
	// if err != nil {
	// 	log.Fatalf("failed find user by id, error: %s", err.Error())
	// }
	// fmt.Println(user)

	user, err := r.insertUser(User{
		Fullname: "test",
		Email:    "test",
	})
	if err != nil {
		log.Fatalf("failed insert user, error: %s", err.Error())
	}
	fmt.Println(user)

}

type repo struct {
	db *sql.DB
}

func (r *repo) FindAllUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, fullname, email, created_at FROM users")

	if err != nil {
		return nil, err
	}
	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt); err != nil {
			log.Fatalf("failed to scan row, error: %s", err.Error())
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repo) FindUserByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT id, fullname, email, created_at FROM users WHERE id = ?", id)
	var user User

	err := row.Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repo) insertUser(user User) (*User, error) {

	timeNow := time.Now().UTC()
	result, err := r.db.Exec("INSERT INTO users(fullname, email, created_at) VALUES(?, ?, ?)", user.Fullname, user.Email, timeNow)
	if err != nil {
		return nil, err
	}
	user.ID, _ = result.LastInsertId()
	user.CreatedAt = timeNow
	return &user, nil
}

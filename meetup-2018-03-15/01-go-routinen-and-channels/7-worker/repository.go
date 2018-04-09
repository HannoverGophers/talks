package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Name string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) Insert(user User) {
	_, err := repo.db.Exec("INSERT INTO users (name) VALUES ($1);", user.Name)
	if err != nil {
		fmt.Printf("error during insert of user %+v: %s\n", user, err)
	}
	<-time.After(time.Millisecond * 100)
}

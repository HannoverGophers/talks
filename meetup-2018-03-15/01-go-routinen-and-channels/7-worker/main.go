package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// create database connection
	db, err := NewDBConnection()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	repo := NewUserRepository(db)
	pool := NewWorkerPool(10)

	// time measurement
	start := time.Now()

	for i := 1; i <= 100; i++ {
		job := NewJob(repo, User{Name: fmt.Sprintf("John Doe %d", i)})
		pool.Add(job)
	}

	// wait until all work is done
	pool.Wait()

	pool.Close()

	fmt.Println("Elapsed: ", time.Now().Sub(start))

	<-time.After(time.Second * 100)
}

type Job struct {
	repo UserRepository
	user User
}

func NewJob(repo UserRepository, user User) Job {
	return Job{
		repo: repo,
		user: user,
	}
}

func (j Job) Do() {
	j.repo.Insert(j.user)
}

func NewDBConnection() (*sql.DB, error) {
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"<username>",
		"",
		"127.0.0.1",
		"5432",
		"meetup",
		"disable",
	)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return db, nil
}

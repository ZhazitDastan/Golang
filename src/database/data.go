package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "qwertyios02"
	dbname   = "test"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//err = createTask(db, "Name1")
	//if err != nil {
	//	log.Fatal(err)
	//}

	tasks, err := readTasks(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Completed: %t\n", task.ID, task.Name, task.Completed)
	}

	err = updateTask(db, 6, true)
	if err != nil {
		log.Fatal(err)
	}

	//err = deleteTask(db, 6)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func createTask(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO tasks(name) VALUES(?)", name)
	return err
}

func readTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func updateTask(db *sql.DB, id int, completed bool) error {
	_, err := db.Exec("UPDATE tasks SET completed = ? WHERE id = ?", completed, id)
	return err
}

func deleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

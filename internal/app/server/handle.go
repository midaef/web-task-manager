package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	connection "packages/internal/app/database"
	"text/template"
	"time"
)

type Task struct {
	Task     string `json:"todo"`
	DateTime string `json:"datetime"`
	ID       uint   `json:"id"`
}

func NewHandle() {
	http.Handle("/resources/static/", http.StripPrefix("/resources/static", http.FileServer(http.Dir("./resources/static/"))))
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/add", addTask)
	http.HandleFunc("/get", getTasks)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	tmpl.Execute(w, nil)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var t Task
	t.Task = r.FormValue("task")
	dt := time.Now()
	t.DateTime = dt.Format("01-02-2006 15:04:05")
	conn := connection.NewConnection()
	defer conn.Close()
	_, err := conn.Exec("insert into task(todo, datetime) values ($1, $2)", t.Task, t.DateTime)
	if err != nil {
		log.Fatal("Connection DB error", err)
	}
	fmt.Println(t)
	jsonTask, _ := json.Marshal(map[string]interface{}{"datetime": t.DateTime})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTask)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	conn := connection.NewConnection()
	defer conn.Close()
	var size int
	conn.QueryRow("select count(*) from task").Scan(&size)
	t := make([]Task, size)
	rows, err := conn.Query("select * from task limit $1", size)
	i := 0
	for rows.Next() {
		_ = rows.Scan(&t[i].Task, &t[i].DateTime, &t[i].ID)
		i++
	}
	if err != nil {
		log.Fatal(err)
	}
	jsonTask, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTask)
}

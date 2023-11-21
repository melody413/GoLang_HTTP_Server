package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	// query := `
	//     CREATE TABLE users (
	//         id INT AUTO_INCREMENT,
	//         username TEXT NOT NULL,
	//         password TEXT NOT NULL,
	//         created_at DATETIME,
	//         PRIMARY KEY (id)
	//     );`

	// _, err = db.Exec(query)
	// if err != nil {
	// 	fmt.Println("Error creating table:", err.Error())
	// 	return
	// }

	//router
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/", queryParamHandler)
	r.HandleFunc("/login", postRequestHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)

	//run server with specific PORT
	http.ListenAndServe(":8080", r)

}

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	paramValue := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Query parameter value: %s", paramValue)
}

func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		fmt.Fprintf(w, "POST request. username: %s, email: %s", username, password)
		storeUserInformation(username, password)
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Request")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Request")
}

func storeUserInformation(username string, password string) {
	db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		fmt.Println("Error opening database:", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
		return
	}

	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Error creating table:", err.Error())
	} else {
		fmt.Println("index", userID)
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error closing database:", err.Error())
		return
	}
}

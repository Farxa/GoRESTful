package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type User struct {
//     ID       int    `json:"id,omitempty"`
//     Username string `json:"username"`
//     Password string `json:"password,omitempty"`
//     Avatar   string `json:"avatar,omitempty"`
// }

var db *sql.DB

func initDB() {
    var err error
    db, err = connectDB() 
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
}

// func createUser(w http.ResponseWriter, r *http.Request) {
//     var user User
//     if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }


//     _, err := db.Exec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3)",
//         user.Username, user.Password, user.Avatar)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(user) 
// }

// func getUserByID(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id := params["id"]

//     var user User
//     err := db.QueryRow("SELECT id, username, avatar FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Avatar)
//     if err != nil {
//         if err == sql.ErrNoRows {
//             http.NotFound(w, r)
//         } else {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return
//     }

//     json.NewEncoder(w).Encode(user)
// }

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    initDB()

    router := mux.NewRouter()
    // router.HandleFunc("/user", createUser).Methods("POST")
    // router.HandleFunc("/user/{id}", getUserByID).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}

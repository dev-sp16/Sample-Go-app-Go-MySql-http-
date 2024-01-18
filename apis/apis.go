package apis

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"screening/db"
)

func handleCreateUserAPI(connection *sql.DB) {
	http.HandleFunc("/createUser", func(writer http.ResponseWriter, request *http.Request) {
		err := db.SaveUser(connection, request.FormValue("name"), request.FormValue("email"))

		if err != nil {
			log.Println("Failed to create the user. ERROR: ", err)
            http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
        }

        writer.Header().Set("Content-Type", "application/json")
        writer.Write([]byte(`{"message": "Created user successfully!"}`))
	})
}

func handleUpdateUserAPI(connection *sql.DB ) {
	http.HandleFunc("/updateUser", func(writer http.ResponseWriter, request *http.Request) {
		err := db.ModifyUser(connection, request.FormValue("name"), request.FormValue("email"), request.FormValue("id"))
	
		if err != nil {
			log.Println("Failed to update the user. ERROR: ", err)
            http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
        }

        writer.Header().Set("Content-Type", "application/json")
        writer.Write([]byte(`{"message": "User updated successfully!"}`))
	})
}

func handleGetUsersAPI(connection *sql.DB ) {
	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		users, err := db.GetAllUsers(connection)

		if err != nil {
			log.Fatal("Failed to retrieve all the users. ERROR: ", err)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(writer, "%v", users)
	})
}

func SetupJsonApi() {
	connection, err := db.InitDB()

	if err != nil {
		log.Fatal("Could not connect to the database.", err)
		return
	}

	handleCreateUserAPI(connection)
	handleUpdateUserAPI(connection)
	handleGetUsersAPI(connection)
}

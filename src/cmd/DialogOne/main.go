package main

import (
	api "DialogOne/handlers"
	"net/http"
	"os"
)

func main() {
	// mysql.CreateTable(db,err)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/v1/messages/text", api.PostTxtMessages)
	http.HandleFunc("/api/v1/messages", api.GetGeneralMessage)
	http.ListenAndServe(":"+port, nil)
}

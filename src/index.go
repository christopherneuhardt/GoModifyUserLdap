package main


import (
	"net/http"
	"handlers"
)


func main() {
	handlers.Handles()
	http.ListenAndServe(":8081", nil)
}

package main

import (
	"fmt"
	"net/http"
	"backend/internal"
)

func main() {
	go internal.HubInstance.Run()

	http.HandleFunc("/ws", internal.ServeWs)

	fmt.Println("✅ Server started on :8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

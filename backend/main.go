package main

import (
	"fmt"
	"net/http"
	"backend/internal"
)

func main() {
	go internal.HubInstance.Run()

	http.HandleFunc("/ws", internal.ServeWs)

	ports := []string{"8080", "8081", "8082"}

	var listener net.Listener
	var err error

	for _, port := range ports {
		listener, err = net.Listen("tcp", "0.0.0.0:"+port)
		if err == nil {
			fmt.Printf("✅ Server started on :%s\n", port)
			break
		}
	}

	if listener == nil {
		log.Fatalf("❌ 全ポートが使用中でした: %v", err)
	}

	log.Fatal(http.Serve(listener, nil))
}

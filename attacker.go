package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	target := "http://localhost:8080"
	paths := []string{"/admin", "/.env", "/config.php", "/wp-login.php"}

	fmt.Println("launching simulated reconnaissance...")

	for i := 0; i < 15; i++ {
		path := paths[i%len(paths)]
		fmt.Printf("sending malicious request to %s%s\n", target, path)

		http.Get(target + path)
		time.Sleep(800 * time.Millisecond) // delay para o log respirar
	}
}

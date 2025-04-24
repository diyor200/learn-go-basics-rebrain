package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const helloWrold = "Hello world!"

func main() {
	http.HandleFunc("/hello", helloWorld)
	fmt.Println("Server started in :8080 address")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received:", r.RequestURI)
	var (
		res string
		n   = rand.Intn(1000)
	)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		if res == "" {
			res = helloWrold
		}

		res += "\n" + helloWrold
	}

	fmt.Fprint(w, res)
	return
}

func MySlowFunction() int {
	a := rand.Intn(100)
	b := rand.Intn(100)
	time.Sleep(time.Millisecond * 100)
	return a + b
}

func MyOptimizedFunction() int {
	a := rand.Intn(100)
	b := rand.Intn(100)
	time.Sleep(time.Millisecond * 10)
	return a + b
}

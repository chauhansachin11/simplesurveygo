//it is a example of handling 10 concurrent requests by go routine  

package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
)

func main() {
	go startServer()
	
    sendRequest := func() {
        resp, _ := http.Get("http://localhost:8080/")
        defer resp.Body.Close()
    }
	start := time.Now()
	
	var wg sync.WaitGroup
	
    chn := make(chan int, 10)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            sendRequest()
            chn <- n
        }(i)
    }
    go func() {
        wg.Wait()
        close(chn)
    }()

    for routineNumber := range chn {
        fmt.Printf("%d ", routineNumber)
    }
    fmt.Println()
    fmt.Println("time:", time.Since(start))
}

func startServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(1 * time.Second)
    })
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
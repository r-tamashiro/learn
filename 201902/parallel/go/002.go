package main

import (
    "fmt"
    "sync"
)


func printHello(wg *sync.WaitGroup) {
    fmt.Println("Hello")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go printHello(&wg)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go printHello(&wg)
    }
    fmt.Println("World")
    wg.Wait()
}

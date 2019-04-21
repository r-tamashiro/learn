package main

import (
    "fmt"
    "sync"
)


func printHello(wg *sync.WaitGroup, i int) {
    // 呼び出し元での指定はできないみたい
    defer wg.Done()
    if i  > 3 {
        return
    }
    fmt.Println("Hello")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go printHello(&wg, i)
    }
    fmt.Println("World")
    wg.Wait()
}

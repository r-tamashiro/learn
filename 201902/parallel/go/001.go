/**
 *
 * 並列処理検証 001
 *
 */
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
    fmt.Println("World")
    wg.Wait()
}

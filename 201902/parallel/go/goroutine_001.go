/**
 *
 * 並列処理検証 002
 * syncを使っての非同期処理 001
 *
 */
package main

import (
    "fmt"
    "sync"
)


func printHello(wg *sync.WaitGroup, i int) {
    fmt.Println("Hello", i)
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

/**
 *
 * 並列処理検証 003
 * syncを使っての非同期処理 002
 * deferを使ってDone前のエラー時の処理
 *
 */
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

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
    "time"
)


func sendHello(ch chan<- string) {
    // ここで重たい処理
    time.Sleep(3 * time.Second)
    ch <- "Hello"
}

func main() {
    ch := make(chan string)
    go sendHello(ch)

    // この間に重たい処理
    msg := <-ch
    fmt.Println(msg)
}

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
    fmt.Println("重たい処理開始")
    // ここで重たい処理
    time.Sleep(3 * time.Second)
    ch <- "Hello"
}

func main() {
    ch := make(chan string)
    go sendHello(ch)

    fmt.Println("それ以外の処理")
    // この間に重たい処理
    msg := <-ch
    fmt.Println(msg)
}

/**
 *
 * 並列処理検証 004
 * channelを使ったポーリング
 * タイムアウトの実装
 *
 */
package main

import (
    "fmt"
    "time"
)

func main() {
    q := make(chan string)

    go func() {
        // 重たい処理
        // 2だと完了する
        time.Sleep(2 * time.Second)
        q <- "完了"
        close(q)
    }()

    select {
    case msg := <-q:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("タイムアウト")
    }

    fmt.Println("何か")
}

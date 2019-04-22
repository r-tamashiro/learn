/**
 *
 * 並列処理検証 003
 * channelを使ったポーリング
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
        time.Sleep(10 * time.Second)
        q <- "完了"
    }()

    n := 0
    loop:
        for {
            select {
            case msg := <-q:
                // 結果がくれば表示して終了
                fmt.Println(msg)
                break loop
            default:
        }

        n++
        fmt.Println("何か", n)
        time.Sleep(1 * time.Second)
    }
}

/**
 *
 * 並列処理検証 005
 * channelを使ったポーリング
 * worker
 * よくわからないがdeadlockする
 : 現状回避方法が不明
 *
 */
package main

import (
    "fmt"
    "time"
    "sync"
)


type result struct {
    url     string
    result  bool
}

func fetchUrl(wg *sync.WaitGroup, q <-chan string, r chan<- result) {
    defer wg.Done()
    for {
        url, ok := <-q
        if !ok {
            return
        }

        fmt.Println("ダウンロード: ", url)
        time.Sleep(3 * time.Second)
        r <- result {
            url:    url,
            result: true,
        }
    }

}

func main() {
    var wg sync.WaitGroup

    q := make(chan string, 5)

    results := make(chan result, cap(q))

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go fetchUrl(&wg, q, results)
    }

    urls := []string {
        "https://ja.wikipedia.org/wiki/Linux",
        "https://ja.wikipedia.org/wiki/Linux%E3%82%AB%E3%83%BC%E3%83%8D%E3%83%AB",
        "https://ja.wikipedia.org/wiki/GNU/Linux%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0",
        "https://ja.wikipedia.org/wiki/Linux%E3%83%87%E3%82%A3%E3%82%B9%E3%83%88%E3%83%AA%E3%83%93%E3%83%A5%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3",
        "https://ja.wikipedia.org/wiki/Linux%E3%83%9E%E3%82%B7%E3%83%B3",
        "https://ja.wikipedia.org/wiki/TOMOYO_Linux",

    }

    for _, url := range urls {
        q <- url
    }
    close(q)

    wg.Wait()
    close(results)
    for r := range results {
        fmt.Println(r)
    }
}

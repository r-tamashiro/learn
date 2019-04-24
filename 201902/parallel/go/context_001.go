package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "time"
)

func doSomething(ctx context.Context) {
    time.Sleep(5 * time.Second)
}

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
    defer cancel()

    go doSomething(ctx)

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, os.Interrupt)

    select {
        case <-sc:
            cancel()
            fmt.Println("interrupted")
        case <-ctx.Done():
            fmt.Println("done:", ctx.Err())
    }
}

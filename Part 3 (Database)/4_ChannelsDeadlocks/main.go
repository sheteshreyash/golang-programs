package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    fmt.Println("Demonstrating safe channel patterns in Go")

    // Buffered channel of ints
    ch := make(chan int, 2)

    // WaitGroup to wait for all goroutines
    var wg sync.WaitGroup

    // sync.Once ensures we only close once
    var once sync.Once
    safeClose := func() {
        once.Do(func() {
            fmt.Println("→ closing channel")
            close(ch)
        })
    }

    // Safe send: recovers if channel is closed
    safeSend := func(val int) {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("⚠️  panic recovered on send(%d): %v\n", val, r)
            }
        }()
        ch <- val
        fmt.Printf("→ sent %d\n", val)
    }

    // 1) Reader: consumes until channel is closed
    wg.Add(1)
    go func() {
        defer wg.Done()
        for v := range ch {
            fmt.Println("📥 reader got:", v)
            time.Sleep(50 * time.Millisecond) // simulate work
        }
        fmt.Println("📤 reader exiting (channel closed)")
    }()

    // 2) Non‑blocking peek (default case)
    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(10 * time.Millisecond) // wait for nothing yet
        select {
        case v, ok := <-ch:
            fmt.Println("👀 non‑blocking peek:", v, ok)
        default:
            fmt.Println("👀 non‑blocking peek: no value")
        }
    }()

    // 3) Writers: send a few values
    wg.Add(1)
    go func() {
        defer wg.Done()
        safeSend(1)
        safeSend(2)
    }()

    // 4) Attempt double‑close and send‑after‑close
    wg.Add(1)
    go func() {
        defer wg.Done()
        // allow previous sends to complete
        time.Sleep(200 * time.Millisecond)
        safeClose()           // first close
        safeClose()           // second close (no panic)
        safeSend(99)          // send after close (recovered)
    }()

    // 5) Nil‑channel demonstration (won’t block the rest)
    wg.Add(1)
    go func() {
        defer wg.Done()
        var nilCh chan int // nil channel
        select {
        case nilCh <- 5:
            fmt.Println("sent on nil")
        case <-time.After(100 * time.Millisecond):
            fmt.Println("⏱ send to nil channel timed out (as expected)")
        }
    }()

    // Wait for all goroutines
    wg.Wait()
    fmt.Println("✅ all done, no deadlocks or panics")
}

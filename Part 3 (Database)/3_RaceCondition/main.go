package main

import (
    "fmt"
    "sync"
)

func main() {
    fmt.Println("Race condition in Golang")

    var (
        wg    = &sync.WaitGroup{}
        mutex = &sync.RWMutex{}
        score = []int{0}
    )

    // You have 4 goroutines, so Add(4)
    wg.Add(4)

    // Writer #1
    go func(wg *sync.WaitGroup, m *sync.RWMutex) {
        defer wg.Done()
        fmt.Println("One R")
        m.Lock()
        score = append(score, 1)
        m.Unlock()
    }(wg, mutex)

    // Writer #2
    go func(wg *sync.WaitGroup, m *sync.RWMutex) {
        defer wg.Done()
        fmt.Println("Two R")
        m.Lock()
        score = append(score, 2)
        m.Unlock()
    }(wg, mutex)

    // Writer #3
    go func(wg *sync.WaitGroup, m *sync.RWMutex) {
        defer wg.Done()
        fmt.Println("Three R")
        m.Lock()
        score = append(score, 3)
        m.Unlock()
    }(wg, mutex)

    // Reader
    go func(wg *sync.WaitGroup, m *sync.RWMutex) {
        defer wg.Done()
        fmt.Println("Reader")
        m.RLock()
        fmt.Println("  snapshot:", score)
        m.RUnlock()
    }(wg, mutex)

    // Wait for all 4 goroutines to finish
    wg.Wait()

    // Now that all appends are done, it's safe to read without a lock
    fmt.Println("Final score slice:", score)
}

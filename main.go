package main

import (
	"fmt"
	"time"
)

func startTimer(minutes int) {
    duration := time.Duration(minutes) * time.Minute
    ticker := time.NewTicker(time.Second)
    done := make(chan bool)

    go func() {
        time.Sleep(duration)
        done <- true
    }()

    for {
        select {
        case <-done:
            fmt.Println("Time's up!")
            return
        case t := <-ticker.C:
            fmt.Printf("\r%v", duration-time.Until(t))
        }
    }
}

func main() {
    var choice string
    fmt.Println("Start Pomodoro timer? (work/break): ")
    fmt.Scanln(&choice)

    if choice == "work" {
        fmt.Println("Starting 25-minute work timer...")
        startTimer(25)
    } else if choice == "break" {
        fmt.Println("Starting 5-minute break timer...")
        startTimer(5)
    } else {
        fmt.Println("Invalid choice")
    }
}

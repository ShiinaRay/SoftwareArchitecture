package main

import (
    "fmt"
    "github.com/ShiinaRay/SoftwareArchitecture"
)

func main() {
    lt := linktable.New[string]()
    lt.PushBack("Alice")
    lt.PushBack("Bob")
    lt.PushBack("Carol")

    fmt.Println("Contents:")
    lt.Iter(func(name string) {
        fmt.Println("-", name)
    })

    for {
        if val, ok := lt.PopFront(); ok {
            fmt.Println("Popped:", val)
        } else {
            break
        }
    }
}

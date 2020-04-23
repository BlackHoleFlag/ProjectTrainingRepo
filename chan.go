package main

import (  
    "fmt"
)

//var wg sync.WaitGroup

func main() {
    c:= make(chan int)
    x:=1
    go count(x,c)

    for msg := range c {
        fmt.Println(msg)
    }
}

func count (thing int, c chan int) {
    for i:=1; i<=5; i++ {
        c <- thing+i
    }
    close(c)
}
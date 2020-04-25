package main

import (  
    "fmt"
    "sync"
)

var wg2 sync.WaitGroup

func main() {
    c:= make(chan int)
    x:=1
   
    go count(x,c)
    go count(x,c)
    go count(x,c)
    for msg := range c {
        fmt.Println(msg)
    }
}

func count (thing int, c chan int) {
   
    for i:=1; i<=10000; i++ {
        c <- thing+i
    }
 
    close(c)
}
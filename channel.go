package main

import (
  "fmt"
  "time"
)

func writer(channel chan<- string, msg string) {
  channel <- msg
}

func reader(channel <-chan string) {
  fmt.Println("Reading Message and Appending hello World....")
  msg := <- channel
  msg = msg +"World"
  fmt.Println(msg)
}

func main() {
  msgChan := make(chan string, 1)

  go reader(msgChan)

  writer(msgChan, "Hello")

  time.Sleep(time.Second * 5)
}



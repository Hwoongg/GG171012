package main

import "fmt"

func main(){
  int a; // 정적타입
  var j string = "Hi Golang!"
  fmt.Println("j: :", j)

  var(
    name string = "machine"
    height int32
    weight float32
    isRunning bool
  )

  height = 250
  weight = 350.56
  isRunning = true
}

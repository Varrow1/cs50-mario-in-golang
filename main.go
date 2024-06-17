package main

import (
	"fmt"
	"strings"
)

func main(){
  fmt.Println("How Many Stairs Do You Want?")
  var i int
  fmt.Scanln(&i)
  if i>0{
    fmt.Println("Alright")
    n := 0
    for n<i{
      n++
      o := strings.Repeat("#", n)
      fmt.Println(o)
    }
  } else {
    fmt.Println("not a valid input")
  }
}

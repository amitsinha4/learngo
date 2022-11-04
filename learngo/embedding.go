package main

import "fmt"

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num=%w", b.num)
}

type container struct {
  base
  str string
}

func main(){
  co := container{base: base{num: 1}, str: "Some Name"}
  
  fmt.Printf("co={num: %v, str:  %v}", co.num, co.str)
  fmt.Println("also num: ", co.base.num)
  fmt.Println("describe: ", co.describe())

}

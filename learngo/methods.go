package main

import "fmt"

type rect struct {
  width, height int
}

type square struct {
  length int
}

func (s *square) area() int {
  return s.length * s.length
}

func (s *square) perim() int {
  return 4 * s.length
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

func main(){
  r:= rect{width: 10, height: 5}

  fmt.Println("Area: ", r.area())
  fmt.Println("perim: ", r.perim())

  rp := &r
  fmt.Println(rp.area())
  fmt.Println(rp.perim())

  s := square{length: 4}
  fmt.Println(s.area())
  fmt.Println(s.perim())
}

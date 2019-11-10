package main

import ("fmt"
        "importarGo")
func anotherfunc()(int,int){
  return 2 , 3
}
func main(){
  fmt.Println("Hello world")

  c:= importarGo.SumarValores(2, 3)
  fmt.Println(c)

  a, b := anotherfunc()
  fmt.Println( a , b )
}

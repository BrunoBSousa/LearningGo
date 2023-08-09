package main

import(
  "fmt"
  "math/rand"
)

func main(){
  c:= [3]string {"Rock","Paper","Scissors"}
var name string
var choise string
fmt.Println("Hi Stanger, what is you name ?")
fmt.Scan(&name)
fmt.Printf("Hi,%v this is a game of Rock, Paper Scissors, please choose one of the following to begin the game\n",name)
fmt.Scan(&choise)
fmt.Println(choise)
//fmt.Print(rand.Intn(3))
r := c[rand.Intn(3)]
//fmt.Println(r)
//fmt.Println(c[rand.Intn(3)])
if choise == r {
fmt.Println("empate")
} else if choise == "Rock" && r== "Paper"{
fmt.Println("perdeste")
}
}


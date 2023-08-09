package main

import(
  "fmt"
  "math/rand"
)

func main(){
  c:= [3]string {"Rock","Paper","Scissors"}
var name string
var choise string


//Asking for user's name and pick for the game
fmt.Println("Hi Stanger, what is you name ?")
fmt.Scan(&name)
fmt.Printf("Hi,%v this is a game of Rock, Paper Scissors, please choose one of the following to begin the game\n",name)
fmt.Scan(&choise)
//Generationg a number from 0 to 3 to use as index of the array
r := c[rand.Intn(3)]
//Doing the logic of the game to see if User won or computer won the game
if choise == r {
fmt.Println("It was a Draw, the computer also choose",r)
} else if choise == "Rock" && r== "Paper"{
fmt.Printf("Sorry you Lost, the computer picked %v \n",r)
}else if choise == "Rock" && r == "Scissors"{
fmt.Printf("Congratulations you Won, the computer picked %v \n",r)
}else if choise == "Paper" && r == "Scissors"{
fmt.Printf("Sorry you Lost, the computer picked %v \n",r)
}else if choise =="Paper" && r == "Rock"{
fmt.Printf("Congratulations you Won, the computer picked %v \n",r)
}else if choise =="Scissors" && r == "Rock"{
fmt.Printf("Sorry you Lost, the computer picked %v \n",r)
}else{
  fmt.Printf("Congratulations you Won, the computer picked %v \n",r)
}
}


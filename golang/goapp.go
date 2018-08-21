package main

import "fmt"

func main() {

// creating vars 
  var user string
  var box string
// assigning value to vars
  fmt.Print("Which user box do you want?:")
  fmt.Scanf("%s", &user)
  fmt.Print("The url of which box do you want?:")
  fmt.Scanf("%s", &box)
//grab json from api call
  fmt.Println("You want the box of", user)
  fmt.Println(box)
// filter json and output urls

}

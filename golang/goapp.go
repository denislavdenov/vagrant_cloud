package main

import (
	"fmt"
//	"encoding/json"
	"net/http"
//	"bytes"
	"io/ioutil"
)



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

var url1 string = "https://app.vagrantup.com/api/v1/box/"

	url1 = url1 + "/" + user + "/" + box

response, err := http.Get(url1)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }



// filter json and output urls


}




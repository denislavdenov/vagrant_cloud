package main

import (
	"fmt"
	"encoding/json"
	"net/http"
//	"bytes"
	"io/ioutil"
	"time"
)



func main() {

// creating vars 
  var user string
  var box string

// response struct
// generared with https://mholt.github.io/json-to-go/

type vagrantcloud struct {
        Tag                 string    `json:"tag"`
        Username            string    `json:"username"`
        Name                string    `json:"name"`
        Private             bool      `json:"private"`
        Downloads           int       `json:"downloads"`
        CreatedAt           time.Time `json:"created_at"`
        UpdatedAt           time.Time `json:"updated_at"`
        ShortDescription    string    `json:"short_description"`
        DescriptionMarkdown string    `json:"description_markdown"`
        DescriptionHTML     string    `json:"description_html"`
        CurrentVersion      struct {
            Version             string    `json:"version"`
            Status              string    `json:"status"`
            DescriptionHTML     string    `json:"description_html"`
            DescriptionMarkdown string    `json:"description_markdown"`
            CreatedAt           time.Time `json:"created_at"`
            UpdatedAt           time.Time `json:"updated_at"`
            Number              string    `json:"number"`
            ReleaseURL          string    `json:"release_url"`
            RevokeURL           string    `json:"revoke_url"`
            Providers           []struct {
                Name        string      `json:"name"`
                Hosted      bool        `json:"hosted"`
                HostedToken interface{} `json:"hosted_token"`
                OriginalURL interface{} `json:"original_url"`
                CreatedAt   time.Time   `json:"created_at"`
                UpdatedAt   time.Time   `json:"updated_at"`
                DownloadURL string      `json:"download_url"`
            } `json:"providers"`
        } `json:"current_version"`
        Versions []struct {
            Version             string    `json:"version"`
            Status              string    `json:"status"`
            DescriptionHTML     string    `json:"description_html"`
            DescriptionMarkdown string    `json:"description_markdown"`
            CreatedAt           time.Time `json:"created_at"`
            UpdatedAt           time.Time `json:"updated_at"`
            Number              string    `json:"number"`
            ReleaseURL          string    `json:"release_url"`
            RevokeURL           string    `json:"revoke_url"`
            Providers           []struct {
                Name        string      `json:"name"`
                Hosted      bool        `json:"hosted"`
                HostedToken interface{} `json:"hosted_token"`
                OriginalURL interface{} `json:"original_url"`
                CreatedAt   time.Time   `json:"created_at"`
                UpdatedAt   time.Time   `json:"updated_at"`
                DownloadURL string      `json:"download_url"`
            } `json:"providers"`
        } `json:"versions"`
    }

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
        data, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        }
        json1 := string(data)
        // filter json and output urls
        var result vagrantcloud
        json.Unmarshal([]byte(json1), &result)
        
        for i := range result.CurrentVersion.Providers {
            println(result.CurrentVersion.Providers[i].DownloadURL)
        }
        
    }
}

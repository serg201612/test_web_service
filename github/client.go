package main

import (
	//"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)


type ForumEntry struct {
	Id      int
	Email   string
	Title   string
	Content string
}

type TopicEntry struct {
	Id      int
	Link   string
	Name   string
}

type ContentEntry struct {
	Id      int
	Content   string
}

type Forum struct {
	forumData []*ForumEntry
	topicData []*TopicEntry
	contentEntry ContentEntry

}





// Command line flags.
var requestUrl = flag.String("request_url", "",
	"Host and port to send the request to")
var requestMethod = flag.String("request_method", "GET",
	"HTTP request method")
var requestData = flag.String("request_data", "",
	"Marshalled JSON data")

// doRequest executes an HTTP request to the given requestUrl using the given
// requestMethod and requestData.
func doRequest(requestMethod, requestUrl,
	requestData string) (*http.Response, error) {
	// These will hold the return value.
	var res *http.Response
	var err error

	
	// Convert method to uppercase for easier checking.
	upperRequestMethod := strings.ToUpper(requestMethod)
	switch upperRequestMethod {
	case "GET":
		// Use the HTTP library Get() method.
		res, err = http.Get(requestUrl)
		//fmt.Printf("!!! res=", res)
		//fmt.Printf("error=", err.Error())

	default:
		// We doń't know how to handle this request.
		return nil, fmt.Errorf(
			"invalid --request_method provided : %s",
				requestMethod)
	}

	return res, err
}

func main() {
	// Parse command line flags.


	flag.Parse()

	if len(*requestUrl) == 0 {
		// We need a request URL.
		fmt.Println("--request_url must be provided")
		return
	}

	
	if len(*requestMethod) == 0 {
		// And we also need a method.
		fmt.Println("--request_method must be provided")
	}

	//requestUrl := "http://127.0.0.1:8000/forum"
	// Execute request (if possible).
	res, err := doRequest(*requestMethod, *requestUrl, *requestData)
	if err != nil {
		// Request failed.
		fmt.Println("Error executing HTTP request :", err)
		return
	}

	// Make sure res.Body is closed whwn we are done.
	defer res.Body.Close()

	// Read body data (i.e. our response).
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Could not read data.
		fmt.Println("Error reading response data :", err)
		return
	}

	// Print received data.
	//fmt.Println(string(responseData))
	
	//var forumEntry ForumEntry
	var forum Forum
	err = json.Unmarshal(responseData, &forum.forumData)
	if err != nil {
		//fmt.Println("Error 1 :", err)
	}else {
	  for _, f := range forum.forumData{
	    fmt.Println("", f)
	  }
	}

	err = json.Unmarshal(responseData, &forum.topicData)
	if err != nil {
		//fmt.Println("Error 2 :", err)
	}else {
	  for _, t := range forum.topicData{
	    fmt.Println("", t)
	  }
	}
	
	err = json.Unmarshal(responseData, &forum.contentEntry.Content)
	if err != nil {
		//fmt.Println("Error 3 :", err)
	}else {
	  fmt.Println("", forum.contentEntry)
	}
	
	
}


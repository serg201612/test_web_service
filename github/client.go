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
	Id     string
	Name   string
	Link   string
}

type TopicEntry struct {
	Id     string
	Name   string
	Link   string
}

type ContentEntry struct {
	Id      string
	Content string
	Link    string
}

type Forum struct {
	forumData []*ForumEntry
	topicData []*TopicEntry
	contentData []*ContentEntry
}






// Command line flags.
var requestUrl = flag.String("request_url", "",
	"Host and port to send the request to")
var requestMethod = flag.String("request_method", "GET",
	"HTTP request method")
var requestData = flag.String("request_data", "", "Marshalled JSON data")

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

	//fmt.Println(*requestUrl)
	//fmt.Println(*requestMethod)
	fmt.Println(*requestData)
	
	
	
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

	
	forum := &Forum{
	make([]*ForumEntry, 0),
	make([]*TopicEntry, 0),
	make([]*ContentEntry, 0),}
		
	err = json.Unmarshal(responseData, &forum.forumData)
	
	if err != nil {
		fmt.Println("Error client forumsdata :", err)
	}else {
	  for _, f := range forum.forumData{
	    if f.Link != "" {
	    fmt.Printf("Link=%s Name=%s \n", f.Link, f.Name)
	    }
	  }
	}

	err = json.Unmarshal(responseData, &forum.contentData)
	if err != nil {
		fmt.Println("Error client contentdata :", err)
	}else {
	  for _, c := range forum.contentData{
	    if c.Content != "" {
	      fmt.Println(c.Content)
	    }
	  }
	}
	
}


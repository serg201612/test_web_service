package crawler

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
	"regexp"
	"./gocrawl"
	"./goquery"
)


var request_url string
var request_type string
var contentData []*ContentEntry
var topicData []*TopicEntry
var forumData []*ForumEntry


func get_http_tag (html_body string, search string, cur int) (bool, int) {

      i := 0
      temp :=""
      _f := false
      if cur+len(search) > len(html_body){return false, 0}
      temp_body := html_body[cur:cur+len(search)]
      start_temp_body := cur
      for _, r := range temp_body {
        c := string(r)
	cur += 1
	if c == search[i:i+1]{
	    i += 1
	    temp += c
	    //if len(temp) > 1{
	    //  fmt.Println(temp)
	    //}
	}else{
	  i = 0 
	  temp = ""
	  start_temp_body = cur
	}
	if temp == search{
	    //fmt.Println(temp)
	    _f = true
	    return _f, start_temp_body
	}   
      }	  
	
    return false, 0

  
}

func getForums(body string) {
    
    find_body := false
    find_href_start := false
    find_href_end := false
    cur := 0
    href_index_start :=cur
    href_index_end :=cur
    
    for _, r := range body {
        c := string(r)
	if c != ""{}
	cur += 1
	
       if find_body == false{
	 find_body, _ = get_http_tag (body, "<div id=\"boardindex_table\">", cur)
       }else{
	  if find_href_start == false{ 
	    find_href_start, href_index_start  = get_http_tag (body, "<a class=\"subject\"", cur)
	  }else{
	    
	    if find_href_end == false{ 
	      find_href_end, href_index_end  = get_http_tag (body, "</a>", cur)
	      if find_href_end == true{
		href_body := body[href_index_start:href_index_end]
		li1 := strings.LastIndex(href_body,"\">")
		li2 := strings.LastIndex(href_body,"name=\"b")
		forum_title := href_body[li1+2:len(href_body)]
		forum_id := href_body[li2+7:li1]+".0"
		//fmt.Println(href_body)
		forumentry := &ForumEntry{forum_id, forum_title, ""}
		forumData = append(forumData, forumentry)
		//fmt.Println(forumEntry)
		find_href_end = false 
		find_href_start = false
	      }	  
	    }
	  }
       }
    }
    
}    


func getTopics(body string) {
    
    find_body := false
    find_href_start := false
    find_href_end := false
    cur := 0
    href_index_start :=cur
    href_index_end :=cur
    
    for _, r := range body {
        c := string(r)
	if c != ""{}
	cur += 1
	
       if find_body == false{
	 find_body, _ = get_http_tag (body, "<div class=\"tborder topic_table\" id=\"messageindex\">", cur)
       }else{
	  if find_href_start == false{ 
	    find_href_start, href_index_start  = get_http_tag (body, "<a href=\"http://bm26rwk32m7u7rec.onion.link/index.php?topic=", cur)
	  }else{
	    if find_href_end == false{ 
	      find_href_end, href_index_end  = get_http_tag (body, "</a>", cur)
	      if find_href_end == true{
		  if !strings.Contains(body[href_index_start:href_index_end], "<img src=") && !strings.Contains(body[href_index_start:href_index_end], "cur_topic_id"){
		    href_body := body[href_index_start:href_index_end]
		    //fmt.Println(href_body)
		    li1 := strings.LastIndex(href_body,"\">")
		    li2 := strings.LastIndex(href_body,"index.php?topic=")
		    topic_title := href_body[li1+2:len(href_body)]
		    topic_id := href_body[li2+16:li1]
		    //fmt.Println(topic_id, topic_title)
		    topicentry := &TopicEntry{topic_id, topic_title, ""}
		    topicData = append(topicData, topicentry)
		  }
		  find_href_end = false 
		  find_href_start = false
	      }	  
	    }
	  }
       }
    }
    
}    
       
func getContent(body string) {
    
    find_body := false
    find_href_start := false
    find_href_end := false
    cur := 0
    href_index_start :=cur
    href_index_end :=cur
    
    for _, r := range body {
        c := string(r)
	if c != ""{}
	cur += 1
	
       if find_body == false{
	 find_body, _ = get_http_tag (body, "<div id=\"forumposts\">", cur)
       }else{
	  if find_href_start == false{ 
	    find_href_start, href_index_start  = get_http_tag (body, "<div class=\"inner\"", cur)
	  }else{
	    if find_href_end == false{ 
	      find_href_end, href_index_end  = get_http_tag (body, "</div>", cur)
	      if find_href_end == true{
		href_body := body[href_index_start:href_index_end]
		li1 := strings.Index(href_body,"\">")
		href_body = href_body[li1+2:len(href_body)]
		href_body = strings.Replace(href_body,"<br />", " ", -1)
		//fmt.Println(href_body)
		re := regexp.MustCompile("<[^>]*>")
		href_body = re.ReplaceAllString(href_body, "")
		contententry := &ContentEntry{"", href_body, ""}
		contentData = append(contentData, contententry)
		find_href_end = false 
		find_href_start = false
	      }	  
	    }
	  }
       }
    }
}    
       


func getboardslist(body string) {
  forumData = make([]*ForumEntry, 0)
  getForums(body)
}

func gettopicslist(body string) {
  topicData = make([]*TopicEntry, 0)
  getTopics(body)
}

func getcontent(body string) {
  contentData = make([]*ContentEntry, 0)
  getContent(body)
  //for _, c := range contentData{
  //  fmt.Println(c)
  //}
}

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//fmt.Printf("Visit: %s\n", ctx.URL())
	if bd, e := ioutil.ReadAll(res.Body); e != nil {
	} else {
	      //_ = ioutil.WriteFile("forums.html", []byte(bd), 0644)
	      if request_type == "forums"{
		getboardslist(string(bd))
	      }
	      if request_type == "topics"{
		gettopicslist(string(bd))
	      }
	      if request_type == "content"{
		getcontent(string(bd))
	      }
	}
	return nil, true
}

var boards map[string]int

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}
	if ctx.URL().Host == "bm26rwk32m7u7rec.onion.link" {
	    if ctx.URL().String() == request_url {
		return true
	  }
	}
	return false
}





func GetForumList(url string, type_request string) []*ForumEntry{
	boards = make(map[string]int)
	ext := &Ext{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = false
	opts.MaxVisits = 100
	goquery.Http_protocol = "https"
	c := gocrawl.NewCrawlerWithOptions(opts)
	request_url = url
	request_type = type_request
	fmt.Println(request_url)
	c.Run(request_url)
	return forumData
}


func GetTopicList(url string, type_request string) []*TopicEntry{
	boards = make(map[string]int)
	ext := &Ext{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = false
	opts.MaxVisits = 100
	goquery.Http_protocol = "https"
	c := gocrawl.NewCrawlerWithOptions(opts)
	request_url = url
	request_type = type_request
	fmt.Println(request_url)
	c.Run(request_url)
	return topicData
}

func GetContent(url string, type_request string) []*ContentEntry{
	boards = make(map[string]int)
	ext := &Ext{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = false
	opts.MaxVisits = 100
	goquery.Http_protocol = "https"
	c := gocrawl.NewCrawlerWithOptions(opts)
	request_url = url
	request_type = type_request
	fmt.Println(request_url)
	c.Run(request_url)
	return contentData
}


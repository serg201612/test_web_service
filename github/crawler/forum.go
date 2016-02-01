package crawler

import (
	//"fmt"
	"sync"
	"../martini"
	
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
	mutex *sync.Mutex
}

func NewForum() *Forum {
	return &Forum{
		make([]*ForumEntry, 0),
		make([]*TopicEntry, 0),
		make([]*ContentEntry, 0),
		new(sync.Mutex),
	}
}

func (g *Forum) GetAllEntries() []*ForumEntry {
  	entries := make([]*ForumEntry, 0)
	entries = GetForumList("http://bm26rwk32m7u7rec.onion.link/", "forums") 
	return entries
}

func (g *Forum) GetAllEntriesFromDB() []*ForumEntry {
	entries := make([]*ForumEntry, 0)
	entries = GetForumListFromDB()
	return entries
}

func (g *Forum) GetAllTopicEntries(params martini.Params) []*TopicEntry {
  	id := params["id"]
	entries := make([]*TopicEntry, 0)
	entries = GetTopicList(id, "http://bm26rwk32m7u7rec.onion.link/index.php?board="+id, "topics") 
	return entries
}

func (g *Forum) GetAllTopicEntriesFromDB(params martini.Params) []*TopicEntry {
  	id := params["id"]
	entries := make([]*TopicEntry, 0)
	entries = GetTopicListFromDB(id) 
	return entries
}

func (g *Forum) GetContentsFromCrawler(params martini.Params) []*ContentEntry {
	entries := make([]*ContentEntry, 0)
  	id := params["id"]
	entries = GetContent(id, "http://bm26rwk32m7u7rec.onion.link/index.php?topic="+id, "content") 
	return entries
}

func (g *Forum) GetContentsFromDB(params martini.Params) []*ContentEntry {
	entries := make([]*ContentEntry, 0)
  	id := params["id"]
	entries = GetContentsFromDB(id) 
	return entries
}

package crawler

import (
	//"fmt"
	"sync"
)

type ForumEntry struct {
	Id      int
	Name   string
	Link   string
}

type TopicEntry struct {
	Id      int
	Name   string
	Link   string
}

type ContentEntry struct {
	Id      int
	Content   string
	Link   string
}

type Forum struct {
	forumData []*ForumEntry
	topicData []*TopicEntry
	contentEntry ContentEntry
	mutex *sync.Mutex
}

func NewForum() *Forum {
	contentEntry := &ContentEntry{0, "", "",}
	return &Forum{
		make([]*ForumEntry, 0),
		make([]*TopicEntry, 0),
		*contentEntry,
		new(sync.Mutex),
	}
}

func (g *Forum) GetAllEntries() []*ForumEntry {
	entries := make([]*ForumEntry, 0)
	forum_entrie := &ForumEntry{1,"First Board Name from crawler","First Board Link from crawler",}
	entries = append(entries, forum_entrie)
	forum_entrie = &ForumEntry{2,"Second Board Name from crawler","Second Board Link from crawler",}
	entries = append(entries, forum_entrie)
	return entries
}

func (g *Forum) GetAllEntriesFromDB() []*ForumEntry {
	entries := make([]*ForumEntry, 0)
	forum_entrie := &ForumEntry{3,"First Board Name from db","First Board Link from db",}
	entries = append(entries, forum_entrie)
	forum_entrie = &ForumEntry{4,"Second Board Name from db","Second Board Link from db",}
	entries = append(entries, forum_entrie)
	return entries
}

func (g *Forum) GetAllTopicEntries() []*TopicEntry {
	entries := make([]*TopicEntry, 0)
	topic_entrie := &TopicEntry{1,"First topic name from crawler","First topic link from crawler",}
	entries = append(entries, topic_entrie)
	topic_entrie = &TopicEntry{2,"Second topic name from crawler","Second topic link from crawler",}
	entries = append(entries, topic_entrie)
	return entries
}

func (g *Forum) GetAllTopicEntriesFromDB() []*TopicEntry {
	entries := make([]*TopicEntry, 0)
	topic_entrie := &TopicEntry{3,"First topic name from db","First topic link from db"}
	entries = append(entries, topic_entrie)
	topic_entrie = &TopicEntry{4,"Second topic name from db","Second topic link from db",}
	entries = append(entries, topic_entrie)
	return entries
}

func (g *Forum) GetTopicContent() string {
	return "Content from crawler"
	
}

func (g *Forum) GetTopicContentFromDB() string {
	return "Content from crawler"
	
}

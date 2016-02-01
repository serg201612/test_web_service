package crawler

import (
        "fmt"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)


type Forum_db struct {
	Id    bson.ObjectId `bson:"_id"`
        Ref string `bson:"ref"`
        Name string `bson:"name"`
}

type Topic_db struct {
	Id    bson.ObjectId `bson:"_id"`
        ForumRef string `bson:"forumref"`
        Ref string `bson:"ref"`
        Name string `bson:"name"`
}

type Content_db struct {
	Id    bson.ObjectId `bson:"_id"`
        TopicRef string `bson:"topicref"`
        Ref string `bson:"ref"`
        Content string `bson:"content"`
}

type Content_db_select struct {
	Id    string
        Ref string 
        Content string 
}

type Topic_db_select struct {
	Id    string
        Ref string 
        Name string 
}

type Forum_db_select struct {
	Id    string
        Ref string 
        Name string 
}


func db_content_select(topic_ref string, table_name string) ([]*ContentEntry) {
	
	contentData = make([]*ContentEntry, 0)

        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
	c := session.DB("test").C(table_name)
  	var results []Content_db_select
  	err = c.Find(bson.M{"topicref": topic_ref}).All(&results)
    
        if err != nil {
	  fmt.Println("error content select: ", err)
        } else {
	  for _, r := range results{
	    contententry := &ContentEntry{"id", r.Content, r.Ref}
	    contentData = append(contentData, contententry)
       	    //fmt.Printf("ref=%s name=%s\n", r.Ref, r.Name, )
	  }
	}  

	return contentData
	
}

func db_topic_select(forum_ref string, table_name string) ([]*TopicEntry) {
	
	topicData = make([]*TopicEntry, 0)

        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
	c := session.DB("test").C(table_name)
  	var results []Topic_db_select
  	err = c.Find(bson.M{"forumref": forum_ref}).All(&results)
    
        if err != nil {
	  fmt.Println("error forum select: ", err)
        } else {
	  for _, r := range results{
	    topicentry := &TopicEntry{"id", r.Name, r.Ref}
	    topicData = append(topicData, topicentry)
       	    //fmt.Printf("ref=%s name=%s\n", r.Ref, r.Name, )
	  }
	}  

	return topicData
	
}

func db_forum_select(table_name string) ([]*ForumEntry) {
	
	forumData = make([]*ForumEntry, 0)

        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
	c := session.DB("test").C(table_name)
  	var results []Forum_db_select
  	err = c.Find(nil).All(&results)
    
        if err != nil {
	  fmt.Println("error forum select: ", err)
        } else {
	  for _, r := range results{
	    forumentry := &ForumEntry{"id", r.Name, r.Ref}
	    forumData = append(forumData, forumentry)
       	    //fmt.Printf("id=%s ref=%s name=%s\n", r.Id, r.Ref, r.Name)
	  }
	}  

	return forumData
	
}


func db_forum_update(ref string, name string, table_name string) {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
        c := session.DB("test").C(table_name)
        forum := Forum_db{}
        err = c.Find(bson.M{"ref": ref}).One(&forum) 
        if err != nil {
		err = c.Insert(&Forum_db{bson.NewObjectId(), ref, name})
		if err != nil {
		  fmt.Println("error forum insert: ", err)
		} else {
		  //fmt.Println("insert forum : ", ref)
		}
        } else {
	  err = c.Update(bson.M{"_id": forum.Id}, bson.M{"ref": ref, "name": name})
	  if err != nil {
	    fmt.Println("error forum update: ", err)
	  } else {
	    //fmt.Println("find and update forum: name=", name)
	  }
	}  
	
}

func db_topic_update(forum_ref string, ref string, name string, table_name string) {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
        c := session.DB("test").C(table_name)
        topic := Topic_db{}
        err = c.Find(bson.M{"ref": ref}).One(&topic) 
        if err != nil {
		err = c.Insert(&Topic_db{bson.NewObjectId(), forum_ref, ref, name})
		if err != nil {
		  fmt.Println("error topic insert: ", err)
		} else {
		  //fmt.Println("insert forum : ", ref)
		}
        } else {
	  err = c.Update(bson.M{"_id": topic.Id}, bson.M{"forumref":forum_ref, "ref": ref, "name": name})
	  if err != nil {
	    fmt.Println("error topic update: ", err)
	  } else {
	    //fmt.Println("find and update forum: name=", name)
	  }
	}  
	
}

func db_content_update(topic_ref string, ref string, content string, table_name string) {
        session, err := mgo.Dial("localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})
        c := session.DB("test").C(table_name)
        topic := Content_db{}
        err = c.Find(bson.M{"topicref": topic_ref}).One(&topic) 
        if err != nil {
		err = c.Insert(&Content_db{bson.NewObjectId(), topic_ref, ref, content})
		if err != nil {
		  fmt.Println("error content insert: ", err)
		} else {
		  //fmt.Println("insert forum : ", ref)
		}
        } else {
	  err = c.Update(bson.M{"_id": topic.Id}, bson.M{"topicref":topic_ref, "ref": ref, "content": content})
	  if err != nil {
	    fmt.Println("error content update: ", err)
	  } else {
	    //fmt.Println("find and update forum: name=", name)
	  }
	}  
	
}

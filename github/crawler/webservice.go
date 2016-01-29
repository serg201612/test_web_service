package crawler

import (
	//"fmt"
	"encoding/json"
	//"io/ioutil"
	"net/http"
	//"strconv"
	"../martini"
)

type WebService interface {
	GetPath() string
	GetBoardsFromCrawler(params martini.Params) (int, string)
	GetTopicsFromCrawler(params martini.Params) (string)
	GetContentFromCrawler(params martini.Params) (string)
	GetBoardsFromDB(params martini.Params) (int, string)
	GetTopicsFromDB(params martini.Params) (string)
	GetContentFromDB(params martini.Params) (string)
}

func RegisterWebService(webService WebService,
	classicMartini *martini.ClassicMartini) {
	path := webService.GetPath()
	classicMartini.Get(path, webService.GetBoardsFromCrawler)
	classicMartini.Get(path+"/boardsfromcrawler", webService.GetBoardsFromCrawler)
	classicMartini.Get(path, webService.GetTopicsFromCrawler)
	classicMartini.Get(path+"/topicsfromcrawler/:id", webService.GetTopicsFromCrawler)
	classicMartini.Get(path, webService.GetContentFromCrawler)
	classicMartini.Get(path+"/contentfromcrawler/:id", webService.GetContentFromCrawler)
	classicMartini.Get(path, webService.GetBoardsFromDB)
	classicMartini.Get(path+"/boardsfromdb", webService.GetBoardsFromDB)
	classicMartini.Get(path, webService.GetTopicsFromDB)
	classicMartini.Get(path+"/topicsfromdb/:id", webService.GetTopicsFromDB)
	classicMartini.Get(path, webService.GetContentFromDB)
	classicMartini.Get(path+"/contentfromdb/:id", webService.GetContentFromDB)
}

func (g *Forum) GetPath() string {
	return "/forum"
}


func (g *Forum) GetBoardsFromCrawler(params martini.Params) (int, string) {
	encodedEntries, err := json.Marshal(g.GetAllEntries())
	if err != nil {
		return http.StatusInternalServerError, "internal WebGetBoardsFromCrawler error"
	}
	return http.StatusOK, string(encodedEntries)
	return http.StatusOK, ""
}

func (g *Forum) GetTopicsFromCrawler(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetAllTopicEntries(params))
	if err != nil {
		return "internal GetTopicsFromCrawler error"
	}
	return string(encodedEntries)
}

func (g *Forum) GetContentFromCrawler(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetContentsFromCrawler(params))
	if err != nil {
		return "internal GetContentFromCrawler error"
	}
	return string(encodedEntries)
}

func (g *Forum) GetBoardsFromDB(params martini.Params) (int, string) {
	encodedEntries, err := json.Marshal(g.GetAllEntriesFromDB())
	if err != nil {
	  return http.StatusInternalServerError, "internal GetBoardsFromDB error"
	}
	return http.StatusOK, string(encodedEntries)
}

func (g *Forum) GetTopicsFromDB(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetAllTopicEntriesFromDB(params))
	if err != nil {
		return "internal GetTopicsFromDB error"
	}
	return string(encodedEntries)
}

func (g *Forum) GetContentFromDB(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetContentsFromDB(params))
	if err != nil {
		return "internal GetContentFromDB error"
	}
	return string(encodedEntries)
}

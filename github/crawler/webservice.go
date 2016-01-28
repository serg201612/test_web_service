package crawler

import (
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
	classicMartini.Get(path+"/topicsfromcrawler", webService.GetTopicsFromCrawler)
	classicMartini.Get(path, webService.GetContentFromCrawler)
	classicMartini.Get(path+"/contentfromcrawler", webService.GetContentFromCrawler)
	classicMartini.Get(path, webService.GetBoardsFromDB)
	classicMartini.Get(path+"/boardsfromdb", webService.GetBoardsFromDB)
	classicMartini.Get(path, webService.GetTopicsFromDB)
	classicMartini.Get(path+"/topicsfromdb", webService.GetTopicsFromDB)
	classicMartini.Get(path, webService.GetContentFromDB)
	classicMartini.Get(path+"/contentfromdb", webService.GetContentFromDB)
}

func (g *Forum) GetPath() string {
	return "/forum"
}


func (g *Forum) GetBoardsFromCrawler(params martini.Params) (int, string) {
	//if len(params) == 0 {
		encodedEntries, err := json.Marshal(g.GetAllEntries())
		if err != nil {
			return http.StatusInternalServerError, "internal WebGetBoardsFromCrawler error"
		}

		return http.StatusOK, string(encodedEntries)
	//}

	return http.StatusOK, ""
	/*
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, "invalid entry id"
	}

	entry, err := g.GetEntry(id)
	if err != nil {
		return http.StatusNotFound, "entry not found"
	}

	encodedEntry, err := json.Marshal(entry)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	return http.StatusOK, string(encodedEntry)
	*/
}

func (g *Forum) GetTopicsFromCrawler(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetAllTopicEntries())
	if err != nil {
		return "internal WebGetTopics error"
	}

	return string(encodedEntries)
}

func (g *Forum) GetContentFromCrawler(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetTopicContent())
	if err != nil {
		return "internal WebGetTopicContent error"
	}

	// Return encoded entries.
	return string(encodedEntries)
}


func (g *Forum) GetBoardsFromDB(params martini.Params) (int, string) {
	if len(params) == 0 {
		encodedEntries, err := json.Marshal(g.GetAllEntriesFromDB())
		if err != nil {
			return http.StatusInternalServerError, "internal WebGetBoardsFromCrawler error"
		}

		return http.StatusOK, string(encodedEntries)
	}

	return http.StatusOK, ""
	/*
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, "invalid entry id"
	}

	entry, err := g.GetEntry(id)
	if err != nil {
		return http.StatusNotFound, "entry not found"
	}

	encodedEntry, err := json.Marshal(entry)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	return http.StatusOK, string(encodedEntry)
	*/
}

func (g *Forum) GetTopicsFromDB(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetAllTopicEntriesFromDB())
	if err != nil {
		return "internal WebGetTopics error"
	}

	return string(encodedEntries)
}

func (g *Forum) GetContentFromDB(params martini.Params) (string) {
	encodedEntries, err := json.Marshal(g.GetTopicContentFromDB())
	if err != nil {
		return "internal WebGetTopicContent error"
	}

	// Return encoded entries.
	return string(encodedEntries)
}

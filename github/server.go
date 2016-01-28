
package main

import (
	"./crawler"
	"./martini"
)

func main() {
	martiniClassic := martini.Classic()
	Crawler := crawler.NewForum()
	crawler.RegisterWebService(Crawler, martiniClassic)
	martiniClassic.Run()
}



package main

import (
	"fmt"
	"github.com/kmnemon/goproject/links"
	"log"
)

/*crawl the url from the website*/

func main() {
	var s []string = make([]string, 10)
	s[0] = "http://movie.douban.com"

	breadthFirst(crawl, s)

}

func breadthFirst(f func(item string) []string, worklist []string){
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items{
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
				}

			}
		}
	}


func crawl(url string) []string{
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list

}




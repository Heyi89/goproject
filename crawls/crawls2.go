package main

import (
	"fmt"
	"github.com/kmnemon/goproject/links"
	"log"
)

func main() {
	worklist := make(chan []string)
	//n caculate the number of worklist
	var n int


	url := []string{
		"http://www.baidu.com",
		"http://movie.douban.com",
		"http://www.baidu.com"}

	n++
	go func() { worklist <- url}()

	seen := make(map[string]bool)
	for ; n>0; n--{
		fmt.Println(n)
		list := <-worklist
		for _, link := range list {
			if !seen[link]{
				seen[link] = true
				n++
				go func(link string){
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}

//a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)


func crawl(url string) []string{
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens //release the token
	if err != nil {
		log.Print(err)
	}
	return list

}
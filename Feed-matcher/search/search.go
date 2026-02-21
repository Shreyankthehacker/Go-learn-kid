package search

import (
	"log"
	_ "log"
	"sync"
	_ "sync"
	"fmt"
)



var matchers = make(map[string]Matcher)


func Register(feedType string , matcher Matcher){
	if _,exists := matchers[feedType];
	exists{
		log.Fatal(feedType,"matcher already exists")
	}
	log.Print("Registering the matcher for the feedtype ",feedType)
	matchers[feedType] = matcher

}


func Run(searchTerm string){
	feeds , err:= RetrieveFeed();
	if err!=nil{
		log.Fatal(err)
	}



	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _,feed := range(feeds){
		matcher,exists := matchers[feed.Type]
		if !exists{
			fmt.Println(matcher)
			matcher = matchers["default"]
		}


		go func(matcher Matcher, feed *Feed){
			Match(matcher,feed,searchTerm,results);
			waitGroup.Done();
		}(matcher,feed)

	}
	go func(){
		waitGroup.Wait();
		close(results)
	}()

	DisplayResults(results)

}
package search


import ("log"
"fmt"
)


type Result struct{  // for result of a search
	Field string 
	Content string 
}


type Matcher interface{
	Search(feed *Feed,searchTerm string) ([]*Result,error)
}



func Match(matcher Matcher , feed *Feed , searchTerm string , results chan<- *Result){
	fmt.Println(">>>",feed,"<<<<")
	

	
	searchResults , err := matcher.Search(feed,searchTerm)
	if err!=nil{
		log.Println(err)
		return 
	}


	for _ , result := range(searchResults){
		results <- result   // write the current results to the channel
	}
	log.Println("done matcher")
}


func DisplayResults(results chan *Result){
	//Channels return one value per iteration, not index + value.
	for result:=range results{
		log.Printf("%s,\n\n\n%s\n\n",result.Field,result.Content);
	}
	// close(results) wrong it will throw panic bcz cant close channel from reciever end 
}
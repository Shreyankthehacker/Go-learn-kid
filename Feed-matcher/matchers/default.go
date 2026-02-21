package matchers

import (
	"feedmatcher/search"
)


type defaultMatcher struct{}


func init(){
	var matcher defaultMatcher
	search.Register("defuault",matcher)
}
func (m defaultMatcher)Search(feed *search.Feed , searchterm string)([]*search.Result , error){
	return nil,nil
}
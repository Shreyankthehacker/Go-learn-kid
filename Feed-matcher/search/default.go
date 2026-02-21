package search


type defaultMatcher struct{}


func init(){
	var matcher defaultMatcher
	Register("defuault",matcher)
}
func (m defaultMatcher)Search(feed *Feed , searchterm string)([]*Result , error){
	return nil,nil
}
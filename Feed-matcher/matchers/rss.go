package matchers

import (
	"encoding/xml"
	"errors"
	"feedmatcher/search"
	"fmt"
	"log"
	"net/http"
	"regexp"
)


type(


item struct {
	XMLName xml.Name `xml:"item"`
	PubDate string `xml:"pubDate"`
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	GUID string `xml:"guid"`
	GeoRssPoint string `xml:"georss:point"`
}

image struct{
	XMLName xml.Name `xml:"image"`
	URL string `xml:"url"`
	Title string `xml:"title"`
	Link string `xml:"link"`
}


channel struct {
	XMLName xml.Name `xml:"channel"`
	PubDate string `xml:"pubDate"`
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	LastBuildDate string `xml:"lastbuildDate"`
	TTL string `xml:"ttl"`
	Language string `xml:"langauge"`
	ManagingEditor string `xml:"managingEditor"`
	WebMaster string `xml:"webMaster"`
	Image image `xml:"image"`
	Item []item `xml:"item"`
}


rssDoc struct{
	XMLName xml.Name `xml:"rss"` 
	Channel channel `xml:"channel"`
}



)

type rssMatcher struct{}


func init(){
	var matcher rssMatcher
	search.Register("rss",matcher);
}



func (matcher rssMatcher) retriever(feed  *search.Feed)(*rssDoc , error){


	if feed.URI==""{
		return nil,errors.New("No rss uri provided")
	}

	resp,err:= http.Get(feed.URI)

	if err!=nil{
		return nil,err
	}

	defer resp.Body.Close();


	if resp.StatusCode!=200{
		return nil,fmt.Errorf("Somewhere response broke with status code %d\n",resp.StatusCode)
	}

	var doc rssDoc
	err = xml.NewDecoder(resp.Body).Decode(&doc)
	return &doc,err

}




func (matcher rssMatcher)Search(feed *search.Feed, searchItem string )([]*search.Result , error){


	fmt.Println("reached here")
	var results []*search.Result

	log.Printf("Search feed type [%s]  site [%s] for uri [%s]",feed.Type , feed.Name , feed.URI)

	doc , err := matcher.retriever(feed)
	if err != nil {
		return nil, err
	}

	for _,channelItem := range doc.Channel.Item{
		matched , err := regexp.MatchString(searchItem,channelItem.Title);
		if err!=nil{
			return nil,err;
		}
		if matched{
			results = append(results,&search.Result{Field: "Title",Content:channelItem.Title})
		}
		

		// for desc 

		matched,err = regexp.MatchString(searchItem,channelItem.Description)
		if err!=nil{
			return nil,err;
		}
		if matched{
			results = append(results,&search.Result{Field: "Description" , Content: channelItem.Description})
		}
	}

	return results,nil



}
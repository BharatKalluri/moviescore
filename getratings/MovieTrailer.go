package getratings

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"net/http"
)

func GetTrailer(mname string) string {
	urlis := "https://www.googleapis.com/youtube/v3/search?part=snippet&q=inception+movie+trailer&key=AIzaSyCXqzqRJUDGsQUO44y-G8PM_qxVpgd66GE"
	var AllVideoArray []string
	res, err := http.Get(urlis)
	if err != nil {
		fmt.Println(err)
	}
	v, err := jason.NewObjectFromReader(res.Body)
	ItemsArr, _ := v.GetObjectArray("items")
	for _, item := range ItemsArr {
		id, _ := item.GetObject("id")
		Vid, _ := id.GetString("videoId")
		AllVideoArray = append(AllVideoArray, Vid)
	}
	MovieTrailerUrl := "https://www.youtube.com/embed/" + AllVideoArray[0]
	return MovieTrailerUrl
}

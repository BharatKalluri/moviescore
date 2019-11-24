package getratings

//GetTrailer takes in a movie name as argument and searches youtube for "<movie name> movie trailer"
//and returns the first result
func GetTrailer(mname string) string {
	return ""
	// TODO: Get a youtube key and use here
	// MovieName := strings.Replace(mname, " ", "_", 9)
	// urlis := "https://www.googleapis.com/youtube/v3/search?part=snippet&q=" + MovieName + "+movie+trailer&key=AIzaSyCXqzqRJUDGsQUO44y-G8PM_qxVpgd66GE"
	// var AllVideoArray []string
	// res, err := http.Get(urlis)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// v, err := jason.NewObjectFromReader(res.Body)
	// if err != nil {
	// 	fmt.Println("An error occurred!")
	// }
	// ItemsArr, _ := v.GetObjectArray("items")
	// for _, item := range ItemsArr {
	// 	id, err := item.GetObject("id")
	// 	LogError(err)
	// 	Vid, err := id.GetString("videoId")
	// 	LogError(err)
	// 	AllVideoArray = append(AllVideoArray, Vid)
	// }
	// movieTrailerURL := "https://www.youtube.com/embed/" + AllVideoArray[0]
	// return movieTrailerURL
}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"encoding/json"
)

//var ratings [10]int;
var ratings = make(map[int][]int)

type rating_struct struct {
	Value int `json:"value"`
	Refid int `json:"refid"`
}

func RatingSimple(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//value , _ := strconv.Atoi(ps.ByName("value"))
	//if value >= 1 && value <=10{
	//	ratings[value-1]++;
	//}


}

func Rating(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var vote rating_struct
	err := decoder.Decode(&vote)
	if err != nil {
		log.Print(err);
	}
	if _, ok := ratings[vote.Refid]; !ok {
		ratings[vote.Refid] = make([]int, 10)
	}

	if vote.Value >= 1 && vote.Value <=10{
		ratings[vote.Refid][vote.Value-1]++;
		fmt.Fprint(w,"Thanks for rating");
	} else {
		fmt.Fprint(w, "Incorrect rating.")
	}
}

func ShowRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	for i := 0; i<10; i++ {
		fmt.Fprintf(w, "Rated with %d: %d \n", i+1, ratings[i])
	}

}



func main() {
	router := httprouter.New()
	router.GET("/ratingsimple/:value", RatingSimple)
	router.GET("/showrating", ShowRating)
	router.POST("/rating",Rating);

	log.Fatal(http.ListenAndServe(":8080", router))
}
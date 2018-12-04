package main

import (
    "log"
	"html/template"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const analyzeUrl = "http://localhost:5000/analyze"

func viewHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, nil)
}

type Listing struct {
    Id int64 `json:"id"`
    Site string `json:"site"`
    Bedrooms int64 `json:"bedrooms"`
    Bathrooms int64 `json:"bathrooms"`
    Parkings float64 `json:"parkings"`
    PropertyType float64 `json:"property_type"`
    Address string `json:"address"`
    Url int64 `json:"url"`
    Price int64 `json:"price"`
    EstimatedPrice string `json:"estimated_price"`
    Diff string `json:"diff"`
}

type AnalyzerAPIResponse struct {
    Listings []Listing `json:"listings"`
}

/*
Example response:
{
	"listings": [
		{
			"site": "domain",
			"bedrooms": 2,
			"bathrooms": 1,
			"parkings": 1,
			"property_type": "house",
			"address": "",
			"url": "",
			"price": 650000,
			"estimated_price": 590000,
			"diff": 9.23	
		}
	]
}
*/

func getListings(body []byte) (*AnalyzerAPIResponse, error) {
    var l = new(AnalyzerAPIResponse)
    err := json.Unmarshal(body, &l)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return l, err
}


func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(analyzeUrl)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
    	if err != nil {
        	panic(err.Error())
	}
	
	l, err := getListings([]byte(body))

	t, _ := template.ParseFiles("result.html")
	
    	t.Execute(w, &l)
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/analyze", analyzeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

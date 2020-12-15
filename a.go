package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RespInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json: "lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func main() {
	ip := flag.String("ip", "", "ip which you want to lookup.")
	flag.Parse()

	geoURL := "http://ip-api.com/json"
	para := "lang=zh-CN"

	resp, err := http.Get(geoURL + "/" + *ip + "?" + para)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var ipInfo RespInfo

	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s, %s, %s, %s\n", ipInfo.Query, ipInfo.Country, ipInfo.City, ipInfo.Isp)
}

package prayer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const endPoint = "http://api.aladhan.com/v1/calendarByCity"

type ApiResponse struct {
	Data []DayInfo `json:"data"`
}

type DayInfo struct {
	Timings Timeing `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}

type Timeing struct {
	Fajr     string `json:"Fajr"`
	Sunrise  string `json:"Sunrise"`
	Dhuhr    string `json:"Dhuhr"`
	Asr      string `json:"Asr"`
	Sunset   string `json:"Sunset"`
	Maghrib  string `json:"Maghrib"`
	Isha     string `json:"Isha"`
	Imsak    string `json:"Imsak"`
	Midnight string `json:"Midnight"`
}

type Date struct {
	Readable  string     `json:"readable"`
	Timestamp string     `json:"timestamp"`
	Gregorian Greogorian `json:"gregorian"`
	Hijri     Hijri      `json:"hijri"`
}

type Greogorian struct {
	Date        string      `json:"date"`
	Format      string      `json:"format"`
	Day         string      `json:"day"`
	Designation Designation `json:"designation"`
}

type Designation struct {
	Abbreviated string `json:"abbreviated"`
	Expanded    string `json:"expanded"`
}

type Hijri struct {
	Date        string        `json:"date"`
	Format      string        `json:"format"`
	Day         string        `json:"day"`
	Designation Designation   `json:"designation"`
	Holidays    []interface{} `json:"holidays"`
}

type Meta struct {
	Latitude                 float64    `json:"latitude"`
	Longitude                float64    `json:"longitude"`
	Timezone                 string     `json:"timezone"`
	Method                   Method     `json:"method"`
	LatitudeAdjustmentMethod string     `json:"latitudeAdjustmentMethod"`
	MidnightMode             string     `json:"midnightMode"`
	School                   string     `json:"school"`
	Offset                   MetaOffset `json:"offset"`
}

type MetaOffset struct {
	Imsak    int `json:"Imsak"`
	Fajr     int `json:"Fajr"`
	Sunrise  int `json:"Sunrise"`
	Dhuhr    int `json:"Dhuhr"`
	Asr      int `json:"Asr"`
	Maghrib  int `json:"Maghrib"`
	Sunset   int `json:"Sunset"`
	Isha     int `json:"Isha"`
	Midnight int `json:"Midnight"`
}

type Method struct {
	ID     int          `json:"id"`
	Name   string       `json:"name"`
	Params MethodParams `json:"params"`
}

type MethodParams struct {
	Fajr int `json:"Fajr"`
	Isha int `json:"Isha"`
}

func GetList(city string, country string) DayInfo {
	params := fmt.Sprintf("?city=%s&country=%s", city, country)
	uri := endPoint + params

	response, err := http.Get(uri)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var apiResponse ApiResponse
	json.Unmarshal(contents, &apiResponse)

	/**
		Because the api returns the whole month, we will loop over the data list until we found today date
		then we will respond the praying times to the user
	**/

	todayDate := time.Now().Format("02-01-2006")
	var todayTimes DayInfo
	found := false
	for _, row := range apiResponse.Data {
		if row.Date.Gregorian.Date == todayDate {
			todayTimes = row
			found = true
			break
		}
	}

	if !found {
		panic("cant find praying times for date " + todayDate)
	}

	return todayTimes
}

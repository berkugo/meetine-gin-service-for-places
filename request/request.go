package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	insert = "http://127.0.0.1:5010/checkin/add"
	getCheckins = "http://127.0.0.1:5010/checkin/get/university/"
	getPlaces = "http://127.0.0.1:5010/places/get/ids"
	getUserCheckins = "http://127.0.0.1:5010/checkin/get/user/"
)
func GetRequestForPersonalCheckins(userid string) (map[string]interface{}, error, int){

	resp, err := http.Get(getUserCheckins + userid)
	return reqJob(resp, err)
}

func SendGetRequestForAll(city string) (map[string]interface{}, error, int) {
	resp, err := http.Get(getCheckins + strings.ToLower(city))
	return reqJob(resp, err)
}

func GetCheckinPlaces(data map[string]interface{}) (map[string]interface{}, error) {

	var placesIdArray []interface{}
	for _, item := range data["message"].(map[string]interface{})["result"].([]interface{}){
		placesIdArray = append(placesIdArray, item.(map[string]interface{})["placeid"])
	}
	jsonValue, _:= json.Marshal(map[string]interface{}{"ids": placesIdArray})
	resp, err := http.Post(getPlaces, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			return result["message"].(map[string]interface{}), nil
		} else {
			return nil, err
		}

	}
}

func reqJob(resp *http.Response, err error) (map[string]interface{}, error, int){

	if resp != nil{
		defer resp.Body.Close()
		if err != nil{
			return map[string]interface{}{"data": "False"}, err, resp.StatusCode

		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return map[string]interface{}{"data": "False"}, err, resp.StatusCode
			} else {
				var result map[string]interface{}
				json.Unmarshal(body, &result)
				if len(result["message"].(map[string]interface{})["result"].([]interface{})) > 0{
					result["message"].(map[string]interface{})["places"], err = GetCheckinPlaces(result)
					return result, nil, resp.StatusCode
				} else {
					return map[string]interface{}{"data": "False"}, err, 404
				}
			}
		}
	} else {
		return map[string]interface{}{"data": "False"}, err, 404
	}
}


func SendPostRequestForInsert(checkin []byte) (string, error, int){

	resp, err := http.Post(insert, "application/json", bytes.NewBuffer(checkin))
	if resp != nil{
		defer resp.Body.Close()
		if err != nil{
			return "FALSE", err, resp.StatusCode

		} else {
			return "OK", nil, resp.StatusCode

		}
	}else{
		return "FALSE", err, 404
	}

}
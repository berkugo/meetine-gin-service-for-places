package models

import (
	"checkin/request"
	"encoding/json"
	"strings"
)

type Checkin struct {
	PlaceId string `json:"placeId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
	Note    string `json:"note" binding:"required"`
	Ctype   int    `json:"type" binding:"required"`
	City    string `json:"city" binding:"required"`
}

func (ins *Checkin) AddCheckin() (string, int) {

	if ins != nil {
		values := map[string]interface{}{"placeid": ins.PlaceId, "uid": ins.UserId, "note": ins.Note, "ctype": ins.Ctype, "city": strings.ToLower(ins.City)}
		jsonValue, _ := json.Marshal(values)
		resp, err, status := request.SendPostRequestForInsert(jsonValue)
		if err != nil {
			return "Error.", status
		} else {
			return resp, status
		}

	} else {
		return "Error.", 404

	}
}

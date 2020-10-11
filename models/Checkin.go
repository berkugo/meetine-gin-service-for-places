package models

var checkinListInstance *CheckinList = &CheckinList{}

type Checkin struct {
	PlaceId string `json:"placeId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
	Note    string `json:"note" binding:"required"`
	Ctype   int    `json:"type" binding:"required"`
}

type CheckinList struct {
	items []*Checkin
}

func GetInstance() *CheckinList {
	return checkinListInstance
}

func (ins *CheckinList) AddCheckin(item *Checkin) {

	ins.items = append(ins.items, item)
}

func (ins *CheckinList) GetCheckins() []*Checkin {
	return ins.items
}

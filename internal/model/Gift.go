package model

type Gifts struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type Record struct {
	Userid    string `json:"userid"`
	PullTime  int64  `json:"pullTime"`
	GiftCodes string `json:"giftCodes"`
}

type GiftCodes struct {
	CreateUserId  string   `json:"createUserId"`
	CreateTime    int64    `json:"createTime"`
	GiftDescribe  string   `json:"giftDescribe"`
	GiftList      *[]Gifts `json:"giftList"`
	GiftCodeType  string   `json:"giftCodeType"`
	GiftPullNum   int      `json:"giftPullNum"`
	Validity      int64    `json:"validity"` //有效期
	GiftPulledNum int      `json:"giftPulledNum"`
	GiftCode      string   `json:"giftCode"`
	recordList    []Record `json:"recordList"`
	GiftPullUser  string   `json:"giftPullUser"` //领取人 限A类礼品码
}

type User struct {
	UserId string
	Name   string
}

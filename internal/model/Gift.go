package model

type Gifts struct {
	Name string `json:"name"` //礼品名
	Num  int    `json:"num"`  //礼品数量
}

type Record struct {
	Userid      string `json:"userid"`      //用户唯一标识id
	PullTime    int64  `json:"pullTime"`    //领取时间戳
	PullTimeStr string `json:"pullTimeStr"` //领取时间字符串
	//GiftCodes string `json:"giftCodes"`
}

type GiftCodes struct {
	CreateUserId  string   `json:"createUserId"`  //创建人唯一标识符id
	CreateTime    int64    `json:"createTime"`    //创建时间戳
	GiftDescribe  string   `json:"giftDescribe"`  //礼品码描述
	GiftList      []Gifts  `json:"giftList"`      //礼品码对应的礼品信息
	GiftCodeType  string   `json:"giftCodeType"`  //礼品码类型
	GiftPullNum   int      `json:"giftPullNum"`   //礼品码可领次数
	ValidityStr   string   `json:"validityStr"`   //礼品码有效时间字符串
	Validity      int64    `json:"validity"`      //有效期时间戳
	GiftPulledNum int      `json:"giftPulledNum"` //礼品码被领取次数
	GiftCode      string   `json:"giftCode"`      //礼品码
	RecordList    []Record `json:"recordList"`    //领取记录
	GiftPullUser  string   `json:"giftPullUser"`  //限制领取人
}

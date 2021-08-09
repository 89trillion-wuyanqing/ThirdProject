package handler

import (
	"ThirdProject/internal/model"
	"ThirdProject/internal/service"
	utils2 "ThirdProject/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"

	"time"
)

type GiftCodeshandler struct {
}

/**
创建礼品码业务处理
*/
func (this *GiftCodeshandler) CreateGiftCodes(giftCodes *model.GiftCodes) model.Result {
	giftCodes.GiftCode = new(utils2.RandomCode).RandomString()
	giftCodes.CreateTime = time.Now().Unix()
	giftCodes.GiftPulledNum = 0
	giftService := service.GiftCodesService{}
	result := giftService.ValPullNum(giftCodes)
	var valErr model.Result
	if result != valErr {
		return result
	}
	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
	}

	r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	if r != nil {
		return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
	}
	return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftCode}
}

/**
获取礼品码信息业务处理
*/
func (this *GiftCodeshandler) GetCiftCodes(giftCode string) model.Result {
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		if r == redis.Nil {
			return model.Result{Code: "214", Msg: "redis中不存在该礼品码", Data: nil}
		} else {
			return model.Result{Code: "215", Msg: "redis获取数据失败", Data: nil}
		}

	}
	giftCodes := &model.GiftCodes{}
	err := json.Unmarshal([]byte(result), giftCodes)
	if err != nil {
		return model.Result{Code: "202", Msg: "后台反序列化出错", Data: nil}
	}
	return model.Result{Code: "200", Msg: "成功", Data: giftCodes}
}

/**
、激活礼品码业务处理
*/
func (this *GiftCodeshandler) ActivateCode(giftCode string, userId string) model.Result {
	//先验证验证码是否存在
	giftCodes := &model.GiftCodes{}
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		if r == redis.Nil {
			return model.Result{Code: "214", Msg: "redis中不存在该礼品码", Data: nil}
		} else {
			return model.Result{Code: "215", Msg: "redis获取数据失败", Data: nil}
		}
	}

	err := json.Unmarshal([]byte(result), giftCodes)
	if err != nil {
		return model.Result{Code: "202", Msg: "后台反序列化出错", Data: nil}
	}
	//先验证礼品码是否过期
	CurrentTime := time.Now().Unix()

	if CurrentTime > giftCodes.Validity {
		return model.Result{Code: "216", Msg: "该礼品码已过期", Data: nil}
	}
	//验证验证码是哪一类验证码
	if giftCodes.GiftCodeType == "A" {
		//查看限制人
		if giftCodes.GiftPullUser == userId {
			//查看可领取次数
			if giftCodes.GiftPullNum >= 1 {
				//领取
				/*giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
				}
				return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftList}*/
				return UpdateGift(giftCodes, userId)
			} else {
				//已领取过
				return model.Result{Code: "217", Msg: "礼品码已领取", Data: nil}
			}

		} else {
			//不是该领取码的限制人
			return model.Result{Code: "218", Msg: "你不可使用该礼品码", Data: nil}
		}
	} else if giftCodes.GiftCodeType == "B" {
		//不限用户  不限次数，用户是否用过
		//先判断可领次数是否可以领取
		//领取礼品
		if giftCodes.GiftPullNum > 0 {
			records := giftCodes.RecordList
			if len(records) <= 0 {
				//没有领取记录，则增加一条领取记录
				//可以领取礼品
				//增加领取记录
				//领取
				/*giftCodes.GiftPullNum -= 1
				giftCodes.GiftPulledNum += 1
				list := giftCodes.RecordList
				m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
				list = append(list, m1)

				giftCodes.RecordList = list

				jsonStr, err := json.Marshal(giftCodes)
				if err != nil {
					return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
				}
				r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
				if r != nil {
					return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
				}
				return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftList}*/
				return UpdateGift(giftCodes, userId)
			} else {
				//有领取记录 使用查看是否领取过
				for i, v := range records {
					if v.Userid == userId {
						//领取列表存在该用户领取记录
						return model.Result{Code: "217", Msg: "礼品码已领取", Data: nil}
						break
					}
					if i == len(records)-1 {
						//可以领取礼品
						//增加领取记录
						//领取
						/*giftCodes.GiftPullNum -= 1
						giftCodes.GiftPulledNum += 1
						list := giftCodes.RecordList
						m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
						list = append(list, m1)
						giftCodes.RecordList = list

						jsonStr, err := json.Marshal(giftCodes)
						if err != nil {
							return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
						}
						r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
						if r != nil {
							return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
						}
						return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftList}*/
						return UpdateGift(giftCodes, userId)
					}
				}

			}

		} else {
			return model.Result{Code: "219", Msg: "该礼品码已被领取完", Data: nil}
		}

	} else if giftCodes.GiftCodeType == "C" {
		records := giftCodes.RecordList
		if len(records) <= 0 {
			//没有领取记录，则增加一条领取记录
			//可以领取礼品
			//增加领取记录
			//领取

			/*giftCodes.GiftPulledNum += 1
			list := giftCodes.RecordList
			m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
			list = append(list, m1)

			giftCodes.RecordList = list
			jsonStr, err := json.Marshal(giftCodes)
			if err != nil {
				return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
			}
			r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
			if r != nil {
				return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
			}
			return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftList}*/
			return UpdateGift(giftCodes, userId)
		} else {
			//有领取记录 使用查看是否领取过
			for i, v := range records {
				if v.Userid == userId {
					//领取列表存在该用户领取记录
					return model.Result{Code: "217", Msg: "礼品码已领取", Data: nil}
					break
				}
				if i == len(records)-1 {
					//可以领取礼品
					//增加领取记录
					//领取

					return UpdateGift(giftCodes, userId)

				}
			}

		}

	}

	return model.Result{Code: "220", Msg: "礼品码无效", Data: nil}
}

func UpdateGift(giftCodes *model.GiftCodes, userId string) model.Result {

	//redis事务开始
	// 开启一个TxPipeline事务
	pipe := utils2.Rdb.TxPipeline()

	// 执行事务操作，可以通过pipe读写redis
	_ = pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)
	if giftCodes.GiftCodeType == "C" {

	} else {
		giftCodes.GiftPullNum -= 1
	}

	giftCodes.GiftPulledNum += 1
	list := giftCodes.RecordList
	m1 := model.Record{Userid: userId, PullTime: time.Now().Unix(), PullTimeStr: time.Now().Format("2006-01-02 15:04:05")}
	list = append(list, m1)
	giftCodes.RecordList = list
	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		return model.Result{Code: "212", Msg: "后台数据序列化出错", Data: nil}
	}
	r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	if r != nil {
		pipe.Discard()
		return model.Result{Code: "213", Msg: "redis存储失败", Data: nil}
	}
	//提交事务
	_, e1 := pipe.Exec()
	if e1 != nil {
		//pipe.Discard()
		fmt.Println(e1.Error())
	}
	return model.Result{Code: "200", Msg: "成功", Data: giftCodes.GiftList}

}

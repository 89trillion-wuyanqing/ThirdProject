package service

import (
	"ThirdProject/internal/model"
	"errors"
)

type GiftCodesService struct {
}

//验证创建礼品码
func (this *GiftCodesService) valPullNum(giftCodes model.GiftCodes) (bool, error) {
	if giftCodes.CreateUserId == "" {
		return false, errors.New("创建人不能为空！")
	}

	if len(giftCodes.GiftList) <= 0 {
		return false, errors.New("礼品内容不能为空！")
	}
	for _, v := range giftCodes.GiftList {
		if v.Name == "" {
			return false, errors.New("礼品内容中礼品名不能为空！")
		}
		if v.Num <= 0 {
			return false, errors.New("礼品内容中礼品数量不能小于0！")
		}
	}

	if giftCodes.Validity == 0 {
		return false, errors.New("有效时间不能为空！")
	}
	if giftCodes.Validity <= giftCodes.CreateTime {
		return false, errors.New("有效时间不能小于创建时间！")
	}
	if giftCodes.GiftCodeType == "" {
		return false, errors.New("请选择礼品码类别！")
	}
	if giftCodes.GiftCodeType == "A" {
		if giftCodes.GiftPullUser == "" {
			return false, errors.New("领取人不能为空！")
		}
		giftCodes.GiftPullNum = 1 //A类码限制一个人领一次

	} else if giftCodes.GiftCodeType == "B" {
		if giftCodes.GiftPullNum == 0 {
			return false, errors.New("可领取次数不能为空！")
		}
	}

	return true, nil
}
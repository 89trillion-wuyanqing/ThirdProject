package handler

import (
	"ThirdProject/internal/model"
	utils2 "ThirdProject/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

type GiftCodeshandler struct {
}

func (this *GiftCodeshandler) CreateGiftCodes(giftCodes model.GiftCodes) (bool, error) {
	giftCodes.GiftCode = new(utils2.RandomCode).RandomString()
	giftCodes.CreateTime = time.Now().Unix()
	giftCodes.GiftPulledNum = 0
	jsonStr, err := json.Marshal(giftCodes)
	if err != nil {
		return false, err
	}
	r := utils2.StringPush(giftCodes.GiftCode, string(jsonStr), 0)
	if r != nil {
		return false, r
	}
	return true, nil
}

func (this *GiftCodeshandler) GetCiftCodes(giftCode string) (model.GiftCodes, error) {
	result, r := utils2.StringPull(giftCode)
	if r != nil {
		return nil, r
	}
	giftCodes := &model.GiftCodes{}
	json.Unmarshal([]byte(result), giftCodes)
}

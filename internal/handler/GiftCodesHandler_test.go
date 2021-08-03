package handler

import (
	"ThirdProject/internal/model"
	utils2 "ThirdProject/internal/utils"
	"fmt"
	"testing"
)

func TestGiftCodeshandler_ActivateCode(t *testing.T) {
	rediserror := utils2.InitClient()
	if rediserror != nil {
		fmt.Println("连接失败")
		t.Fatal("redis服务连接失败")
	}
	giftHandler := GiftCodeshandler{}

	result := giftHandler.ActivateCode("5F589BDI", "100001")

	t.Log(result)
}

func TestGiftCodeshandler_CreateGiftCodes(t *testing.T) {
	rediserror := utils2.InitClient()
	if rediserror != nil {
		fmt.Println("连接失败")
		t.Fatal("redis服务连接失败")
	}
	giftHandler := GiftCodeshandler{}
	gifts := []model.Gifts{}
	gifts = append(gifts, model.Gifts{Name: "士兵", Num: 11})
	gifts = append(gifts, model.Gifts{Name: "金币", Num: 11})
	var giftCodes = &model.GiftCodes{GiftCodeType: "A", GiftPullUser: "100001", GiftList: gifts, CreateUserId: "10001", ValidityStr: "2021-08-12 02:03:45"}
	result := giftHandler.CreateGiftCodes(giftCodes)
	t.Log(result)

}

func TestGiftCodeshandler_GetCiftCodes(t *testing.T) {
	rediserror := utils2.InitClient()
	if rediserror != nil {
		fmt.Println("连接失败")
		t.Fatal("redis服务连接失败")
	}
	giftHandler := GiftCodeshandler{}

	result := giftHandler.GetCiftCodes("5F589BDI")
	t.Log(result)

}

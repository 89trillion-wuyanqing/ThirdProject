package ctrl

import "github.com/gin-gonic/gin"

type GiftCodeController struct {
}

func (this *GiftCodeController) CreateGiftCodes() gin.HandlerFunc {
	return func(context *gin.Context) {
		rarity, _ := context.GetPostForm("createUserId")
	}
}

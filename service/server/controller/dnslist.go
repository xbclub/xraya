package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xbclub/xraya/common"
	"github.com/xbclub/xraya/db/configure"
	"github.com/xbclub/xraya/server/service"
)

func PutDnsList(ctx *gin.Context) {
	var data struct {
		Internal string `json:"internal"`
		External string `json:"external"`
	}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	if len(data.Internal) == 0 && len(data.External) != 0 {
		common.ResponseError(ctx, logError("internal dns servers cannot be empty"))
		return
	}
	internal, err := service.RefineDnsList(data.Internal)
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("internal dns servers: %w", err)))
		return
	}
	external, err := service.RefineDnsList(data.External)
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("external dns servers: %w", err)))
		return
	}
	if err = configure.SetInternalDnsList(&internal); err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	if err = configure.SetExternalDnsList(&external); err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func GetDnsList(ctx *gin.Context) {
	common.ResponseSuccess(ctx, gin.H{
		"internal": configure.GetInternalDnsListNotNil(),
		"external": configure.GetExternalDnsListNotNil(),
	})
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	"time"
)

type CsOperatorSearch struct {
	operator.CsOperator
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CsOperatorApisReq struct {
	OperatorID uint   `json:"operatorID" form:"operatorID"`
	ApiIDs     []uint `json:"apiIds" form:"apiIds"`
}

type CsOperatorMenusReq struct {
	OperatorID uint   `json:"operatorID" form:"operatorID" `
	MenuIDs    []uint `json:"menuIds" form:"menuIds"`
}

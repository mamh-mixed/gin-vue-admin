package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type CsOperatorApi struct {
}

var csOperatorService = service.ServiceGroupApp.OperatorServiceGroup.CsOperatorService

// CreateCsOperator 创建CsOperator
// @Tags CsOperator
// @Summary 创建CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperator true "创建CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperator/createCsOperator [post]
func (csOperatorApi *CsOperatorApi) CreateCsOperator(c *gin.Context) {
	var csOperator operator.CsOperator
	tenantID := utils.GetTenantID(c)
	err := c.ShouldBindJSON(&csOperator)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csOperator.TenantID = tenantID
	if err := csOperatorService.CreateCsOperator(&csOperator); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsOperator 删除CsOperator
// @Tags CsOperator
// @Summary 删除CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperator true "删除CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperator/deleteCsOperator [delete]
func (csOperatorApi *CsOperatorApi) DeleteCsOperator(c *gin.Context) {
	var csOperator operator.CsOperator
	err := c.ShouldBindJSON(&csOperator)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorService.DeleteCsOperator(csOperator); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsOperatorByIds 批量删除CsOperator
// @Tags CsOperator
// @Summary 批量删除CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csOperator/deleteCsOperatorByIds [delete]
func (csOperatorApi *CsOperatorApi) DeleteCsOperatorByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := csOperatorService.DeleteCsOperatorByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsOperator 更新CsOperator
// @Tags CsOperator
// @Summary 更新CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperator true "更新CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperator/updateCsOperator [put]
func (csOperatorApi *CsOperatorApi) UpdateCsOperator(c *gin.Context) {
	var csOperator operator.CsOperator
	err := c.ShouldBindJSON(&csOperator)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorService.UpdateCsOperator(csOperator); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsOperator 用id查询CsOperator
// @Tags CsOperator
// @Summary 用id查询CsOperator
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operator.CsOperator true "用id查询CsOperator"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperator/findCsOperator [get]
func (csOperatorApi *CsOperatorApi) FindCsOperator(c *gin.Context) {
	var csOperator operator.CsOperator
	err := c.ShouldBindQuery(&csOperator)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsOperator, err := csOperatorService.GetCsOperator(csOperator.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsOperator": recsOperator}, c)
	}
}

// GetCsOperatorList 分页获取CsOperator列表
// @Tags CsOperator
// @Summary 分页获取CsOperator列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operatorReq.CsOperatorSearch true "分页获取CsOperator列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperator/getCsOperatorList [get]
func (csOperatorApi *CsOperatorApi) GetCsOperatorList(c *gin.Context) {
	var pageInfo operatorReq.CsOperatorSearch
	tenantID := utils.GetTenantID(c)
	err := c.ShouldBindQuery(&pageInfo)
	pageInfo.TenantID = tenantID
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csOperatorService.GetCsOperatorInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (csOperatorApi *CsOperatorApi) SetOperatorApis(c *gin.Context) {
	var operatorApis operatorReq.CsOperatorApisReq
	err := c.ShouldBindJSON(&operatorApis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorService.SetOperatorApis(operatorApis); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (csOperatorApi *CsOperatorApi) SetOperatorMenus(c *gin.Context) {
	var operatorMenus operatorReq.CsOperatorMenusReq
	err := c.ShouldBindJSON(&operatorMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorService.SetOperatorMenus(operatorMenus); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (csOperatorApi *CsOperatorApi) CreateCsOperatorAdmin(c *gin.Context) {
	operatorStr := c.Query("operatorID")
	operatorID, _ := strconv.Atoi(operatorStr)
	tenantID := utils.GetTenantID(c)
	if err := csOperatorService.CreateOperatorAdmin(tenantID, uint(operatorID)); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (csOperatorApi *CsOperatorApi) GetApisByOperatorID(c *gin.Context) {
	id := c.Query("operatorID")
	if apis, err := csOperatorService.GetOperatorApis(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(apis, c)
	}
}

func (csOperatorApi *CsOperatorApi) GetMenusByOperatorID(c *gin.Context) {
	id := c.Query("operatorID")
	if menus, err := csOperatorService.GetOperatorMenus(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(menus, c)
	}
}

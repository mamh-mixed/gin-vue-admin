package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type CsTenantApi struct {
}

var csTenantService = service.ServiceGroupApp.TenantServiceGroup.CsTenantService

// CreateCsTenant 创建CsTenant
// @Tags CsTenant
// @Summary 创建CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenant true "创建CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenant/createCsTenant [post]
func (csTenantApi *CsTenantApi) CreateCsTenant(c *gin.Context) {
	var csTenant tenant.CsTenant
	err := c.ShouldBindJSON(&csTenant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(csTenant, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.CreateCsTenant(&csTenant); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (csTenantApi *CsTenantApi) CreateCsTenantAdmin(c *gin.Context) {
	tenantStr := c.Query("tenantID")
	tenantID, _ := strconv.Atoi(tenantStr)
	if err := csTenantService.CreateTenantAdmin(uint(tenantID)); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsTenant 删除CsTenant
// @Tags CsTenant
// @Summary 删除CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenant true "删除CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenant/deleteCsTenant [delete]
func (csTenantApi *CsTenantApi) DeleteCsTenant(c *gin.Context) {
	var csTenant tenant.CsTenant
	err := c.ShouldBindJSON(&csTenant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.DeleteCsTenant(csTenant); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsTenantByIds 批量删除CsTenant
// @Tags CsTenant
// @Summary 批量删除CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csTenant/deleteCsTenantByIds [delete]
func (csTenantApi *CsTenantApi) DeleteCsTenantByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.DeleteCsTenantByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsTenant 更新CsTenant
// @Tags CsTenant
// @Summary 更新CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenant true "更新CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csTenant/updateCsTenant [put]
func (csTenantApi *CsTenantApi) UpdateCsTenant(c *gin.Context) {
	var csTenant tenant.CsTenant
	err := c.ShouldBindJSON(&csTenant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(csTenant, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.UpdateCsTenant(csTenant); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsTenant 用id查询CsTenant
// @Tags CsTenant
// @Summary 用id查询CsTenant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenant.CsTenant true "用id查询CsTenant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csTenant/findCsTenant [get]
func (csTenantApi *CsTenantApi) FindCsTenant(c *gin.Context) {
	var csTenant tenant.CsTenant
	err := c.ShouldBindQuery(&csTenant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsTenant, err := csTenantService.GetCsTenant(csTenant.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsTenant": recsTenant}, c)
	}
}

// GetCsTenantList 分页获取CsTenant列表
// @Tags CsTenant
// @Summary 分页获取CsTenant列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenantReq.CsTenantSearch true "分页获取CsTenant列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenant/getCsTenantList [get]
func (csTenantApi *CsTenantApi) GetCsTenantList(c *gin.Context) {
	var pageInfo tenantReq.CsTenantSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csTenantService.GetCsTenantInfoList(pageInfo); err != nil {
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

func (csTenantApi *CsTenantApi) SetTenantApis(c *gin.Context) {
	var tenantApis tenantReq.CsTenantApisReq
	err := c.ShouldBindJSON(&tenantApis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.SetTenantApis(tenantApis); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (csTenantApi *CsTenantApi) SetTenantMenus(c *gin.Context) {
	var tenantMenus tenantReq.CsTenantMenusReq
	err := c.ShouldBindJSON(&tenantMenus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantService.SetTenantMenus(tenantMenus); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (csTenantApi *CsTenantApi) GetApisByTenantID(c *gin.Context) {
	id := c.Query("tenantID")
	if apis, err := csTenantService.GetTenantApis(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(apis, c)
	}
}

func (csTenantApi *CsTenantApi) GetMenusByTenantID(c *gin.Context) {
	id := c.Query("tenantID")
	if menus, err := csTenantService.GetTenantMenus(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(menus, c)
	}
}

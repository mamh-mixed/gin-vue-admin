package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsTenantRouter struct {
}

// InitCsTenantRouter 初始化 CsTenant 路由信息
func (s *CsTenantRouter) InitCsTenantRouter(Router *gin.RouterGroup) {
	csTenantRouter := Router.Group("csTenant").Use(middleware.OperationRecord())
	csTenantRouterWithoutRecord := Router.Group("csTenant")
	var csTenantApi = v1.ApiGroupApp.TenantApiGroup.CsTenantApi
	{
		csTenantRouter.POST("createCsTenant", csTenantApi.CreateCsTenant)             // 新建CsTenant
		csTenantRouter.DELETE("deleteCsTenant", csTenantApi.DeleteCsTenant)           // 删除CsTenant
		csTenantRouter.DELETE("deleteCsTenantByIds", csTenantApi.DeleteCsTenantByIds) // 批量删除CsTenant
		csTenantRouter.PUT("updateCsTenant", csTenantApi.UpdateCsTenant)              // 更新CsTenant
		csTenantRouter.PUT("setTenantApis", csTenantApi.SetTenantApis)                // 系统分配租户api
		csTenantRouter.PUT("setTenantMenus", csTenantApi.SetTenantMenus)              // 系统分配租户菜单
		csTenantRouter.GET("getApisByTenantID", csTenantApi.GetApisByTenantID)        // 系统获取租户api
		csTenantRouter.GET("getMenusByTenantID", csTenantApi.GetMenusByTenantID)      // 系统获取租户菜单
		csTenantRouter.POST("createCsTenantAdmin", csTenantApi.CreateCsTenantAdmin)   // 系统创建租户管理员
	}
	{
		csTenantRouterWithoutRecord.GET("findCsTenant", csTenantApi.FindCsTenant)       // 根据ID获取CsTenant
		csTenantRouterWithoutRecord.GET("getCsTenantList", csTenantApi.GetCsTenantList) // 获取CsTenant列表
	}
}

package tenant

import (
	"errors"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"strconv"
)

type CsTenantService struct {
}

var casbinServer = new(system2.CasbinService)
var menuServer = new(system2.MenuService)

func (csTenantService *CsTenantService) CreateTenantIDCasbin(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetCasbinName(tenantID, 0)).AutoMigrate(&adapter.CasbinRule{})
	return err
}

func (csTenantService *CsTenantService) CreateTenantIDAuthMenu(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthMenuTableName(tenantID, 0)).AutoMigrate(&system.SysAuthorityMenu{})
	return err
}

func (csTenantService *CsTenantService) CreateTenetUserTableName(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserTableName(tenantID, 0)).AutoMigrate(&system.SysUserTable{})
	return err
}

func (csTenantService *CsTenantService) CreateTenetAuthsTable(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthsTable(tenantID, 0)).AutoMigrate(&system.SysAuthorityTable{})
	return err
}

func (csTenantService *CsTenantService) CreateUserAuthorityTable(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserAuthorityTableName(tenantID, 0)).AutoMigrate(&system.SysUserAuthority{})
	return err
}

var sysUserServer = new(system2.UserService)
var sysAuthorityServer = new(system2.AuthorityService)

// CreateCsTenant 创建CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) CreateCsTenant(csTenant *tenant.CsTenant) (err error) {
	csTenant.OnlyKey, _ = uuid.NewV4()
	if global.GVA_DB.First(&tenant.CsTenant{}, "username = ?", csTenant.Username).Error == nil {
		return errors.New("商户已注册")
	}
	err = global.GVA_DB.Create(csTenant).Error
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenetUserTableName(csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenetAuthsTable(csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateUserAuthorityTable(csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenantIDCasbin(csTenant.ID)
	if err != nil {
		return err
	}
	err = csTenantService.CreateTenantIDAuthMenu(csTenant.ID)
	return err
}

func (cstenantService *CsTenantService) CreateTenantAdmin(tenantID uint) (err error) {
	user := &system.SysUser{
		Username: "admin",
		Password: "123456",
	}
	user.TenantID = tenantID
	user.NickName = "租户管理员"
	user.AuthorityId = 888
	user.Authorities = []system.SysAuthority{{AuthorityId: 888}}
	_, err = sysUserServer.Register(*user, tenantID, 0)
	if err != nil {
		return err
	}

	auth := &system.SysAuthority{
		AuthorityId:   888,
		AuthorityName: "租户管理员",
		ParentId:      utils.Pointer(uint(0)),
		DefaultRouter: "dashboard",
	}

	_, err = sysAuthorityServer.CreateAuthority(*auth, tenantID, 0)
	if err != nil {
		return err
	}
	apis, err := cstenantService.GetTenantApis(strconv.Itoa(int(tenantID)))
	if err != nil {
		return err
	}
	var baseApis []system.SysApi

	var apiIDs []uint
	for _, api := range apis {
		apiIDs = append(apiIDs, api.ApiId)
	}

	global.GVA_DB.Find(&baseApis, "id in ?", apiIDs)

	var casbinInfo []systemReq.CasbinInfo
	for _, api := range baseApis {
		casbinInfo = append(casbinInfo, systemReq.CasbinInfo{
			Path:   api.Path,
			Method: api.Method,
		})
	}

	err = casbinServer.UpdateCasbin(888, casbinInfo, tenantID, 0)
	if err != nil {
		return err
	}
	menus, err := cstenantService.GetTenantMenus(strconv.Itoa(int(tenantID)))
	var menuIDs []uint
	for _, menu := range menus {
		menuIDs = append(menuIDs, menu.MenuId)
	}
	var baseMenus []system.SysBaseMenu
	global.GVA_DB.Find(&baseMenus, "id in ?", menuIDs)
	err = menuServer.AddMenuAuthority(baseMenus, 888, tenantID, 0)
	return err
}

// DeleteCsTenant 删除CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) DeleteCsTenant(csTenant tenant.CsTenant) (err error) {
	err = global.GVA_DB.Delete(&csTenant).Error
	return err
}

// DeleteCsTenantByIds 批量删除CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) DeleteCsTenantByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]tenant.CsTenant{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCsTenant 更新CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) UpdateCsTenant(csTenant tenant.CsTenant) (err error) {
	err = global.GVA_DB.Save(&csTenant).Error
	return err
}

// GetCsTenant 根据id获取CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) GetCsTenant(id uint) (csTenant tenant.CsTenant, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csTenant).Error
	return
}

// GetCsTenantInfoList 分页获取CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) GetCsTenantInfoList(info tenantReq.CsTenantSearch) (list []tenant.CsTenant, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tenant.CsTenant{})
	var csTenants []tenant.CsTenant
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.StartEndTime != nil && info.EndEndTime != nil {
		db = db.Where("end_time BETWEEN ? AND ? ", info.StartEndTime, info.EndEndTime)
	}
	if info.Allocation != nil {
		db = db.Where("allocation = ?", info.Allocation)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&csTenants).Error
	return csTenants, total, err
}

func (csTenantService *CsTenantService) SetTenantApis(tenantApis tenantReq.CsTenantApisReq) error {
	var csTenantApis []tenant.CsTenantApis
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("tenant_id = ?", tenantApis.TenantID).Delete(&csTenantApis).Error
		if err != nil {
			return err
		}
		for _, apiId := range tenantApis.ApiIDs {
			csTenantApis = append(csTenantApis, tenant.CsTenantApis{
				TenantID: tenantApis.TenantID,
				ApiId:    apiId,
			})
		}
		return global.GVA_DB.Create(&csTenantApis).Error
	})
}

func (csTenantService *CsTenantService) GetTenantApis(id string) (csTenantApis []tenant.CsTenantApis, err error) {
	err = global.GVA_DB.Where("tenant_id = ?", id).Find(&csTenantApis).Error
	return
}

func (csTenantService *CsTenantService) SetTenantMenus(tenantMenus tenantReq.CsTenantMenusReq) error {
	var csTenantMenus []tenant.CsTenantMenus
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("tenant_id = ?", tenantMenus.TenantID).Delete(&csTenantMenus).Error
		if err != nil {
			return err
		}
		for _, menuID := range tenantMenus.MenuIDs {
			csTenantMenus = append(csTenantMenus, tenant.CsTenantMenus{
				TenantID: tenantMenus.TenantID,
				MenuId:   menuID,
			})
		}
		return global.GVA_DB.Create(&csTenantMenus).Error
	})
}

func (csTenantService *CsTenantService) GetTenantMenus(id string) (csTenantMenus []tenant.CsTenantMenus, err error) {
	err = global.GVA_DB.Where("tenant_id = ?", id).Find(&csTenantMenus).Error
	return
}

func (csTenantService *CsTenantService) GetTenantIDByOnlyKey(onlyKey string) uint {
	var csTenant tenant.CsTenant
	global.GVA_DB.Where("only_key = ?", onlyKey).First(&csTenant)
	return csTenant.ID
}

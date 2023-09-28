package operator

import (
	"errors"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	tenant2 "github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"strconv"
)

var casbinServer = new(system2.CasbinService)
var menuServer = new(system2.MenuService)

type CsOperatorService struct {
}

var sysUserServer = new(system2.UserService)
var sysAuthorityServer = new(system2.AuthorityService)

func (csOperatorService *CsOperatorService) CreateOperatorIDCasbin(tenantID uint, operatorID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetCasbinName(tenantID, operatorID)).AutoMigrate(&adapter.CasbinRule{})
	return err
}

func (csOperatorService *CsOperatorService) CreateOperatorIDAuthMenu(tenantID uint, operatorID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthMenuTableName(tenantID, operatorID)).AutoMigrate(&system.SysAuthorityMenu{})
	return err
}

func (csOperatorService *CsOperatorService) CreateTenetUserTableName(tenantID uint, operatorID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserTableName(tenantID, operatorID)).AutoMigrate(&system.SysUserTable{})
	return err
}

func (csOperatorService *CsOperatorService) CreateTenetAuthsTable(tenantID uint, operatorID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthsTable(tenantID, operatorID)).AutoMigrate(&system.SysAuthorityTable{})
	return err
}

func (csOperatorService *CsOperatorService) CreateUserAuthorityTable(tenantID uint, operatorID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserAuthorityTableName(tenantID, operatorID)).AutoMigrate(&system.SysUserAuthority{})
	return err
}

// CreateCsOperator 创建CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) CreateCsOperator(csOperator *operator.CsOperator) (err error) {
	if global.GVA_DB.First(&operator.CsOperator{}, "username = ?", csOperator.Username).Error == nil {
		return errors.New("商户已注册")
	}
	var tenant tenant2.CsTenant
	err = global.GVA_DB.First(&tenant, "id = ?", csOperator.TenantID).Error
	if err != nil {
		return err
	}
	csOperator.TenantOnlyKey = tenant.OnlyKey
	csOperator.OnlyKey, _ = uuid.NewV4()
	err = global.GVA_DB.Create(csOperator).Error
	if err != nil {
		return err
	}

	err = csOperatorService.CreateTenetUserTableName(csOperator.TenantID, csOperator.ID)
	if err != nil {
		return err
	}

	err = csOperatorService.CreateTenetAuthsTable(csOperator.TenantID, csOperator.ID)
	if err != nil {
		return err
	}

	err = csOperatorService.CreateUserAuthorityTable(csOperator.TenantID, csOperator.ID)
	if err != nil {
		return err
	}

	err = csOperatorService.CreateOperatorIDCasbin(csOperator.TenantID, csOperator.ID)
	if err != nil {
		return err
	}
	return csOperatorService.CreateOperatorIDAuthMenu(csOperator.TenantID, csOperator.ID)
}

func (csOperatorService *CsOperatorService) CreateOperatorAdmin(tenantID uint, operatorID uint) (err error) {
	user := &system.SysUser{
		Username: "admin",
		Password: "123456",
	}
	user.TenantID = tenantID
	user.NickName = "商户管理员"
	user.AuthorityId = 888
	user.Authorities = []system.SysAuthority{{AuthorityId: 888}}
	_, err = sysUserServer.Register(*user, tenantID, operatorID)
	if err != nil {
		return err
	}

	auth := &system.SysAuthority{
		AuthorityId:   888,
		AuthorityName: "商户管理员",
		ParentId:      utils.Pointer(uint(0)),
		DefaultRouter: "dashboard",
	}

	_, err = sysAuthorityServer.CreateAuthority(*auth, tenantID, operatorID)

	apis, err := csOperatorService.GetOperatorApis(strconv.Itoa(int(operatorID)))
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

	err = casbinServer.UpdateCasbin(888, casbinInfo, tenantID, operatorID)
	if err != nil {
		return err
	}
	menus, err := csOperatorService.GetOperatorMenus(strconv.Itoa(int(operatorID)))
	var menuIDs []uint
	for _, menu := range menus {
		menuIDs = append(menuIDs, menu.MenuId)
	}
	var baseMenus []system.SysBaseMenu
	global.GVA_DB.Find(&baseMenus, "id in ?", menuIDs)
	err = menuServer.AddMenuAuthority(baseMenus, 888, tenantID, operatorID)

	return err
}

// DeleteCsOperator 删除CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) DeleteCsOperator(csOperator operator.CsOperator) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperator{}).Where("id = ?", csOperator.ID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&csOperator).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCsOperatorByIds 批量删除CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) DeleteCsOperatorByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperator{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&operator.CsOperator{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCsOperator 更新CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) UpdateCsOperator(csOperator operator.CsOperator) (err error) {
	err = global.GVA_DB.Save(&csOperator).Error
	return err
}

// GetCsOperator 根据id获取CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) GetCsOperator(id uint) (csOperator operator.CsOperator, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csOperator).Error
	return
}

// GetCsOperatorInfoList 分页获取CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) GetCsOperatorInfoList(info operatorReq.CsOperatorSearch) (list []operator.CsOperator, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&operator.CsOperator{})
	var csOperators []operator.CsOperator
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TenantID != 0 {
		db = db.Where("tenant_id = ?", info.TenantID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&csOperators).Error
	return csOperators, total, err
}

func (csOperatorService *CsOperatorService) SetOperatorApis(operatorApis operatorReq.CsOperatorApisReq) error {
	var csOperatorApis []operator.CsOperatorApis
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("operator_id = ?", operatorApis.OperatorID).Delete(&csOperatorApis).Error
		if err != nil {
			return err
		}
		for _, apiId := range operatorApis.ApiIDs {
			csOperatorApis = append(csOperatorApis, operator.CsOperatorApis{
				OperatorID: operatorApis.OperatorID,
				ApiId:      apiId,
			})
		}
		return global.GVA_DB.Create(&csOperatorApis).Error
	})
}

func (csOperatorService *CsOperatorService) SetOperatorMenus(operatorMenus operatorReq.CsOperatorMenusReq) error {
	var csOperatorMenus []operator.CsOperatorMenus
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("operator_id = ?", operatorMenus.OperatorID).Delete(&csOperatorMenus).Error
		if err != nil {
			return err
		}
		for _, menuID := range operatorMenus.MenuIDs {
			csOperatorMenus = append(csOperatorMenus, operator.CsOperatorMenus{
				OperatorID: operatorMenus.OperatorID,
				MenuId:     menuID,
			})
		}
		return global.GVA_DB.Create(&csOperatorMenus).Error
	})
}

func (csOperatorService *CsOperatorService) GetOperatorIDByOnlyKey(onlyKey string) uint {
	var csOperator operator.CsOperator
	global.GVA_DB.Where("only_key = ?", onlyKey).First(&csOperator)
	return csOperator.ID
}

func (csOperatorService *CsOperatorService) GetOperatorApis(id string) (csOperatorApis []operator.CsOperatorApis, err error) {
	err = global.GVA_DB.Where("operator_id = ?", id).Find(&csOperatorApis).Error
	return
}

func (csOperatorService *CsOperatorService) GetOperatorMenus(id string) (csOperatorMenus []operator.CsOperatorMenus, err error) {
	err = global.GVA_DB.Where("operator_id = ?", id).Find(&csOperatorMenus).Error
	return
}

package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority, tenantID uint, operatorID uint) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	tableName := utils.GetAuthsTable(tenantID, operatorID)
	if !errors.Is(global.GVA_DB.Table(tableName).Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.GVA_DB.Table(tableName).Create(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse, tenantID uint, operatorID uint) (authority system.SysAuthority, err error) {
	// TODO: 找表

	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId}, tenantID, operatorID)
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}

	var btns []system.SysAuthorityBtn

	err = global.GVA_DB.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityId).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		err = global.GVA_DB.Create(&btns).Error

		if err != nil {
			return
		}
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId, tenantID, operatorID)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths, tenantID, operatorID)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority, tenantID, operatorID)
	}
	return copyInfo.Authority, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority, tenantID uint, operatorID uint) (authority system.SysAuthority, err error) {
	tableName := utils.GetAuthsTable(tenantID, operatorID)
	err = global.GVA_DB.Table(tableName).Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority, tenantID uint, operatorID uint) (err error) {
	// TODO: 找表
	tableName := utils.GetAuthsTable(tenantID, operatorID)
	userTable := utils.GetUserTableName(tenantID, operatorID)
	if errors.Is(global.GVA_DB.Table(tableName).Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Table(userTable)
	}).First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Table(userTable).Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Table(tableName).Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GVA_DB.Table(tableName).Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		var ids []uint
		for i := range auth.SysBaseMenus {
			ids = append(ids, auth.SysBaseMenus[i].ID)
		}
		authMenuTable := utils.GetAuthMenuTableName(tenantID, operatorID)
		err = global.GVA_DB.Table(authMenuTable).Where("sys_base_menu_id in (?)", ids).Delete(&[]system.SysAuthorityMenu{}).Error
		if err != nil {
			return
		}
		// err = db.Association("SysBaseMenus").Delete(&auth)
	}
	userAuthTable := utils.GetUserAuthorityTableName(tenantID, operatorID)
	err = global.GVA_DB.Table(userAuthTable).Delete(&[]system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	if err != nil {
		return
	}
	// TODO: 按钮关联暂时未完成
	//err = global.GVA_DB.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ?", auth.AuthorityId).Error
	//if err != nil {
	//	return
	//}
	authorityId := strconv.Itoa(int(auth.AuthorityId))
	CasbinServiceApp.ClearCasbin(tenantID, operatorID, 0, authorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo, tenantID uint, operatorID uint) (list interface{}, total int64, err error) {
	tableName := utils.GetAuthsTable(tenantID, operatorID)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Table(tableName)
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k], tableName)
	}
	return authority, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: sa system.SysAuthority, err error

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority, tenantID uint, operationID uint) error {
	tableName := utils.GetAuthMenuTableName(tenantID, operationID)
	global.GVA_DB.Table(tableName).Where("sys_authority_authority_id = ?", auth.AuthorityId).Delete(&[]system.SysAuthorityMenu{})
	var authMenus []system.SysAuthorityMenu
	for _, v := range auth.SysBaseMenus {
		authMenus = append(authMenus, system.SysAuthorityMenu{AuthorityId: strconv.Itoa(int(auth.AuthorityId)), MenuId: strconv.Itoa(int(v.ID))})
	}
	return global.GVA_DB.Table(tableName).Create(authMenus).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority, tableName string) (err error) {
	err = global.GVA_DB.Table(tableName).Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k], tableName)
		}
	}
	return err
}

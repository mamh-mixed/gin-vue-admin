package system

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

func (userService *UserService) Register(u system.SysUser, tenantID uint, operatorID uint) (userInter system.SysUser, err error) {
	var user system.SysUser
	userTable := utils.GetUserTableName(tenantID, operatorID)
	var auths []system.SysAuthority
	for _, v := range u.Authorities {
		auths = append(auths, v)
	}
	userAuthTable := utils.GetUserAuthorityTableName(tenantID, operatorID)
	userTableName := utils.GetUserTableName(tenantID, operatorID)
	u.TenantID = tenantID
	u.OperatorID = operatorID
	if !errors.Is(global.GVA_DB.Table(userTableName).Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	u.Authorities = nil
	if len(auths) > 0 {
		u.AuthorityId = auths[0].AuthorityId
	}
	err = global.GVA_DB.Table(userTable).Create(&u).Error
	if err != nil {
		return u, err
	}
	var userAuths []system.SysUserAuthority
	for _, v := range auths {
		userAuths = append(userAuths, system.SysUserAuthority{
			SysUserId:               u.ID,
			SysAuthorityAuthorityId: v.AuthorityId,
		})
	}

	err = global.GVA_DB.Table(userAuthTable).Create(&userAuths).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser, tenantID uint, operatorID uint) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	tableName := utils.GetUserTableName(tenantID, operatorID)
	AuthTableName := utils.GetAuthsTable(tenantID, operatorID)
	userAuthTableName := utils.GetUserAuthorityTableName(tenantID, operatorID)
	var userAuthorities []system.SysUserAuthority
	err = global.GVA_DB.Table(userAuthTableName).Where("sys_user_id = ?", u.ID).Find(&userAuthorities).Error
	if err != nil {
		return nil, err
	}
	authIds := make([]uint, 0, len(userAuthorities))
	for _, v := range userAuthorities {
		authIds = append(authIds, v.SysAuthorityAuthorityId)
	}

	auths := make([]system.SysAuthority, 0, len(authIds))
	err = global.GVA_DB.Table(AuthTableName).Where("authority_id in (?)", authIds).Find(&auths).Error

	var user system.SysUser
	err = global.GVA_DB.Table(tableName).Where("username = ?", u.Username).
		Preload("Authority", func(db *gorm.DB) *gorm.DB {
			return db.Table(AuthTableName)
		}).First(&user).Error
	user.Authorities = auths
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user, tenantID, operatorID)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string, tenantID uint, operatorID uint) (userInter *system.SysUser, err error) {
	var user system.SysUser
	userTableName := utils.GetUserTableName(tenantID, operatorID)
	if err = global.GVA_DB.Table(userTableName).Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Table(userTableName).First(&user).Update("password", user.Password).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo, tenantID uint, operatorID uint) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	tableName := utils.GetUserTableName(tenantID, operatorID)
	authTableName := utils.GetAuthsTable(tenantID, operatorID)
	db := global.GVA_DB.Table(tableName)
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).
		Preload("Authority", func(db *gorm.DB) *gorm.DB {
			return db.Table(authTableName)
		}).Find(&userList).Error

	for i := range userList {
		userAuthTableName := utils.GetUserAuthorityTableName(tenantID, operatorID)
		var userAuthorities []system.SysUserAuthority
		err = global.GVA_DB.Table(userAuthTableName).Where("sys_user_id = ?", userList[i].ID).Find(&userAuthorities).Error
		if err != nil {
			return nil, 0, err
		}
		authIds := make([]uint, 0, len(userAuthorities))
		for _, v := range userAuthorities {
			authIds = append(authIds, v.SysAuthorityAuthorityId)
		}

		err = global.GVA_DB.Table(authTableName).Where("authority_id in (?)", authIds).Find(&userList[i].Authorities).Error
	}

	return userList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint, tenantID uint, operatorID uint) (err error) {
	userAuthTable := utils.GetUserAuthorityTableName(tenantID, operatorID)
	assignErr := global.GVA_DB.Table(userAuthTable).Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	userTable := utils.GetUserTableName(tenantID, operatorID)
	err = global.GVA_DB.Table(userTable).Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint, tenantID uint, operatorID uint) (err error) {
	userAuthTable := utils.GetUserAuthorityTableName(tenantID, operatorID)
	userTable := utils.GetUserTableName(tenantID, operatorID)
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Table(userAuthTable).Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Table(userAuthTable).Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Table(userTable).Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int, tenantID uint, operatorID uint) (err error) {
	var user system.SysUser
	userAuthTable := utils.GetUserAuthorityTableName(tenantID, operatorID)
	userTableName := utils.GetUserTableName(tenantID, operatorID)
	err = global.GVA_DB.Table(userTableName).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Table(userAuthTable).Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser, tenantID uint, operatorID uint) error {
	userTableName := utils.GetUserTableName(tenantID, operatorID)
	return global.GVA_DB.Table(userTableName).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req system.SysUser, tenantID uint, operatorID uint) error {
	userTableName := utils.GetUserTableName(tenantID, operatorID)
	return global.GVA_DB.Table(userTableName).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(userID, tenantID uint, operatorID uint) (user system.SysUser, err error) {
	var reqUser system.SysUser
	tableName := utils.GetUserTableName(tenantID, operatorID)
	authTableName := utils.GetAuthsTable(tenantID, operatorID)
	userAuthTableName := utils.GetUserAuthorityTableName(tenantID, operatorID)

	err = global.GVA_DB.Table(tableName).
		Preload("Authority", func(db *gorm.DB) *gorm.DB {
			return db.Table(authTableName)
		}).First(&reqUser, "id = ?", userID).Error

	var userAuthorities []system.SysUserAuthority
	err = global.GVA_DB.Table(userAuthTableName).Where("sys_user_id = ?", reqUser.ID).Find(&userAuthorities).Error
	if err != nil {
		return user, err
	}
	authIds := make([]uint, 0, len(userAuthorities))
	for _, v := range userAuthorities {
		authIds = append(authIds, v.SysAuthorityAuthorityId)
	}

	err = global.GVA_DB.Table(authTableName).Where("authority_id in (?)", authIds).Find(&reqUser.Authorities).Error

	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser, tenantID, operatorID)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int, tenantID uint, operatorID uint) (user *system.SysUser, err error) {
	var u system.SysUser
	tableName := utils.GetUserTableName(tenantID, operatorID)
	err = global.GVA_DB.Table(tableName).Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string, tenantID uint, operatorID uint) (user *system.SysUser, err error) {
	var u system.SysUser
	tableName := utils.GetUserTableName(tenantID, operatorID)
	if err = global.GVA_DB.Table(tableName).Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint, tenantID uint, operatorID uint) (err error) {
	tableName := utils.GetUserTableName(tenantID, operatorID)
	err = global.GVA_DB.Table(tableName).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

package tenant

type CsTenantMenus struct {
	ID       uint `gorm:"primarykey"` // 主键ID
	TenantID uint `json:"tenantID" form:"tenantID" gorm:"column:tenant_id;comment:租户id;"`
	MenuId   uint `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:菜单id;"`
}

// TableName CsTenant 表名
func (CsTenantMenus) TableName() string {
	return "cs_tenant_menus"
}

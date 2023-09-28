package operator

type CsOperatorMenus struct {
	ID         uint `gorm:"primarykey"` // 主键ID
	OperatorID uint `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:运营商id;"`
	MenuId     uint `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:菜单id;"`
}

// TableName CsTenant 表名
func (CsOperatorMenus) TableName() string {
	return "cs_operator_menus"
}

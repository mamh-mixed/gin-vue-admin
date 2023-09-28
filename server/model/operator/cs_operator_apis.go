package operator

type CsOperatorApis struct {
	ID         uint `gorm:"primarykey"` // 主键ID
	OperatorID uint `json:"operatorId" form:"operatorId" gorm:"column:operator_id;comment:运营商id;"`
	ApiId      uint `json:"apiId" form:"apiId" gorm:"column:api_id;comment:api id;"`
}

// TableName CsTenant 表名
func (CsOperatorApis) TableName() string {
	return "cs_operator_apis"
}

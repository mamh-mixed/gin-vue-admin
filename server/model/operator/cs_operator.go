// 自动生成模板CsOperator
package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
	"time"
)

// CsOperator 结构体
type CsOperator struct {
	global.GVA_MODEL
	Username      string    `json:"username" form:"username" gorm:"column:username;comment:商户名;size:191;"`
	Nickname      string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:商户昵称;size:191;"`
	OnlyKey       uuid.UUID `json:"onlyKey" form:"onlyKey" gorm:"column:only_key;comment:商户唯一标识;type:char(36);size:36;"`
	TenantOnlyKey uuid.UUID `json:"tenantOnlyKey" form:"tenantOnlyKey" gorm:"column:tenant_only_key;comment:租户唯一标识;type:char(36);size:36;"`
	//Logo                 string     `json:"logo" form:"logo" gorm:"column:logo;comment:租户logo;"`
	EndTime              *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"`
	Allocation           *int       `json:"allocation" form:"allocation" gorm:"column:allocation;comment:;size:19;"`
	AllocationProportion *float64   `json:"allocationProportion" form:"allocationProportion" gorm:"column:allocation_proportion;comment:;size:22;"`
}

// TableName CsOperator 表名
func (CsOperator) TableName() string {
	return "cs_operator"
}

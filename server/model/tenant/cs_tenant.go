// 自动生成模板CsTenant
package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
	"time"
)

// CsTenant 结构体
type CsTenant struct {
	global.GVA_MODEL
	OnlyKey              uuid.UUID  `json:"onlyKey" form:"onlyKey" gorm:"column:only_key;comment:租户唯一标识;type:char(36);size:36;"`
	Username             string     `json:"username" form:"username" gorm:"column:username;comment:租户名;"`
	Nickname             string     `json:"nickname" form:"nickname" gorm:"column:nickname;comment:租户昵称;"`
	Logo                 string     `json:"logo" form:"logo" gorm:"column:logo;comment:租户logo;"`
	EndTime              *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"`
	Allocation           *int       `json:"allocation" form:"allocation" gorm:"column:allocation;comment:;"`
	AllocationProportion *float64   `json:"allocationProportion" form:"allocationProportion" gorm:"column:allocation_proportion;comment:;"`
}

// TableName CsTenant 表名
func (CsTenant) TableName() string {
	return "cs_tenant"
}

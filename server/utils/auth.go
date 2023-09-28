package utils

import "fmt"

func GetCasbinName(tenantID uint, operatorID uint) string {
	if tenantID != 0 {
		if operatorID != 0 {
			return fmt.Sprintf("casbin_rule_tenant_%d_operator_%d", tenantID, operatorID)
		}
		return fmt.Sprintf("casbin_rule_tenant_%d", tenantID)
	}
	return "casbin_rule"
}

func GetAuthMenuTableName(tenantID uint, operatorID uint) string {
	if tenantID != 0 {
		if operatorID != 0 {
			return fmt.Sprintf("cs_authority_menus_tenant_%d_operator_%d", tenantID, operatorID)
		}
		return fmt.Sprintf("cs_authority_menus_tenant_%d", tenantID)
	}
	return "sys_authority_menus"
}

func GetUserAuthorityTableName(tenantID uint, operatorID uint) string {
	if tenantID != 0 {
		if operatorID != 0 {
			return fmt.Sprintf("cs_user_authority_tenant_%d_operator_%d", tenantID, operatorID)
		}
		return fmt.Sprintf("cs_user_authority_tenant_%d", tenantID)
	}
	return "sys_user_authority"
}

func GetUserTableName(tenantID uint, operatorID uint) string {
	if tenantID != 0 {
		if operatorID != 0 {
			return fmt.Sprintf("cs_user_tenant_%d_operator_%d", tenantID, operatorID)
		}
		return fmt.Sprintf("cs_user_tenant_%d", tenantID)
	}
	return "sys_users"
}

func GetAuthsTable(tenantID uint, operatorID uint) string {
	if tenantID != 0 {
		if operatorID != 0 {
			return fmt.Sprintf("cs_authorities_tenant_%d_operator_%d", tenantID, operatorID)
		}
		return fmt.Sprintf("cs_authorities_tenant_%d", tenantID)
	}
	return "sys_authorities"
}

package daemons

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
)


func RoleSystemApiEnforcerDaemon(policies []common.PolicyAction) (err error) {
	en := common.Enforcer
	for _, rule := range policies {
		en.AddPolicy(rule.PType, rule.Address, rule.Method)

	}
	err = en.SavePolicy()
	if err == nil {
		en.LoadPolicy()

	}
	return
}

//UserOrRoleEnforcerDaemon enforcer daemon will modify policy and reload
func UserOrRoleEnforcerDaemon(groupPolicyActions []common.GroupPolicyAction) {
	en := common.Enforcer
	if len(groupPolicyActions) > 0 {
		for _, gpa := range groupPolicyActions {
			switch gpa.Action {
			case "delete":
				en.RemoveGroupingPolicy(gpa.UserOrRole, gpa.Role)
			case "add":
				en.AddGroupingPolicy(gpa.UserOrRole, gpa.Role)

			}
		}
		en.SavePolicy()
		en.LoadPolicy()
	}
}

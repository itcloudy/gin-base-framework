package middles

import (
	"encoding/json"
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

//CasbinJwtAuthorize returns the authorizer, uses a Casbin enforcer as input
func CasbinJwtAuthorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
			//c.Abort is must
			return
		}
	}
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	userId := c.GetInt(common.LOGIN_USER_ID)
	roles := c.GetStringSlice(common.LOGIN_USER_ROLES)
	isAdmin := c.GetBool(common.LOGIN_IS_ADMIN)
	if isAdmin {
		return true
	}
	roles = append(roles, fmt.Sprintf("user_%d", userId))
	authOk := false
	path := c.Request.URL.Path
	method := c.Request.Method
	return true
	for _, key := range roles {
		if a.enforcer.Enforce(key, path, method) {
			authOk = true
			break
		}
	}
	return authOk
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	tokenValid := c.GetBool(common.TOKEN_VALID)
	var res common.ResponseJson
	if tokenValid {
		c.Writer.WriteHeader(http.StatusForbidden)
		res = common.ResponseJson{
			Code:    common.FORBIDDEN,
			Data:    nil,
			Message: "you have no right for this action",
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		res = common.ResponseJson{
			Code:    common.TOKENERR,
			Data:    nil,
			Message: "token error or expire",
		}
	}

	m, _ := json.Marshal(res)
	c.Writer.Write([]byte(m))
	c.Abort()
}

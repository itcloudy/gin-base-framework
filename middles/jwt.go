package middles

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"go.uber.org/zap"
	"io/ioutil"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	jwtPrefix = "Bearer"
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type JwtClaims struct {
	*jwt.StandardClaims
	Name    string
	Role    []string
	UserId  int
	IsAdmin bool
}

// Read the key files before starting http handlers
func InitKeys() {

	signBytes, err := ioutil.ReadFile(common.ServerInfo.JwtPriKeyPath)
	if err != nil {
		common.Logger.Fatal("jwt private key file read failed", zap.Error(err))
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		common.Logger.Fatal("jwt private key file parse failed", zap.Error(err))
	}

	verifyBytes, err := ioutil.ReadFile(common.ServerInfo.JwtPubKeyPath)
	if err != nil {
		common.Logger.Fatal("jwt public key file read failed", zap.Error(err))
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		common.Logger.Fatal("jwt public key file parse failed", zap.Error(err))
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(name string, role []string, userId int, isAdmin bool) string {

	claims := JwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * common.ServerInfo.TokenExpireSecond).Unix(),
		},
		name,
		role,
		userId,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		common.Logger.Error("token generator failed", zap.Error(err))
		return ""
	}

	return common.StringsJoin(jwtPrefix, " ", ss)
}

//JwtAuthorize parse jwt info
func JwtAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		headToken := c.GetHeader("Authorization")
		if headToken != "" {
			headToken = string(headToken[len(jwtPrefix)+1:])
			var jclaim = &JwtClaims{}
			token, err := jwt.ParseWithClaims(headToken, jclaim, func(*jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})
			if err != nil {
				common.Logger.Error("token parse failed")
			}
			if token.Valid {
				jwtClaims := token.Claims.(*JwtClaims)
				c.Set(common.LOGIN_USER_ID, jwtClaims.UserId)
				c.Set(common.LOGIN_USER_NAME, jwtClaims.Name)
				c.Set(common.LOGIN_USER_ROLES, jwtClaims.Role)
				c.Set(common.LOGIN_IS_ADMIN, jwtClaims.IsAdmin)
				c.Set(common.TOKEN_VALID, true)
			} else {
				c.Set(common.LOGIN_USER_ID, 0)
				c.Set(common.LOGIN_USER_NAME, "")
				c.Set(common.LOGIN_USER_ROLES, []string{})
				c.Set(common.LOGIN_IS_ADMIN, false)
				c.Set(common.TOKEN_VALID, false)
				c.Set(common.LOGIN_COMPANY_ID, 0)
				c.Set(common.LOGIN_COMPANY_REQUEST_NO, "")
				common.Logger.Warn("token invalid")
			}
		}

	}
}

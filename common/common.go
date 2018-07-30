package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"strings"
)

type ResponseJson struct {
	Code    uint        `yaml:"code" json:"code"`   //response code
	Data    interface{} `yaml:"data" json:"data"`   //response data
	Message string      `yaml:"message" json:"msg"` //response message
}

type PolicyAction struct {
	PType   string // type ：role_$ user_$
	Address string //request address
	Method  string //request method :DELETE GET  PUT POST

}

type GroupPolicyAction struct {
	Action     string //action :delete add
	UserOrRole string // user or role :user_$,role_$
	Role       string //role  role_$
}

//GenResponse genrate reponse ,json format
func GenResponse(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

// StringsJoin string array join
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}
func Join2String(split string, strs ...interface{}) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		var str interface{}
		switch str.(type) {
		case string:
			b.WriteString(str.(string))
		case int:
			b.WriteString(fmt.Sprintf("%d", str.(int)))
		case int64:
			b.WriteString(fmt.Sprintf("%d", str.(int)))

		}
	}
	str = b.String()
	return str

}
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

//String2Int string to int
func String2Int(str string, defVal int) int {
	if in, err := strconv.Atoi(str); err != nil {
		return defVal
	} else {
		return in
	}
}

func Validate(model interface{}) error {
	var (
		validate *validator.Validate
		err      error
	)
	validate = validator.New()
	err = validate.Struct(model)
	return err
}

func GetVisitorInfo(c *gin.Context) map[string]interface{} {
	var (
		remoteAddr  []string
		visitorInfo map[string]interface{}
	)
	visitorInfo = make(map[string]interface{})
	remoteAddr = strings.Split(c.Request.RemoteAddr, ":")
	if len(remoteAddr) == 2 {
		visitorInfo["ip"] = remoteAddr[0]
	}
	visitorInfo["header"] = c.Request.Header
	visitorInfo["url"] = c.Request.URL
	visitorInfo["content-length"] = c.Request.ContentLength
	visitorInfo["host"] = c.Request.Host
	visitorInfo["uri"] = c.Request.RequestURI

	return visitorInfo
}

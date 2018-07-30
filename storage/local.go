package storage

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileLocalStorage(file multipart.File, filename string, params ...interface{}) (localFileName string, err error) {
	var (
		bys    []byte
		exists bool
	)

	databuf := bytes.NewBuffer(bys)
	for {
		var oneBuff = make([]byte, 1024)
		oneLen, _ := file.Read(oneBuff)
		if oneLen > 0 {
			databuf.Write(oneBuff)
		} else {
			break
		}
	}
	file.Seek(0, 0)
	hash := md5.New()
	hash.Write(databuf.Bytes())
	hashName := hex.EncodeToString(hash.Sum(nil))
	rootPath, _ := os.Getwd()
	uploadPath := common.FileUploadInfo.Path
	sufixList := strings.Split(filename, ".")
	sufix := sufixList[len(sufixList)-1]
	localFileName = common.StringsJoin(hashName, ".", sufix)
	filePath := path.Join(rootPath, uploadPath, localFileName)
	exists, _ = PathExists(path.Join(rootPath, uploadPath))

	if !exists {
		err := os.Mkdir(path.Join(rootPath, uploadPath), os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err

	}

	return
}

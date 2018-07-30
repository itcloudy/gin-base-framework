package storage

import (
	"os"
)

//qiniu upload response struct
type QiNiuPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
	Domain string
}

// 获取
// 文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

/*

func QiNiuServerUpload(w http.ResponseWriter, r *http.Request, data *apiData, logger *log.Entry) (err error) {
	//get paramters
	_, pr_fix, err := checkEcosystem(w, data, logger)
	qiniuAccessKey := &model.StateParameter{}
	qiniuSecretKey := &model.StateParameter{}
	qiniuBucketKey := &model.StateParameter{}
	qiniuDomainKey := &model.StateParameter{}
	qiniuAccessKey.SetTablePrefix(pr_fix)
	qiniuSecretKey.SetTablePrefix(pr_fix)
	qiniuBucketKey.SetTablePrefix(pr_fix)
	qiniuDomainKey.SetTablePrefix(pr_fix)
	accessKey, err := qiniuAccessKey.Get(nil, "qiniu_accessKey")
	secretKey, err := qiniuSecretKey.Get(nil, "qiniu_secretKey")
	bucketKey, err := qiniuBucketKey.Get(nil, "qiniu_bucket")
	domainKey, err := qiniuDomainKey.Get(nil, "qiniu_domain")
	if !accessKey || !secretKey || !bucketKey || !domainKey {
		if !accessKey {
			logger.WithFields(log.Fields{"type": consts.NotFound, "key": "qiniu_accessKey"}).Error("qiniu_accessKey parameter not found")

		}
		if !secretKey {
			logger.WithFields(log.Fields{"type": consts.NotFound, "key": "qiniu_secretKey"}).Error("qiniu_secretKey parameter not found")

		}
		if !bucketKey {
			logger.WithFields(log.Fields{"type": consts.NotFound, "key": "qiniu_bucket"}).Error("qiniu_bucket parameter not found")

		}
		if !domainKey {
			logger.WithFields(log.Fields{"type": consts.NotFound, "key": "qiniu_domain"}).Error("qiniu_domain parameter not found")

		}
		return errorAPI(w, err, http.StatusBadRequest)
	}

	//upload file handler

	file, handler, err := r.FormFile("upload_file") //name的字段
	if err != nil {
		return
	}

	defer file.Close()
	var bys []byte
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
	if err != nil {
		return
	}

	hash := md5.New()
	hashName := hex.EncodeToString(hash.Sum(nil))
	fileName := hashName + "_" + handler.Filename

	putPolicy := storage.PutPolicy{
		Scope:      qiniuBucketKey.Value,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(qiniuAccessKey.Value, qiniuSecretKey.Value)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := QiNiuPutRet{}

	err = formUploader.Put(context.Background(), &ret, upToken, fileName, databuf, int64(databuf.Len()), nil)
	if err != nil {
		return
	}
	ret.Domain = qiniuDomainKey.Value
	data.result = &ret
	return
}

*/

package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"oss-tool/model"
	"time"
)
type AliOSS struct {
	Account model.OSSAccount
}

func (a *AliOSS) UploadFile(filePath string) (url string, err error) {
	// 创建OSSClient实例
	client, err := oss.New(a.Account.Endpoint, a.Account.AccessKeyId, a.Account.AccessKeySecret)
	if err != nil {
		return "", err
	}
	// 获取存储空间
	bucket, err := client.Bucket(a.Account.BucketName)
	// 批量上传
	objectDir := ""
	if a.Account.ObjectName == "" {
		objectDir = fmt.Sprintf("%d%d/", time.Now().Year(), time.Now().Month())
	} else {
		objectDir = fmt.Sprintf("%s/", a.Account.ObjectName)
	}
	// 上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg
	objectFileName := objectDir + parseFileName(filePath)
	// 上传单个文件
	// 由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt
	err = bucket.PutObjectFromFile(objectFileName, filePath)
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf("https://%s.%s/%s", a.Account.BucketName, a.Account.Endpoint, objectFileName)
	return url, nil
}

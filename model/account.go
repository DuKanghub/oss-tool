package model

type OSSAccount struct {
	Platform string		// 平台: ali,
	Endpoint string		// 云服务器所在地域
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string	  // 用来存放上传文件的桶
	ObjectName		string   // 用来存放上传文件的object目录
}

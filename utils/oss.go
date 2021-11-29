package utils

import (
	"oss-tool/model"
)

type OSS interface {
	UploadFile(string) (string, error)
}

func NewOSS(a model.OSSAccount) OSS {
	switch a.Platform {
	case "ali":
		return &AliOSS{a}
	default:
		return &AliOSS{a}
	}
}
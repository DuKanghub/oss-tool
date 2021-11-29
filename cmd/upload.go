/*
Copyright © 2021 DuKang <dukang@dukanghub.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"oss-tool/model"
	"oss-tool/utils"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传",
	Long: `上传文件到OSS`,
	Run: func(cmd *cobra.Command, args []string) {
		// 如果使用upload子命令，必须在这里取flag变量的值，才能取到，在root.go下取到的是空的。
		account := model.OSSAccount{
			Platform: platForm,
			Endpoint: endPoint,
			AccessKeyId: accessKeyId,
			AccessKeySecret: accessKeySecret,
			BucketName: bucketName,
			ObjectName: objectName,
		}
		if len(args) <= 0 {
			fmt.Println("请在参数位传入文件名")
			return
		}
		oss := utils.NewOSS(account)
		fmt.Printf("一共需上传文件%d个\n", len(args))
		for i, file := range args {
			url, err := oss.UploadFile(file)
			if err != nil {
				fmt.Printf("上传第%d个文件时失败，错误：%+v\n", i+1, err)
			} else {
				fmt.Printf("上传成功，地址为：%s\n", url)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}

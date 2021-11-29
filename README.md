# 基于cobra的OSS命令行工具示例
# 目前功能

- 支持的平台：阿里云OSS

- 支持的OSS操作：上传文件

# 使用

进入目录根目录，编译
```shell
go build
```
如果是windows平台，编译后生成`oss-tool.exe`，linux平台编译后生成`oss-tool`

子命令：目前只有一个`upload`

选项(`flags`):
- `--endpoint`(简写：`-e`): OSS地域，必传
- `--platform`(简写：`-p`): OSS平台, 可选值有：ali, 不传默认是ali
- `--access_key_id`(简写：`-k`): accessKeyId，必传
- `--access_key_secret`(简写：`-s`): accessKeySecret，必传
- `--bucket_name`(简写：`-b`): 用来存放上传文件的桶名字，必传
- `--object_name`(简写：`-o`): OSS文件夹名字，可选，默认是当前年月日

上传文件：
```shell
# 不指定oss平台，默认是ali，即阿里云OSS
oss-tool upload -k ${access_key_id} -s ${access_key_secret} -e ${endpoint} -b ${bucket_name} ${filepath}
```


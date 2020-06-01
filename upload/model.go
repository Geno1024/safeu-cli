package upload

// 上传相关数据结构
import (
	"io"
	"net/http"
	"os"
)

// 上传策略

// 阿里云 OSS 接口返回 的 uuid
type UploadResponse struct {
	UUID string
}

// 上传文件对应状态
const (
	UploadFileReadyCode   = iota // 0
	UploadFileSuccessCode        // 1
	UploadFileFailedCode         // 2
)

// 上传实例
// 一次上传策略对应一个上传实例
// 一个上传实例包含多个文件上传
type Instance struct {
	UploadPolicy UploadPolicy
	UploadFiles  []UploadFile
	Owner        string
	Recode       string
	CSRF         string // csrf token
	Cookie       string // csrf cookie
	Password     string // 密码
	DownCount    int    // 下载次数
	ExpireTime   int    // 过期时间(小时)
}

// 上传策略
type UploadPolicy struct {
	AccessID  string
	Host      string
	Expire    int64
	Signature string
	Policy    string
	Dir       string
	Callback  string
}

// 上传文件
type UploadFile struct {
	File           *os.File        // 文件本体
	UploadResponse *UploadResponse // 上传到OSS返回的结构体
	StatusCode     int             // 状态码

	Client *http.Client         // http Client
	Url    string               // 上传地址
	Values map[string]io.Reader // 上传结构体
}

// 工具函数区
// 上传文件完毕发送Finish请求
// example : {"files":["0f652be1-394b-43f6-95bf-948de1520d0c","5561366d-e18c-4e48-8ae9-ec46f0a70ecf"]}
type FinishRequest struct {
	Files []string `json:"files"`
}

// Finish 请求返回结果
type FinishResponse struct {
	Owner  string `json:"owner"`
	Recode string `json:"recode"`
}

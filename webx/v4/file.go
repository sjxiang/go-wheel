package webx

import (
	"io"
	"mime/multipart"
	"os"
)

// 文件上传
type FileUploader struct {
	// form-data 对应与文件在表单中的字段名字
	FileField   string
	// 目录
	Dst         string 
	// 计算目标路径（自定义生成）
	DstPathFunc func(fh *multipart.FileHeader) string
}


func (f FileUploader) Handle() HandleFunc {
	return func(ctx *Context) {
		// 读取数据
		src, header, err := ctx.Req.FormFile(f.FileField)
		if err != nil {
			ctx.Resp.WriteHeader(400)
			ctx.Resp.Write([]byte("上传失败"))
			return 
		}

		// 创建目标文件
		dst, err := os.OpenFile(f.DstPathFunc(header), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			ctx.Resp.WriteHeader(500)
			ctx.Resp.Write([]byte("打开文件失败"))
			return 
		}
		defer dst.Close()

		// 拷贝数据
		_, err = io.CopyBuffer(dst, src, nil)
		if err != nil {
			ctx.Resp.WriteHeader(500)
			ctx.Resp.Write([]byte("拷贝失败"))
			return 
		}
	}
}


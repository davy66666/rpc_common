package miniox

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"rabbitmq/common/utils/idx"
	"rabbitmq/common/utils/imagex"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Client struct {
	conn        *minio.Client
	Bucket      string        // 桶名
	Location    string        // 地点/服务
	Timeout     time.Duration // 超时
	ContentType string
}
type Options struct {
	Expires    int64 // 过期时间-单位分钟(默认30分钟)
	Width      int   // 指定宽
	Height     int   // 指定高
	IsDownload bool  // 是否下载。图片默认False.返回预览链接地址。文件或者下载图片可以为Ture为下载连接
}

type MinioConf struct {
	Address     string // 地址-必填
	Username    string // 账号-必填
	Password    string // 密码-必填
	Bucket      string // 桶名-必填
	Token       string // 默认无
	Location    string // 地点/服务
	Timeout     int    // 超时
	ContentType string // 文件类型
	TLS         bool   // 是否开启TLS(HTTPS)
}

func MustMinioClient(c MinioConf) *Client {

	conn, err := minio.New(c.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Username, c.Password, ""),
		Secure: c.TLS,
	})

	if err != nil {
		log.Fatalln(err)
	}
	if c.Bucket == "" {
		panic("no bucket name")
	}
	ok, err := conn.BucketExists(context.Background(), c.Bucket)
	if err != nil {
		log.Fatalln(err)
	}

	timeout := 10 * time.Second
	if c.Timeout != 0 {
		timeout = time.Duration(c.Timeout) * time.Second
	}

	if !ok {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err = conn.MakeBucket(ctx, c.Bucket, minio.MakeBucketOptions{Region: c.Location})
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &Client{
		conn:        conn,
		Bucket:      c.Bucket,
		Location:    c.Location,
		Timeout:     timeout,
		ContentType: c.ContentType,
	}
}

// Upload 上传/更新文件或者图片
func (c *Client) Upload(ctx context.Context, f io.Reader, name string, size int64) error {
	// 将文件上传到Minio服务器
	_, err := c.conn.PutObject(ctx, c.Bucket, name, f, size, minio.PutObjectOptions{
		ContentType: c.ContentType,
	})

	return errors.WithStack(err)
}

// GetUrl 获取地址
func (c *Client) GetUrl(ctx context.Context, name string, ops Options) string {

	params := url.Values{}
	expires := 30 * time.Minute

	if !ops.IsDownload {
		if ops.Width != 0 && ops.Height != 0 {
			params.Set("resize", fmt.Sprintf("%dx%d", ops.Width, ops.Height))
		}
		params.Set("response-content-disposition", "inline")
		params.Set("response-content-type", "image/jpeg")
	}

	if ops.Expires != 0 {
		expires = time.Duration(ops.Expires) * time.Minute
	}
	res, err := c.conn.PresignedGetObject(ctx, c.Bucket, name, expires, params)
	if err != nil {
		logx.Error(err)
		return ""
	}

	return res.String()
}

// GetUrls 批量获取地址
func (c *Client) GetUrls(ctx context.Context, names []string, ops Options) map[string]string {

	params := url.Values{}
	if !ops.IsDownload {
		if ops.Width != 0 && ops.Height != 0 {
			params.Set("resize", fmt.Sprintf("%dx%d", ops.Width, ops.Height))
		}
		params.Set("response-content-disposition", "inline")
		params.Set("response-content-type", "image/jpeg")
	}
	expires := 30 * time.Minute
	if ops.Expires != 0 {
		expires = time.Duration(ops.Expires) * time.Minute
	}

	var data = make(map[string]string)
	for _, name := range names {
		res, err := c.conn.PresignedGetObject(ctx, c.Bucket, name, expires, params)
		if err != nil {
			logx.Error(err)
			continue
		}
		data[name] = res.String()
	}

	return data
}

// DelObject 删除对象
func (c *Client) DelObject(ctx context.Context, names ...string) error {

	for _, name := range names {
		err := c.conn.RemoveObject(ctx, c.Bucket, name, minio.RemoveObjectOptions{})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// UploadFile 上传文件
func (c *Client) UploadFile(ctx context.Context, r *http.Request, prefix string) (string, error) {
	file, head, err := r.FormFile("file")
	if err != nil {
		return "", errors.WithStack(err)
	}

	defer file.Close()
	name := idx.GenUUID()
	tem := strings.Split(head.Filename, ".")
	suffix := tem[len(tem)-1]
	c.ContentType = imagex.SwitchContentType(suffix)
	if len(tem) > 1 {
		name += "." + suffix
	}

	if prefix != "" {
		name = prefix + "/" + name
	}
	err = c.Upload(ctx, file, name, head.Size)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return name, nil
}

package httputil

import (
	"net/http"
	"net"
	"time"
	"strings"
	"net/url"
	"github.com/liuchonglin/go-utils"
	"encoding/json"
	"os"
	"mime/multipart"
	"bytes"
	"io"
)

const (
	FormContentType      = "application/x-www-form-urlencoded;charset=utf-8"
	JsonContentType      = "application/json;charset=utf-8"
	MultipartContentType = "multipart/form-data"
)

// http请求类
type httpRequest struct {
	// 网关
	Url string `json:"url"`
	// 数据类型
	ContentType string `json:"contentType"`
	// 请求头
	Header map[string]string `json:"header"`
	// 请求体
	Body map[string]string `json:"body"`
	// cookie数组
	Cookies []*http.Cookie `json:"cookies"`
	// http客户端
	Client *http.Client `json:"client"`
}

// 创建http请求
func NewHttpRequest(url string) *httpRequest {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 2 * time.Second,
		}).DialContext,
	}

	client := &http.Client{Transport: transport}
	return &httpRequest{
		Url:         url,
		ContentType: FormContentType,
		Client:      client,
	}
}

// 设置数据类型
func (h *httpRequest) SetContentType(contentType string) *httpRequest {
	if utils.IsEmpty(contentType) {
		contentType = FormContentType
	}
	h.ContentType = contentType
	return h
}

// 设置请求头
func (h *httpRequest) SetHeader(header map[string]string) *httpRequest {
	h.Header = header
	return h
}

// 设置请求体
func (h *httpRequest) SetBody(body map[string]string) *httpRequest {
	h.Body = body
	return h
}

// 设置请求超时时间，这里的超时指的是连接超时
func (h *httpRequest) SetTimeout(timeout time.Duration) *httpRequest {
	h.Client.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: timeout,
		}).DialContext,
	}
	return h
}

// 设置cookies
func (h *httpRequest) SetCookies(cookies []*http.Cookie) *httpRequest {
	h.Cookies = cookies
	return h
}

// Get请求
func (h *httpRequest) Get() (*http.Response, error) {
	return h.request(http.MethodGet, nil)
}

// Post请求
func (h *httpRequest) Post() (*http.Response, error) {
	return h.request(http.MethodPost, nil)
}

// 文件上传
func (h *httpRequest) UploadFile(fileName, path string) (*http.Response, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	formFileW, err := w.CreateFormFile("fileName", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFileW, file)
	if err != nil {
		return nil, err
	}

	for k, v := range h.Body {
		w.WriteField(k, v)
	}

	h.ContentType = MultipartContentType
	return h.request(http.MethodPost, body)
}

// 请求
func (h *httpRequest) request(method string, body io.Reader) (*http.Response, error) {
	// 处理请求体
	if h.ContentType == FormContentType { // form表单
		values := url.Values{}
		for k, v := range h.Body {
			values.Add(k, v)
		}
		body = strings.NewReader(values.Encode())
	} else if h.ContentType == JsonContentType { // json格式
		jsonData, err := json.Marshal(h.Body)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(jsonData))
	}

	// 创建请求
	req, err := http.NewRequest(method, h.Url, body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// 添加请求头
	for k, v := range h.Header {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}

	// 添加Cookie
	for _, v := range h.Cookies {
		req.AddCookie(v)
	}

	return h.Client.Do(req)
}

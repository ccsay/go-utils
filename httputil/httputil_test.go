package httputil

import (
	"net/http"
	"reflect"
	"testing"
	"time"
	"fmt"
	"github.com/liuchonglin/go-tools/token"
	"liuchonglin.com/baseCms/controllers/contextUtil"
	"github.com/satori/go.uuid"
)

func TestNewHttpRequest(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *httpRequest
	}{
		{
			name: "ok",
			args: args{
				url: "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHttpRequest(tt.args.url)
			fmt.Printf("httputil.TestNewhttpRequest() = %v\n", got)
		})
	}
}

func TestHttpRequest_Get(t *testing.T) {
	get1 := NewHttpRequest("https://www.cnblogs.com/mafeng/p/7068837.html")

	token, err := jsonWebToken.New(&jsonWebToken.TokenConfig{Issuer: "liu"}).CreateToken(contextUtil.CreateTokenData("admin", "13000008888", 1))
	if err != nil {
		panic(err)
	}
	get2 := NewHttpRequest("http://localhost:8080/v1/manager/get/1").SetHeader(map[string]string{"token": token, "traceId": uuid.Must(uuid.NewV1()).String()})
	type fields struct {
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *http.Response
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Url:         get1.Url,
				ContentType: get1.ContentType,
				Header:      get1.Header,
				Body:        get1.Body,
				Client:      get1.Client,
			},
			wantErr: false,
		}, {
			name: "RESTful get",
			fields: fields{
				Url:         get2.Url,
				ContentType: get2.ContentType,
				Header:      get2.Header,
				Body:        get2.Body,
				Client:      get2.Client,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			got, err := h.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("httpRequest.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("httpRequest.Get() = %v\n", got)
		})
	}
}

func TestHttpRequest_Post(t *testing.T) {
	headerMap := map[string]string{"traceId": uuid.Must(uuid.NewV1()).String()}
	bodyMap := map[string]string{"username": "admin", "password": "123456"}
	post1 := NewHttpRequest("http://localhost:8080/v1/login/login").SetHeader(headerMap).SetBody(bodyMap)
	post2 := NewHttpRequest("http://localhost:8080/v1/login/login").SetHeader(headerMap).SetBody(bodyMap).SetContentType(ContentTypeJson)
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *http.Response
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Url:         post1.Url,
				ContentType: post1.ContentType,
				Header:      post1.Header,
				Body:        post1.Body,
				Client:      post1.Client,
			},
			wantErr: false,
		}, {
			name: "json",
			fields: fields{
				Url:         post2.Url,
				ContentType: post2.ContentType,
				Header:      post2.Header,
				Body:        post2.Body,
				Client:      post2.Client,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			got, err := h.Post()
			if (err != nil) != tt.wantErr {
				t.Errorf("httpRequest.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("httpRequest.Post() = %v\n", got)
		})
	}
}

func TesthttpRequest_request(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	type args struct {
		method string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			got, err := h.request(tt.args.method, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpRequest.request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest.request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TesthttpRequest_SetContentType(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	type args struct {
		contentType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *httpRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			if got := h.SetContentType(tt.args.contentType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest.SetContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TesthttpRequest_SetHeader(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	type args struct {
		header map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *httpRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			if got := h.SetHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest.SetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TesthttpRequest_SetBody(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	type args struct {
		body map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *httpRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			if got := h.SetBody(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest.SetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TesthttpRequest_SetTimeout(t *testing.T) {
	type fields struct {
		Method      string
		Url         string
		ContentType string
		Header      map[string]string
		Body        map[string]string
		Client      *http.Client
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *httpRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpRequest{
				Url:         tt.fields.Url,
				ContentType: tt.fields.ContentType,
				Header:      tt.fields.Header,
				Body:        tt.fields.Body,
				Client:      tt.fields.Client,
			}
			if got := h.SetTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest.SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

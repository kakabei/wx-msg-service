package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// Handle 参考chatgpt实现
func LogHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("Failed to read request body: %v", err)
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		logx.WithContext(r.Context()).Debugf("[Request]: %s %s %+v %s", r.Method, r.RequestURI, r.Header, body)

		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}
		next(recorder, r)

		logx.WithContext(r.Context()).Debugf("[Response]: %s %s %s cost:%d", r.Method, r.RequestURI, string(recorder.body), time.Since(startTime).Milliseconds())

	}
}

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}

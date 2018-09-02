package gin

import (
	"io/ioutil"
	"net/http"

	"peckergo/api/model"
	"peckergo/api/utils/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// WriteGinJSON gin 写入 JSON 数据
func WriteGinJSON(c *gin.Context, code int, v interface{}) {
	c.Render(code, JSON{Data: v})
}

// BindGinJSON gin 读取 JSON 数据
func BindGinJSON(c *gin.Context, v interface{}) error {
	return c.MustBindWith(v, jsonBinding{c})
}

// 重写jsonBinding

type jsonBinding struct {
	c *gin.Context
}

func (jsonBinding) Name() string {
	return "json"
}

func (j jsonBinding) Bind(req *http.Request, obj interface{}) error {
	u, _ := j.c.Get("user")
	var user model.User
	if u != nil {
		user = u.(model.User)
	}
	if req.Method == "GET" || &user == nil || user.ID == 0 {
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(obj); err != nil {
			return err
		}
	} else {
		body, err := ioutil.ReadAll(j.c.Request.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
		model.NewLogManagement(&model.LogManagement{
			UserID: user.ID,
			Path:   j.c.Request.RequestURI,
			Method: j.c.Request.Method,
			Body:   string(body),
		})
	}
	return validate(obj)
}

func validate(obj interface{}) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

// 重写 render.JSON
type (
	// JSON 重写JSON
	JSON struct {
		Data interface{}
	}
)

var jsonContentType = []string{"application/json; charset=utf-8"}

// Render JSON 渲染
func (r JSON) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType 写入 JSON 头
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON 写JSON
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write([]byte(jsonBytes))
	return nil
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

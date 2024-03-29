package response

import (
	`bytes`
	`encoding/json`
	`errors`
	`github.com/chaodoing/app/assets`
	`github.com/kataras/iris/v12`
	`html/template`
	`log`
)

type Response struct {
	ctx  iris.Context
	data interface{}
}

// renderHTML 渲染成html
func (r Response) renderHTML() (html string, err error) {
	var data []byte
	data, err = assets.Asset("code/index.html")
	if err != nil {
		return
	}
	tpl, err := template.New("json").Parse(string(data))
	if err != nil {
		return
	}
	data, err = json.Marshal(r.data)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	theme := r.ctx.URLParamDefault("theme", "vs-dark")
	err = tpl.Execute(buf, map[string]string{
		"Title": "JSON",
		"Json":  string(data),
		"Theme": theme,
	})
	return buf.String(), err
}

func (r Response) Xml() (err error) {
	//r.ctx.Gzip(true)
	err = r.ctx.XML(r.data, iris.XML{Indent: "\t", Prefix: ""})
	return
}

func (r Response) Json() (err error) {
	err = r.ctx.JSON(r.data, iris.JSON{Indent: "\t", Prefix: ""})
	return
}

func (r Response) Data(status uint32, message string, data interface{}) Response {
	r.data = Data{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return r
}

// Send 发送数据
func (r Response) Send() (err error) {
	if r.data == nil {
		return errors.New("输出的数据不能为空")
	}
	html, err := r.renderHTML()
	if err != nil {
		return err
	}
	r.ctx.Negotiation().JSON(r.data).XML(r.data).HTML(html).EncodingGzip()
	_, err = r.ctx.Negotiate(nil)
	return
}

func (r Response) Set(ctx iris.Context, data interface{}) {
	r.ctx = ctx
	r.data = data
	err := r.Send()
	if err != nil {
		log.Println(err)
	}
}

// O 输出内容
func (r Response) O(ctx iris.Context, status uint32, message string, data interface{}) {
	r.ctx = ctx
	err := r.Data(status, message, data).Send()
	if err != nil {
		log.Println(err)
	}
}

// Pagination 输出分页内容
func (r Response) Pagination(ctx iris.Context, status uint32, message string, data interface{}, total uint64, page uint, limit uint) {
	r.ctx = ctx
	r.data = Pagination{
		Status:  status,
		Message: message,
		Total:   total,
		Page:    page,
		Limit:   limit,
		Data:    data,
	}
	err := r.Send()
	if err != nil {
		log.Println(err)
	}
}

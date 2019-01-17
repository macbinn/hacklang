package builtin

import (
	"encoding/json"
	"fmt"
	"github.com/macbinn/hacklang/value"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type WebContext struct {
	Req *http.Request
	w http.ResponseWriter
	Body *String
	JsonBody value.Object
	Query value.Object
	data map[string]value.Object
}

func (w *WebContext) Get(name string) value.Object {
	switch name {
	case "json":
		return NewFunction("webContext.json", func(args ...value.Object) value.Object {
			w.Json(Convert(args[0]))
			return nil
		})
	case "resp":
		return NewFunction("webContext.resp", func(args ...value.Object) value.Object {
			code := args[0].(*Number).Int
			contentType := args[1].(*String).S
			body := args[2].(*String).S
			w.Resp(code, contentType, body)
			return nil
		})
	case "static":
		return NewFunction("webContext.static", func(args ...value.Object) value.Object {
			file := args[0].(*String).S
			bs, err := ioutil.ReadFile(file)
			if err != nil {
				w.Resp(500, "", "")
				return nil
			}
			w.Resp(200, "text/html", string(bs))
			return nil
		})
	case "header": return NewFunction("webContext.header", func(args ...value.Object) value.Object {
		name := args[0].(*String).S
		return NewString(w.Req.Header.Get(name))
	})
	case "body": return w.readBody()
	case "jsonBody": return w.ReadJson()
	case "query": return w.ReadQuery()
	default:
		return w.data[name]
	}
}

func (w *WebContext) Set(name string, value value.Object) {
	w.data[name] = value
}

func (w *WebContext) Repr() string {
	return "<WebContext>"
}

func (w *WebContext) error(err error) {
	http.Error(w.w, err.Error(), 500)
}

func (w *WebContext) Json(v interface{}) {
	w.w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w.w).Encode(v)
	if err != nil {
		w.error(err)
	}
}

func (w *WebContext) ReadJson() value.Object {
	if w.Req.Method != http.MethodPost {
		return nil
	}
	if w.JsonBody == nil {
		var m map[string]interface{}
		err := json.NewDecoder(w.Req.Body).Decode(&m)
		if err != nil {
			log.Printf("read json error %s", err)
			w.error(err)
			return nil
		}
		w.JsonBody = Obj(m)
	}
	return w.JsonBody
}

func (w *WebContext) ReadQuery() value.Object {
	if w.Query == nil {
		query := w.Req.URL.Query()
		m := map[string]interface{}{}
		for name, vaules := range query {
			m[name] = vaules[0]
		}
		w.Query = Obj(m)
	}
	return w.Query
}

func (w *WebContext) Resp(code int, contentType string, body string) {
	w.w.Header().Set("Content-Type", contentType)
	w.w.WriteHeader(code)
	_, err := fmt.Fprint(w.w, body)
	if err != nil {
		w.error(err)
	}
}

// Body return cached body string, it can be called with multi times
func (w *WebContext) readBody() *String {
	if w.Body == nil {
		bs, err := ioutil.ReadAll(w.Req.Body)
		if err != nil {
			w.error(err)
			return nil
		}
		w.Body = NewString(string(bs))
	}
	return w.Body
}

func NewWebContext(w http.ResponseWriter, r *http.Request) *WebContext {
	return &WebContext{
		w: w,
		Req: r,
		data: map[string]value.Object{},
	}
}

type Router struct {
	Pattern string
	Regexp *regexp.Regexp
	Get func(...value.Object) value.Object
	Post func(...value.Object) value.Object
}

var routers = map[string]Router{}

func Route(args ...value.Object) value.Object {
	param := args[0].(*Map).Val
	var getHandler, postHandler func(...value.Object) value.Object
	pattern := param["url"].(*String).S
	get, ok := param["get"]
	if ok {
		getHandler = get.(*Function).fn
	}
	post, ok := param["post"]
	if ok {
		postHandler = post.(*Function).fn
	}
	p, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	router := Router{
		Pattern: pattern,
		Regexp: p,
		Get: getHandler,
		Post: postHandler,
	}
	routers[pattern] = router
	return nil
}

var prepareHandlers []func(...value.Object) value.Object

func Prepare(args ...value.Object) value.Object {
	handler := args[0].(*Function).fn
	prepareHandlers = append(prepareHandlers, handler)
	return nil
}

type handler struct {

}

func prepare(ctx *WebContext) {
	for _, handler := range prepareHandlers {
		handler(ctx)
	}
}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routerFound := false
	for _, router := range routers {
		matched := router.Regexp.FindStringSubmatch(r.URL.Path)
		if matched == nil {
			continue
		}
		//log.Printf("matched: %v", matched)
		routerFound = true
		ctx := NewWebContext(w, r)
		args := []value.Object{
			ctx,
		}
		for _, s := range matched[1:] {
			n, err := strconv.Atoi(s)
			if err == nil {
				args = append(args, NewNumber(n))
			} else {
				args = append(args, NewString(s))
			}
		}
		var handler func(...value.Object) value.Object
		switch r.Method {
		case http.MethodGet:
			handler = router.Get
		case http.MethodPost:
			handler = router.Post
		default:
			w.WriteHeader(405)
		}
		if handler != nil {
			prepare(ctx)
			handler(args...)
		}
	}
	if !routerFound {
		w.WriteHeader(404)
	}
}

func Run(args ...value.Object) value.Object {
	addr := args[0].(*String).S
	http.ListenAndServe(addr, handler{})
	return nil
}

var WebExports = NewMap(map[string]value.Object{
	"route": NewFunction("web.route", Route),
	"run": NewFunction("web.run", Run),
	"prepare": NewFunction("web.prepare", Prepare),
})

func init() {
	GlobalScope.Register("web", WebExports)
}
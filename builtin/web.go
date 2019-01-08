package builtin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type WebContext struct {
	Req *http.Request
	w http.ResponseWriter
	Body *String
	JsonBody Object
	Query Object
}

func (w *WebContext) Get(name string) Object {
	switch name {
	case "json":
		return NewFunction("webContext.json", func(args ...Object) Object {
			w.Json(Convert(args[0]))
			return nil
		})
	case "resp":
		return NewFunction("webContext.resp", func(args ...Object) Object {
			code := args[0].(*Number).Int
			contentType := args[1].(*String).s
			body := args[2].(*String).s
			w.Resp(code, contentType, body)
			return nil
		})
	case "body":
		return w.readBody()
	case "jsonBody":
		return w.ReadJson()
	case "query":
		return w.ReadQuery()
	}
	return nil
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

func (w *WebContext) ReadJson() Object {
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

func (w *WebContext) ReadQuery() Object {
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

type Router struct {
	Pattern string
	Regexp *regexp.Regexp
	Get func(...Object) Object
	Post func(...Object) Object
}

var routers = map[string]Router{}

func Route(args ...Object) Object {
	pattern := args[0].(*String).s
	getHandler := args[1].(*Function).fn
	postHandler := args[2].(*Function).fn
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

type handler struct {

}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routerFound := false
	for _, router := range routers {
		if !router.Regexp.MatchString(r.URL.Path) {
			continue
		}
		routerFound = true
		ctx := &WebContext{
			Req: r,
			w: w,
		}
		switch r.Method {
		case http.MethodGet:
			router.Get(ctx)
		case http.MethodPost:
			router.Post(ctx)
		default:
			w.WriteHeader(405)
		}
	}
	if !routerFound {
		w.WriteHeader(404)
	}
}

func Run(args ...Object) Object {
	addr := args[0].(*String).s
	http.ListenAndServe(addr, handler{})
	return nil
}

var WebExports = NewMap(map[string]Object{
	"route": NewFunction("web.route", Route),
	"run": NewFunction("web.run", Run),
})
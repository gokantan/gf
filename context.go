package gf

import (
	"fmt"
	"github.com/goframework/gf/cfg"
	"github.com/goframework/gf/ext"
	"github.com/goframework/gf/sessions"
	"net/http"
	"database/sql"
)

type Context struct {
	w              http.ResponseWriter
	r              *http.Request
	vars           map[string]interface{}
	isSelfResponse bool

	Config         *cfg.Cfg
	RouteVars      map[string]ext.VarType
	FinishFilter   bool
	RedirectPath   string
	RedirectStatus int
	Session        *sessions.Session
	ViewBases      []string
	View           string
	ViewData       map[string]interface{}
	UrlPath        string
	Method         string
	IsGetMethod    bool
	IsPostMethod   bool
	Form           map[string][]string

	DB             *sql.DB
}

func (ctx *Context) Redirect(path string) {
	ctx.RedirectPath = path
	ctx.RedirectStatus = http.StatusFound
}

func (ctx *Context) RedirectPermanently(path string) {
	ctx.RedirectPath = path
	ctx.RedirectStatus = http.StatusMovedPermanently
}

func (ctx *Context) Set(key string, value interface{}) {
	ctx.vars[key] = value
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	val, ok := ctx.vars[key]
	return val, ok
}

func (ctx *Context) GetString(key string) (string, bool) {
	val, ok := ctx.vars[key]
	if ok {
		str, ok := val.(string)
		return str, ok
	}
	return "", false
}

func (ctx *Context) GetInt(key string) (int, bool) {
	val, ok := ctx.vars[key]
	if ok {
		intval, ok := val.(int)
		return intval, ok
	}
	return 0, false
}

func (ctx *Context) GetBool(key string) (bool, bool) {
	val, ok := ctx.vars[key]
	if ok {
		boolval, ok := val.(bool)
		return boolval, ok
	}
	return false, false
}

func (ctx *Context) Write(data []byte) (int, error) {
	ctx.isSelfResponse = true
	return ctx.w.Write(data)
}

func (ctx *Context) WriteS(output string) {
	ctx.isSelfResponse = true
	fmt.Fprint(ctx.w, output)
}

func (ctx *Context) Writef(format string, content ...interface{}) {
	ctx.isSelfResponse = true
	fmt.Fprintf(ctx.w, format, content...)
}

func (ctx *Context) ClearSession() {
	for k := range ctx.Session.Values {
		delete(ctx.Session.Values, k)
	}
}

func (ctx *Context) NewSession() {
	ctx.ClearSession()

	var err error
	ctx.Session, err = mCookieStore.New(ctx.r, SERVER_SESSION_ID)
	if err != nil {
		http.Error(ctx.w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctx *Context) GetResponseWriter() http.ResponseWriter {
	return ctx.w
}

func (ctx *Context) AddResponseHeader(key string, value string) {
	ctx.w.Header().Add(key, value)
}

package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *HttpRouter {
	return &HttpRouter{
		router: mux.NewRouter(),
	}
}

type HttpMW func(http.HandlerFunc) http.HandlerFunc

type Router interface {
	GET(pattern string, handler http.HandlerFunc, mw ...HttpMW)
	POST(pattern string, handler http.HandlerFunc, mw ...HttpMW)
	PUT(pattern string, handler http.HandlerFunc, mw ...HttpMW)
	DELETE(pattern string, handler http.HandlerFunc, mw ...HttpMW)
	GetServeHTTP() http.HandlerFunc
}

type HttpRouter struct {
	router *mux.Router
}

// func panicCatcher(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var response responses.Response

// 		defer func() {
// 			if err := recover(); err != nil {
// 				response = responses.InternalError
// 				response.Message = err.(string)
// 				response.WriteJson(w)
// 				return
// 			}
// 		}()

// 		h(w, r)
// 	}
// }

func (r *HttpRouter) ChainsMiddlewares(method string, handler http.HandlerFunc, mw ...HttpMW) http.HandlerFunc {

	// middleware chains
	mws := make([]HttpMW, 0)
	// mws = append(mws, panicCatcher)
	mws = append(mws, mw...)

	for _, httpMW := range mws {
		handler = httpMW(handler)
	}
	return handler
}

// func Handle404(w http.ResponseWriter, r *http.Request) {
// 	var resp responses.Response
// 	defer resp.WriteJson(w)

// 	resp.Code = 404
// 	resp.Message = "Путь не найден"
// }

func (r *HttpRouter) Handle(method string, pattern string, handler http.HandlerFunc, mw ...HttpMW) {
	handler = r.ChainsMiddlewares(method, handler, mw...)
	r.router.HandleFunc(pattern, handler).Methods(method, "OPTIONS")
	// r.router.NotFoundHandler = http.HandlerFunc()
}

func (r *HttpRouter) GET(pattern string, handler http.HandlerFunc, mw ...HttpMW) {
	r.Handle("GET", pattern, handler, mw...)
}

func (r *HttpRouter) POST(pattern string, handler http.HandlerFunc, mw ...HttpMW) {
	log.Println(handler)
	r.Handle("POST", pattern, handler, mw...)
}

func (r *HttpRouter) PUT(pattern string, handler http.HandlerFunc, mw ...HttpMW) {
	r.Handle("PUT", pattern, handler, mw...)
}

func (r *HttpRouter) DELETE(pattern string, handler http.HandlerFunc, mw ...HttpMW) {
	r.Handle("DELETE", pattern, handler, mw...)
}

func (r *HttpRouter) GetServeHTTP() http.HandlerFunc {
	return r.router.ServeHTTP
}

func (s *HttpRouter) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	log.Println(request)
	s.router.ServeHTTP(write, request)
}

// func (s *HttpRouter) ConnectSwagger(sw http.Handler) {
// 	s.router.PathPrefix("/docs").Handler(sw)
// }
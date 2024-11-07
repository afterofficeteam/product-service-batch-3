package routes

import (
	"codebase-service/config"
	"codebase-service/util/middleware"
	"log"
	"net/http"
	"strings"
	"time"

	product "codebase-service/handlers/products"
	user "codebase-service/handlers/users"

	"github.com/spf13/viper"
)

type Routes struct {
	Router  *http.ServeMux
	User    *user.Handler
	Product *product.Handler
}

func URLRewriter(baseURLPath string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, baseURLPath)

		next.ServeHTTP(w, r)
	}
}

func (r *Routes) SetupBaseURL() {
	baseURL := viper.GetString("BASE_URL_PATH")
	if baseURL != "" && baseURL != "/" {
		r.Router.HandleFunc(baseURL+"/", URLRewriter(baseURL, r.Router))
	}
}

func (r *Routes) SetupRouter() {
	r.Router = http.NewServeMux()
	r.SetupBaseURL()
	r.userRoutes()
	r.productRoutes()
}

func (r *Routes) userRoutes() {
	r.Router.HandleFunc("POST /signup", middleware.ApplyMiddleware(r.User.SignUpByEmail, middleware.EnabledCors, middleware.LoggerMiddleware()))
	r.Router.Handle("POST /signin", middleware.ApplyMiddleware(r.User.SignInByEmail, middleware.EnabledCors, middleware.LoggerMiddleware()))
}

func (r *Routes) productRoutes() {
	r.Router.HandleFunc("GET /products/{id}", middleware.ApplyMiddleware(r.Product.GetProduct, middleware.EnabledCors, middleware.LoggerMiddleware()))
	r.Router.HandleFunc("GET /products", middleware.ApplyMiddleware(r.Product.GetProducts, middleware.EnabledCors, middleware.LoggerMiddleware()))

	r.Router.HandleFunc("POST /products", middleware.ApplyMiddleware(r.Product.CreateProduct, middleware.GetUserId, middleware.EnabledCors, middleware.LoggerMiddleware()))
	r.Router.HandleFunc("DELETE /products/{id}", middleware.ApplyMiddleware(r.Product.DeleteProduct, middleware.GetUserId, middleware.EnabledCors, middleware.LoggerMiddleware()))
}

func (r *Routes) Run(port string) {
	r.SetupRouter()

	log.Printf("[Running-Success] clients on localhost on port :%s", port)
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "localhost:" + port,
		WriteTimeout: config.WriteTimeout() * time.Second,
		ReadTimeout:  config.ReadTimeout() * time.Second,
	}

	log.Panic(srv.ListenAndServe())
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"models"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	// 接收Json的api
	r.HandleFunc("/articles/add", ArticleHandler_Add).Methods("POST")

	//----- More advanced options -----
	//match path prefixes
	// r.PathPrefix("/products/")

	//HTTP methods
	// r.Methods("GET", "POST")

	//URL schemes
	// r.Schemes("https")

	//header values
	// r.Headers("X-Requested-With", "XMLHttpRequest")

	//query values: "http://news.domain.com/articles/technology/42?key=value"
	// r.Queries("key", "value")

	//use a custom matcher function
	// r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
	// 	return r.ProtoMajor == 0
	// })

	//combine several matchers in a single route
	// r.HandleFunc("/products", ProductsHandler).
	// 	Host("www.example.com").
	// 	Methods("GET").
	// 	Schemes("http")

	//假设我们有多个网址只能在主机为www.example.com时匹配。 为该主机创建路由并从中获取“子路由器”：
	// r := mux.NewRouter()
	// s := r.Host("www.example.com").Subrouter()
	// s.HandleFunc("/products/", ProductsHandler)
	// s.HandleFunc("/products/{key}", ProductHandler)
	// s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	//构建URL
	// r := mux.NewRouter()
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).
	// 	Name("article")
	// url, err := r.Get("article").URL("category", "technology", "id", "42")
	// //上面的等同于"/articles/technology/42"

	//日志框架uber-go/zap
	//日志框架的配置
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "./api.log"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	fmt.Println("api server startup ok! port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello, golang!\n")
	w.Write([]byte("hello, golang!\n"))
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(vars["category"])
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(vars["id"])
	fmt.Fprintf(w, "id: %v\n", vars["id"])
}

func ArticleHandler_Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body) //定义json converter
	var article models.Article
	err := decoder.Decode(&article) //反序列化并赋值

	if err != nil {
		panic(err)
	}

	fmt.Println(article.Title)
	fmt.Fprintf(w, "title: %v\n", article.Title)
}

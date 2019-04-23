package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type shortUrl struct {
	Path string `yaml: path`
	Url  string `yaml: url`
}

func main() {
	var shortUrlList []shortUrl
	v := viper.New()
	// add config path, this function could be called multi times to add multi path
	v.AddConfigPath("./")
	// add file, without the extention name
	v.SetConfigName("redirect")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to read config, err:%s", err))
	}
	err = v.UnmarshalKey("data", &shortUrlList)
	if err != nil {
		panic(fmt.Sprintf("Failed to decoder data into urlMapList:%s", err))
	}
	urlMap := buildMap(shortUrlList)

	// Init a default mux, for those unset shortUtl will always to mux
	mux := defaultMux()
	handleFunc := MapHandler(urlMap, mux)
	http.ListenAndServe(":8080", handleFunc)
}

func buildMap(urls []shortUrl) map[string]string {
	ret := make(map[string]string, len(urls))
	for _, pu := range urls {
		ret[pu.Path] = pu.Url
	}
	return ret
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

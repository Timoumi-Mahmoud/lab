package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func main() {
	fmt.Println("vim-go")
	//template.Must(parse))
	re := regexp.MustCompile(`\d+`)

	fmt.Println(re.MatchString("12345"))
	data := Must(os.ReadFile("reg.go"))
	fmt.Println(string(data))

}

//#############################################################//

func MustQuery(r *http.Request, key string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		panic(fmt.Sprintf("Query parameter '%s' is required", key))
	}
	return val
}

// #############################################################//
func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}
	return val
}

// #############################################################//
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

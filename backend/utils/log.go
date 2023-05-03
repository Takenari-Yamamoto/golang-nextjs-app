package utils

import (
	"fmt"
	"net/http"
)

func LogRequestHeaders(r *http.Request) {
	for name, values := range r.Header {
		// リクエストヘッダー内のすべての値を表示する
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}

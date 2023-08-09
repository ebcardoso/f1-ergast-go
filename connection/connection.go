package connection

import (
	"fmt"
	"net/http"

	"github.com/ebcardoso/f1-ergast-go/utils"
)

func Get(path string, query string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s.json%s", utils.ROOT_URL, path, query)

	resp, err := http.Get(url)
	return resp, err
}

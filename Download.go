package fm

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(_url, file_name string) (file_name2 string, err error) {
	reqest, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Mobile Safari/537.36")
	//处理返回结果
	res, err := client.Do(reqest)
	if err != nil {
		return
	}

	if !strings.Contains(file_name, ".") {
		file_name += ".jpg"
	}
	f, err := os.Create(file_name)

	if err != nil {
		return
	}
	_, err = io.Copy(f, res.Body)
	file_name2 = file_name
	return
}

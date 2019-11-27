package fm

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(_url, file_path, file_name string) (file_name2 string, err error) {
	b, err := PathExists(file_path)
	if err != nil {
		return
	}
	if !b {
		err = os.MkdirAll(file_path, os.ModePerm)
		if err != nil {
			return
		}
	}
	file_ := file_path + file_name
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
		file_ += ".jpg"
	}
	f, err := os.Create(file_)

	if err != nil {
		return
	}
	_, err = io.Copy(f, res.Body)
	file_name2 = strings.ReplaceAll(file_, file_path, "")
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

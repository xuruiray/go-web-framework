package config

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
)

func Init(file string) error {

	var (
		err error
	)

	err = LoadConfig(file, &MainConfig)
	if err != nil {
		return err
	}

	return nil
}

// LoadConfig 解析json文件
func LoadConfig(filename string, result interface{}) error {
	var (
		bytes []byte
		err   error
	)

	path := strings.Join([]string{ConfigBasePath, filename}, "/")

	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	re1 := regexp.MustCompile("^#.*$")
	lines := strings.Split(string(bytes), "\n")
	var loc []int
	var cnt string
	for _, line := range lines {
		loc = re1.FindStringIndex(line)
		if len(loc) == 0 {
			cnt += line
		}
	}

	err = json.Unmarshal([]byte(cnt), result)
	return err
}

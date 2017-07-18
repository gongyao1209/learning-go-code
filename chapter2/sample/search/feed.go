package search

import (
	"encoding/json"
	"os"
)

//本地文件的绝对路径
const dataFile = "/Users/gongyao/workspace/go_code1/code/chapter2/sample/data/data.json";
//const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.	打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once	当函数返回时，关闭文件
	// the function returns.
	// defer 关键字，安排随后的函数调用在函数返回时才执行
	defer file.Close()

	// Decode the file into a slice of pointers	将文件解码到一个切片里面
	// to Feed values.	这个切片每一项都是一个指向一个 Feed类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	// 这个函数不需要检查错误，调用者会这么做
	return feeds, err
}

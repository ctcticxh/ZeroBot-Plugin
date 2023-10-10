package yuantu

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/FloatTech/floatbox/file"
	"github.com/FloatTech/floatbox/web"
	"github.com/tidwall/gjson"
)

var api = "https://api.lolicon.app/setu/v2" + "?tag=" + url.QueryEscape("魔法少女まどか☆マギカ")
var imgIndexToPath = make(map[int]string)
var imgMap = make(map[string]string)

// 该爬虫每30分钟爬取一次图片
func spider() {
	readJson()
	writeJson()
	for {
		imageUrl, err := getimgurl(api)
		fmt.Println(len(imgMap))
		if err != nil {

		}
		resp, err := http.Get(imageUrl)
		if err != nil {
			// fmt.Println("ERROR:", err)
		}

		imageBase := CalculateHash(imageUrl)
		if imgMap[imageUrl] == "" {
			imgMap[imageUrl] = baseURL + imageBase
			imgIndexToPath[len(imgIndexToPath)] = baseURL + imageBase + imageUrl[len(imageUrl)-4:]
			fmt.Println(imageUrl)

			file.DownloadTo(imageUrl, base+imageBase+imageUrl[len(imageUrl)-4:])
		}

		resp.Body.Close()
		writeJson()
		time.Sleep(time.Duration(1) * time.Minute)
	}
}

func getimgurl(url string) (string, error) {
	data, err := web.GetData(url)
	if err != nil {
		return "", err
	}
	json := gjson.ParseBytes(data)
	if e := json.Get("error").Str; e != "" {
		return "", errors.New(e)
	}
	var imageurl string
	if imageurl = json.Get("data.0.urls.original").Str; imageurl == "" {
		return "", errors.New("未找到相关内容, 换个tag试试吧")
	}
	return strings.ReplaceAll(imageurl, "i.pixiv.cat", "i.pixiv.re"), nil
}

// 计算图片地址的哈希值
func CalculateHash(hashValue string) string {
	h := sha1.New()
	h.Write([]byte(hashValue))
	hashValue = hex.EncodeToString(h.Sum(nil))
	return hashValue
}

func writeJson() error {
	jsonData, err := json.Marshal(imgMap)
	if err != nil {
		return err
	}

	// 将 JSON 数据写入本地文件
	err = os.WriteFile(base+"imgMap.json", jsonData, 0644)
	if err != nil {
		return err
	}
	jsonData, err = json.Marshal(imgIndexToPath)
	if err != nil {
		return err
	}

	// 将 JSON 数据写入本地文件
	err = os.WriteFile(base+"imgIndexToPath.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readJson() error {
	// 从本地文件读取 JSON 数据
	jsonData, err := os.ReadFile(base + "imgMap.json")
	if err != nil {
		return err
	}
	// 将 JSON 数据解码为哈希表
	err = json.Unmarshal(jsonData, &imgMap)
	if err != nil {
		return err
	}

	// 从本地文件读取 JSON 数据
	jsonData, err = os.ReadFile(base + "imgIndexToPath.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 将 JSON 数据解码为哈希表
	err = json.Unmarshal(jsonData, &imgIndexToPath)
	if err != nil {
		return err
	}

	return nil
}

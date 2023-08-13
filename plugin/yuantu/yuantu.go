// 圆图
package yuantu

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	engine = control.Register("chat", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "来份圆图",
		Help:             "圆图\n- 来份圆图\n- 圆图十连\n- 更新图库\n",
	})
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func init() { // 插件主体
	engine.OnFullMatch("来份圆图").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("还没做好哦~"))
		})
	engine.OnFullMatch("更新图库").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			start := 1
			end := 10
			working(start, end)
		})
}

// 获取一个网页所有的内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])

	}
	return

}

// 写入文件
func SaveJokeFile(url string) {
	captcha := strconv.Itoa(int(hash(url)))
	//保存在本地的地址
	path := "file/" + captcha + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	defer f.Close()

	//读取url的信息
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//写入文件
		f.Write(buf[:n])
	}

}

func working(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d \n", start, end)

	page := make(chan int) //设置多线程

	for i := start; i <= end; i++ {
		go SpidePage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}
}

func SpidePage(i int, page chan int) {
	//网站每一页的改变
	url := "https://wall.alphacoders.com/by_sub_category.php?id=297132&name=Honkai+Impact+3rd+Wallpapers&page=" + strconv.Itoa(i*1)
	//读取这个页面的所有信息
	result, err := HttpGet(url)
	//判断是否出错，并打印信息
	if err != nil {
		fmt.Println("SpidePage err:", err)
	}

	//正则表达式提取信息
	str := `<div class="thumb-container-big " id="thumb_(?s:(.*?))">`
	//解析、编译正则
	ret := regexp.MustCompile(str)
	//提取需要信息-每一个图片的数字
	urls := ret.FindAllStringSubmatch(result, -1)

	for _, jokeURL := range urls {
		//组合每个图片的url
		joke := `https://wall.alphacoders.com/big.php?i=` + jokeURL[1]

		//爬取图片的url
		tuUrl, err := SpideJokePage(joke)
		if err != nil {
			fmt.Println("tuUrl err:", err)
			continue
		}

		SaveJokeFile(tuUrl)

	}

	//防止主go程提前结束
	page <- i
}

// 爬取图片放大的页面
func SpideJokePage(url string) (tuUrl string, err error) {
	//爬取网站的信息
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		fmt.Println("SpidePage err:", err)
	}

	str := `<img class="main-content" src="(?s:(.*?))"`
	//解析、编译正则
	ret := regexp.MustCompile(str)
	//提取需要信息-每一个段子的url
	alls := ret.FindAllStringSubmatch(result, -1)
	for _, temTitle := range alls {
		tuUrl = temTitle[1]
		break
	}

	return
}

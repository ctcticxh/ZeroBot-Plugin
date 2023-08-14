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
	"github.com/gocolly/colly/v2"
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
			//SpidePage()
			go main()
		})
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", h.Text, link)
		c.Visit(h.Request.AbsoluteURL(link))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://wall.alphacoders.com/by_sub_category.php?id=297132&name=Honkai+Impact+3rd+Wallpapers")
}

// 获取一个网页所有的内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	os.WriteFile("test.txt", body, 0664)
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
	path := "E:/imgs/" + captcha + ".jpg"
	fmt.Println(path)
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
func SpidePage() {
	fmt.Println("爬爬爬1")
	url := "https://wall.alphacoders.com/by_sub_category.php?id=297132&name=Honkai+Impact+3rd+Wallpapers"
	//读取这个页面的所有信息
	result, err := HttpGet(url)
	//判断是否出错，并打印信息
	if err != nil {
		fmt.Println("SpidePage err:", err)
	}
	fmt.Println("result :", len(result))

	//正则表达式提取信息
	str := `<div class="thumb-container-big " id="thumb_(?s:(.*?))" itemprop="associatedMedia" itemscope itemtype="http://schema.org/ImageObject">`
	fmt.Println("爬爬爬2")
	//解析、编译正则
	ret := regexp.MustCompile(str)
	//提取需要信息-每一个图片的数字
	urls := ret.FindAllStringSubmatch(result, -1)
	fmt.Println("爬爬爬3")
	fmt.Println(len(urls))
	for _, jokeURL := range urls {
		//组合每个图片的url
		joke := `https://wall.alphacoders.com/big.php?i=` + jokeURL[1]
		fmt.Println(joke)
		//爬取图片的url
		tuUrl, err := SpideJokePage(joke)
		fmt.Println("在爬了")
		if err != nil {
			fmt.Println("tuUrl err:", err)
			continue
		}
		fmt.Println(tuUrl)
		SaveJokeFile(tuUrl)

	}
}

// 爬取图片放大的页面
func SpideJokePage(url string) (tuUrl string, err error) {
	//爬取网站的信息
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		fmt.Println("SpidePage err:", err)
	}

	str := `<meta itemprop="image" content="(.*)">`
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

package chatgpt1

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	proxyURL           = "https://open.aiproxy.xyz/v1/"
	modelGPT3Dot5Turbo = "gpt-3.5-turbo"
	MyKey              = "471548b54e6625f0a4915a390a9bd5db"
	URL                = "https://api.a20safe.com/api.php?api=36&key=%s&text=%s&gptkey=%s"
	gptkey             = "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"
)

var (
	engine = control.Register("chatgpt1", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "chatgpt1",
		Help: "-@bot chatgpt [对话内容]\n" +
			"- 添加预设xxx xxx\n" +
			"- 设置(默认)预设xxx\n" +
			"- 删除本群预设\n" +
			"- 查看预设列表\n" +
			"- 余额查询\n" +
			"- (私聊发送)设置OpenAI apikey [apikey]\n" +
			"- (私聊发送)删除apikey\n" +
			"- (群聊发送)(授权|取消)(本群|全局)使用apikey\n" +
			"注:先私聊设置自己的key,再授权群聊使用,不会泄露key的\n",
	})
	pre = ""
)

func init() {
	engine.OnRegex(`^#([\s\S]*)$`).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			args := ctx.State["regex_matched"].([]string)[1]
			question := pre + args
			data, err := web.GetData(fmt.Sprintf(URL, MyKey, question, gptkey))
			if err != nil {
				ctx.SendChain(message.Text("Error:", err))
			}
			reply := gjson.ParseBytes(data).Get("data.#.reply").Array()[0]
			ctx.SendChain(message.Text(reply))
		})
	/*engine.OnRegex(`^设置\s*OpenAI\s*apikey\s*(.*)$`, zero.OnlyPrivate).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("保存apikey成功"))
	})
	engine.OnFullMatch("删除apikey").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("保存apikey成功"))
		})*/
}

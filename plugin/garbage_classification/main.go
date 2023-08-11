// 垃圾分类
package garbageclassification

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const url = "https://api.a20safe.com/api.php?api=49&key=471548b54e6625f0a4915a390a9bd5db&word=%v"

func init() {
	engine := control.Register("garbageclassification", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "垃圾分类",
		Help:             "xxx是什么垃圾\n",
	})
	engine.OnFullMatch("疯狂星期四").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		data, err := web.GetData(url)
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return
		}
		ctx.SendChain(message.Text(gjson.ParseBytes(data).Get("data.#.result").Array()[0]))

	})
	engine.OnSuffix("是什么垃圾").Handle(func(ctx *zero.Ctx) {
		args := ctx.State["args"].(string)
		data, err := web.GetData(fmt.Sprintf(url, args))
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return
		}
		result := gjson.ParseBytes(data).Get("data.#.result").Array()[0].String()
		if result == "未找到结果" {
			ctx.SendChain(message.Text(result))
			return
		}
		ctx.SendChain(message.Text(args, "是", result))
	})
}

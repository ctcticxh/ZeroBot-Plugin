// Package kfccrazythursday 疯狂星期四
package kfccrazythursday

import (
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	crazyURL = "https://api.a20safe.com/api.php?api=7&key=471548b54e6625f0a4915a390a9bd5db"
)

func init() {
	engine := control.Register("kfccrazythursday", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "疯狂星期四",
		Help:             "疯狂星期四\n",
	})
	engine.OnFullMatch("疯狂星期四").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		data, err := web.GetData(crazyURL)
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return
		}
		ctx.SendChain(message.Text(gjson.ParseBytes(data).Get("data.#.result").Array()[0]))
		//fmt.Println(gjson.ParseBytes(data).Get("data"))
		//fmt.Println(gjson.ParseBytes(data).Get("data.#.result").Array()[0])
	})
}

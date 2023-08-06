// 历史上的今天
package history_day

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type user struct {
	Name string `flag:"n"`
	Age  int    `flag:"a"`
}

const (
	history_day_url = "https://api.suwanya.cn/api/lishi?format=%v"
)

func init() {
	engine := control.Register("history_day", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "历史上的今天（bug版）",
		Help:             "历史上的今天（bug版）\n- 历史上的今天",
	})
	engine.OnFullMatch("历史上的今天").Handle(func(ctx *zero.Ctx) {
		//matched := ctx.State["matched"].(string)
		//ctx.SendChain(message.Text("完全匹配的匹配词: ", matched))
		format := "json"
		data, err := web.GetData(fmt.Sprint(history_day_url, format))
		if err != nil {
			ctx.SendChain(message.Text("Error:", err))
		}
		ctx.SendChain(message.Text(string(data)))
	})
}

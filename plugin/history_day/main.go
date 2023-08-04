// 历史上的今天
package history_day

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type user struct {
	Name string `flag:"n"`
	Age  int    `flag:"a"`
}

const (
	history_day_url = "https://api.suwanya.cn/api/lishi"
)

func init() {
	engine = control.Register("history_day", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "历史上的今天",
		Help:             "历史上的今天\n- 历史上的今天",
	})
	engine.OnFullMatch("历史上的今天").Handle(func(ctx *zero.Ctx) {
		matched := ctx.State["matched"].(string)
		ctx.SendChain(message.Text("完全匹配的匹配词: ", matched))
		data, err := web.GetData(fmt.Sprint(history_day_url))
		if err != nil {
			ctx.SendChain(message.Text("Error:", err))
		}
		ctx.SendChain(message.JSON(data))
	})
}
func main() {
	zero.RunAndBlock(&zero.Config{
		NickName:      []string{"吼姆拉"},
		CommandPrefix: "/",
		SuperUsers:    []int64{3574736597},
		Driver: []zero.Driver{
			driver.NewWebSocketClient("ws://127.0.0.1:6700/", ""),
		},
	}, nil)
}

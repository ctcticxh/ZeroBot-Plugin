// 历史上的今天
package history_day

import (
	"fmt"
	"strconv"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
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
		data, err := web.GetData(fmt.Sprintf(history_day_url, format))
		if err != nil {
			ctx.SendChain(message.Text("Error:", err))
		}
		content_array := gjson.ParseBytes(data).Get("content").Array()
		str := "历史上的今天：\n"
		for i := 0; i < len(content_array); i++ {
			str += strconv.Itoa(i+1) + ". " + content_array[i].String()
			if i < len(content_array)-1 {
				str += "\n"
			}
		}
		ctx.SendChain(message.Text(str))
	})
}

// Package answer 解答之书
package answer

import (
	"fmt"
	"os"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	engine = control.Register("answer", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "解答之书",
		Help:             "解答 [xxx]",
	})
)

func init() { // 插件主体
	engine.OnPrefix("解答").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := os.Open("answer.json")

			if err != nil {
				ctx.SendChain(message.Text("Error:", err))
			}

			defer data.Close()

			fmt.Println(data)

		})
}

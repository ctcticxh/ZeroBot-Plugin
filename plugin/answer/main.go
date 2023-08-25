// Package answer 解答之书
package answer

import (
	"math/rand"
	"os"

	"github.com/FloatTech/floatbox/file"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	engine = control.Register("answer", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "解答之书",
		Help:              "解答 [xxx]",
		PrivateDataFolder: "answer",
	})

	base    = engine.DataFolder()
	baseurl = "file:///" + file.BOTPATH + "/" + base
)

func init() { // 插件主体
	engine.OnPrefix("解答").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := os.ReadFile(baseurl + "answer.json")

			if err != nil {
				ctx.SendChain(message.Text("Error:", err))
			}

			//answer := gjson.ParseBytes(data).Get("#.info")

			i := rand.Intn(134)

			ctx.SendChain(message.At(ctx.Event.UserID), message.Text(gjson.ParseBytes(data).Get("#.info").Array()[i]))

			ctx.SendChain(message.Text(baseurl))
		})
}

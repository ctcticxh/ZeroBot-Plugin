// Package yuantu 圆图
package yuantu

import (
	"math/rand"

	"github.com/FloatTech/floatbox/file"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	en = control.Register("yuantu", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "来份圆图",
		Help:              "圆图\n- 来份圆图\n- 圆图十连\n- 更新图库\n",
		PrivateDataFolder: "yuantu",
	})
	base    = en.DataFolder()
	baseURL = "file:///" + file.BOTPATH + "/" + base
)

func init() {
	en.OnFullMatch("来份圆图").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			url := imgIndexToPath[rand.Intn(len(imgIndexToPath))]
			ctx.SendChain(message.Image(url))
		})
	en.OnFullMatch("更新圆图").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ch := make(chan int)
			// ctx.SendChain(message.Text("开始更新圆图啦~"))
			go spider()
			var discard = 10
			discard += 100
			discard = <-ch
		})
}

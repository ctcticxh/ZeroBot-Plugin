// 赚钱
package earn

import (
	"math/rand"

	"github.com/FloatTech/AnimeAPI/wallet"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	engine = control.Register("earn", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "赚钱",
		Help:              "- \n- 打工\n- 乞讨",
		PrivateDataFolder: "earn",
	})
)

// 技能CD记录表
type cdsheet struct {
	Time    int64  // 时间
	GroupID int64  // 群号
	UserID  int64  // 用户
	ModeID  string // 技能类型
}

func init() {
	engine.OnFullMatch("打工", zero.OnlyGroup).Limit(ctxext.LimitByGroup).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			uid := ctx.Event.UserID
			add := rand.Intn(11)
			go func() {
				err := wallet.InsertWalletOf(uid, add)
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			}()
			if add != 0 {
				ctx.SendChain(message.Text("本次打工获得了：", add, " 个atri币哦~"))
			} else {
				ctx.SendChain(message.Text("本次打工一无所得哦~"))
			}
		})
	engine.OnFullMatch("乞讨", zero.OnlyGroup).Limit(ctxext.LimitByGroup).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			uid := ctx.Event.UserID
			add := rand.Intn(11)
			go func() {
				err := wallet.InsertWalletOf(uid, add)
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
			}()
			if add != 0 {
				ctx.SendChain(message.Text("本次乞讨获得了：", add, " 个atri币哦~"))
			} else {
				ctx.SendChain(message.Text("本次乞讨一无所得哦~"))
			}
		})

}

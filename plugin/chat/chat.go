// Package chat 对话插件
package chat

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	pokes = [...][2]string{
		{"请不要戳", " >_<"},
		{"喂(#`O′) 戳", "干嘛！"},
		{"再戳", "就要生气啦！"},
		{"再戳", "，就要诅咒你啦(╯‵□′)╯︵┻━┻"},
		{"一直戳", "你一定是变态对吧？！对吧？！"}}
	poke   = rate.NewManager[int64](time.Minute*5, 8) // 戳一戳
	engine = control.Register("chat", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "基础反应, 群空调",
		Help:             "chat\n- [BOT名字]\n- [戳一戳BOT]\n- 空调开\n- 空调关\n- 群温度\n- 设置温度[正整数]",
	})
)

func init() { // 插件主体
	// 被喊名字
	engine.OnFullMatch("", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			time.Sleep(time.Second * 1)
			ctx.SendChain(message.Text(
				[]string{
					nickname + "在此，有何贵干~",
					"(っ●ω●)っ在~",
					"这里是" + nickname + "(っ●ω●)っ",
					nickname + "不在呢~",
				}[rand.Intn(4)],
			))
		})
	// 戳一戳
	engine.On("notice/notify/poke", zero.OnlyToMe).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			nickname_order := rand.Intn(len(zero.BotConfig.NickName))
			var nickname = zero.BotConfig.NickName[nickname_order]
			/*switch {
			case poke.Load(ctx.Event.GroupID).AcquireN(3):
				// 5分钟共8块命令牌 一次消耗3块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("请不要戳", nickname, " >_<"))
			case poke.Load(ctx.Event.GroupID).Acquire():
				// 5分钟共8块命令牌 一次消耗1块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("喂(#`O′) 戳", nickname, "干嘛！"))
			default:
				// 频繁触发，不回复
			}*/
			poke_order := rand.Intn(len(pokes))
			var t_poke = pokes[poke_order]
			ctx.SendChain(message.Text(t_poke[0], nickname, t_poke[1]))
			ctx.SendChain(message.Poke(ctx.Event.UserID))

		})
	// 群空调
	var AirConditTemp = map[int64]int{}
	var AirConditSwitch = map[int64]bool{}
	engine.OnFullMatch("空调开").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = true
			ctx.SendChain(message.Text("❄️哔~"))
		})
	engine.OnFullMatch("空调关").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = false
			delete(AirConditTemp, ctx.Event.GroupID)
			ctx.SendChain(message.Text("💤哔~"))
		})
	engine.OnRegex(`设置温度(\d+)`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				temp := ctx.State["regex_matched"].([]string)[1]
				AirConditTemp[ctx.Event.GroupID], _ = strconv.Atoi(temp)
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
	engine.OnFullMatch(`群温度`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
	engine.OnRegex(`^(.*)(呜呜)+(.*)$`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("哭什么，没用的东西，给你一拳！"))
		})
	engine.OnSuffix("家人们").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("谁是你家人？"))
		})
	engine.OnPrefix("家人们").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("谁是你家人？"))
		})
	engine.OnRegex(`^我是(.*)$`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			sta := ctx.State["regex_matched"].([]string)[1]
			ctx.SendChain(message.Text("好的，你是", sta))
		})
	engine.OnPrefix("我觉得").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		thought := ctx.State["args"].(string)
		ctx.SendChain(message.Text("是的，", thought))
	})
	engine.OnPrefix("复读").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		content := ctx.State["args"].(string)
		ctx.SendChain(message.Text(content))
	})

	engine.OnFullMatch("图图图").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		index := rand.Intn(10000)
		true_index := index + 1300000
		url := "https://images7.alphacoders.com/131/" + strconv.Itoa(true_index) + ".jpeg"
		rsp, err := http.Get(url)
		if err != nil {
			ctx.SendChain(message.Text("ERROR", err))
		}
		defer rsp.Body.Close()
		for rsp.StatusCode == http.StatusNotFound {
			fmt.Println("查找中")
			index := rand.Intn(10000)
			true_index := index + 1300000
			url := "https://images7.alphacoders.com/131/" + strconv.Itoa(true_index) + ".jpeg"
			rsp, err = http.Get(url)
			if err != nil {
				ctx.SendChain(message.Text("ERROR", err))
			}
			defer rsp.Body.Close()
		}
		ctx.SendChain(message.Image(url))
		fmt.Println(url)
	})

}

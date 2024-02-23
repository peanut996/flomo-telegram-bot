package main

var (
	SuccessTip = "🟢 发送成功"

	FailTip = "🔴 发送失败"

	StartTip = "🤖 欢迎使用Flomo机器人"

	PingTip = "🏓 Pong!"

	AuthTip = "🔒 请先发送你的flomo api，进行绑定。\n发送格式：/bind https://flomoapp.com/xxx/xxxx/xxxxxx/ \n\n✨ 如何获取 api? https://flomoapp.com/minesource=incoming_webhook \n\n🔗 更多帮助：https://help.flomoapp.com/advance/extension/tgbot.html"

	BindTip = "🟢 绑定成功\n发送文字，即可保存到Flomo"

	InvalidFlomoAPITip = "🔴 绑定失败\n请检查你的flomo api是否正确"

	UnBindTip = "🟢 解绑成功"
)

var (
	TgBotToken string
)

var (
	TgBotEnvKey = "TG_BOT_TOKEN"
)

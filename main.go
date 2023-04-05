package main

import (
	. "apihandler/pkg/qiita"
	"os"
)

func main() {
	discord := os.Getenv("DISCORD_WH_1")

	GetQiitaTrend(discord)
}

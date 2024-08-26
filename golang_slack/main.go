package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/slack-go/slack"
)

func main() {

	// 日付を生成する
	now := time.Now()
	year := now.Year()
	month := now.Month()

	// 指定した月の月初と月末を取得
	firstDay := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)

	fmt.Println("月初:", firstDay.Format("2006/01/02"))
	fmt.Println("月末:", lastDay.Format("2006/01/02"))

	// 日付と時刻をフォーマットして変数化
	start_date := firstDay.Format("2006/01/02")
	to_date := lastDay.Format("2006/01/02")

	// 途中で止める
	//os.Exit(0)

	// URLを生成する
	baseURL := "https://connpass.com/search/"
	// クエリパラメータを作成
	params := url.Values{}
	params.Add("q", "sre")
	params.Add("start_from", start_date)
	params.Add("start_to", to_date)
	params.Add("sort", "")

	// ベースURLにクエリパラメータを追加
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 結果を表示
	fmt.Println("生成されたURL:", fullURL)

	// 途中で止める
	//os.Exit(0)

	// 投稿する文章生成
		comment_template := fmt.Sprintf(
			"来月のSREセミナー情報は<%s|こちら>\n興味があるセミナーがないかチェックしてみてね:eyes:",
			fullURL,
		)

	// アクセストークンを使用してクライアントを生成する
	tkn := "{slackAppsToken}"
	c := slack.New(tkn)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage(
		"#random",
		slack.MsgOptionText(comment_template, false),
		slack.MsgOptionDisableLinkUnfurl(),
		slack.MsgOptionDisableMediaUnfurl(),
	)
	if err != nil {
		panic(err)
	}
}

package file_bot

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func upload_file() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4462018505568-5973904069431-akh8FN883tYeR7cRVu9qUjMV")
	os.Setenv("CHANNEL_ID", "C04D8U36JV7")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"DS.pdf", "CC57.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}

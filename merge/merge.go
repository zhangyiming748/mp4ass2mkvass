package merge

import (
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"mp4ass2mkvass/util"
	"os"
	"os/exec"
)

func MkvWithAss(mp4, srt, mkv string) {
	cmd := exec.Command("ffmpeg", "-i", mp4, "-i", srt, "-c:v", "libx265", "-ac", "1", "-c:s", "ass", mkv)
	log.Printf("生成的命令: %s\n", cmd.String())

	mi := FastMediaInfo.GetStandMediaInfo(mp4)
	frame := mi.Video.FrameCount
	if err := util.ExecCommand(cmd, frame); err != nil {
		log.Fatalf("程序运行出错:%v\n", err)
	} else {
		os.Remove(mp4)
		os.Remove(srt)
	}
}

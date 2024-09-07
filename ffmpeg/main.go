package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ffmpeg")
  // ffmpeg -t "00:00:05" -i "rtsp://admin:password@192.168.88.111:554/ISAPI/Streaming/Channels/101" "./room.mp4"
	cmd.Args = append(
		cmd.Args,
    "-t", "00:00:05",
    "-i", "rtsp://admin:password@192.168.1.111:554/ISAPI/Streaming/Channels/101",
    "./room.mp4",
	)

	fmt.Println("cmd", cmd, cmd.Args)

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println("Output:", string(out))
}

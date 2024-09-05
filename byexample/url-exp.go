package main

import (
	"fmt"
	"net/url"
)

func urlExp() {
	fmt.Println("URL Experiments")

  // rtsp://admin:password@192.168.88.105:554/ISAPI/Streaming/Channels/101
	url := url.URL{
    Scheme: "rtsp",
    User: url.UserPassword("admin", "pass"),
    Path: "ISAPI/Streaming/Channels/101",
    // http://admin:password@192.168.88.105/ISAPI/Streaming/channels/101/picture?snapShotImageType=JPEG
    RawQuery: "picture?snapShotImageType=JPEG",
    Host: "192.168.88.79:552",
  }

  fmt.Println("URL: ", url.String())
}

package main

import (
	"flag"
	"log"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "localhost:8000", "host to connect to/listen on")
	flag.Parse()
	log.Printf("MCP server listening on %s", url)
	// 开启多个协议，控制退出
	runToolServer(url)
}

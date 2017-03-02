package main

import "github.com/lizhengqiang/hot"

func main() {
	docs := &hot.Docs{LocalRoot: "./docs", GitTarget: "git@github.com:lizhengqiang/hot.git"}
	docs.ReloadDocs()
}

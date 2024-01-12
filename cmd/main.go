package main

import (
	"flag"

	"github.com/lflxp/lflxp-static/pkg"
	"github.com/lflxp/lflxp-static/pkg/sync"
)

var (
	portHttpStatic                string
	staticPortHttpStatic          string
	pathHttpStatic                string
	isVideo                       bool
	pagesize                      int
	types                         string
	raw, debug                    bool
	src, dest, clean, cleanString string
	username, password            string
)

func init() {
	flag.StringVar(&portHttpStatic, "p", "9090", "service port")
	flag.StringVar(&staticPortHttpStatic, "P", "9091", "static service port")
	flag.StringVar(&pathHttpStatic, "f", "./", "load catalogs")
	flag.StringVar(&types, "t", ".avi,.wma,.rmvb,.rm,.mp4,.mov,.3gp,.mpeg,.mpg,.mpe,.m4v,.mkv,.flv,.vob,.wmv,.asf,.asx", "Filter video types, multiple separated by commas")
	flag.BoolVar(&isVideo, "v", false, "Switch to video mode or not")
	flag.BoolVar(&raw, "r", false, "Switch to Raw HTTP URL")
	flag.BoolVar(&debug, "debug", false, "是否开启断点续传debug日志")
	flag.IntVar(&pagesize, "c", 20, "Videos per page")
	flag.StringVar(&src, "src", "", "复制文件原文件或文件夹")
	flag.StringVar(&dest, "dest", "", "复制文件目标文件或文件夹")
	flag.StringVar(&clean, "clean", "", "删除文件或文件夹,如：/tmp")
	flag.StringVar(&cleanString, "contains", "", "删除文件名包含的内容，如：.temp")
	flag.StringVar(&username, "username", "", "用户名")
	flag.StringVar(&password, "password", "", "密码")
	flag.Parse()
}

func main() {
	if src != "" && dest != "" {
		err := sync.LocalDirSync(src, dest, debug)
		if err != nil {
			panic(err)
		}
	} else if clean != "" && cleanString != "" {
		err := sync.Clean(clean, cleanString)
		if err != nil {
			panic(err)
		}
	} else {
		api := pkg.Apis{
			StaticPort: staticPortHttpStatic,
			Port:       portHttpStatic,
			Path:       pathHttpStatic,
			Types:      types,
			IsVideo:    isVideo,
			PageSize:   pagesize,
			Raw:        raw,
			Username:   username,
			Password:   password,
		}
		err := api.Execute()
		if err != nil {
			panic(err)
		}
	}
}

package main

import (
	"flag"

	"github.com/lflxp/lflxp-static/pkg"
)

var (
	portHttpStatic       string
	staticPortHttpStatic string
	pathHttpStatic       string
	isVideo              bool
	pagesize             int
	types                string
	raw                  bool
)

func init() {
	flag.StringVar(&portHttpStatic, "p", "9090", "service port")
	flag.StringVar(&staticPortHttpStatic, "P", "9091", "static service port")
	flag.StringVar(&pathHttpStatic, "f", "./", "load catalogs")
	flag.StringVar(&types, "t", ".avi,.wma,.rmvb,.rm,.mp4,.mov,.3gp,.mpeg,.mpg,.mpe,.m4v,.mkv,.flv,.vob,.wmv,.asf,.asx", "Filter video types, multiple separated by commas")
	flag.BoolVar(&isVideo, "v", false, "Switch to video mode or not")
	flag.BoolVar(&raw, "r", false, "Switch to Raw HTTP URL")
	flag.IntVar(&pagesize, "c", 20, "Videos per page")
	flag.Parse()
}

func main() {
	api := pkg.Apis{
		StaticPort: staticPortHttpStatic,
		Port:       portHttpStatic,
		Path:       pathHttpStatic,
		Types:      types,
		IsVideo:    isVideo,
		PageSize:   pagesize,
		Raw:        raw,
	}
	err := api.Execute()
	if err != nil {
		panic(err)
	}
}

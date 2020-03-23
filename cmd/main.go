package main

import (
	"flag"

	"github.com/lflxp/lflxp-static/pkg"
)

var (
	portHttpStatic string
	pathHttpStatic string
	isVideo        bool
	pagesize       int
	types          string
)

func init() {
	flag.StringVar(&portHttpStatic, "p", "9090", "service port")
	flag.StringVar(&pathHttpStatic, "f", "./", "load catalogs")
	flag.StringVar(&types, "t", ".avi,.wma,.rmvb,.rm,.mp4,.mov,.3gp,.mpeg,.mpg,.mpe,.m4v,.mkv,.flv,.vob,.wmv,.asf,.asx", "Filter video types, multiple separated by commas")
	flag.BoolVar(&isVideo, "v", false, "Switch to video mode or not")
	flag.IntVar(&pagesize, "c", 20, "Videos per page")
	flag.Parse()
}

func main() {
	api := pkg.Apis{
		Port:     portHttpStatic,
		Path:     pathHttpStatic,
		Types:    types,
		IsVideo:  isVideo,
		PageSize: pagesize,
	}

	err := api.Check()
	if err != nil {
		panic(err)
	}

	api.Execute()
}

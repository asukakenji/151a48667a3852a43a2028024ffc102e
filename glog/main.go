package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	isCustom := false
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "logtostderr" {
			isCustom = true
		}
	})
	if !isCustom {
		flag.Set("logtostderr", "true")
	}
	glog.Info("Hello")
}

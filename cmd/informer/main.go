/*
modification history
--------------------
2018/9/7, by lovejoy, create
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/lovejoy/informer/cmd/informer/app"
	"github.com/golang/glog"
)

const (
	APISERVER_VERSION = "Informer v0.1.1"
)

func main() {
	showVersion := flag.Bool("version", false, "show version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("%v\n", APISERVER_VERSION)
		os.Exit(0)
	}
	defer glog.Flush()
	app.Run()
}

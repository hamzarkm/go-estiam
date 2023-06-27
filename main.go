package main

import (
	"flag"
	"imgexo/filter"
	"imgexo/task"
)

func main() {

	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	t := task.NewChanTask(*srcDir, *dstDir, f, 16)
	t.Process()

}

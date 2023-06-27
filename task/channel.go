package task

import (
	"fmt"
	"imgexo/filter"
	"path"
	"path/filepath"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {

	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, res chan<- string) {
	for j := range jobs {
		fmt.Printf("worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		res <- j.dst
	}
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	res := make(chan string, size)

	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, res)
	}

	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-res)
	}

	return nil
}

package main

/**
5-3
 */
type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longlived bool
	Dispatcher fn
	Executor fn
}

func NewRunner(size int,longlived bool,d fn,e fn) *Runner {
	return &Runner{
		Controller:make(chan string,1),
		Error:make(chan string,1),
		Data:make(chan interface{},size),
		longlived:longlived,
		dataSize:size,
		Dispatcher:d,
		Executor:e,
	}
}
func (r *Runner) startDispatch() {
	//go 匿名函数写法  第一个括号是参数类型，第二个括号是参数传值
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	//go selector use
	for true {
		select {
		//	一旦r.Controller:有内容，则写入c
			case c:= <-r.Controller:
				if c == READY_TO_DISPATCH {
					err := r.Dispatcher(r.Data)
					if err != nil {
						r.Error <- CLOSE
					}else {
						r.Controller <- READY_TO_EXECUTE
					}
				}
				if c == READY_TO_EXECUTE {
					err := r.Executor(r.Data)
					if err != nil {
						r.Error <- CLOSE
					}else {
						r.Controller <- READY_TO_DISPATCH
					}
				}
			case e:= <-r.Error:
				if e == CLOSE {
					return
				}
			default:
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
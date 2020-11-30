package pool

// App 协程池本体
type App struct {
	// 协程总数
	Routine int
	// 私有计数器
	counter int
	// 协程完成消息通道
	doneMessage chan int
	// 任务开启前执行函数
	beforeFunc AppFunc
	// 任务结束后执行函数
	afterFunc AppFunc
	// 任一协程结束时触发函数
	onceDoneFunc AppFunc
	// 设置任务区间
	task [2]int
}

// RunFunc 协程运行时函数
type RunFunc func(a *App, index int)

// AppFunc 协程开关函数
type AppFunc func(a *App)

// New 创建协程池
func New(routine int) *App {
	return &App{
		Routine:     routine,
		counter:     routine,
		doneMessage: make(chan int, routine),
	}
}

// do 执行协程运行时函数
func (a *App) do(doFunc RunFunc) {
	for i := 0; i < a.Routine; i++ {
		go func(index int) {
			if doFunc != nil {
				doFunc(a, index)
			}
			a.doneMessage <- -1
		}(i)
	}
}

// Before 设置任务开启前执行函数
func (a *App) Before(beforeFunc AppFunc) *App {
	if a.beforeFunc == nil && beforeFunc != nil {
		a.beforeFunc = beforeFunc
	}
	return a
}

// After 设置任务结束后执行函数
func (a *App) After(afterFunc AppFunc) *App {
	if a.afterFunc == nil && afterFunc != nil {
		a.afterFunc = afterFunc
	}
	return a
}

// OnceDone 设置任一协程结束时触发函数
func (a *App) OnceDone(onceDoneFunc AppFunc) *App {
	if a.onceDoneFunc == nil && onceDoneFunc != nil {
		a.onceDoneFunc = onceDoneFunc
	}
	return a
}

// Counter 获取剩余协程数
func (a *App) Counter() int {
	return a.counter
}

// Done 判断是否结束所有协程
func (a *App) Done() bool {
	return a.counter == 0
}

// Run 协程调度
func (a *App) Run(doFunc RunFunc) *App {

	// 开始时机函数
	if a.beforeFunc != nil {
		a.beforeFunc(a)
	}

	// 开启协程
	go a.do(doFunc)

	// 检测所有协程是否结束
	for {

		if -1 == <-a.doneMessage {
			// 计数器减一
			a.counter--

			// 触发函数
			if a.onceDoneFunc != nil {
				a.onceDoneFunc(a)
			}

			if a.counter == 0 {
				// 关闭通道
				close(a.doneMessage)
				break
			}
		}

	}

	// 结束时机函数
	if a.afterFunc != nil {
		a.afterFunc(a)
	}

	return a
}

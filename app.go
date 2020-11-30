package pool

// App 协程池本体
type App struct {
	Routine int
	// 私有计数器
	counter int
	// 协程完成消息通道
	doneMessage chan int

	// 任务开启前执行函数
	beforeFunc AppFunc

	// 任务结束后执行函数
	afterFunc AppFunc

	// 设置任务区间
	task [2]int
}

// RunFunc 协程运行时函数
type RunFunc func(app *App, index int)

// AppFunc 协程开关函数
type AppFunc func(app *App)

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

// Run 协程调度
func (a *App) Run(doFunc RunFunc) *App {

	if a.beforeFunc != nil {
		a.beforeFunc(a)
	}

	// 开启协程
	go a.do(doFunc)

	// 检测所有协程是否结束
	for a.counter != 0 {

		if -1 == <-a.doneMessage {
			// 计数器减一
			a.counter--
		}

	}

	if a.afterFunc != nil {
		// 关闭通道
		close(a.doneMessage)
		a.afterFunc(a)
	}

	return a
}

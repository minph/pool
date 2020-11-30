# pool

简易 Golang 协程并发任务库

`import "github.com/minph/pool"`

# 示例

``` go
package main

import (
	"fmt"

	"github.com/minph/pool"
)

func main() {
	p := pool.New(10)

	data := make(chan string)

	go p.Run(func(app *pool.App, index int) {

		msg := fmt.Sprintf("这是协程%v", index)
		data <- msg
	})

	for !p.Done() {
		msg := <-data
		fmt.Println(msg)
	}
}

```

# 概览

- [type App](#App)
  - [func New(routine int) \*App](#New)
  - [func (a *App) After(afterFunc AppFunc) *App](#App.After)
  - [func (a *App) Before(beforeFunc AppFunc) *App](#App.Before)
  - [func (a \*App) Counter() int](#App.Counter)
  - [func (a \*App) Done() bool](#App.Done)
  - [func (a *App) OnceDone(onceDoneFunc AppFunc) *App](#App.OnceDone)
  - [func (a *App) Run(doFunc RunFunc) *App](#App.Run)
  - [func (a *App) SetTask(from, to int) *App](#App.SetTask)
  - [func (a \*App) Task(index int) (int, int)](#App.Task)
- [type AppFunc](#AppFunc)
- [type RunFunc](#RunFunc)

# 详情

## <a name="App">type</a> App

```go
type App struct {
    // 协程总数
    Routine int
    // contains filtered or unexported fields
}

```

App 协程池本体

### <a name="New">func</a> New

```go
func New(routine int) *App
```

New 创建协程池

### <a name="App.After">func</a> (\*App) After

```go
func (a *App) After(afterFunc AppFunc) *App
```

After 设置任务结束后执行函数

### <a name="App.Before">func</a> (\*App) Before

```go
func (a *App) Before(beforeFunc AppFunc) *App
```

Before 设置任务开启前执行函数

### <a name="App.Counter">func</a> (\*App) Counter

```go
func (a *App) Counter() int
```

Counter 获取剩余协程数

### <a name="App.Done">func</a> (\*App) Done

```go
func (a *App) Done() bool
```

Done 判断是否结束所有协程

### <a name="App.OnceDone">func</a> (\*App) OnceDone

```go
func (a *App) OnceDone(onceDoneFunc AppFunc) *App
```

OnceDone 设置任一协程结束时触发函数

### <a name="App.Run">func</a> (\*App) Run

```go
func (a *App) Run(doFunc RunFunc) *App
```

Run 协程调度

### <a name="App.SetTask">func</a> (\*App) SetTask

```go
func (a *App) SetTask(from, to int) *App
```

SetTask 设置任务区间

### <a name="App.Task">func</a> (\*App) Task

```go
func (a *App) Task(index int) (int, int)
```

Task 获取子任务区间

## <a name="AppFunc">type</a> AppFunc

```go
type AppFunc func(a *App)
```

AppFunc 协程开关函数

## <a name="RunFunc">type</a> RunFunc

```go
type RunFunc func(a *App, index int)
```

RunFunc 协程运行时函数

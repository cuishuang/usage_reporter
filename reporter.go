package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

var (

	// 要检测的进程名称;默认为当前进程
	pName string
	// 时间间隔；默认1s
	timeInterval int64
	// 是否去除其他信息，仅保留最基本信息
	pure bool

	// 要检测的进程id
	pid int
)

func main() {

	//go func() {
	//	fmt.Println(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", "6060"), nil))
	//}()

	flag.StringVar(&pName, "n", "", "要监测的进程名称（默认为当前进程）")
	flag.Int64Var(&timeInterval, "t", 1, "时间间隔（默认1s）")

	flag.BoolVar(&pure, "p", false, "pure output |  是否去除其他信息，仅保留基本信息")

	flag.IntVar(&pid, "pid", 0, "要监测的进程id")

	flag.Parse()

	fmt.Println("用户输入的参数数量:", flag.NFlag())

	fmt.Println("pure is:", pure)

	if len(pName) == 0 && pid == 0 {
		pid = os.Getpid()
		fmt.Println("当前程序的进程号为：", pid)
	} else {

		// 优先用pid,否则通过pName去获取pid
		if pid == 0 {
			// 执行 ps 命令查找进程
			cmd := exec.Command("ps", "aux")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}

			// 遍历输出的进程列表，查找进程名称为 pName 的进程
			for _, line := range strings.Split(string(output), "\n") {
				if strings.Contains(line, pName) && !strings.Contains(line, "grep") {
					// 进程名称包含 pName，且不是 grep 命令的输出
					// 获取进程PID
					fields := strings.Fields(line)
					pid, _ = strconv.Atoi(fields[1])
					fmt.Printf("进程 %s 的PID为 %d\n", pName, pid)
				}
			}
		}

	}

	fmt.Println("要监测的进程号为:", pid)

	go func() {

		p, _ := process.NewProcess(int32(pid))
		for {
			cpuPercent, _ := p.CPUPercent()
			memPercent, _ := p.MemoryPercent()
			fmt.Printf("该进程的cpu占用率:%.3f,内存占用:%.3f, 时间:%s\n", cpuPercent, memPercent, time.Now().Format("2006-01-02 15:04:05"))
			println("---------------分割线------------------")
			time.Sleep(time.Duration(timeInterval) * time.Second)
		}
	}()

	select {}
	//ch := make(chan int, 0)
	//
	//ch <- 1

}

## usage_reporter


上报某个进程占用的cpu和内存比例,例如

```bash
要监测的进程号为: 71956
该进程的cpu占用率:1.382,内存占用:0.043, 时间:2020-05-16 23:28:41
---------------分割线------------------
该进程的cpu占用率:0.762,内存占用:0.045, 时间:2020-05-16 23:28:42
---------------分割线------------------
该进程的cpu占用率:0.533,内存占用:0.045, 时间:2020-05-16 23:28:43
---------------分割线------------------
该进程的cpu占用率:0.414,内存占用:0.045, 时间:2020-05-16 23:28:44
```

<br>

### 安装

<br>

`go install github.com/cuishuang/usage_reporter@latest` 或  `go get -u github.com/cuishuang/usage_reporter`


Mac上可使用brew

<br>

### 参数

<br>


- 不带任何参数,`usage_reporter`将每隔1s输出当前进程占用的cpu和内存

- pid, 指定要监测的进程id。例如`usage_reporter -pid 16789`,监测pid为16789这个进程的cpu和内存占用
- n,  指定要监测的进程名称（如果同时指定pid和n，以pid为准）。例如`usage_reporter -n chrome`，监测chrome这个进程的cpu和内存占用;注意有可能匹配到多个，以ps aux第一个有效进程为准
- t, 时间间隔(默认1s)。例如`usage_reporter -pid 16789 -t 10`,监测pid为16789这个进程的cpu和内存占用，每隔10s打点一次
- p, 是否去除其他信息，仅保留基本信息，默认为false




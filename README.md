### GO 学习笔记

本人快速迁移学习 Go 语言的日程记录；

在学一门语言的过程之中，我们不仅需要学习语言本身如何使用，更要学习配套的工具链；

---



### 环境配置

```shell
# 关于Go环境配置可能需要修改下面两个配置文件（Mac/Linux)
(base) >> go git:(main) ✗ code ~/.bash_profile  # 用户级别的配置文件，适用于当前用户
(base) >> go git:(main) ✗ code /etc/profile		# 系统级别的配置文件，适用于全部用户

# 使用命令查看所有环境
(base) >> go git:(main) ✗ go env

# 针对性的查看某个环境变量，首先使用下面的命令的查看Go下载目录
(base) >> go git:(main) ✗ go env GOPATH	
/Users/Taoist/go

(base) >> go git:(main) ✗ cd /Users/Taoist/go	# 然后切入目录查看
(base) >> go git:(main) ✗ ls					# 查看当前目录包含哪些文件夹
bin pkg															

# 查看发现只有 bin、pkg 两个文件夹，此时不得不说早期的 GOPATH 机制
# 	1.1 早期必须在此目录下面新开一个 src 文件，然后代码必须写在 src 里面；
# 	1.2 早期手动下载的外部模块也需要手动配置放入 src 里面；

# 而在 Go 1.11.x 版本引入了 modules 管理机制之后，
#	2.1 项目依赖模块全部通过 go.mod 下载，然后放在 pkg/mod 里面；
#	2.1 项目目录也不再具有限制，能够写在任意位置；
```

---



### 配置文件

更加具体详细的内容，等有需要的时候，再查阅 [这篇](https://juejin.cn/post/7182091980099289147) 博客即可；

我们在新手阶段的会遇到使用 OJ 去刷语法题的需求，针对每道题，我们可能都将编写一个包含 main 函数的代码问题，又或者我们跑一跑编写教材上面的，可能会引言非标准库的 demo 代码，如果我们将这些的代码全部放在一个文件下面，并且编写了相应的 go.mod/go.work，极有可能出错。

因为具备  go.mod/go.work 文件的当前目录会被当做是一个工程文件 (当成一个 module)，一个工程文件仅允许存在一个 main 入口，因而我们针对每道算法题编写的代码会因为重复声明 main 而有红色波浪线 (VS Code 错误提醒)，为此我们可在 ***当前目录*** 创建 *go.work*，***再在某个子文件夹*** 之中 创建 *go.mod* 文件，尽管这种做法非常不符合 Go 语言哲学，但在初学阶段，这种图个方便的做法也是可以接受的。

| 配置文件 |                             作用                             |
| :------: | :----------------------------------------------------------: |
|  go.mod  | 老版本的 Go 项目通过 GOPATH 模式管理，项目代码要写在 Go 安装目录下面的 src， Go 通过 go.mod 管理 modules，其作用有点类似于Java Maven 里面的 pom.xml，或说 Python pip/conda 里面的 requirement.txt |
|  go.sum  | 创建 go.mod之后自动生成的文件。对于使用者来说，这个文件的内容，基本就是天书，文件逐行列举了需要导入的模块、模块版本号码，以及模块相应版本的哈希值 |
| go.work  | 如果一个文件目录包含 go.mod， 则其 可以通过 use 关键字加入 go.work 文件，使其变成module，一个 module 只允许一个 main 函数作为入口，有时候我们并不编写大项目，而是仅仅针对算法问题编写若干份短小的代码，这些代码文件都有 main 函数，此时可以规定当前 workshop 下面的一个子目录，将其变为 module，比如 ./dependencies，除此文件夹内部以外的其他 go 代码均可正常使用 run 或 build 运行、编译，而且亦可享用导入的非标准模块 |

```go
/*
 * 下面给出的go.work本项目的文件
 * 只有包含 go.mod 文件的目录才能使用 use 命令添加进来
*/
go 1.21.0

use (
	./dependencies
)
```

---







### 高级主题

Go号称现代版 C/C++，使用下来感觉就是， 程序的语法非常接近 C/C++，但是又可以像是 Java、Python 那样方便的引入外部模块，同时Go相较于 C/C++ 又多出了很多避免程序员炫技的限制。整体看下来，感觉是一门不错的语言。能做数值计算、数据分析、深度学习、网络爬虫、网络服务器等等。实现这些高级主题，需要深入学习 Go [函数](https://go.dev/tour/methods/1)、[泛型](https://go.dev/tour/generics/1)、[并发](https://go.dev/tour/concurrency/1) 这些特性

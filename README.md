# neptune 一个 redis 连接池 GRPC 服务端

#### 介绍
不管项目是大是小，名字一定要响；对这个项目就是 “海王星”；
本项目可以提供大家学习 GRPC 服务的编写与熟悉，了解连接池的实现；
至于本项目能解决多少问题，其实并不能解决太多问题；如果大家对改项目感兴趣，欢迎大胆尝试；


#### 客户端说明
这里有个 php-neptune客户端，可以在php下使用该连接池服务；
大家都了解，在php下调用redis使用短连接会反复的创建连接/关闭连接，并发请求过大时会产生大量 TIME_WAIT 文件句柄数一旦达到上限，新的请求就无法被处理了；
使用长连接时，redis 长连接托管给 php-fpm ；这里只要php-fpm不死，连接就不会消失，这里缺少了动态管理的支持；

连接池则可以解决这个问题，但是，但是... GRPC也是TCP连接，悲伤了；所以该项目的存在似乎只能是学习了；
如果想解决php短连接的问题，似乎有两种方案 1.换其他语言；2.利用 swoole 搭建http服务使php项目已一个常驻进程运行；


#### 安装教程
依赖安装,项目中用到了一些包，直接安装即可；

被墙了百度下好多方法安装；
```
go get google.golang.org/grpc
go get golang.org/x/net
go get github.com/zhuCheer/pool
go get github.com/zhuCheer/neptune

```


#### 使用说明
直接到项目目录下
```
 go run main.go --host=127.0.0.1:50033
```

#### 说明
如果你不想过多了解源码细节，可以直接到 Release 中下载编译好的程序，windows/linux
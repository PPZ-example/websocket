# websocket example
golang <--> js in browser

## TODO
+ 读写不能并行（close 被锁住） -> 读锁、写锁 分开
+ 测试 websocket 读写异常类型（有些错误发生之后，ws 还能用，只是这次异常了；有些则 ws 不能用了，要重连）

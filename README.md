# 基于go的开机启动服务管理

# 打包编译成exe
> go build -o auto.exe

# 使用方式
1. 配置启动命令 
> 按照service.json的配置填写要开机启动的命令, 配置文件必须和exe执行文件同级目录下
2. 注册服务
> auto.exe install
> 
> 测试windows服务中出现我们的应用，将服务设置为自动后，每次开机都会把配置文件中的命令一个个执行
3. 移除服务
> auto.exe remove
> 
> 不想自启动可以卸载服务,或者修改配置

4. 配置详解
```
"name": 命令名称,
"start": 执行命令,
"stop": 停止命令，暂时用不上,
"status": 是否执行,
"description": 描述
```
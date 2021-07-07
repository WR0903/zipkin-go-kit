# zipkin 链路追踪在go-kit微服务中应用
在分布式系统中，由于异步调用的存在，通过日志埋点的方式有时候解决线上问题不是很高效，zipkin可以追踪每次调用的过程。  
本项目是将zipkin应用在go-kit微服务。

## 使用方法  
```docker run -d -p 9411:9411 openzipkin/zipkin```运行zipkin；  
```consul agent -dev```启动consul   

进入gateway目录，运行```go run main.go```
打开浏览器```http://127.0.0.1:9090/string-service/health``` 虽然浏览器没有反应，      
但是浏览器打开```http://127.0.0.1:9411/zipkin```就可以看到这个调用过程，效果如下图：
![image](https://github.com/WR0903/zipkin-go-kit/blob/master/gatewayzipkin.png)   

进入server目录，运行```go run *.go```,然后在运行```go run *.go -service.port=9010 -grpc=:9007```,这样就启动了两个实例化的服务，打开浏览器```http://127.0.0.1:8500/ui```,可以看到如下图：  
![image](https://github.com/WR0903/zipkin-go-kit/blob/master/consulzipkin.png)   

此时，用浏览器重新访问```http://127.0.0.1:9090/string-service/health```,就能正常返回了。

另外可以用postman发送请求：   
![image](https://github.com/WR0903/zipkin-go-kit/blob/master/concat.png)  
此时可以在zipkin上看到调用链路：   
![image](https://github.com/WR0903/zipkin-go-kit/blob/master/gateway.png) 

在这个项目中，server是业务项目的微服务，gateway是自己写的一个简单的网关。客户端请求都通过网关，然后才到实例化的server

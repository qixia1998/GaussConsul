# GaussConsul
> Go Zero使用Consul做为微服务注册发现

* `git clone https://github.com/qixia1998/GaussConsul.git`
* `cd GaussConsul`
* `go mod tidy`
*  修改etc目录下配置文件

### gauss 
> 1. `go run rpc/gauss.go -f rpc/etc/gauss.yaml`
> 2. `go run api/gauss.go -f api/etc/gauss.yaml`
### user
> 1. `go run rpc/user.go -f rpc/etc/user.yaml`
> 2. `go run api/user.go -f api/etc/user.yaml`

### Consul
> 1. `docker run -d --name consul -p 8500:8500 -p 8600:8600/udp -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}' consul agent -server -bootstrap-expect 1 -ui -bind=0.0.0.0`
> 2. `consul agent -dev`
> 3.  查看 localhost:8500
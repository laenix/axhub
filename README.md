# Axhub 更新提醒到飞书
## 起因  
为了应对产品经理修改产品图时信息不同步的情况，开发了此Axhub转飞书提醒的小组件。
## 用法
### 源文件运行
修改位于 ./config/config.yaml.sample 的配置文件，将之重命名为config.yaml
go run ./cmd/axhub/main.go
### docker
docker build . -t axhub-lark:latest
docker run --name axhub-lark -v /主机目录/:/opt/config/ axhub-lark:latest

# Axhub notice to Lark
## Cause
In order to deal with the situation that the information is not synchronized when the product manager modifies the product diagram,the Axhub notice to lark reminder widget is developed
## Usage
### source file run
Change the sample configuration file located in ./config/config.yaml.sample,rename it to config.yaml
### docker
docker build . -t axhub-lark:latest
docker run --name axhub-lark -v /host/:/opt/config/ axhub-lark:latest
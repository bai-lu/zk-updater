# zk-updater


## 项目特点

1. 使用golnag 开发, 利用golang原生的channel 机制实现消息通知, 更新事件生产消费
2. go mod 管理依赖, 源码依赖的都可以通过go get远程包, 运行环境无需依赖
3. 实现飞书通知

## 编译

执行build.sh cross 交叉编译linux环境二进制程序, 生成文件在 dist目录

## 部署

支持容器化部署

## 组件

1. cmd 程序主入口
2. internal
   1. config: 此包import后自动执行init 加载配置文件获取节点信息
   2. node: 核心程序, 负责巡检和监听及执行更新事件
   3. notify: 开启消息队列, 有消息过来发送给飞书api
   4. utils/kcagent: kms加密agent的调用封装

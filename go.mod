module zk-updater

require (
	git.n.github.com/golang/thriftlib v0.0.0
	github.com/apache/thrift v0.0.0-20151001171628-53dd39833a08
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/sirupsen/logrus v1.8.1
)

replace git.n.github.com/golang/thriftlib v0.0.0 => ./thriftlib

go 1.16

module github.com/qhzhyt/tiops-go-sdk

go 1.16

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-git/go-billy/v5 v5.2.0
	github.com/go-git/go-git/v5 v5.3.0
	github.com/golang/protobuf v1.5.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852
	github.com/robfig/cron/v3 v3.0.1
	go.uber.org/atomic v1.4.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.28.0
	tiops v0.0.0
)

replace tiops v0.0.0 => ./

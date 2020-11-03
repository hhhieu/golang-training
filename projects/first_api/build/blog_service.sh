confd -confdir configs/confd -config-file blog_service.toml -onetime -backend env
go run ./cmd/blog_service

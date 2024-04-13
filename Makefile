lint:
	golangci-lint --exclude-use-default=false --out-format tab run ./...

dockerbuild:
	docker build -t 96solutions/neurography:latest .

gogen:
	go generate ./...

test:
	make gogen && go test ./...

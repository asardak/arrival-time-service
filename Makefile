BIN=$(PWD)/bin

bin/swagger:
	GO111MODULE=off go get -d github.com/go-swagger/go-swagger
	cd $(GOPATH) && GOBIN=$(BIN) go install github.com/go-swagger/go-swagger/cmd/swagger

bin/mockery:
	GO111MODULE=off GOBIN=$(BIN) go get github.com/vektra/mockery/.../

bin/golangci-lint:
	GO111MODULE=off GOBIN=$(BIN) go get github.com/golangci/golangci-lint/cmd/golangci-lint

generate: bin/swagger bin/mockery
	bin/swagger generate client -t ./pkg/car-service -f ./pkg/car-service/swagger.yml
	bin/swagger generate client -t ./pkg/predict-service -f ./pkg/predict-service/swagger.yml
	bin/mockery -dir ./internal -all -output ./internal/test/mocks

run:
	go run ./cmd/arrival-time-service

test: bin/golangci-lint
	bin/golangci-lint run
	go test --tags=db ./...

.PHONY: generate run test
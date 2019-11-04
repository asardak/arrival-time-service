TMP_DIR:=/tmp/swagger-go-$(shell head -n 1 /dev/urandom | md5)
BIN=$(PWD)/bin

bin/swagger:
	mkdir $(TMP_DIR)
	git clone https://github.com/go-swagger/go-swagger $(TMP_DIR)
	cd $(TMP_DIR) && GOBIN=$(BIN) go install ./cmd/swagger

generate: bin/swagger
	$(BIN)/swagger generate client -t ./pkg/car-service -f ./pkg/car-service/swagger.yml
	$(BIN)/swagger generate client -t ./pkg/predict-service -f ./pkg/predict-service/swagger.yml

run:
	go run ./cmd/arrival-time-service

.PHONY: generate run
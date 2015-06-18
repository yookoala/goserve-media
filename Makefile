build: test
	go build

test:
	cd server && TEST_DIR="../_example" go test

.PHONY:
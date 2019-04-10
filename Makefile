BUILD_PATH=../../bin
BUILD_BIN=nestor

clean:
	@rm -f $(BUILD_PATH)/$(BUILD_BIN)

build: clean
	@go build -o $(BUILD_PATH)/$(BUILD_BIN)

start: build
	@$(BUILD_PATH)/$(BUILD_BIN) $(ARGS)


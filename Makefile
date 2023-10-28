.PHONY: build-linux build-darwin build-windows build-release-artifacts

build-release-artifacts:
	make build-linux && \
	make build-darwin && \
	make build-windows && \
	zip schedule-notifier.linux_amd64.zip schedule-notifier_linux_amd64 && \
	zip schedule-notifier.darwin_arm64.zip schedule-notifier_darwin_arm64 && \
	zip schedule-notifier.windows_amd64.zip schedule-notifier.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./schedule-notifier_linux_amd64 -v

build-darwin:
	GOOS=darwin GOARCH=arm64 go build -o ./schedule-notifier_darwin_arm64

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./schedule-notifier.exe -v


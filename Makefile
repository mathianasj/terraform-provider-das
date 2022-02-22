install-mac-deps:
	brew tap go-swagger/go-swagger && brew install go-swagger

build:
	go build -o terraform-provider-das main.go
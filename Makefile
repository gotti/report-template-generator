statik:
	statik -f -src ./template

build: statik
	go build -o report-gen cmd/main.go

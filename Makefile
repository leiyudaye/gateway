ROOT=`pwd`
PBPATH=./protobuf
all:
	go build -o /bin/main.bin main.go
pb:
	echo $(ROOT)
	echo $(PBPATH)
	protoc  -I $(PBPATH) --go_out=$(PBPATH)  $(PBPATH)/*.proto

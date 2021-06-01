ROOT=`pwd`
PBPATH=./protobuf
PBDST=./protobuf_go
all:
	go build -o /bin/main.bin main.go
pb:
	echo $(ROOT)
	echo $(PBPATH)
	protoc  -I $(PBPATH) --go_out=plugins=grpc:$(PBDST)  $(PBPATH)/*.proto

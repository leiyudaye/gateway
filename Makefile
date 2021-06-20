PBPATH=./protobuf
PBDST=./protobuf_go
PBFILE=$(shell ls ./protobuf)
all:
	go build -o /bin/main.bin main.go
pb: 
	# @for NAME in $(PBFILE);do\
	# 	echo $(NAME);\
	# done  
	protoc  -I $(PBPATH) --go_out=plugins=grpc:$(PBDST)  $(PBPATH)/user.proto
	protoc  -I $(PBPATH) --go_out=plugins=grpc:$(PBDST)  $(PBPATH)/flowcount.proto

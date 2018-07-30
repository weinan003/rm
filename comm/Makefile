all : comm.proto
	protoc -I ./ --go_out=plugins=grpc:./ ./comm.proto

clean:
	rm ./comm.pb.go

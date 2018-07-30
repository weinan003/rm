all: main.go
	make -C comm
	go build -o rmserver $^

clean:
	rm ./comm/comm.pb.go rmserver

testv1:
	go test v1/server_test.go v1/server.go -v -race
.PHONY: testv1

runv1:
	go run ./v1/main.go ./v1/server.go -v
.PHONY: runv1


testv2:
	go test v2/server_test.go v2/server.go -v -race
.PHONY: testv2

runv2:
	go run ./v2/main.go ./v2/server.go -v
.PHONY: runv2

curlv2:
	curl http://localhost:5000/player/Floyd -v
.PHONY: curlv2



testv3:
	go test v3/server_test.go v3/server.go -v -race
.PHONY: testv3

runv3:
	go run ./v3/main.go ./v3/server.go -v
.PHONY: runv3

curlv3:
	curl http://localhost:5000/player/Apollo -v
.PHONY: curlv3
 

testv4:
	go test v4/server_test.go v4/server.go -v -race
.PHONY: testv4

runv4:
	go run ./v4/main.go ./v4/server.go -v
.PHONY: runv4

curlv4:
	curl http://localhost:5000/player/Apollo -v
.PHONY: curlv4


testv5:
	go test v5/server_test.go v5/server.go -v -race
.PHONY: testv5

runv5:
	go run ./v5/main.go ./v5/server.go -v
.PHONY: runv5

curlv5:
	curl http://localhost:5000/player/Apollo -v
.PHONY: curlv5


testv6:
	go test v6/server_test.go v6/server.go -v -race
.PHONY: testv6

runv6:
	go run ./v6/main.go ./v6/server.go -v
.PHONY: runv6

curlv6:
	curl http://localhost:5000/player/Apollo -v
.PHONY: curlv6



testv7:
	go test v7/server_test.go v7/server.go -v -race
.PHONY: testv7

runv7:
	go run ./v7/main.go ./v7/server.go -v
.PHONY: runv7

curlv7:
	curl http://localhost:5000/player/Apollo -v
.PHONY: curlv7



testv8:
	go test v8/server_test.go v8/server_integration_test.go v8/server.go v8/main.go  -v -race
.PHONY: testv7

runv8:
	go run ./v8/main.go ./v8/server.go -v
.PHONY: runv7

postv8:
	curl -X POST http://localhost:5000/player/Pepper -v
.PHONY: postv8

getv8:
	curl -X GET http://localhost:5000/player/Pepper -v
.PHONY: getv8

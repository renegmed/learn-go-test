
testv0:	
	go test v0/*.go  -v -race
.PHONY: testv0

runv0:
	go run ./v0/main.go ./v0/server.go ./v0/InMemoryPlayerStore.go -v
.PHONY: runv0

postv0:
	curl -X POST http://localhost:5000/player/Pepper -v
.PHONY: postv0

getv0:
	curl -X GET http://localhost:5000/player/Pepper -v
.PHONY: getv0


testv1:	
	go test v1/*.go  -v -race
.PHONY: testv1

runv1:
	go run ./v1/main.go ./v1/server.go ./v1/InMemoryPlayerStore.go -v
.PHONY: runv1

postv1:
	curl -X POST http://localhost:5000/league-v
.PHONY: postv1

getv1:
	curl -X GET http://localhost:5000/league -v
.PHONY: getv1


testv2:	
	go test v2/*.go  -v -race
.PHONY: testv2
# runv2:
# 	go run ./v2/main.go ./v2/server.go ./v2/InMemoryPlayerStore.go -v
# .PHONY: runv2

# postv2:
# 	curl -X POST http://localhost:5000/league-v
# .PHONY: postv2

# getv2:
# 	curl -X GET http://localhost:5000/league -v
# .PHONY: getv2


testv3:	
	go test v3/*.go  -v -race
.PHONY: testv3


testv4:	
	go test v4/*.go  -v -race
.PHONY: testv4

testv5:	
	go test v5/*.go  -v -race
.PHONY: testv5

testv6:	
	go test v6/*.go  -v -race
.PHONY: testv6



testv7:	
	go test v7/*.go  -v -race
.PHONY: testv7
runv7:
	go run ./v7/main.go ./v7/server.go ./v7/InMemoryPlayerStore.go -v
.PHONY: runv7

postv7:
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/Peter -v
.PHONY: postv7

getv7:
	curl -X GET http://localhost:5000/league -v
.PHONY: getv7

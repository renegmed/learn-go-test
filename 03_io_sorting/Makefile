
testv0:	
	go test v0/*.go  -v -race
.PHONY: testv0

runv0:
	go run ./v0/main.go ./v0/server.go ./v0/InMemoryPlayerStore.go -v
.PHONY: runv0

postv0:
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/Peter -v
.PHONY: postv0

getv0:
	curl -X GET http://localhost:5000/league -v
.PHONY: getv0



testv1:	
	go test v1/*.go  -v -race
.PHONY: testv1


testv2:	
	go test v2/*.go  -v -race
.PHONY: testv2


testv3:	
	go test v3/*.go  -v -race
.PHONY: testv3

testv4:	
	go test v4/*.go  -v -race
.PHONY: testv4

testv5:	
	# created and used new type League for []Player thus able to crate new method for League
	go test v5/*.go  -v -race
.PHONY: testv5




testv6:	
	# store data, wins for new players
	# remove InMemoryPlayerStore as replaced by database
	# improve dbase performance - to read the file when the program starts up and only need to
	#   update the file when dat changes
	go test v6/*.go  -v -race
.PHONY: testv6

runv6:
	go run ./v6/main.go ./v6/server.go ./v6/FileSystemStore.go ./v6/league.go -v
.PHONY: runv6

postv6:
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/Peter -v
.PHONY: postv6

getv6:
	curl -X GET http://localhost:5000/league -v
.PHONY: getv6

postnewv6:
	curl -X POST http://localhost:5000/players/Pepper -v	
.PHONY: postnewv6

dbfilev6:
	cat db.game.json
.PHONY: dbfilev6


testv7:	
	# able to handle when a new data was smaller than what was there before  
	# like loading the data into memory, the file's record was reduced.
	# Need to sync the memory and the physical file contents
	# Solution is to separate out the concern of the kind of data we write from writing.
	go test v7/*.go  -v -race
.PHONY: testv7



testv8:	
	# error handling
	# works with empty file
	go test v8/*.go  -v -race
.PHONY: testv8


testv9:	
	# sorting - players sorted by their scores, from highest to lowest
	go test v9/*.go  -v -race
.PHONY: testv9

testv1:	
	go test v1/*.go -v -race
	# go test ./... -v -race
.PHONY: testv1

run-cliv1: 
	go run v1/cmd/cli/main.go 
.PHONY: run-cliv1


run-serverv1:
	#cd v1/cmd/webserver
	go run v1/cmd/webserver/main.go 
.PHONY: run-serverv1


postv1:
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/John -v
	curl -X POST http://localhost:5000/players/Peter -v
.PHONY: postv1

getv1:
	curl -X GET http://localhost:5000/league -v
.PHONY: getv1

postnewv1:
	curl -X POST http://localhost:5000/players/Pepper -v	
.PHONY: postnewv1

dbfilev1:
	cat v1/cmd/webserver/game.db.json
.PHONY: dbfilev1

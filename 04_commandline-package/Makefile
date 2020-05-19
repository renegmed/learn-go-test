
# v0 is a copy from previous section's latest version v9
testv0:	
	# sorting - players sorted by their scores, from highest to lowest
	go test v0/*.go  -v -race
.PHONY: testv0

testv1:	
	go test v1/*.go  -v -race
.PHONY: testv1


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

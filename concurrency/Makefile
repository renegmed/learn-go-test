testv1:
	go test v1/CheckWebsites_test.go v1/CheckWebsites.go -v 
.PHONY: testv1

bench-testv1:
	go test v1/CheckWebsites_benchmark_test.go v1/CheckWebsites.go -bench=.
.PHONY: bench-testv1


testv2:
	go test v2/CheckWebsites_test.go v2/CheckWebsites.go -v -race
.PHONY: testv2

bench-testv2:
	go test v2/CheckWebsites_benchmark_test.go v2/CheckWebsites.go -bench=.
.PHONY: bench-testv2
 

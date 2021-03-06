
init-project:
	go mod init github.com/renegmed/learn-go-test
.PHONY: init-project

test:	
	$(MAKE) -C 01_http_server testv1
	$(MAKE) -C 01_http_server testv2
	$(MAKE) -C 01_http_server testv3
	$(MAKE) -C 01_http_server testv4
	$(MAKE) -C 01_http_server testv5
	$(MAKE) -C 01_http_server testv6
	$(MAKE) -C 01_http_server testv7
	$(MAKE) -C 01_http_server testv8 
	$(MAKE) -C 02_json_routing_embedding testv1
	$(MAKE) -C 02_json_routing_embedding testv2
	$(MAKE) -C 02_json_routing_embedding testv3
	$(MAKE) -C 02_json_routing_embedding testv4
	$(MAKE) -C 02_json_routing_embedding testv5
	$(MAKE) -C 02_json_routing_embedding testv6
	$(MAKE) -C 02_json_routing_embedding testv7
	$(MAKE) -C 03_io_sorting testv1
	$(MAKE) -C 03_io_sorting testv2
	$(MAKE) -C 03_io_sorting testv3
	$(MAKE) -C 03_io_sorting testv4
	$(MAKE) -C 03_io_sorting testv5
	$(MAKE) -C 03_io_sorting testv6
	$(MAKE) -C 03_io_sorting testv7
	$(MAKE) -C 03_io_sorting testv8
	$(MAKE) -C 03_io_sorting testv9
	# $(MAKE) -C 04_commandline-package testv1
	$(MAKE) -C 05_time testv1
	$(MAKE) -C arrays-slices testv1
	$(MAKE) -C arrays-slices testv2
	$(MAKE) -C arrays-slices testv3
	$(MAKE) -C arrays-slices testv4
	$(MAKE) -C concurrency testv1
	$(MAKE) -C concurrency testv2
	$(MAKE) -C context testv1
	$(MAKE) -C context testv2
	$(MAKE) -C dependency-injection testv1
	$(MAKE) -C iteration testv1
	$(MAKE) -C mocking testv1
	$(MAKE) -C pointers-error testv1
	$(MAKE) -C reflection testv1
	$(MAKE) -C reflection testv2
	$(MAKE) -C reflection testv3
	$(MAKE) -C reflection testv4
	$(MAKE) -C reflection testv5
	$(MAKE) -C reflection testv6
	$(MAKE) -C reflection testv7
	$(MAKE) -C reflection testv8
	$(MAKE) -C reflection testv9
	$(MAKE) -C reflection testv91
	$(MAKE) -C reflection testv10
	$(MAKE) -C reflection testv11
	$(MAKE) -C select testv1
	$(MAKE) -C select testv1
	$(MAKE) -C structs-methods-interfaces testv1
	$(MAKE) -C structs-methods-interfaces test-rectangle
	$(MAKE) -C sync testv1
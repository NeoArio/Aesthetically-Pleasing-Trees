PROJECT_NAME :=  Aesthetic_Garden

.PHONY: build run

build:
	## Build
	@echo "  >  Building binary..."
	- go build
	@echo "  >  Build done ..."

run:
	- ./main $(input)

clean:
	- rm -f main

test:
	- go test *.go -v
	- @echo "Tests Run Finished"

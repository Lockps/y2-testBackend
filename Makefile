build:
	@go build -o ./bin/acid-reflex ./Backend/

run: build
	@./bin/acid-reflex
# clear:
# 	@clear

# build:
# 	@go build -o totion .

# run: clear build
# 	@./totion

clear:
	@cls || clear

build:
	@go build -o bin/totion ./cmd/totion

run: clear build
	@./bin/totion

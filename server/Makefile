gql_gen:
	gqlgen generate

run:
	figlet Shadow Chat
	go install github.com/cespare/reflex@latest
	reflex -r '\.go$$' -s go run main.go

lint:
	golangci-lint run

fix:
	golangci-lint run --fix
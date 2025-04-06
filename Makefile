.PHONY: dev build clean generate run

# Development with live reload
dev:
	air

# Generate templ files
generate:
	templ generate

# Build the application
build: generate
	go build -o ./bin/app .

# Run without live reload
run: generate
	go run .


# Clean temporary files
clean:
	rm -rf tmp/
	rm -rf bin/
# Build the project (release version)
go build -ldflags=-w

# Run the program on different test files
./strings -t=data/dna-text.txt -p=data/dna-pattern.txt
./strings -t=data/war-peace-text.txt -p=data/war-peace-pattern.txt
./strings -t=data/lorem-ipsum-text.txt -p=data/lorem-ipsum-pattern.txt

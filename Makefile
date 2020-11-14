

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = eko
PKG := example.com/EKO


all: $(EXEC)

$(EXEC):
	mkdir -p ./bin
	GO111MODULE=on go build -o ./bin/$(EXEC) ./cmd/main.go

run: 
	./bin/$(EXEC)

clean: 
	rm -f ./bin/*
	rm -r ./bin
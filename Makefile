$SRC = entity/entity.go operator/get.go operator/post.go
$SRC += parser/parser.go

build: 
	go build -o go-http main.go $(SRC)

clean:
	rm go-http

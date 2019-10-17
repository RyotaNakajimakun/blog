#Basic makefile

default: build

build: vet
	go build

doc:
	@godoc -http=:6060 -index

lint:
	@golint ./...

debug:
	@reflex -c reflex.conf

run: build
	./blog && cp ./blog ./tmp/blog

test:
	@go test ./...

vet:
	@go vet ./...

clean:
	rm ./blog

migration:
	touch ./system/migration/`date '+%Y%m%d%H%I%s'`migration.go

psqlstart:
	docker run --name my-db -p 5432:5432 -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=secret -d postgres:11.4

fresh: vet
	go build && rm tmp/blog && mv blog tmp/blog && fresh


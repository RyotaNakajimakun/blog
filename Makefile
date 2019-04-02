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
	postgres -D /usr/local/var/postgres &

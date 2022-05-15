run:
	rm -vfr generated/*
	rm -vfr models/*
	#unset GOMOD
	go get github.com/99designs/gqlgen@v0.17.5
	go get github.com/99designs/gqlgen/internal/imports@v0.17.5
	go get github.com/99designs/gqlgen/codegen/config@v0.17.5
	go get github.com/99designs/gqlgen/internal/imports@v0.17.5
	go run github.com/99designs/gqlgen --verbose
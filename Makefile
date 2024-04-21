gen:
	@echo "Getting gql gen..."
	go get github.com/99designs/gqlgen
	@echo "Generating..."
	go run github.com/99designs/gqlgen generate
	@echo "Running server..."
	go run server.go 

run:
	@echo "Running server..."
	gow run server.go 

dockerip:
	docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql
 	

dockerbash:
	docker exec -it mysql bash

dockerlogs:
	docker logs mysql

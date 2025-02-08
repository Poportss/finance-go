# Build the container image using Podman
build-image:
	podman build -t poports/finance .

# Run the application with Podman Compose
run-app:
	podman compose --file ./devops/compose.yaml up --detach

lint:
	go fmt -n ./...

unit_test:
	go test	./...
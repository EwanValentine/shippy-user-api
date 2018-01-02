build:
	docker build -t shippy-user-api .

run:
	docker run --net="host" \
		-p 8080 \
		-e SERVICE_ADDRESS=:50051 \
		-e API_ADDRESS=:8080 \
		shippy-user-api

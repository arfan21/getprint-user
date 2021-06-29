build-dev:
	docker build -f dev.Dockerfile -t getprint-service-user-dev .

build-prod:
	docker build -f prod.Dockerfile -t getprint-service-user-prod .


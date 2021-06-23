build-dev:
	docker build -f dev.Dockerfile -t getprint-service-user-dev .

build-prod:
	docker build -f prod.Dockerfile -t getprint-service-user-prod .

test-repo:
	go test -timeout 30s -run ^TestMySQLUserTest$$ github.com/arfan21/getprint-user/app/repository/mysql -v  

test-service:
	clear
	go test -timeout 30s -run ^TestUserServices$$ github.com/arfan21/getprint-user/app/services -v
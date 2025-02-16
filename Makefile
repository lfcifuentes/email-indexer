
dowload_email_data:
	chmod +x scripts/etl.sh
	./scripts/etl.sh
up:
	docker-compose up -d --build

stop:
	docker-compose stop

status: 
	docker-compose ps

logs:
	docker-compose logs

clean:
	docker-compose down -v --remove-orphans

build: ## Build docker image
	docker-compose build

docs_generate:
	@echo "Generando documentacion"
	swag init -g main.go -d backend -o backend/docs
	@echo "Documentacion generada"

docs_format:
	cd backend
	@echo "Formateando documentacion"
	swag fmt
	@echo "Documentacion formateada"

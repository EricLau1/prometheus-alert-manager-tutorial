up:
	docker-compose up

down:
	docker-compose down

clear: down
	docker volume rm docker_prometheus_data_tutorial && \
    docker volume rm docker_alertmanager_data_tutorial && \
    rm -rf data
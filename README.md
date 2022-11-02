# go-rest

### Command to Clear All Database Data and Docker Service

    rm -rf external/database/data && \
    rm -rf external/database/db_backup && \
    docker rm db && \
    docker rm api

### Prune Docker System

    docker system prune

### Start Docker Service

    docker-compose up --build

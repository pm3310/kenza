package cli

const envCloud = `
# RabbitMQ
RABBITMQ_DATA_VOLUME=/home/ubuntu/kenza/data/rabbitmq/storage

# POSTGRES
POSTGRES_DB=kenza
POSTGRES_HOST=db
POSTGRES_USER=kenza
POSTGRES_PASSWORD=kenza
POSTGRES_DATA_VOLUME=/home/ubuntu/kenza/data/postgres/storage

# DB MIGRATIONS
KENZA_DB_MIGRATIONS_HOST_PATH=./db/migrations
KENZA_DB_MIGRATIONS_CONTAINER_PATH=/kenza/data/db/migrations

# AWS HOST CONFIG LOCATIONS
KENZA_AWS_CONFIG_HOST_PATH=~/.aws/config
KENZA_AWS_CREDENTIALS_HOST_PATH=~/.aws/credentials

# API
KENZA_API_PORT=:8080
KENZA_API_HOST=localhost
KENZA_API_AWS_CONFIG_CONTAINER_PATH=/home/api/.aws/config
KENZA_API_AWS_CREDENTIALS_CONTAINER_PATH=/home/api/.aws/credentials
KENZA_API_AWS_PROFILE=default

# WORKER
KENZA_WORKER_AWS_CONFIG_CONTAINER_PATH=/root/.aws/config
KENZA_WORKER_AWS_CREDENTIALS_CONTAINER_PATH=/root/.aws/credentials
KENZA_WORKER_AWS_PROFILE=default

# JOB LOGS
KENZA_JOB_LOGS_HOST_PATH=/home/ubuntu/kenza/data/jobs/logs
KENZA_JOB_LOGS_CONTAINER_PATH=/kenza/data/jobs/logs

# KENZA SERVICE CONTAINERS
KENZA_VERSION=latest # overriden by the "kenza start" command to use the binary's semver version
KENZA_CONTAINER_REGISTRY=aikenza
`

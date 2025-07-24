# Go-Fiber Trendyol Intern Project

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Fiber](https://img.shields.io/badge/Fiber-00C7B7?logo=fiber&logoColor=white&style=for-the-badge)
![GORM](https://img.shields.io/badge/GORM-FFCA28?logo=go&logoColor=white&style=for-the-badge)
![Kafka](https://img.shields.io/badge/Kafka-231F20?logo=apachekafka&logoColor=white&style=for-the-badge)
![Sarama](https://img.shields.io/badge/Sarama-5A29E4?style=for-the-badge)
![Elasticsearch](https://img.shields.io/badge/Elasticsearch-005571?logo=elasticsearch&logoColor=white&style=for-the-badge)
![Debezium](https://img.shields.io/badge/Debezium-EA601B?logo=debezium&logoColor=white&style=for-the-badge)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge)
![Zookeeper](https://img.shields.io/badge/Zookeeper-FF6800?logo=apachezookeeper&logoColor=white&style=for-the-badge)




---

## Getting Started

### Prerequisites

- Docker and Docker Compose
- PGAdmin 4 or another PostgreSQL client

### Quick Start

1. Clone this repository.
2. Make sure you have a `.env` file in the root directory with the following contents:

    ```env
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=123123123
    DB_NAME=products
    DB_PORT=5432
    DB_SSLMODE=disable
    DB_TIMEZONE=Asia/Shanghai

    ```

3. Navigate to the `docker/services` directory:

    ```sh
    cd docker/services
    ```

4. Start all Docker services:

    ```sh
    docker compose up
    ```

5. Initialize the Kafka connector for Debezium by making a POST request to the Debezium service (from inside the `docker/services` directory):

    ```sh
    curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" 127.0.0.1:8083/connectors/ --data "@debezium.json"
    ```

6. To monitor Kafka topics with kafkacat, run:

    ```sh
    docker run --tty --network postgres_debezium_cdc_default confluentinc/cp-kafkacat kafkacat -b kafka:9093 -C -t postgres.public.products
    ```

7. Once everything is running, you can update the data in the PostgreSQL database using an external tool like PG Admin 4. After updating the database, you should receive an update in Kafka, which will then be written to the Elasticsearch database.

---

## Notes

- All Docker and connector commands must be run inside the `docker/services` folder.
- Ensure your `.env` file is properly configured before starting.
- This project is for learning purposes only and not intended for production use.

---

## License

This project does not carry a formal license yet.

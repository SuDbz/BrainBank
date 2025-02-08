# Run Kafka 
```yaml

version: '3'

services:
  zookeeper:
    image: bitnami/zookeeper:latest
    container_name: zookeeper
    environment:
      - ALLOW_ANONYMOUS_LOGIN=yes
    ports:
      - "2181:2181"
    networks:
      - kafka-net

  kafka:
    image: bitnami/kafka:latest
    container_name: kafka
    environment:
      - KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
      - KAFKA_ADVERTISED_LISTENER_URI=kafka:9093
      - KAFKA_LISTENER_SECURITY_PROTOCOL=PLAINTEXT
      - KAFKA_LISTENER_PORT=9093
    ports:
      - "9093:9093"
    networks:
      - kafka-net
    depends_on:
      - zookeeper

  kafka-manager:
    image: wurstmeister/kafka-manager
    container_name: kafka-manager
    environment:
      - ZK_HOSTS=zookeeper:2181
    ports:
      - "9000:9000"
    networks:
      - kafka-net
    depends_on:
      - kafka

networks:
  kafka-net:
    driver: bridge

```

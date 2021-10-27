FROM scratch

COPY RabbitMQSample.viper.yaml /etc/RabbitMQSample/RabbitMQSample.viper.yaml
COPY rabbitmq-sample /usr/local/bin/rabbitmq-sample

CMD ["rabbitmq-sample", "run"]

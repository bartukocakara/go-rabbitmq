## Setup (Windows)

- Install RabbitMQ driver to your local machine with https://www.rabbitmq.com/install-windows.html
- Then go to http://localhost:15672 link u:guest p:guest
- You can see your queue list on Queues tab http://localhost:15672/#/queues


## Commands

- First you need to open consumer for waiting queue group with this command
```
go run consumer.go
```
- Second you can publish queues to consume.So that process will be finished successfully

<img src="/rabbitmq-ss.png" alt="Alt text" title="Rabbitmq ss" width="800">



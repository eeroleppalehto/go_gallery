# Instruction

Command to run the docker container:

```bash
docker run --name gollery-mysql -v ./init.sql:/docker-entrypoint-initdb.d/init.sql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=******** -d mysql
```

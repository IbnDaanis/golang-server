.PHONY: run_mysql

run_mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:8.0.31

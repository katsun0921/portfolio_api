#!/bin/sh

echo "MySQL, Migrate"

migrate -source file://migrations/ -database 'mysql://root:@tcp(127.0.0.1:3306)/blog_db' down 1

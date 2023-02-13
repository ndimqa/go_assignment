# go_assignment


# notes
1) purchasing
2) publishing
3) welcome page
4) commenting
5) Roles: admin, client, seller 

Также в описании проекта есть "Code should be clean, maintainable and follow SOLID principles. Also, containers have to be used (orchestrate by the use of docker compose)"

В плане структуризации файловой системы полагаю можно 
1) декомпозировать structs.go на разные файлы и с более связанными с фичами названиями файлов
2) saving_in_db.go хранит все операции записи в базу данных, думаю это не совсем хорошо

Также вместо базы данных используется json local storage, а нужно использовать gorm ORM to store data into sql db

В контексте assignments, надо будет туда добавить дб, структуризировать файлы, сделать код более структурным, используя docker compose поднять контейнеры проекта, сервера и дб

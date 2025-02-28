### Установка:
**Внимание! Не работает с версией postgresql = 17, использовалась последняя стабильная версия 15.***

Удалить /vendor (если есть)
````
go install ariga.io/atlas/cmd/atlas@latest
go mod download
````

### Работа с миграциями
**Никаких diff миграций на prod и test, если только создавать новую БД**
#### Создать diff
````
make db-diff
````
#### Применить миграции
````
make db-migrate
````
#### Статус миграций
````
make db-migration-status
````
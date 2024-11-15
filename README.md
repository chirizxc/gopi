# GoPI - REST API для работы с GIF.

## Пример использования:

### 1. Загрузка GIF
```bash
curl -X POST http://localhost:1111/save -F "file=@<path>"
```
### 2. Получение всех путей и `uuid` GIF-изображений.
```bash
curl http://localhost:1111/gifs
```
### 3. Получение конкретного GIF по `uuid` или `alias`
```bash
curl http://localhost:1111/gif/<uuid or alias>
```

## TODO:

- [x] Перенести логику работы с БД в другой модуль 
- [x] Переписать структуру БД
- [ ] Сделать для каждого `UUID` alias
- [x] Добавить логгер [???](https://t.me/c/2420815282/926)
- [ ] Добавить авторизацию
- [ ] Добавить метод для удаления GIF
- [ ] Добавить тесты
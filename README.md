# GoPI - REST API для работы с GIF.

## Пример использования:

### 1. Загрузка GIF
```bash
curl -X POST http://localhost:1111/create -F "gif=@<path>/assets/xd.gif"
```
### 2. Получение всех путей и `uuid` GIF-изображений.
```bash
curl http://localhost:1111/gifs
```
### 3. Получение конкретного GIF по UUID
```bash
curl curl http://localhost:1111/files/<uuid>
```
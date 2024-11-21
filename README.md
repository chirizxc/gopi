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
### 4. Удаление GIF по `uuid` или `alias`
```bash
curl -X DELETE http://localhost:1111/delete/<uuid or alias> -u user:pass
```

## Документация
```bash
bun install --dev && bun run docs:dev
```

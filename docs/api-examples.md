# Примеры работы с API GoPI

<script setup>
import { useData } from 'vitepress'
const { site, theme, page, frontmatter } = useData()
</script>

<style>
.api-endpoint {
  font-family: Arial, sans-serif;
  font-weight: bold;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  display: inline-block;
}

.get {
  background-color: #8F00FF;
}

.post {
  background-color: #49cc90;
}

.put {
  background-color: #fca130;
}

.delete {
  background-color: #f93e3e;
}
</style>

## API Команды

### 1. <span class="api-endpoint post"> POST</span> Загрузка GIF
```bash
curl -X POST http://localhost:1111/save -F "file=@<path>"
```

### 2. <span class="api-endpoint get"> GET</span> Получение всех UUID GIF-изображений
```bash
curl http://localhost:1111/gifs
```

### 3. <span class="api-endpoint get"> GET</span> Получение конкретного GIF по UUID или Alias
```bash
curl http://localhost:1111/gif/<uuid or alias>
```

### 4. <span class="api-endpoint delete"> DELETE</span> Удаление GIF по UUID или Alias
```bash
curl -X DELETE http://localhost:1111/delete/<uuid or alias> -u user:pass
```

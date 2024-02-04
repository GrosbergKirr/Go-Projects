# Приложение реализующее систему обработки транзакций платёжной системы
**В качестве роутера использов алась библиотека chi. Все остальные элементы из стандартной библиотеки Golang.**

**Запуск с помощью Docker compose:**
```
docker compose up --force-recreate --build
```

**Для локального запуска в в файле internal/config/config.go в функции Mustload в 25 строке написать:**
```
os.Setenv("CONFIG_PATH", "./config/local.yaml")
```


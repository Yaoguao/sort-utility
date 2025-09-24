# sort-utility

Упрощённый аналог UNIX-утилиты sort, написанный на Go.
Работает с файлами

## Особенности

- Сортировка по колонкам (`-k N`)
- Числовая сортировка (`-n`)
- Человеко-читаемая сортировка (`-h`) — например, `10K`, `5M`
- Сортировка по месяцам (`-M`)
- Обратный порядок (`-r`)
- Удаление дубликатов (`-u`)
- Проверка отсортированности (`-c`)

## Пример использования

Файл с такими водными данными
```
ssd 1 Samsung 500G
ssd 100 Samsung 1024G
USB-Flash 3 Kingston 6000M
hdd 2 Samsung 2000G
hdd 100 Samsung 1024G
```

```
➜  go-utility git:(master) ✗ go run . -k 2 test.txt
ssd 1 Samsung 500G
ssd 100 Samsung 1024G
hdd 100 Samsung 1024G
hdd 2 Samsung 2000G
USB-Flash 3 Kingston 6000M
```

```
➜  go-utility git:(master) ✗ go run . -k 2 -n test.txt 
ssd 1 Samsung 500G
hdd 2 Samsung 2000G
USB-Flash 3 Kingston 6000M
ssd 100 Samsung 1024G
hdd 100 Samsung 1024G
```

```
➜  go-utility git:(master) ✗ go run . -k 4 -h -r test.txt  
hdd 2 Samsung 2000G
hdd 100 Samsung 1024G
ssd 100 Samsung 1024G
ssd 1 Samsung 500G
USB-Flash 3 Kingston 6000M
```
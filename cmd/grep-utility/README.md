# grep-utility

Упрощённый аналог UNIX-утилиты grep, написанный на Go.
Фильтрует текстовый поток и выводит строки, соответствующие заданному шаблону (подстроке или регулярному выражению).
Работает с файлами

## Особенности

| Флаг   | Описание                                                                 |
| ------ | ------------------------------------------------------------------------ |
| `-A N` | Вывести N строк после каждой найденной строки (контекст после).          |
| `-B N` | Вывести N строк до каждой найденной строки (контекст до).                |
| `-C N` | Вывести N строк контекста вокруг найденной строки (`-A N -B N`).         |
| `-c`   | Выводить только количество совпадающих строк.                            |
| `-i`   | Игнорировать регистр.                                                    |
| `-v`   | Инвертировать фильтр: вывести строки, не содержащие шаблон.              |
| `-F`   | Воспринимать шаблон как фиксированную строку, а не регулярное выражение. |
| `-n`   | Выводить номер строки перед совпадением.                                 |


## Пример использования

Файл с такими водными данными
```
Error: failed to start service
Warning: low disk space
Info: user logged in
error: file not found
DEBUG: connection established
Info: request received
ERROR: timeout while connecting
Note: update available
info: shutting down
error: failed to allocate memory
Success: backup completed
Warning: CPU temperature high
Info: update installed
Debug: cache cleared
error: network unreachable
ERROR: failed to load module
Info: user logged out
note: maintenance scheduled
Warning: disk almost full
Debug: session terminated
error: permission denied
```

```
➜  go-utility git:(master) ✗ go run ./cmd/grep-utility -i "error" testfile.txt
Error: failed to start service
error: file not found
ERROR: timeout while connecting
error: failed to allocate memory
error: network unreachable
ERROR: failed to load module
error: permission denied

```

```
➜  go-sort-utility git:(master) ✗ go run ./cmd/grep-utility -C 2 -F -n "user logged in" testfile.txt
1:Error: failed to start service
2:Warning: low disk space
3:Info: user logged in
4:error: file not found
5:DEBUG: connection established
```

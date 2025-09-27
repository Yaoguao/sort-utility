# cut-utility

Упрощённый аналог UNIX-утилиты cut, написанный на Go.
Позволяет вырезать определённые поля (колонки) из строк текста, используя разделитель.

## Особенности

| Флаг | Описание                                                 |
|------|----------------------------------------------------------|
| `-f N` | Указание номеров полей (колонок), которые нужно вывести. |
| `-d` | Указание разделителя (по умолчанию табуляция `\t`).      |
| `-s` | Игнорировать строки, которые не содержат разделителя.    | |


## Пример использования

Файл с такими водными данными
```
id;name;age;city
1;Alice;23;Berlin
2;Bob;35;Paris
DATASET
3;Charlie;29;London
4;David;41;Rome
5;Eve;30;Madrid
```

```
➜  go-utility git:(master) ✗ go run ./cmd/cut-utility -f 1 -d ";" < testfile.txt 
id
1
2
DATASET
3
4
5
```

```
➜  go-utility git:(master) ✗ go run ./cmd/cut-utility -f 1,2,4 -s -d ";" < testfile.txt
id;name;city
1;Alice;Berlin
2;Bob;Paris
3;Charlie;London
4;David;Rome
5;Eve;Madrid
```

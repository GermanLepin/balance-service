## Описание ТЗ:

*Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.*

Полное описание задания находится по ссылке https://github.com/avito-tech/autumn-2021-intern-assignment 

## Инструкция по разворачиванию приложения:

Клонируйте репозиторий с помощью git

    https://github.com/GermanLepin/Tech_task.git

Запуск проекта:

    make start-project

Автоматически поднимается сервер и база данных
Сервер вернет *Ping OK!* Это будет означать, что сервер запущен и коннект с базой данных установлен. 
Также будет выполнена миграция по созданию таблиц users и descriptions.

В основном для тестирования API я использую Postman или Insomnia.

|Ключ| Тип данных| Описание|
|---|---|---|
|id | int64| является положительным уникальным идентификатором пользователя |
|amount| float64| сумма зачисления/списания средств на счет/с счета является положительным и может содержать только два значения после точки|

Реализован метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить. POST метод.

    localhost:9000/up-balance

*Добавить в тело запроса(JSON формат):*
```
  {
    "id":"1",
    "amount":"10590.55"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id": 1,
    "top up an amount": 10590.55
  }
```

Метод списания средств с баланса. Принимает id пользователя и сколько средств списать. Patch метод.

    localhost:9000/writing-off

*Добавить в тело запроса(JSON формат):*
```
  {
    "id":"1",
    "amount":"590.55"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id": 1,
    "writing off an amount": 590.55
  }
```

Метод перевода средств от пользователя к пользователю. Принимает id пользователя с которого нужно списать средства и id пользователя которому должны зачислить средства, а также сумму. Patch метод.

    localhost:9000/user-to-user

*Добавить в тело запроса(JSON формат):*
```
  {
    "id1":"1",
    "id2":"2",
    "amount":"1700"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id sender": 1,
    "writing off an amount": 1700,
    "user id recipient": 2
  }
```

Метод получения текущего баланса пользователя. Принимает id пользователя. Баланс всегда в рублях. GET метод.

    localhost:9000/balance-info

*Добавить в тело запроса(JSON формат):*
```
  {
    "id":"1"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id": 1,
    "balance": 8300
  }
```

# Дополнительное задание 1:

Эффективные менеджеры захотели добавить в наши приложения товары и услуги в различных от рубля валютах. Необходима возможность вывода баланса пользователя в отличной от рубля валюте.

Задача: добавить к методу получения баланса доп. параметр. Пример: ?currency=USD. Если этот параметр присутствует, то мы должны конвертировать баланс пользователя с рубля на указанную валюту. Данные по текущему курсу валют можно взять отсюда https://exchangeratesapi.io/ или из любого другого открытого источника.

Примечание: напоминаем, что базовая валюта которая хранится на балансе у нас всегда рубль. В рамках этой задачи конвертация всегда происходит с базовой валюты.

Метод получения текущего баланса пользователя в иностранной валюте. Принимает id пользователя и currency=USD. GET метод.

      localhost:9000/balance-info?convert&currency=USD

*Добавить в тело запроса(JSON формат):*
```
  {
    "id":"1"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id": 1,
    "balance": 117.98
  }
```
(Ответ всегда будет разным, так как, к сожалению, или к счастью, курс валюты не константа)

# Дополнительное задание 2:

Пользователи жалуются, что не понимают за что были списаны (или зачислены) средства.

Задача: необходимо предоставить метод получения списка транзакций с комментариями откуда и зачем были начислены/списаны средства с баланса. Необходимо предусмотреть пагинацию и сортировку по сумме и дате.

|Ключ| Тип данных| Описание|
|---|---|---|
|id | int64| является положительным уникальным идентификатором пользователя |
|amount| float64| сумма зачисления/списания средств на счет/с счета является положительным и может содержать только два значения после точки|
|description | string| описание транзакции за что были списаны/зачислены средства |
|sender_receiver | string| информация о откуда/куда были списаны/зачислены средства |
|refilll | string| при указании в этом поле T - означает, что было пополнение баланса пользователя, при указании F - означает, что было списание с баланса пользователя,|

Добавление записи с описанием. POST метод.

    localhost:9000/descriptions/add  

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1",
    "amount": "6780",
    "description": "Покупка наушников",
    "sender_receiver": "Avito",
    "refilll": "F"
  }
```

*Ответ запроса(JSON формат):*
```
  {
    "user id": 1,
    "balance at moment": 1520,
    "amount": 6780,
    "description of transaction": "Покупка наушников",
    "sender or receiver": "Avito",
    "refilll the balance": "F"
  }
```

Необходимо добавить несколько произвольных записей в БД. POST метод.

    localhost:9000/description/add  

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1",
    "amount": "5490",
    "description": "Продажа куртки",
    "sender_receiver": "Avito",
    "refilll": "T"
  }
```

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1",
    "amount": "1270",
    "description": "Покупка книги",
    "sender_receiver": "Avito",
    "refilll": "F"
  }
```

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "2",
    "amount": "7490",
    "description": "Продажа зеркала",
    "sender_receiver": "Avito",
    "refilll": "T"
  }
```

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "2",
    "amount": "3270",
    "description": "Покупка крана",
    "sender_receiver": "Avito",
    "refilll": "F"
  }
```

Получение описаний транзакцйий(description)

Параметры тела запроса:
|Ключ| Тип данных| Описание| Признак |
|---|---|---|---|
|id | int64| является положительным уникальным идентификатором пользователя | Не обезательно |
|sort_by | string| в данном параметре необходимо указать по какому полю провести сортировку, 
  сейчас доступны два поля created_at(дата создания) и amount(сумма операции)| Не обезательно |
|order_by | string| в данном поле необходимо указать в каком порядке необходимо сделать сортировку, по умолчанию всегда указан ASC(от меньшего к большему), можно указать desc(от большего к меньшему) | Не обезательно |

Если тело JSON будет пыстым, то в ответ вернется список всех записей (description) отсортированный по дате создания описания, и в порядке от ранее созданного к более позднему.

Получение всех записей (description). Метод GET.
    
    localhost:9000/descriptions/get

*Ответ запроса(JSON формат):*
```
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1520,"user id":1,"created at":"2021-10-28T01:17:08.660784Z","refill":"F"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":7010,"user id":1,"created at":"2021-10-28T01:32:37.122076Z","refill":"T"}
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5740,"user id":1,"created at":"2021-10-28T01:32:46.778208Z","refill":"F"}
{"id":4,"sender receiver":"Avito","amount":7490,"description":"Продажа зеркала","balance at moment":9190,"user id":2,"created at":"2021-10-28T01:32:56.434473Z","refill":"T"}
{"id":5,"sender receiver":"Avito","amount":3270,"description":"Покупка крана","balance at moment":5920,"user id":2,"created at":"2021-10-28T01:33:08.176709Z","refill":"F"}
```

Сортировка всех записей по полю (amount)сумма. Метод GET.

    localhost:9000/descriptions/get

*Добавить в тело запроса(JSON формат):*
```
  {
    "sort_by":"amount"
  }
```

*Ответ запроса(JSON формат):*
```
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5430,"user id":1,"created at":"2021-12-12T16:00:43.849871Z","refill":"F"}
{"id":5,"sender receiver":"Avito","amount":3270,"description":"Покупка крана","balance at moment":3270,"user id":3,"created at":"2021-12-12T16:00:58.846923Z","refill":"T"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":6700,"user id":1,"created at":"2021-12-12T16:00:32.103028Z","refill":"T"}
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1210,"user id":1,"created at":"2021-12-12T15:38:53.812268Z","refill":"F"}
{"id":4,"sender receiver":"Avito","amount":7490,"description":"Продажа зеркала","balance at moment":9190,"user id":2,"created at":"2021-12-12T16:00:51.421637Z","refill":"T"}
```

Сортировка всех записей по полю (amount)сумма от большеего к меньшему. Метод GET.

    localhost:9000/descriptions/get

*Добавить в тело запроса(JSON формат):*
```
  {
    "sort_by":"amount",
	  "order_by":"desc"
  }  
```

*Ответ запроса(JSON формат):*
```
{"id":4,"sender receiver":"Avito","amount":7490,"description":"Продажа зеркала","balance at moment":9190,"user id":2,"created at":"2021-12-12T16:00:51.421637Z","refill":"T"}
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1210,"user id":1,"created at":"2021-12-12T15:38:53.812268Z","refill":"F"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":6700,"user id":1,"created at":"2021-12-12T16:00:32.103028Z","refill":"T"}
{"id":5,"sender receiver":"Avito","amount":3270,"description":"Покупка крана","balance at moment":3270,"user id":3,"created at":"2021-12-12T16:00:58.846923Z","refill":"T"}
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5430,"user id":1,"created at":"2021-12-12T16:00:43.849871Z","refill":"F"}
```

Получение записей определеного пользователя. GET метод.

    localhost:9000/descriptions/get

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1"
  }
```

*Ответ запроса(JSON формат):*
```
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1520,"user id":1,"created at":"2021-10-28T01:17:08.660784Z","refill":"F"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":7010,"user id":1,"created at":"2021-10-28T01:32:37.122076Z","refill":"T"}
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5740,"user id":1,"created at":"2021-10-28T01:32:46.778208Z","refill":"F"}
```

Получение записей определеного пользователя с сортировкой по полю amount(сумма). GET метод.

    localhost:9000/descriptions/get

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1",
    "sort_by":"amount"
  }
```

*Ответ запроса(JSON формат):*
```
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5740,"user id":1,"created at":"2021-10-28T01:32:46.778208Z","refill":"F"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":7010,"user id":1,"created at":"2021-10-28T01:32:37.122076Z","refill":"T"}
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1520,"user id":1,"created at":"2021-10-28T01:17:08.660784Z","refill":"F"}
```
Получение записей определеного пользователя. GET метод.

    localhost:9000/descriptions/get

*Добавить в тело запроса(JSON формат):*
```
  {
    "id": "1",
    "sort_by":"amount",
    "order_by":"desc"
  }
```

*Ответ запроса(JSON формат):*
```
{"id":1,"sender receiver":"Avito","amount":6780,"description":"Покупка наушников","balance at moment":1520,"user id":1,"created at":"2021-10-28T01:17:08.660784Z","refill":"F"}
{"id":2,"sender receiver":"Avito","amount":5490,"description":"Продажа куртки","balance at moment":7010,"user id":1,"created at":"2021-10-28T01:32:37.122076Z","refill":"T"}
{"id":3,"sender receiver":"Avito","amount":1270,"description":"Покупка книги","balance at moment":5740,"user id":1,"created at":"2021-10-28T01:32:46.778208Z","refill":"F"}
```

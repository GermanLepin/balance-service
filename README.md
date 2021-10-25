## Описание ТЗ:

*Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.*

Полное описание задания находится по ссылке https://github.com/avito-tech/autumn-2021-intern-assignment 

На дааный момент выполнены все пункты основного и дополнительного задания, кроме unit/интеграционных тестов 

## Инструкция по разворачиванию приложения:

Склонируйте репозиторий с помощью git

    https://github.com/GermanLepin/Tech_task.git

Стартуем docker-compose.yaml: 

    docker-compose up -d 

Использя makefile создайте базу данных и две таблицы

    make create_database 
    make create-users-table 
    make create-description-table

Запустите сервер командой 

    go run cmd/app/main.go  

Сервер вернет *Ping OK!* Это будет означать, что сервер запущен и конект с базой даных установле.

В основном для тестиролвания API я использую Postman или Insomnia.

|Ключ| Тип данных| Описание|
|---|---|---|
|id | int64| является положительным уникальным идентификатором пользователя |
|amount| float64| сумма зачисления/списания средств на счет/с счета является положительным и может содержать только два значения после точки|

Реализован метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить. POST метод.

    localhost:9000/up-balance?id=1&amount=15000.00

Ответ запроса:
```
  {
    "user id": 1,
    "top up an amount": 15000
  }
```

Метод списания средств с баланса. Принимает id пользователя и сколько средств списать. Patch метод.

    localhost:9000/writing-off?id=1&amount=2600.50

Ответ запроса:
```
  {
    "user id": 1,
    "writing off an amount": 2600.5
  }
```

Метод перевода средств от пользователя к пользователю. Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму. Patch метод.

    localhost:9000/user-to-user?id1=1&id2=2&amount=5500.50

Ответ запроса:
```
  {
    "user id sender": 1,
    "writing off an amount": 5500.5,
    "user id recipient": 2
  }
```

Метод получения текущего баланса пользователя. Принимает id пользователя. Баланс всегда в рублях. GET метод.

    localhost:9000/balance-info?id=1

Ответ запроса:
```
  {
    "user id": 1,
    "balance": 6899
  }
```

# Доп. задание 1:

Эффективные менеджеры захотели добавить в наши приложения товары и услуги в различных от рубля валютах. Необходима возможность вывода баланса пользователя в отличной от рубля валюте.

Задача: добавить к методу получения баланса доп. параметр. Пример: ?currency=USD. Если этот параметр присутствует, то мы должны конвертировать баланс пользователя с рубля на указанную валюту. Данные по текущему курсу валют можно взять отсюда https://exchangeratesapi.io/ или из любого другого открытого источника.

Примечание: напоминаем, что базовая валюта которая хранится на балансе у нас всегда рубль. В рамках этой задачи конвертация всегда происходит с базовой валюты.

Метод получения текущего баланса пользователя в иностраной валюте. Принимает id пользователя и currency=USD или currency=RUB. GET метод.

      localhost:9000/balance-info/convert?id=1&currency=USD

Ответ запроса:
```
  {
    "user id": 1,
    "balance": 98.08
  }
```

# Доп. задание 2:

Пользователи жалуются, что не понимают за что были списаны (или зачислены) средства.

Задача: необходимо предоставить метод получения списка транзакций с комментариями откуда и зачем были начислены/списаны средства с баланса. Необходимо предусмотреть пагинацию и сортировку по сумме и дате.

|Ключ| Тип данных| Описание|
|---|---|---|
|id | int64| является положительным уникальным идентификатором пользователя |
|amount| float64| сумма зачисления/списания средств на счет/с счета является положительным и может содержать только два значения после точки|
|description | string| описание транзакции за что были списаны/зачислены средства |
|sender_receiver | string| информация о откуда/куда были списаны/зачислены средства |
|refill | string| при указании в этом поле T - означает, что было поплнение баланса пользователя, при указании F - означает, что было списание с баланса пользователя,|

Добавлении записи с описанием. POST метод.

    localhost:9000/description/add?id=1&amount=6780&description=Покупка наушников&sender_receiver=Avito&refill=F

Ответ запроса:
```
{
  "user id": 1,
  "balance at moment": 119,
  "amount": 6780,
  "description of transaction": "Покупка наушников",
  "sender or receiver": "Avito",
  "refill the balance": "F"
}
```

Получения всех записей по id определенного/определенных пользователя/пользователей. GET метод.

    localhost:9000/description/get-user?id=1

Ответ запроса:
```
  {
    "Id": 1,
    "SenderReceiver": "Avito",
    "Amount": 6780,
    "Description": "Покупка наушников",
    "BalanceAtMoment": 119,
    "UserId": 1,
    "CreatedAt": "2021-10-25T04:42:53.904609Z",
    "Refil": "F"
  }
```

Получения всех записей. GET метод.

    localhost:9000/description/get-all

Ответ запроса:
```
  {
    "Id": 1,
    "SenderReceiver": "Avito",
    "Amount": 6780,
    "Description": "Покупка наушников",
    "BalanceAtMoment": 119,
    "UserId": 1,
    "CreatedAt": "2021-10-25T04:42:53.904609Z",
    "Refil": "F"
  }
```

Для следующих методов нужно добавить несколько записей. POST метод.

    localhost:9000/description/add?id=2&amount=7980&description=Продажа куртки&sender_receiver=Avito&refill=T
    localhost:9000/description/add?id=2&amount=1780&description=Покупка книги&sender_receiver=Avito&refill=F
    localhost:9000/description/add?id=1&amount=4780&description=Продажа зеркала&sender_receiver=Avito&refill=T
    localhost:9000/description/add?id=3&amount=19780&description=Продажа телефона&sender_receiver=Avito&refill=T
    localhost:9000/description/add?id=2&amount=2980&description=Покупка крана&sender_receiver=Avito&refill=F
    localhost:9000/description/add?id=1&amount=3510&description=Покупка колонки&sender_receiver=Avito&refill=F

Получения всех записей. GET метод.

    localhost:9000/description/get-all


Ответ запроса:
```
{"Id":1,"SenderReceiver":"Avito","Amount":6780,"Description":"Покупка наушников","BalanceAtMoment":119,"UserId":1,"CreatedAt":"2021-10-25T04:42:53.904609Z","Refil":"F"}
{"Id":2,"SenderReceiver":"Avito","Amount":7980,"Description":"Продажа куртки","BalanceAtMoment":13480.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:35.558063Z","Refil":"T"}
{"Id":3,"SenderReceiver":"Avito","Amount":1780,"Description":"Покупка книги","BalanceAtMoment":11700.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:47.364289Z","Refil":"F"}
{"Id":4,"SenderReceiver":"Avito","Amount":4780,"Description":"Продажа зеркала","BalanceAtMoment":4899,"UserId":1,"CreatedAt":"2021-10-25T05:12:56.169917Z","Refil":"T"}
{"Id":5,"SenderReceiver":"Avito","Amount":19780,"Description":"Продажа телефона","BalanceAtMoment":19780,"UserId":3,"CreatedAt":"2021-10-25T05:13:11.61466Z","Refil":"T"}
{"Id":6,"SenderReceiver":"Avito","Amount":2980,"Description":"Покупка крана","BalanceAtMoment":8720.5,"UserId":2,"CreatedAt":"2021-10-25T05:13:32.135379Z","Refil":"F"}
{"Id":7,"SenderReceiver":"Avito","Amount":3510,"Description":"Покупка колонки","BalanceAtMoment":1389,"UserId":1,"CreatedAt":"2021-10-25T05:13:57.901295Z","Refil":"F"}
```

Предусмотрена сортировка по сумме и дате как и для какого-то конкретного/конкретных пользователя/пользователей, так и для всех пользователей, также есть способ сортировки от большего к меньшему(desc)

Сортировка от меньшего к большему с указанием пользователя. Метод GET.

    localhost:9000/description/get-user/sort?by=amount&id=1   

Ответ запроса:
```
{"Id":7,"SenderReceiver":"Avito","Amount":3510,"Description":"Покупка колонки","BalanceAtMoment":1389,"UserId":1,"CreatedAt":"2021-10-25T05:13:57.901295Z","Refil":"F"}
{"Id":4,"SenderReceiver":"Avito","Amount":4780,"Description":"Продажа зеркала","BalanceAtMoment":4899,"UserId":1,"CreatedAt":"2021-10-25T05:12:56.169917Z","Refil":"T"}
{"Id":1,"SenderReceiver":"Avito","Amount":6780,"Description":"Покупка наушников","BalanceAtMoment":119,"UserId":1,"CreatedAt":"2021-10-25T04:42:53.904609Z","Refil":"F"}
```

Сортировка от большего к меньшему с указанием пользователя. Метод GET.

    localhost:9000/description/get-user/sort_by?desc=amount&id=1  

Ответ запроса:
```
{"Id":1,"SenderReceiver":"Avito","Amount":6780,"Description":"Покупка наушников","BalanceAtMoment":119,"UserId":1,"CreatedAt":"2021-10-25T04:42:53.904609Z","Refil":"F"}
{"Id":4,"SenderReceiver":"Avito","Amount":4780,"Description":"Продажа зеркала","BalanceAtMoment":4899,"UserId":1,"CreatedAt":"2021-10-25T05:12:56.169917Z","Refil":"T"}
{"Id":7,"SenderReceiver":"Avito","Amount":3510,"Description":"Покупка колонки","BalanceAtMoment":1389,"UserId":1,"CreatedAt":"2021-10-25T05:13:57.901295Z","Refil":"F"}
 ```

Сортировка от меньшего к большему всех пользователя. Метод GET.

    localhost:9000/description/get-all/sort?by=amount&id=1     

Ответ запроса:
```
{"Id":3,"SenderReceiver":"Avito","Amount":1780,"Description":"Покупка книги","BalanceAtMoment":11700.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:47.364289Z","Refil":"F"}
{"Id":6,"SenderReceiver":"Avito","Amount":2980,"Description":"Покупка крана","BalanceAtMoment":8720.5,"UserId":2,"CreatedAt":"2021-10-25T05:13:32.135379Z","Refil":"F"}
{"Id":7,"SenderReceiver":"Avito","Amount":3510,"Description":"Покупка колонки","BalanceAtMoment":1389,"UserId":1,"CreatedAt":"2021-10-25T05:13:57.901295Z","Refil":"F"}
{"Id":4,"SenderReceiver":"Avito","Amount":4780,"Description":"Продажа зеркала","BalanceAtMoment":4899,"UserId":1,"CreatedAt":"2021-10-25T05:12:56.169917Z","Refil":"T"}
{"Id":1,"SenderReceiver":"Avito","Amount":6780,"Description":"Покупка наушников","BalanceAtMoment":119,"UserId":1,"CreatedAt":"2021-10-25T04:42:53.904609Z","Refil":"F"}
{"Id":2,"SenderReceiver":"Avito","Amount":7980,"Description":"Продажа куртки","BalanceAtMoment":13480.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:35.558063Z","Refil":"T"}
{"Id":5,"SenderReceiver":"Avito","Amount":19780,"Description":"Продажа телефона","BalanceAtMoment":19780,"UserId":3,"CreatedAt":"2021-10-25T05:13:11.61466Z","Refil":"T"}
```

Сортировка от большего к меньшему всех пользователя. Метод GET.

    localhost:9000/description/get-all/sort_by?desc=amount&id=1  

Ответ запроса:
```
{"Id":5,"SenderReceiver":"Avito","Amount":19780,"Description":"Продажа телефона","BalanceAtMoment":19780,"UserId":3,"CreatedAt":"2021-10-25T05:13:11.61466Z","Refil":"T"}
{"Id":2,"SenderReceiver":"Avito","Amount":7980,"Description":"Продажа куртки","BalanceAtMoment":13480.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:35.558063Z","Refil":"T"}
{"Id":1,"SenderReceiver":"Avito","Amount":6780,"Description":"Покупка наушников","BalanceAtMoment":119,"UserId":1,"CreatedAt":"2021-10-25T04:42:53.904609Z","Refil":"F"}
{"Id":4,"SenderReceiver":"Avito","Amount":4780,"Description":"Продажа зеркала","BalanceAtMoment":4899,"UserId":1,"CreatedAt":"2021-10-25T05:12:56.169917Z","Refil":"T"}
{"Id":7,"SenderReceiver":"Avito","Amount":3510,"Description":"Покупка колонки","BalanceAtMoment":1389,"UserId":1,"CreatedAt":"2021-10-25T05:13:57.901295Z","Refil":"F"}
{"Id":6,"SenderReceiver":"Avito","Amount":2980,"Description":"Покупка крана","BalanceAtMoment":8720.5,"UserId":2,"CreatedAt":"2021-10-25T05:13:32.135379Z","Refil":"F"}
{"Id":3,"SenderReceiver":"Avito","Amount":1780,"Description":"Покупка книги","BalanceAtMoment":11700.5,"UserId":2,"CreatedAt":"2021-10-25T05:12:47.364289Z","Refil":"F"}
```

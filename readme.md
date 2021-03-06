[![Build Status](https://travis-ci.org/sr-2020/eva-gateway.svg?branch=master)](https://travis-ci.org/sr-2020/eva-auth)
# Gateway

- [Деплой платформы](#deploy)
	- [Локально Linux](#deploy-linux)
	- [Docker](#deploy-docker)
- [Локальная установка](#localsetup)
- [Пользователи](#users)
	- [Регистрация](#registration)
	- [Авторизация](#authorization)
	- [Авторизационный токен](#authtoken)
	- [Профиль](#profile)
	- [Список пользователей со статусами](#usersList)
- [Позиционирование](#position)
	- [Отправка уровня слышимости маячков](#sendbeacons)
	- [Список событий](#positions)
- [Биллинг](#billing)
	- [Состояние баланса](#balance)
	- [Список трансферов](#transfers)
	- [Создать трансфер](#transfer)


## <a name="deploy"></a> Деплой платформы
Для внесения изменений на продакшен нужно отредактировать файл `docker-compose.yml`:

Для редактирования переменнных окружения с конфигурацией и секретами необходимо создать файл `ansible/.vault_pass` в который нужно положить пароль.
Узнать пароль можно у владельца репозитория.

### <a name="deploy-linux"></a> Локально Linux
Установить `ansible` и выполнить следующую команду:
```
cd ansible
ansible-vault edit services/platform/secrets
```

### <a name="deploy-docker"></a> Docker
Установить `docker` и выполнить следующие команды:
```
docker run -it --rm -v $(pwd)/ansible:/ansible sr2020/ansible ansible-vault edit services/platform/secrets
```

Закоммитеть изменения в файлах `docker-compose.yml` и `ansible/services/platform/secrets`. Файл `ansible/.vault_pass` находится в `.gitignore`.
При пуше в любую ветку репозитория изменения применятся к проду.

## <a name="localsetup"></a> Локальная установка
Для локальной установки и тестирования нужно выполнить:
```
make install
make up
make test
```
Команда test может пройти не сразу, так как приложение запускается в асинхронном режиме. Нужно будет подождать 1-2 минуты и запустить команду `make test` еще раз.

## <a name="users"></a> Пользователи
#### <a name="registration"></a> Регистрация
Регистрация осуществляется через POST запрос на http://gateway.evarun.ru/api/v1/auth/register

Тело запроса:
```
{
  "email": "example@example.com",
  "password": "hunter2",
  "name": "John Doe"
}
```
Тело ответа:
```
{
  "id": 1,
  "api_key": "MmVDellSdUpKa0h5MFBDdjN1QnlVbEVC"
}
```

Пример:
```
curl -X POST "http://gateway.evarun.ru/api/v1/auth/register" -H "Content-Type: application/json" -d "{\"email\":\"example@example.com\",\"password\":\"hunter2\",\"name\":\"John Doe\"}"
```


#### <a name="authorization"></a> Авторизация
Авторизация осуществляется через POST запрос на http://gateway.evarun.ru/api/v1/auth/login

Тело запроса:
```
{
  "email": "example@example.com",
  "password": "hunter2"
}
```
Тело ответа:
```
{
  "id": 1,
  "api_key": "MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC"
}
```

Пример:
```
curl -X POST "http://gateway.evarun.ru/api/v1/auth/login" -H "Content-Type: application/json" -d "{\"email\":\"example@example.com\",\"password\":\"hunter2\"}"
```

#### <a name="authtoken"></a> Авторизационный токен
Авторизационный токен `api_key` необходим для работы с API.

В каждый момент для конкретного пользователя валиден только один токен (полученный при последнем логине или сразу после регистрации).

К каждому запросу, требующему авторизацию должен добавляться заголовок `Authorization` формата Bearer Token:

```Authorization: Bearer <api_key>```

Авторизационный токен пользователя не являющегося администратором позволяет выполнять действия связанные только с изменением его профиля.

Все действия для получения общей информации не требуют использования авторизационного токена.

#### <a name="profile"></a> Профиль
Получение информации о текущем пользователе осуществляется через GET запрос на http://gateway.evarun.ru/api/v1/auth/profile с авторизационным токеном `api_key`.

Этот кейс может быть полезен, когда нужно получить информацию только по одному конкретному авторизованному пользователю, вместо того, чтобы грузить весь список пользователей.

Так же данные в этом роуте не кэшируются.

Тело ответа:
```
{
  "id": 1,
  "admin": true,
  "name": "Api Tim Cook",
  "status": "free",
  "created_at": "2019-03-24 21:08:00",
  "updated_at": "2019-03-24 21:08:30"
}
```

Пример:
```
curl -X GET "http://gateway.evarun.ru/api/v1/auth/profile" -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC"
```

Редактирование информации о текущем пользователе осуществляется через PUT запрос на http://gateway.evarun.ru/api/v1/auth/profile с авторизационным токеном `api_key`

Тело запроса:
```
{
  "email": "api-test@email.com",
  "password": "secret",
  "name": "Api Tim Cook",
  "status": "free"
}
```
Тело ответа:
```
{
  "id": 1,
  "admin": true,
  "name": "Api Tim Cook",
  "status": "free",
  "created_at": "2019-03-24 21:08:00",
  "updated_at": "2019-03-24 21:08:30"
}
```

Пример:
```
curl -X PUT "http://gateway.evarun.ru/api/v1/auth/profile" -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC" -H "Content-Type: application/json" -d "{\"email\":\"api-test@email.com\",\"password\":\"secret\",\"name\":\"Api Tim Cook\",\"status\":\"free\"}"```
```

#### <a name="usersList"></a> Список пользователей со статусами

Получение информации о статусах всех пользователей осуществляется через GET запрос на http://gateway.evarun.ru/api/v1/users

Данные в этом списке кэшируются на 1 секунду методом автоматического прогревания кэша крон-скриптом.

Полезные поля которые можно отобразить:
 - `name` (имя указанное при регистрации)
 - `status` (статус выбранный пользователем)
 - `created_at` (время когда пользователь зарегистрировался в системе)
 - `updated_at` (время когда пользователь в последний раз обновлял свой профиль)

Тело ответа:
```
[
    {
      "id": 1,
      "admin": true,
      "name": "Api Tim Cook",
      "status": "free",
      "created_at": "2019-03-24 21:08:00",
      "updated_at": "2019-03-24 21:08:30"
    },
    ...
]
```

Пример:
```
curl -X GET "http://gateway.evarun.ru/api/v1/users"
```

## <a name="position"></a>Позиционирование
#### <a name="sendbeacons"></a> Отправка уровня слышимости маячков
Отправка уровня слышимости маячков осуществляется через POST запрос на http://gateway.evarun.ru/api/v1/position/positions

Тело запроса:
```
{
  "beacons": [
    {
      "ssid": "beacon1",
      "bssid": "b0:0a:95:9d:00:0a",
      "level": -50
    }
  ]
}
```

Тело ответа:
```
{
  "id":1001,
  "user_id":1
  "location_id":2,
  "created_at":"2019-10-23 19:46:11"
}
```

Пример:
```
curl -X POST "http://gateway.evarun.ru/api/v1/position/positions" -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC" -H "Content-Type: application/json" -d "{\"beacons\":[{\"ssid\":\"beacon1\",\"bssid\":\"b0:0a:95:9d:00:0a\",\"level\":-50}]}"
```

#### <a name="positions"></a> Список событий
Получение списка событий осуществляется через GET запрос на http://gateway.evarun.ru/api/v1/position/positions

Можно использовать сортировку по полю, указать список необходимых полей для вывода и выставить фильтр, например, по конкретному пользователю.
Так же можно указать параметры `limit` и `offset`.

Пример:
```
http://gateway.evarun.ru/api/v1/position/positions?sort=-id&select=user_id,beacons&filter[user_id]=1&limit=10&offset=100
```

## <a name="billing"></a> Биллинг
#### <a name="balance"></a> Состояние баланса
Получение состояния баланса осуществляется через GET запрос на http://gateway.evarun.ru/api/v1/billing/balance

Пример:
```
curl -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC" "http://gateway.evarun.ru/api/v1/billing/balance"
```

#### <a name="transfers"></a> Список трансферов
Получение списка трансферов осуществляется через GET запрос на http://gateway.evarun.ru/api/v1/billing/transfers

Пример:
```
curl -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC" "http://gateway.evarun.ru/api/v1/billing/transfers"
```

#### <a name="transfe"></a> Создать трансфер
Создание трансфера осуществляется через POST запрос на http://gateway.evarun.ru/api/v1/billing/transfer

Тело запроса:
```
{
  "sin_to": 131,
  "amount": 90,
  "comment": "Test transfer"
}
```

Пример:
```
curl -X POST -H "Authorization: Bearer MmVDDllSdUpKa0h5MFBDdjN1QnlVbEVC" -H "Content-Type: application/json" "http://gateway.evarun.ru/api/v1/billing/transfer" -d "{\"sin_to\":131,\"amount\":90,\"comment\":\"Test transfer\"}"
```

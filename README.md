# CWS [![Go](https://github.com/sharkboy-j/CWS/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/sharkboy-j/CWS/actions/workflows/go.yml)
СWS - corpse whore searcher

Если вкратце, небольшой скрипт для выявления мертвых раздач с рутрекера на вашем торент клиенте

# Как работает:
1) Берем СИДИРУЕМЫЕ разадчи из ващего торент клиента
2) Фильтруем их по наличию у них трекера rutracker
3) Берем хеш каждой раздачи и стучимся на апишку рутрекер с вопросом "а есть ли у вас раздача с таким хешем?"
4) Если нету - скорее всего раздача удалена\поглощена\обновлена\бог его знает что там еще
5) Пытаемся вытянуть из этого торента комментарий(в котором обычно хранится ссылка на раздачу на трекере)
6) Вываливаем вам весь список дохлых куртизанок(что делать с ними дальше, решать вам)
7) Шлем уведомление в телегу если настроили бота

# roadmap:
- [x] добавить уведомления в телегу
- [x] добавить команду ручной проверки
- [x] завернуть эту красоту в докер
- [x] запуск по таймеру
- [ ] добавить поддержку других клиентов по типу transmission\deluge\etc..
- [ ] удалить всё к херам сабачим, переписать заново

# Как запускать:
1) Скачали [отсюда](https://github.com/sharkboy-j/CWS/releases) последний релиз (конкретно файл cws.exe)
2) Положили в нужную папку
3) Рядом с бинарем положили [config.json](https://github.com/sharkboy-j/CWS/blob/main/config.json)
4) Отредактировали config
5) Открываем терминал\консоль, переходим в папку с бинарем, запускаем
6) ![image](https://github.com/sharkboy-j/CWS/assets/13855710/0133bbeb-77b3-42b4-997c-638f3b96c759)

> [!CAUTION]
> Как запускать через докер гуглим сами. Либо юзайте докер файл и билдите сами, либо используйте готовый пакет https://github.com/sharkboy-j/CWS/pkgs/container/cws

> Список env для докера такой-же как и [config.json](https://github.com/sharkboy-j/CWS/blob/main/config.json) 


# Как настроить отправку сообщений в телегу:
1) Пишем в телеге боту https://t.me/BotFather
2) Там для меня всё интуитивно понятно, а если вам нет то читаем https://core.telegram.org/bots/tutorial#obtain-your-bot-token
3) После создания, бот телеги вышлет ссылку на вашего только что созданного бота, вместе с токеном. Один раз переходим по этой ссылке и активируем бота через сообщение /start, иначе он не сможет вам писать сообщения
4) Запихните токен в config
5) Пишем боту https://t.me/userinfobot
6) Он отправит вам инфу по вашему аккаунту где ID будет ChatId. Запихиваем его в конфиг


# Как настроить config?
```
{
"qb_host": "тут указываем IP торент клиента",
"qb_port": тут порт клиента, без ковычек,
"ssl": тут пишем false\true в зависимости от того используете ли вы https,
"qb_login": "логин торент клиента",
"qb_password": "пароль торент клиента",
"rutracker_api_token": "ваш api токен, можно взять в профиле пользователя",
"telegram_token": "токен телеграм бота, как поулчить описано выше",
"telegram_chat_id": ID чата в котором будет жить ваш бот(как получить описано выше),
"only_manual_check": true\false без кавычек. Если хотим что бы проверка запускалась переодически, пишем false.
Если хотим что бы проверка запускалась только когда вы напишите боту /check то пишем true
"duration_seconds": частота выполнения проверок, указывается в секундах. если only_manual_check указан true, то можно забить на это поле
}
```


# Список команд бота(будет дополнятся по мере необходимости)

- `/check` запускает ручную проверку

<img width="607" alt="image" src="https://github.com/user-attachments/assets/367fdb36-cfae-4e10-8e47-c6f21acf81d8">

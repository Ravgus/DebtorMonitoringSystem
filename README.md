Простий сервіс для моніторингу вашої заборгованості в єдиному реєстрі боржників України.

Вкажіть ваші данні для пошуку у реєстрі в .env.local або .env файлі, данні вашої gmail скриньки для відправленя email сповіщень у разі виявлення заборгованості, добавте скрипт до шедулєру і чекайте лист щастя :)

Приклад данних у .env
---

```dotenv
LAST_NAME=Іван
MIDDLE_NAME=Іванович
FIRST_NAME=Ванько

BIRTH_DATE=18.09.1960
SMTP_USER_NAME=ivan@gmail.com
SMTP_PASSWORD=password
SMTP_PORT=587
```

Також вам треба вказати Api Key з capsolver.com в CAP_SOLVER_API_KEY, для того щоб скрипт міг успішно вирішити капчу.

Опціонально, ви можете вказати Proxy для роботи з capsolver.com в PROXY в форматі:

```env
PROXY=ip:port:username:password
```
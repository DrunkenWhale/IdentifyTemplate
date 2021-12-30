# Simple Docs

### Register

url : /api/auth/register

|param|type|value|
|-----|----|----|
|name|string|username|
|password|string|password|
|mailbox|string|mailbox|

response:

```json
{
  "status": 1,
  "message": "Succeed",
  "data": {}
}
```

|param|type|value|
|-----|----|-----|
|status|bool|1 / 0|
|message|string|"Succeed" / "InvalidArgument" /"UnknownPasswordContentType" / ""UserExist"|
|data| object| null|

### Login

url : /api/auth/login

|param|type|value|
|-----|----|-----|
|mailbox|string|mailbox|
|password|string|password|

response:

```json
{
  "status": 1,
  "message": "Succeed",
  "data": {
    "token": "114514 1919810"
  }
}
```

|param|type|value|
|-----|----|-----|
|status|bool|1 / 0|
|message|string|"Succeed" / "InvalidArgument" / "MailboxUnMatchPassword"|
|data| object| {"token" : tokenString}|


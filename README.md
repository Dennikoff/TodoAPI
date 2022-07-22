# ToDo Api
1. [What libs were used](#Libs)


2. [Network](#Network)


3. [Database](#Database)

## What libs were used <a name="Libs"></a>
1. github.com/gorilla/mux<br/>
   For routing and network communication
2. github.com/gorilla/sessions<br/>
   For sessions and cookie generation and usage
3. database/sql<br/>
   Internal go package for DB connection and usage
4. github.com/BurnedSushi/toml<br/>
   For Config generation
5. github.com/go-ozzo/ozzo-validation<br/>
   For data correction and control
6. github.com/sirupsen/logrus<br/>
   For information logging
7. github.com/stretchr/testify<br/>
   For tests
8. golang.org/x/crypto<br/>
   For password encrypting

## Network<a name="Network"></a>

Todo API work on localhost:8080.<br/>
Postman web service was used for network testing.
Todo API provide some features:<br/>
### Create User
You can create new user:

`POST localhost:8080/create`

Request Body should have such json structure:

```
{
   "email": "your@email.org",
   "password": "your_password"
} 
```

If you type incorrect data (for example: "youremail.com") 
validation will return error

### Log In
You can Log In to the system:

`POST localhost:8080/login`

Request Body should have the same json structure as create:

```
{
   "email": "your@email.org",
   "password": "your_password"
} 
```
After that you will get ***Set-Cookie*** Header and your 
browser will set session cookie. And with this cookie 
you will have permission to 

`localhost:8080/private`

### Who am I
You can easily find out what user logged in:

`GET localhost:8080/private/whoami`

The User taken from **request.Context**

### Create TODO
You can create todo:

`POST localhost:8080/private/create`

Request Body should have such json structure:

```
{
   "header": "Your Header",
   "text": "Your text"
} 
```

### Get TODOs
You can get your tasks:

`GET localhost:8080/private/get`

This request returns slice of todos:
```
{{"id": int, "user_id": int, "header": "string",
  "text": "string", "created_date": time.Time}, ... }
```
## Database<a name="Database"></a>
 
For Data storage I use PostgreSQL. I have two tables:

#### User

|  id   |   email   |  password (encrypted)  |
|:-----:|:---------:|:----------------------:|
|  int  |  varchar  |        varchar         |

#### Todo

|  id   |  user_id  |  header   |  text  |  created_date  |
|:-----:|:---------:|:---------:|:------:|:--------------:|
|  int  |    int    |  varchar  |  text  |      date      |

More information you can find in [Sample](./sample)






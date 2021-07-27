# GTA Open Web

Website / Control Panel made for [GTA Open](https://github.com/PatrickGTR/gta-open) gamemode

Preview of the website ca be seen here: https://imgur.com/a/tN38jGx

## Back end API

- Written in **Go** and **MySQL**

## Endpoints

### User

**POST** - /user/ `User authentication (Login)` recieves a json data
```json
{
    "username": "username",
    "password": "password"
}
```

**GET** - /user/ `Grabs all the registered account, statistics, and item data`

**GET** - /user/{userid} `Grabs account, statistics and item data`

### Server

**GET** - /server/stats `Grabs the data dependong on option specified, available options are 1 -> most kills, 2 -> most money, 3 -> most deaths`

**GET** - /server/banlist `Retrieves all the banned accounts`

### Media
**POST** - /media/ `Creates a new media entry`
```json
{
    "youtubeLink": "Link",
    "title": "Title"
}
```

**POST** - /media/add_views `increments views of a post`
```json
{
    "mediaid": "id",
}
```

**POST** - /media/comment/ `Creates a new comment entry`
```json
{
    "mediaid": "id",
    "comment": "hello world"
}
```

**GET** - /media/comment/{id} `Retrieves all the comments of the specified media id`

**GET** - /media/ `Retrieves all data of a posted media`

**GET** - /media/{id} `Retrieves the data of the specified media`

**GET** - /media/trending/ `Retrieves the data of the 'trending' post, takes a parameter q(uery), available options are hottest and newest`


## Front end

- Written in **React & Next JS**

# How to build

## Development

Fill and add the following environment variables

```js
SECRET_KEY=""
MYSQL_USERNAME=""
MYSQL_DATABASE=""
MYSQL_PASSWORD=""
MYSQL_SERVER=""
ENV="ENV"
```

and run...
```batch
cd front-end
npm run dev

cd ..

cd back-end
go run main.go
```

## Production

Fill and add the following environment variables

```js
SECRET_KEY=""
MYSQL_USERNAME=""
MYSQL_DATABASE=""
MYSQL_PASSWORD=""
MYSQL_SERVER=""
ENV="PROD"
```

and run...
```batch
docker-compose up
```

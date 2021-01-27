# GTA Open Web

Website / Control Panel made for [GTA Open](https://github.com/PatrickGTR/gta-open) a hobby project of mine.

## Back end

- Written in **Go**

### Restful API Structure

**POST** - /user/ `User authentication (Login)`

**takes in form-data**

```html
username: value; password: value
```

**GET** - /user/ `Grabs all the registered account, statistics, and item data`

**GET** - /user/uid `Grabs account, statistics and item data`

Currenly WIP ... more endpoints will be added soon!

## Front end

- Written in **React & Next JS**

Work in progress.

# How to build

## Development mode

```batch
cd front-end
npm run dev

cd ..

cd back-end
go run main.go
```

## Production deployment

```batch
WIP
```

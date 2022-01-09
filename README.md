''

## Here we go

#### 1. Database

This application uses MySQL 8.0+ (download here: https://dev.mysql.com/downloads/mysql/).

Config username and password, config your MySQL port, and name (which should be 'inventory'). 

After above are all set, open a command terminal window, execute:

```
mysql -u root -p
```

Execute `.sql` file using command line:

```
CREATE DATABASE inventory CHARACTER SET utf8 COLLATE utf8_bin;
```

Then use the created database:

```
USE inventory
```

Execute the `.sql` file, which is in the root directory, with `source`, using absolute path

```
SOURCE xxx/xxx/xxx/inventory.sql
```

Finally, change the connection URI in `config/config.go`, format should be `username:password@tcp(127.0.0.1:port)`



### 2. application

#### Go

Go to `https://go.dev/dl/` to download.

#### Run

Under the project's root directory, run in terminal: 

```
go run ./cmd/server/server.go
```

After your see this line:

```
[GIN-debug] Listening and serving HTTP on :4000
```

You know the application is now running successfully.



### 3. To test

1. Please use chrome, and google `Talend API Tester`, then download it.
2. open it in extension



#### 3.1Test CREATE

choose POST

```
http://localhost:4000/create
```

in body input box, copy and paste this:

```
{
	"comment": "5555",
	"from_location": "NY",
	"current_location": "NY",
	"to_location": "NY",
	"original_price": 4500,
	"current_price": 860,
	"weight": 34034,
	"url": "6324",
	"name": "sword"
}
```

#### 3.2 Test Retrieve

choose GET

```
http://localhost:4000/get-list?start=1&size=10
```

#### 3.3 Test Update

choose UPDATE

```
http://localhost:4000/update?query-id=1
```

in body input box, copy and paste this:

```
{
	   	"comment": "Good job"
}
```

#### 3.4 Test Delete

choose GET

```
http://localhost:4000/delete?query-id=1
```

#### 3.5 Test uploading image

Choose POST

```
http://localhost:4000/upload_test
```

On the right upper of `BODY` input box, choose `Form`

Click the button 'add form parameter', input 'photo'; and in dropdown, choose 'File'.

And click the 'choose file' button  to choose your photo.

Near 'add form parameter', choose `multipart/form-data` in dropdown.



Hope you enjoy!
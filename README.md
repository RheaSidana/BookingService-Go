<h1> Clone the repo </h1>

```
git clone https://github.com/RheaSidana/BookingService-Go.git
```

<h1> Run the command:</h1>

```
go mod tidy
```

<h1>3. DB operations:</h1>
<h3><<<< INSTALLING >>>><br/>
    &emsp;a. MAC <br/>
    &emsp;&emsp;1. install postgres:</h3>

```
brew install postgresql
```

<h3>&emsp;&emsp;2. start/stop postgres service:</h3>

```
brew services start/stop postgresql
```

<h1>1. Clone the repo </h1>

```
git clone https://github.com/RheaSidana/BookingService-Go.git
```

<h1>2. Run the command:</h1>

```
go mod tidy
```

<h1>3. DB operations:</h1>
<h3><<<< INSTALLING >>>><br/>
    &emsp;a. MAC <br/>
    &emsp;&emsp;1. install postgres:</h3>

```
brew install postgresql
```

<h3>&emsp;&emsp;2. start/stop postgres service:</h3>


```
brew services start/stop postgresql
```

<h3>&emsp;&emsp;3.</h3>

```
psql postgres
```

<h3>&emsp;&emsp;4. add postgres password:</h3>

```
\password {password};
```



<h3>&emsp;&emsp;3.</h3>

```
psql postgres
```

<h3>&emsp;a. WINDOWS <br/> 
    &emsp;&emsp;1. install postgres: </h3>
<a href="https://www.postgresql.org/download/windows/">[Link text Here]</a>

<h3>&emsp;&emsp;2. port: 5432 (defaut), user: postgres (defaut)</h3>
<h3>&emsp;&emsp;3. add postgres password</h3>
<h3>&emsp;&emsp;4. open psql sql shell</h3>
<br/>
<h3> <<<< CREATING >>>> <br/>&emsp;&emsp;5. create orm db: </h3>

```
CREATE database booking_service;
```

<h3>  <<<< CONNECTING >>>> <br/>&emsp;&emsp; 6. connect to db: </h3>

```
\c "db"
```

<h3>&emsp;&emsp;7. Edit .env file with postgres details</h3>
<br/>


<h1>4. Migrate Tables: </h1>

```
go run .\migrations\migrate.go
```

<h3>&emsp;&emsp;View DB Table Schemas: </h3>

```
\d "tablename"
```

<h1>5. Seed Data to the Table </h1>

```
go run .\dataSeeding\seedData.go
```


<h1>6. Run the application </h1>

```
go run .
```


<h1>7. Call APIs </h1>
Postman: 

```
https://api.postman.com/collections/28378586-85f81d43-b0a0-4ee4-9cf8-34ec0aad13db?access_key=PMAT-01H4P7D5A1283YZQBZXWMZ02YR
```


# refree
golang web application 

# configuration
example of config.toml 
```
[Server]
Address = "localhost:8080"
ReadTimeout = 10
WriteTimeout = 600
Static = "public"
Logsettingfile = "log/seelog.xml"

[DBconfig]
Host = "localhost"
User = "USER"
Dbname = "DATABASE"
Sslmode = "disable"
Password = "******"
Initialize = true
Test = true
```

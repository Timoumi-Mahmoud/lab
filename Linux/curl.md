# Curl  
- is a programming library & a command line tool that is used to transfer muliple data types to and from a server.

- Use many protocols tom move information:
HTTP
Telnet
FTP
IMAP
POP3
SCP
LDAP
SMTP
MQTT ...


## Use case:
1. download a file 
2. check status 
3. FTP client
4. troubleshooting connections
5. test APIs
6. send SMS
7. copy as curl from the browser
.....

```bash 
curl -o file.html http://google.com
```
```bash
curl -ILS http://yahoo.com | grep HTTP 
```
## Fun Usage
```bash
curl ifconfig.me
curl http://wttr.in/Tunisia
curl http://v2.wttr.in/Tunisia
curl qrenco.de/Hello-World
curl parrot.live
```

------------------------------------------------------------------------------------
## Test APIs
#### GET:
```bash
 curl localhost:4000/v1/healthcheck
 curl -v localhost:4000/v1/healthcheck -X GET
 curl -s localhost:4000/v1/healthcheck -o /dev/null -w '%{response_code}'
```

```bash
 curl -I localhost:4000/v1/healthcheck  -X GET
```

#### POST
 curl -X POST -d'{email:bara@gmail, password: secretPWD}' localhost:4000/v1/login -H 'content-type: application/json'
 curl -d'{email:bara@gmail, password: secretPWD}' localhost:4000/v1/login (only works for POST for other method like patch and put I need to add -X PUT)
 curl -d @file_name.json  localhost:4000/v1/login
```
#### PUT
```bash
 curl -X PUT -d'{first_name:bara, last_name:hacker, email:bara@gmail, password: secretPWD}' localhost:4000/v1/user/create
```

#### DELETE
```bash
 curl -X DELETE localhost:4000/v1/user/delete/1
```

#### HEAD
```bash
 curl -I localhost:4000/v1/healthcheck
```

#### Autorization

```bash
 curl localhost:4000/v1/admin -u admin:secret
 curl localhost:4000/v1/admin -H  "Autorization: Basic $(echo -n admin:secret | base64)"
```



--------
s: silent
v: verbose
I: inspect
H: Header
X: request
O: remote-name
d: data
E: certificate
L: location 
D: dump-header

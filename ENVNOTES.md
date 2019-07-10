# USERetl

## Statistics: Difference between 1 & 2 branch architectures

- Ultsys user database ETL to local database instance: writing concurrently to 2 tables.  
- Processed: 3.324M records or approximately 4.5GB.  
- 1 Branches: 1 user query and 1 loop through the 3.324M records.   
- 2 Branches: 1 user query and 2 loops through the 3.324M records (6.648M records).  

### Statistics for ETL of 3,324,305 rows of Ultsys user data

### 1 BRANCH

- Total time: 297.342837 secs (4 mins 57.3 secs).  
- Stage 1: 297.122859 secs (4 mins 57.1 secs) : sent 4,451,661,716 (4.452 GB).  
- Stage 2: 290.869976 secs (4 mins 50.9 secs) : recd 4,451,661,716 (4.452 GB): sent 681,690,227 (681.7 MB).  
- Stage 3:  
    - Writer 1: 207.586308 secs (3 mins 27.6 secs) : recd 681,690,227 (681.7 MB).  
    - Writer 2: 238.688619 secs (3 mins 58.7 secs) : recd 681,690,227 (681.7 MB).  
- Data not utilized: 3.770GB.  

### 2 Branches

- Total time: 721.6150995 secs (12 mins 1.62 secs).  
- Stage 1:   
    - Reader 1: 719.352100 secs (11 mins 59.4 secs) : sent 4,451,705,503 (4.452 MB).  
    - Reader 2: 721.214085 secs (12 mins 1.21 secs) : sent 4,451,705,503 (4.452 MB). 
- Stage 2:
    - Transformer 1: 719.098156 secs (11 mins 59.1 secs) : recd 4,451,705,503 : sent 360,131,723 (4.452 MB/360.1 MB).  
    - Transformer 2: 720.904329 secs (12 mins 0.90 secs) : recd 4,451,705,503 : sent 331,473,313 (4.452 MB/331.5 MB).  
- Satge 3:  
    - Writer 1: 341.293178 secs (5 mins 28.5 secs) : recd 360,131,723 (360.1 MB).  
    - Writer 2: 328.535333 secs (5 mins 23.8 secs) : recd 331,473,313 (331.5 MB).  
- Data not utilized: 3.760GB.  

## Environment variables:

See the build_environment directory.

### Local Server: 

This is the server that Golang will pass the tunneling request to.

- LOCAL_IP: 127.0.0.1
- LOCAL_PORT: 3307

### Local MySQL Server:

The local MySQL server connection (not required for tunneling).

- LOCAL_SQL_SERVER: "root"
- LOCAL_SQL_PASSWORD: "Password" 
- LOCAL_SQL_IP: 127.0.0.1
- LOCAL_SQL_PORT: 3306
- LOCAL_SQL_SCHEMA: "nusa_user"

### Intermediate Server:

The intermediate server that will do the port forwarding from 3307 at 127.0.0.1 
to 3306 at 199.167.77.248.

- INTERMEDIATE_SERVER_USER: "apeters"
- INTERMEDIATE_SERVER_PASSWORD: "/home/apeters/.ssh/id_rsa" 
- INTERMEDIATE_SERVER_IP: 199.167.77.245
- INTERMEDIATE_SERVER_PORT: 22

### Remote MySQL Server:

This is where the tunnel terminates:

- REMOTE_SQL_SERVER_USER: "datastudy"
- REMOTE_SQL_SERVER_PASSWORD: "Password" 
- REMOTE_SQL_SERVER_IP: 199.167.77.248
- REMOTE_SQL_SERVER__PORT: 3306
- REMOTE_SQL_SERVER__SCHEMA: "ultsys"

## Tunnel

To build a connection to a machine like DEV0 a tunnel will have to be created. 
The bash command to use looks like:

- ssh -i /home/user/.ssh/id_rsa -L3307:remote_server_ip:3306 root@intermediate_server_ip
- ssh -i /home/apeters/.ssh/id_rsa -L3307:199.167.77.248:3306 apeters@199.167.77.245
    - Worked 2019-06-10

Golang sends a request to `127.0.0.1` at port `3307` (-L3307) and this gets sent 
through `SSH` on the local machine through the tunnel to the intermediate server 
`SSH` connection. The intermediate server in this case is 
`DEV0`: `root@intermediate_server_ip`. `SSH` is using a private key (`id_rsa`) 
in this case. The tunnel knows to forward this on to the remote server. The 
remote server is the MySQL replica at  `remote_server_ip`. This is the port 
forwarding: 3307 on local host gets forwarded to 3306 on 248. This is how Golang 
sends a request from local host to the replica MySQL database.

Here is an example for the tunnel shell command. It can be issued at `/home/user`
 or just `/`:

Here are the database connections used in the `USERetl`. In this case we are 
taking data from the replica database and passing it into the local database. 
We might instead, take the data and drop it into a CSV, as well, and in that 
there is no need for the local database connection structure.

```
util.Database{
    Host: "127.0.0.1",
    Port: 3306,
    UserID: "root",
    PassWord: "Password",
    Database: "nusa_user",
    ParseTime: true
}


util.Database{
    Host: "127.0.0.1",
    Port: 3307,
    UserID: "datastudy",
    PassWord: "Password",
    Database: "ultsys",
    ParseTime: true}

```

### Local Database Connection
```
util.Database{
    Host: "127.0.0.1",
    Port: 3306,
    UserID: "root",
    PassWord: "Password",
    Database: "nusa_user",
    ParseTime: true
}
```

### Tunnel Database Connection
```
util.Database{
    Host: "127.0.0.1",
    Port: 3307,
    UserID: "datastudy",
    PassWord: "Password",
    Database: "ultsys",
    ParseTime: true}
```

# wow
Word of wisdom TCP server implementation

Start client and server: 
```
docker-compose up
```

Start server in background:
```
docker-compose up server -d
```

Start client:
```
docker-compose up client
```

Stop and remove docker container:
```
docker-compose down
```


# PoW algorithm description
For this particular task i created a simple PoW algorithm.
Server just create a random integer between DIFFICULTY/2 and DIFFICULTY env varible, then calculates sha1 hash for it and sends it to client. 
Client calculates hashes from 0 to DIFFICULTY and when it finds equal hash sends integer to server. Server calculates hash of integer provided and if they are equal returns random quote.

With default difficulty delay is pretty optimal and if someone will try to DDOS it will take significant amount of their processor powers. 

# Tests
There is some test, but coverage is around 50%

# Possible improvements 
1. Store quotes in DB or JSON file
2. Better test coverage, especially for server and client
3. Separate connection handlers and use PoW as a middleware
4. Adjusting difficulty based on RPS or something similar and send difficulty to client in challenge request message
# plan
this is where i plan (? sure)
# goal min: crud api
GET    /users
GET    /users/:id
POST   /users
PATCH  /users
DELETE /users/:id
### stretch goal to include todos as well (likely just go to gin at this point)
GET    /users/{id}/todos
POST   /todos
PATCH  /todos/{id}
DELETE /todos/{id}

connect to db

open up socket and listen for connections
parse req and answer with json

error handling bad requests

auth?
*/

# goals
just btw reminder to go from tcp -> http protocol (start at tcp level) - could go as low as syn syn-ack ack?
1. connect to db (done)
2. open port and accept connections (done)
2b. properly parse and print said connection (done)
3. parse connection and do some action (here)

what is left: 
    need to answer in marshalled json not just []bytes
    wire dbpool into handlers (likely refactor later)
    read the actual body
    routing
    write sql queries


4. return something to the client (hardcoded but works)

 - should i include /api in my endpoints? probably
# After CRUD API complete
After completion, fully test with go tests (however the easiest way/go way would be to auto test api endpoints)
also ci/cd as well 

# Claude helper plan
The path to "finished CRUD API"
In order, one at a time:

1. Wire dbpool into handlers — pass it through handleConnection down into verb functions
2. Read the actual body — Content-Length header → io.ReadFull
3. Add routing — a simple map or if/else on url inside each verb function
4. Add encoding/json — marshal your User struct into the response body, set Content-Type: application/json and correct Content-Length
5. Write actual SQL — SELECT, INSERT, UPDATE, DELETE against your users table with parameterized queries
6. 404 and error responses — default cases, malformed request handling
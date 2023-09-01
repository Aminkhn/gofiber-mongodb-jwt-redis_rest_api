
# Simple User managment REST API (GoFiber - Mongodb - JWT - Redis)

#### A simple user managment REST API
using:
- Golang programming language
- Fiber as Back-End framework  
- JWT (Json Web Token) for authorization
- Redis for storing blacklisted jwt tokens (Securing jwt weakness)
- Mongodb as base database

## User Entity Model
To get an overview of the user model:
```go
  type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bbson:"name"`
	Family    string             `json:"family" bson:"family"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time
}
```
## API Reference
### API Health Check
```http
  GET /api/HealthChecker
```
### Authentication
#### Login
```http
  POST /auth/login
```
Input data format for login request:
```json
{
  "identity": "<either username or email>",
  "password": "your password"
}
```
#### Logout
```http
  POST /auth/logout
```
### User CRUD 
 all CRUD operations in this project need authentication so before this section you should login and have your own authentication token and pass it  through HTTP request Header `Authorization` with the format mentioned below.
| key | Type     | value                      |
| :-------- | :------- | :------------------------------- |
| `Authorization` | `string` |  "Brearer TOKEN"  **Required**.|

- #### operations without id parameter
##### Create user
```http
  POST /api/user
```
Here is an example of the input data structure to create a user:
```json
 {
    "name": "John",
    "family": "Doe",
    "username": "JoDoe",
    "email": "JoDoe@gmail.com",
    "password": "123456789"
}
```
 ##### Get all users
```http
  GET /api/user 
```
Here is an example of a user's output data structure:
```json
{
  "id": "64f2480a33d5061ff203633f",
  "name": "John",
  "family": "Doe",
  "username": "JoDoe",
  "password": "***Hashed password***",
  "email": "JoDoe@gmail.com",
  "CreatedAt": "2023-09-01T20:22:34.046Z"
}
```

- #### operations with id parameter
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` |  id belong to desired userID  **Required**. |
##### Get user by Id
```http
  GET /api/user/id
```
##### Edit user by Id
```http
  PUT /api/user/id
```
##### Edit user by Id
```http
  PATCH /api/user/id
```
##### Delete user by Id
```http
  DELETE /api/user/id
```
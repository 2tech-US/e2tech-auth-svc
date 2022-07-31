# E2TechEcommerce API document
This is a documentation guide for the Authentication Service.

## Database

- note: `createdAt` `updatedAt` is auto generate in everys table

### account
| Column        | Type           |   |
| ------------- |:-------------:| -----:|
| id     | Long | Primary key, auto generate | 
| phone    | String      |   unique, [8,15] character |
| password | String      |    Strong password |


### role
| Column        | Type           |   |
| ------------- |:-------------:| -----:|
| id     | Long | Primary key, auto generate | 
| name    | String      |   unique |
| description | String      |   |


### account_roles
| Column        |
| ------------- |
| account_id    |
| role_id       |


## API (V1)

### Auth

#### Login:
```
POST /api/v1/auth/login
```
Authorization token: `None`

Body parameters available (`Json`):
- `phone`: `required` `string` `not empty` `length=8-15`
- `password`: `required`

#### Register: Driver, passenger
```
POST /api/v1/auth/register/passenger
POST /api/v1/auth/register/driver
```
Authorization token: `None`

Body parameters available (`Json`):
- `phone`: `required` `string` `length=8-15`
- `password`: `required`

#### Register: Call center employee
```
POST /api/v1/auth/admin/register/call-center
```
Authorization token: `admin`

Body parameters available (`Json`):
- `phone`: `required` `string` `length=8-15`
- `password`: `required`


### Account

#### Get Account
```
GET /api/v1/account/{phone}
```
Authorization token: `any role`:
- `Token.payload.phone` = `req.params` or Token.payload.roles have `admin` role 

#### Change Password
```
PUT /api/v1/account/{phone}
```
Authorization token: `any role`:
- `Token.payload.phone` = `req.params` or Token.payload.roles have `admin` role

Body parameters available (`Json`):
- `oldPassword`: `required` `string` `length=8-15`
- `newPassword`: `required` `string` `length=8-15`

#### Get Accounts
```
GET /api/v1/account/admin/
```
Authorization token: `admin`

Query parameters available (Pagination):
- `limit`: `optional` `integer`  `min=1`, limit the result
- `page`: `optional` `integer` `min=1`, page requested, need `limit` to work
- `sort`: `optional` `string`  name of the field to sort results with `default is 'createdAt`
- `order_by`: `optional` `must be either asc or desc`, only work it `sort` is specified, default to `desc`
- `query`: `optional` `string`, search by phone's number

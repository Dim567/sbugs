# sbugs
## Debugging and Problem-Solving (Code Review Task)

### Problems

### 1. Goroutine creation inside request handler
- Issue: request handler already runs in separate goroutine and creation of additional gouroutine inside in this case it's just wasting of the system resources
- Fix: run request handler logic without creation of goroutine

### 2. SQL injection vulnerability in the createUser handler
- Issue: if we send SQL statement as username, it can damage db system
- Fix: use parametrized SQL query

### 3. getUsers returns all users from database
- Issue: request handler returns all records for users which can impact performance in case of very big users count
- Fix: limit number of users per request (with offset)

### 4. Lack of logs
- Issue: if error happens, it's impossible to track what was the issue
- Fix: add logging

### 5. Sensitive information in error response
- Issue: if error happens, clients can potentially see sensitive information in error description
- Fix: remove error details from response to client

### 6. Database requests errors is not handled properly
- Issue: not all database requests errors handled, which can lead to server failure
- Fix: add database request errors handlers

## Additional improvement that can be implemented
### 7. Database settings hardcoded
- Issue: it is not possible to change database configuration without code modifications
- Fix: provide mechanism to get database configuration from config file/cli options/... (not implemented)

### 8. Everyone can add and get list of users
- Issue: there is no restriction of who can read and modify data in the database
- Fix: provide authentication and authorization (not implemented)

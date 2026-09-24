# Calculator App

This application implements a calculator application which performs calculation through a Go backend API and displays calculator interface as a React application.

## Setup instructions
To run application locally run docker comopse:
```docker
docker compose up -d
```
The application will be available under `http://localhost:5173/`

## Design rationale
A simple calculator might work in a following way:
1. User inputs left operand.
2. User inputs operation with a right operand if appropriate and executes operation.
3. Operands are transformed according the operation and stored as a left operand.
4. Repeat from 2.

Such design restricts calculator to performing only a single operation at a time.
Furthermore, with the constraint of the application to perform operations on the backend, this design demands a separate endpoint for every type of operation.

A more advanced calculator pipeline might look in the following way:
1. User input expression.
2. Expression is split into lexical tokens.
3. Tokens are assembled into a syntax tree according to the precedence.
4. Syntax tree is evaluated.

This design allows for execution of more complex expression containing multiple operations and parenthesis.
Additionally, it allows minimizing API to a single request containing the expression and expecting the result or an error message as a response.

For this reason the latter option is chosen for this application.

The component architecture looks the following way:
1. Frontend
    - Allows entering an expression.
    - Sends the expression through to the API.
    - Displays result, or an error message.
    - Allows using the result as a part of the next expression.
2. Backend
    - Receives an expression.
    - Performs lexical and syntactic analysis.
    - Based on the analysis determines expression's value or an error.
    - Responses with the value or an error message.

### API design:
Single endpoint: `POST /calculate` 

POST method is chosen for two reasons:
1. The body is a more convenient way to send expressions, that might contain '+' signs, which in a GET query string without URL encoding would be treated differently. GET requests with a body are possible, but are not encouraged by the standard and might be forbidden by some systems.
2. Endpoint does not change server side state in any way, but might do so in the future, for example if calculation history should be saved in a user session.

Endpoint expects a JSON body in the following format:
```json
{
    "expression": string // The expression to be calculated
}
```

Endpoint response has a 200(OK) or 400(BAD REQUEST) status code and a body of the following format:
```json
{
	"status_code"`: integer // Seperate non-http statu code
	"expression"`: string // The calculated expression
	"result"`: string // Result of the calculated expression if no errors incountered or an empty string
	"error"`: string // Error message if errors incountered or an empty string
}
```
When encountering errors which should be displayed to the user response has status code 200 and 400 otherwise.

Possible status codes:
- 0 = Success 
- 1 = InvalidInput
- 2 = DivisionByZero
- 3 = NegativeSqrt


### API usage:

Simple expression:
```bash
curl --location 'localhost/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "3+4-5"}'
```
response (200):
```json
{
    "status_code": 0,
    "expression": "3+4-5",
    "result": "2",
    "error": ""
}
```

Complex expression:
```bash
curl --location 'localhost/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "3+(4+3^2/(sqrt(3.123))-5"}'
```
response (200):
```json
{
    "status_code": 0,
    "expression": "3+(4+3^2/(sqrt(3.123)))-5",
    "result": "7.092798780987829",
    "error": ""
}
```


Expression with deletion by zero:
```bash
curl --location 'localhost/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "1/(2-2)"}'
```
response (200):
```json
{
    "status_code": 2,
    "expression": "1/(2-2)",
    "result": "",
    "error": "Cannot devide by zero"
    }
```

Expression with invalid syntax:
```bash
curl --location 'localhost/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "1 + 2 +"}'
```
response (400):
```json
{
    "status_code": 1,
    "expression": "1 + 2 +",
    "result": "",
    "error": "unexpected token \"\""
}
```


## Test Coverage
Test coverage reports for both sides are available in ./coverprofiles/ folder

<img src="https://github.com/gatinhodev/verifycat/assets/135276762/8f4b2368-3733-4864-8c28-c7719edf0ece" alt="verifycat-api" width="100" height="100">

# VerifyCat  

**VerifyCat** is a versatile validation API designed to handle various types of validations, including CPF (Brazilian ID number), CNPJ (Brazilian legal entity number), URL, email, and credit card numbers. 

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Overview

VerifyCat is built using the [Gin](https://github.com/gin-gonic/gin) framework for handling HTTP requests and providing a fast and efficient web server. The project is structured to accommodate different validation types, each implemented in a separate file within the `validate` package.

## Architecture

The main entry point of the application is the `verifycat_api.go` file. It sets up the Gin router and defines the endpoint for validation at `/validate`. The validation logic is delegated to specific handlers in the `validate` package.

The `validate` package contains individual files (`cpf.go`, `cnpj.go`, `url.go`, `email.go`, and `creditcard.go`) for each validation type. These files house the validation logic and request handling specific to their validation type.

## API Endpoints

### `POST /validate`

This endpoint supports validation for various types of data. The payload should be in JSON format, containing the `type` (validation type) and `value` (data to be validated).

- **Request Payload Example:**
  ```json
  {
    "type": "cpf",
    "value": "123.456.789-09"
  }
  ```

- **Response Example:**
  ```json
  {
    "isValid": true,
    "message": "CPF"
  }
  ```

### CURL Request Example

```bash
curl -X POST http://localhost:8080/validate -H "Content-Type: application/json" -d '{"type": "cpf", "value": "123.456.789-09"}'
```

- **CURL Response Example:**
  ```json
  {
    "isValid": true,
    "message": "CPF"
  }
  ```

### Supported Validation Types

- `cpf`: Brazilian ID number
- `cnpj`: Brazilian legal entity number
- `url`: URL
- `email`: Email address
- `creditcard`: Credit card number

## RESTful Architecture

The API VerifyCat follows the principles of REST, including:

- **Identifiable Resources:** Each validation type (CPF, CNPJ, URL, etc.) is treated as an identifiable resource. Operations are performed on these resources through specific URLs.

- **Standard HTTP Operations:** CRUD operations (Create, Read, Update, Delete) are mapped to standard HTTP operations. In this code, the main operation is validation, performed through a POST request to the `/validate` resource.

- **Statelessness:** Each client request to the server contains all the information needed to understand and process the request. There is no dependency on intermediate states between requests.

- **Resource Representation:** Resources are represented in JSON in the body of HTTP responses. The response is a representation of the current state of the resource (e.g., whether a CPF is valid or not).

- **HATEOAS (Hypermedia As The Engine Of Application State):** While the provided code does not explicitly include links to other resources in the HATEOAS style, the concept is embedded in the general principle that the client interacts with the API through representations of resources and states provided in responses.

## Usage

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/verifycat.git
   cd verifycat
   ```

2. **Run the Application:**
   ```bash
   go run verifycat_api.go
   ```

3. **Make API Requests:**
   - Use your preferred API client (e.g., cURL, Postman) to send POST requests to `http://localhost:8080/validate` with the appropriate payload.

## Contributing

Feel free to contribute by opening issues, providing feedback, or submitting pull requests.

## License

This project is licensed under the AGPL V3 License - see the [LICENSE](LICENSE) file for details.

# API resources

All endpoints must have `/Content-Type` header with `application/json` value.

## Merchants API

### Ping

<details>
 <summary><code>GET</code> <code><b>/running</b></code> <code>(Checks if the service is healthy)</code></summary>

#### Parameters

> None

#### Responses

##### HTTP Code 200

```json
Welcome to merchants api..!!
```

</details>



### Sign Up

Allows to register a new merchant for payments processing.

<details>
  <summary><code>POST</code> <code><b>/api/v1/merchants/sign-up</b></code></summary>

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | name            |  required | string      | Name of merchant                              |
  > | user.username   |  required | string      | Username to be able to log in in the future   |
  > | user.password   |  required | string      | User password                                 |
  > | user.email      |  required | string      | Email user                                    |

  #### Responses

  ##### HTTP Code 200

  Successful payment

  ```json
  {
    "data": {
      "merchant_code": "67b281fc-52f1-41cc-bd21-0737713fbf75",
      "name": "Sample merchant",
      "user": {
        "username": "sample_merchant",
        "email": "sample_merchant@email.com"
      }
    },
    "status": "ok"
  }
  ```
  ##### HTTP Code 400
  Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```
</details>

### Login

Return token for refund and query payments.
<details>
  <summary><code>POST</code> <code><b>/api/v1/merchants/login</b></code></summary>

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | username   |  required | string      | Username   |
  > | password   |  required | string      | User password                                 |

  #### Responses

  ##### HTTP Code 200

  Successful login

  ```json
  {
    "data": {
      "username": "hzambrano",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk2NDg2OTgsInVzZXJuYW1lIjoiaHphbWJyYW5vIn0.1rWmYm1zglb-Xi00FoK1xhhxozdpbNjUhqenv-hzv94",
      "message": "you are logged..!"
    },
    "status": "ok"
  }
  ```
##### HTTP Code 400
Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```
</details>


## Payments API

### Ping

<details>
  <summary><code>GET</code> <code><b>/running</b></code> <code>(Checks if the service is healthy)</code></summary>

  #### Parameters

  > None

  #### Responses

  ##### HTTP Code 200

  ```json
  Welcome to payments api..!!
  ```
</details>




### Create a payment order

Return payment order information for process charges and refunds.
<details>
  <summary><code>POST</code> <code><b>/api/v1/payments/checkouts/{merchant_code}</b></code></summary>

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | amount   |  required | numeric      | Amount to perform a payment   |
  > | description   |  required | string      | Description of payment                                 |
  > | currency   |  required | string      | Currency to perform a payment                                |
  > | customer.dni   |  required | string      | Dni of customer                                |
  > | customer.name   |  required | string      | Name of customer                                |
  > | customer.email   |  required | string      | Email of customer                                |
  > | customer.phone   |  required | string      | Phone of customer                                |
  > | customer.address   |  optional | string      | Address of customer                                |

  ```json
  {
    "amount": 15.75,
    "description": "Sample payment",
    "currency": "USD",
    "customer": {
        "dni": "123456",
        "name": "Sample Customer",
        "email": "customer@email.com",
        "phone": "213213",
        "address": ""
    }
}
  ```

  #### Responses

  ##### HTTP Code 200

  Successful payment checkout

  ```json
  {
    "data": {
        "payment_code": "245592bc-ee36-4ff6-a919-3bc731584db4",
        "amount": 50.75,
        "description": "Sample payment",
        "currency": "USD",
        "status": "pending",
        "natural_expiration_process": "2024-03-05 13:24:17",
        "bank_name": "simulator",
        "customer": {
            "dni": "123456",
            "name": "FSample Customer",
            "email": "customer@email.com",
            "phone": "213213",
            "address": ""
        }
    },
    "status": "ok"
}
```
        
##### HTTP Code 400
  Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```
</details>



### Charge payment

Api proxy for process a charge with the bank.
<details>
  <summary><code>POST</code> <code><b>/api/v1/payments/process/{payment_code}</b></code></summary>

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | card_number   |  required | string      | Customer credit card. Simulator accept: card_success, card_insufficient_founds, card_incorrect, card_bad_request, card_server_error   |

  ```json
  {
    "card_number": "card_success"
  }
  ```

  #### Responses

  ##### HTTP Code 200

  Succesful payment charge

  ```json
  {
    "data": {
        "status": "succeeded",
        "code": 1000,
        "message": "payment processed successullfy",
        "reference": "aba14e54-3738-411b-81ed-be249ea7d2f2"
    },
    "status": "ok"
}
```
##### HTTP Code 400
Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```
</details>


### Refund payment

Api proxy for process a refund with the bank.
<details>
  <summary><code>POST</code> <code><b>/api/v1/payments/refunds/{payment_code}</b></code></summary>

  #### Headers

  > | name            |  Example                                   |
  > |-----------------|-----------|
  > | Authorization   |  Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk2NDg2OTgsInVzZXJuYW1lIjoiaHphbWJyYW5vIn0.1rWmYm1zglb-Xi00FoK1xhhxozdpbNjUhqenv-hzv94   |

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | refund_case   |  required | string      | Configurate the responses of bank simulator. Accept: refund_success, refund_already_refunded, refund_incorrect, refund_bad_request, refund_server_error   |

  ```json
  {
    "card_number": "card_success"
  }
  ```

  #### Responses

  ##### HTTP Code 200

  Succesful refunded checkout

  ```json
  {
    "data": {
        "status": "refunded",
        "code": 1000,
        "message": "refunded processed successullfy",
        "reference": "aba14e54-3738-411b-81ed-be249ea7d2f2"
    },
    "status": "ok"
}
  ```
##### HTTP Code 400
Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```

  ##### HTTP Code 400
  Missing Information.

  ```json
  {
    "status_code": "failed",
    "message": "message description"
  }
  ```

  ##### HTTP Code 409
  Payment is already refunded, pending o failed.

  ```json
  {
    "status_code": "failed",
    "message": "could not process. payment is status."
  }
  ```

  ##### HTTP Code 500

  ```json
  {
    "status_code": "failed",
    "message": "message description"
  }
  ```
</details>


### Query payment

Get payment information.
<details>
  <summary><code>GET</code> <code><b>/api/v1/payments/{payment_code}</b></code></summary>

  #### Headers

  > | name            |  type     | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | Authorization   |  required | Bearer token_sample   |

  #### Responses

  ##### HTTP Code 200

  Payment found

  ```json
  {
    "data": {
        "payment_code": "245592bc-ee36-4ff6-a919-3bc731584db4",
        "amount": 50.75,
        "description": "Sample payment description",
        "currency": "USD",
        "status": "succedeed",
        "natural_expiration_process": "2024-03-05 13:24:17",
        "failure_reason": "",
        "bank_reference": "aba14e54-3738-411b-81ed-be249ea7d2f2",
        "bank_name": "simulator"
    },
    "status": "ok"
}
  ```
##### HTTP Code 400
Something happened.

  ```json
  {
    "status": "failed",
    "message": "error detail"
  }
  ```
</details>
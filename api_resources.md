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
Welcome to merchant api..!!
```

</details>



### Sing Up

Allows to register a new merchant for payments processing.

<details>
  <summary><code>POST</code> <code><b>/api/v1/merchants/sing-up</b></code></summary>

  #### Parameters

  > | name            |  type     | data type   | description                                   |
  > |-----------------|-----------|-------------------------|-----------------------------------|
  > | name            |  required | string      | Name of merchant                              |
  > | user.username   |  required | string      | Username to be able to log in in the future   |
  > | user.password   |  required | string      | User password                                 |
  > | user.email      |  required | string      | Email user                                    |

  #### Responses

  ##### HTTP Code 200

  Succesful payment

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
  Bad request. Sent information are incorrect.

  ```json
  {
    "status": "failed",
    "message": "data sent is invalid"
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
  User alreay exits.

  ```json
  {
    "status_code": "failed",
    "message": "username sample_merchant is alreay exits"
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

  Succesful login

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
  Bad request. Sent information are incorrect.

  ```json
  {
    "status": "failed",
    "message": "data sent is invalid"
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

  ##### HTTP Code 404
  User not found.

  ```json
  {
    "status_code": "failed",
    "message": "user sample_merchant not found"
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

  Succesful payment checkout

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
  Bad request. Sent information are incorrect.

  ```json
  {
    "status": "failed",
    "message": "data sent is invalid"
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

  ##### HTTP Code 500

  ```json
  {
    "status_code": "failed",
    "message": "message description"
  }
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

  Succesful payment checkout

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
  Bad request. Sent information are incorrect.

  ```json
  {
    "status": "failed",
    "message": "data sent is invalid"
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

  ##### HTTP Code 500

  ```json
  {
    "status_code": "failed",
    "message": "message description"
  }
  ```
</details>



### Process charge payment

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

  Succesful payment checkout

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
  Bad request. Sent information are incorrect.

  ```json
  {
    "status": "failed",
    "message": "data sent is invalid"
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

  ##### HTTP Code 500

  ```json
  {
    "status_code": "failed",
    "message": "message description"
  }
  ```
</details>










### Query payment

<details>
 <summary><code>GET</code> <code><b>/{transaction_id}</b></code> <code>(Queries a payment given its transaction_id)</code></summary>

#### Parameters

> | name            |  type     | data type               | description                                              |
> |-----------------|-----------|-------------------------|----------------------------------------------------------|
> | id              |  required | string (path parameter) | Identifier to the given transaction_id                    |

#### Responses

##### HTTP Code 200

```json
{
  "transaction_id": "TXN_01HP06ZRSNFDPKN3ZBSWS4Z0KT",
  "status": "succeeded",
  "description": "Sample transaction",
  "payment_provider": "stripe",
  "amount": 2000,
  "currency": "eur",
  "type": "charge",
  "additional_fields": {
      "charge_id": "ch_3OgwgvGVGHB8I6rc1Etj264n",
      "payment_intent_id": "pi_3OgwgvGVGHB8I6rc1ZC8RNGK"
  }
}
```

##### HTTP Code 400

```json
{
  "code": "resource_not_found",
  "status_code": 404,
  "message": "Resource 'transaction' not found"
}
```

##### HTTP Code 500

```json
{
  "code": "invalid_server_error",
  "status_code": 500,
  "message": "Internal server error"
}
```

</details>

### Refund payment

**Disclaimer**: When a payment is refunded its initial status is intentionally set to `pending`. In order to mock use case where it takes X amount of time to charge a payment. Therefore, the final status will be given by the event received by webhooks.

<details>
 <summary><code>POST</code> <code><b>/{transaction_id}</b></code> <code>(Refunds a payment given its transaction_id)</code></summary>

#### Parameters

> | name            |  type     | data type               | description                                              |
> |-----------------|-----------|-------------------------|----------------------------------------------------------|
> | id              |  required | string (path parameter) | Identifier to the given transaction_id                    |

#### Responses

##### HTTP Code 200

Succesful refund 

```json
{
  "transaction_id": "TXN_01HP06ZRSNFDPKN3ZBSWS4Z0KT",
  "status": "pending",
  "description": "Sample transaction",
  "payment_provider": "stripe",
  "amount": 2000,
  "currency": "eur",
  "type": "refund",
  "additional_fields": {
      "charge_id": "ch_3OgwgvGVGHB8I6rc1Etj264n",
      "payment_intent_id": "pi_3OgwgvGVGHB8I6rc1ZC8RNGK",
      "refund_id": "re_3OgwgvGVGHB8I6rc1rBOb2uO"
  }
}
```

##### HTTP Code 400

```json
{
  "code": "invalid_request",
  "status_code": 400,
  "message": "Invalid request: charge already refunded"
}
```

##### HTTP Code 500

```json
{
  "code": "invalid_server_error",
  "status_code": 500,
  "message": "Internal server error"
}
```

</details>

## Online Payment Webhooks

### Ping

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>(Checks if the service is healthy)</code></summary>

#### Parameters

> None

#### Responses

##### HTTP Code 200

```json
Hello World!
```

</details>







### Query payment

<details>
 <summary><code>GET</code> <code><b>/{transaction_id}</b></code> <code>(Queries a payment given its transaction_id)</code></summary>

#### Parameters

> | name            |  type     | data type               | description                                              |
> |-----------------|-----------|-------------------------|----------------------------------------------------------|
> | id              |  required | string (path parameter) | Identifier to the given transaction_id                    |

#### Responses

##### HTTP Code 200

```json
{
  "transaction_id": "TXN_01HP06ZRSNFDPKN3ZBSWS4Z0KT",
  "status": "succeeded",
  "description": "Sample transaction",
  "payment_provider": "stripe",
  "amount": 2000,
  "currency": "eur",
  "type": "charge",
  "additional_fields": {
      "charge_id": "ch_3OgwgvGVGHB8I6rc1Etj264n",
      "payment_intent_id": "pi_3OgwgvGVGHB8I6rc1ZC8RNGK"
  }
}
```

##### HTTP Code 400

```json
{
  "code": "resource_not_found",
  "status_code": 404,
  "message": "Resource 'transaction' not found"
}
```

##### HTTP Code 500

```json
{
  "code": "invalid_server_error",
  "status_code": 500,
  "message": "Internal server error"
}
```

</details>

### Refund payment

**Disclaimer**: When a payment is refunded its initial status is intentionally set to `pending`. In order to mock use case where it takes X amount of time to charge a payment. Therefore, the final status will be given by the event received by webhooks.

<details>
 <summary><code>POST</code> <code><b>/{transaction_id}</b></code> <code>(Refunds a payment given its transaction_id)</code></summary>

#### Parameters

> | name            |  type     | data type               | description                                              |
> |-----------------|-----------|-------------------------|----------------------------------------------------------|
> | id              |  required | string (path parameter) | Identifier to the given transaction_id                    |

#### Responses

##### HTTP Code 200

Succesful refund 

```json
{
  "transaction_id": "TXN_01HP06ZRSNFDPKN3ZBSWS4Z0KT",
  "status": "pending",
  "description": "Sample transaction",
  "payment_provider": "stripe",
  "amount": 2000,
  "currency": "eur",
  "type": "refund",
  "additional_fields": {
      "charge_id": "ch_3OgwgvGVGHB8I6rc1Etj264n",
      "payment_intent_id": "pi_3OgwgvGVGHB8I6rc1ZC8RNGK",
      "refund_id": "re_3OgwgvGVGHB8I6rc1rBOb2uO"
  }
}
```

##### HTTP Code 400

```json
{
  "code": "invalid_request",
  "status_code": 400,
  "message": "Invalid request: charge already refunded"
}
```

##### HTTP Code 500

```json
{
  "code": "invalid_server_error",
  "status_code": 500,
  "message": "Internal server error"
}
```

</details>

## Online Payment Webhooks

### Ping

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>(Checks if the service is healthy)</code></summary>

#### Parameters

> None

#### Responses

##### HTTP Code 200

```json
Hello World!
```

</details>




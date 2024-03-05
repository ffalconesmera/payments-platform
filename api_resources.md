# API resources

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

All endpoints must have `/Content-Type` header with `application/json` value.

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
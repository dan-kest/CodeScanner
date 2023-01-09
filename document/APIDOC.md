# API Documentation

#### List repositories
<details>
<summary><code>GET</code> <code>/api/repo/list</code></summary>

##### Query Parameters

| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| page      | No       | Number | Page shown           |
| item      | No       | Number | Item count per page  |

##### Responses

`200` Success
```json
{
  "status": "OK",
  "data": {
    "page": 1,
    "item_count": 5,
    "total_count": 1,
    "item_list": [
      {
        "id": "ee592d7a-9022-11ed-92be-0242c0a81002",
        "name": "Code Scanner",
        "url": "https://github.com/dan-kest/CodeScanner.git",
        "scan_status": "Success",
        "timestamp": "2023-01-09T13:39:54Z"
      }
    ]
  }
}
```
| Name        | Type     | Description                   |
|-------------|----------|-------------------------------|
| page        | Number   | Current page                  |
| item_count  | Number   | Item shown per page           |
| total_count | Number   | Total item count (all pages)  |
| id          | String   | Repository ID                 |
| name        | String   | Repository Name               |
| url         | String   | Repository URL                |
| scan_status | String   | Latest scan status            |
| timestamp   | String   | Latest scan status timestamp  |

`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

______________________________

#### Get a repository
<details>
<summary><code>GET</code> <code>/api/repo/{{id}}</code></summary>

##### Path Parameters

| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| id        | Yes      | UUID   | Repository ID        |

##### Responses

`200` Success
```json
{
  "status": "OK",
  "data": {
    "id": "ee592d7a-9022-11ed-92be-0242c0a81002",
    "name": "Code Scanner",
    "url": "https://github.com/dan-kest/CodeScanner.git",
    "scan_status": "Success",
    "timestamp": "2023-01-09T13:39:54Z",
    "findings": [
      {
        "type": "public_key",
        "ruleId": "G402",
        "location": {
          "path": "/config/rule.yaml",
          "positions": {
            "begin": {
              "line": 5
            }
          }
        },
        "metadata": {
          "description": "Exposed sensitive information.",
          "severity": "HIGH"
        }
      }
    ]
  }
}
```
| Name        | Type     | Description                   |
|-------------|----------|-------------------------------|
| id          | String   | Repository ID                 |
| name        | String   | Repository Name               |
| url         | String   | Repository URL                |
| scan_status | String   | Latest scan status            |
| timestamp   | String   | Latest scan status timestamp  |
| findings    | Array    | Array of success scan results |

`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

______________________________

#### Scan a repository
<details>
<summary><code>POST</code> <code>/api/repo/scan</code></summary>

##### Request Body

```json
{
  "id": "ee592d7a-9022-11ed-92be-0242c0a81002"
}
```
| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| id        | Yes      | UUID   | Repository ID        |

##### Responses

`200` Success
```json
{
  "status": "OK",
}
```
`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

______________________________

#### Create a repository
<details>
<summary><code>POST</code> <code>/api/repo</code></summary>

##### Request Body

```json
{
  "name": "Code Scanner",
  "url": "https://github.com/dan-kest/CodeScanner.git"
}
```
| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| name      | Yes      | String | Repository Name      |
| url       | Yes      | String | Repository URL       |

##### Responses

`200` Success
```json
{
  "status": "OK",
  "data": "ee592d7a-9022-11ed-92be-0242c0a81002"
}
```
| Name        | Type     | Description                 |
|-------------|----------|-----------------------------|
| data        | String   | Created repository ID       |

`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

______________________________

#### Update a repository
<details>
<summary><code>PUT</code> <code>/api/repo/{{id}}</code></summary>

##### Path Parameters

| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| id        | Yes      | UUID   | Repository ID        |

##### Request Body

```json
{
  "name": "Code Scanner",
  "url": "https://github.com/dan-kest/CodeScanner.git"
}
```
| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| name      | No       | String | Repository Name      |
| url       | No       | String | Repository URL       |

##### Responses

`200` Success
```json
{
  "status": "OK",
}
```
`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

______________________________

#### Delete a repository
<details>
<summary><code>DELETE</code> <code>/api/repo/{{id}}</code></summary>

##### Path Parameters

| Name      | Required | Type   | Description          |
|-----------|----------|--------|----------------------|
| id        | Yes      | UUID   | Repository ID        |

##### Responses

`200` Success
```json
{
  "status": "OK",
}
```
`400` Invalid Request
```json
{
  "status": "ERROR",
  "message": "example invalid message"
}
```
`500` Error
```json
{
  "status": "ERROR",
  "message": "example error message"
}
```

</details>

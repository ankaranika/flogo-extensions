
# REST file - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/restfile
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/restfile
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["GET", "POST"]
    },
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "type",
      "type": "string",
      "required": true,
      "allowed": ["text", "audio/wav", "audio/raw"]
    }
  ],
  "outputs": [
    {
      "name": "status",
      "type": "bool"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| method | The HTTP method to invoke (Allowed values are GET, POST) |
| uri | The URI of the service to invoke |
| pathParams | path parameters |
| type | The type of the file to GET/POST |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| status | Status of the service |

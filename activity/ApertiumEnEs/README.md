
# ApertiumEnEs - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/ApertiumEnEs
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/ApertiumEnEs
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "ip",
      "type": "string"
    },
    {
      "name": "req_id",
      "type": "string"
    },
    {
      "name": "english",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "spanish",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| ip | The ip of the request sender |
| req_id | The request id from the specified sender |
| english | The initial text |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| spanish | The translated text |


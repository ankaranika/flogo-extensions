
# pocketsphinx - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/pocketsphinx
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/pocketsphinx
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
      "name": "speech",
      "type": "any"
    }
  ],
  "outputs": [
    {
      "name": "text",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| ip | The ip of the request sender |
| req_id | The request id of the specified sender |
| speech | The speech to be transcribed |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| text | The transcribed text |

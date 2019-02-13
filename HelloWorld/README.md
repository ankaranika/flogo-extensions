# 	HelloWorld - Activity

## 	Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/HelloWorld
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/HelloWorld
```

##	Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "salutation",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| name | The name to be saluted |
| salutation | The kind of salutation |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| result | The salutation to name |

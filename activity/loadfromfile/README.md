
# Load From File - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/loadfromfile
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/loadfromfile
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "path",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "content",
      "type": "any"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| path | The path of the file to be loaded |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| content | The file's content |

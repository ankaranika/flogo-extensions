
# eSpeak-NG - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/espeak
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/espeak
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "text",
      "type": "string"
    }
  ],
  "outputs":[
    {
      "name": "speech",
      "type": "any"
    }
  ]
}
```

## Inputs
| Input   | Description    |
|:----------|:---------------|
| text | The text to be spoken |

## Outputs
| Output | Description |
|:----------|:---------------|
| speech | The spoken text |

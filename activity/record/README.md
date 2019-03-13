
# Record - Activity

## Installation

```bash
flogo install github.com/ankaranika/flogo-extensions/activity/record
```
Link for flogo web:
```bash
https://github.com/ankaranika/flogo-extensions/activity/record
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "audiotype",
      "type": "string"
    }
  ],
  "outputs":[
    {
      "name": "recording",
      "type": "any"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| audiotype | The desired type for the audio |

## Outputs
| Output   | Description    |
|:----------|:---------------|
| recording | The recorded audio in the selected type |

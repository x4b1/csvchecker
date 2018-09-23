Csvchecker is a CSV structure and data validator written in pure Go

It can be used as a command or a package to use in your Go project.

# Installation
```
go get github.com/xabi93/csvchecker
```

# Usage
As mentioned this tool can be used as a package to use in your go project or as a command.

## Package
```go
checker := csvchecker.NewChecker(',', true)

checker.AddColum(
	csvchecker.NewColumn(
		1,
		csvchecker.NewStringValidation(
			true,
			csvchecker.NewRangeValidation(1, 5),
		),
	),
)

reader := strings.NewReader(`column1;colum2
value1;value2
;value3`)

errors := checker.Check(reader)
```

## Command
The command usage is simple just:
```
csvchecker {configuration file path} {csv to validate}
```

Configuration file must have the folling structure:

```json
{
    "separator": ";",
    "withHeader": true,
    "columns": [
        {
            "position": 1,
            "validation": {}
        }
    ]
}
```

### string validator
```json
{
    "type": "string",
    "allowEmpty": true,
    "range": {
        "min": 2,
        "max": 5
    }
}
```

### number validator
```json
{
    "type": "number",
    "allowEmpty": true,
    "range": {
        "min": 2,
        "max": 5
    }
}
```
### regexp validator
```json
{
    "type": "regexp",
    "regexp": "/+d/"
}
```

### list values validator
```json
{
    "type": "list",
    "allowEmpty": true,
    "list": [
        "test"
    ]
}
```

# Contribute
[Contributions](https://github.com/xabi93/csvchecker/issues) are more than welcome
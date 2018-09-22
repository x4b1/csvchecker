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
checker := csvchecker.NewChecker(",", true)

checker.AddColum(csvchecker.NewColumn(1, &csvchecker.StringValidation{
    AllowEmpty: true,
    InRange:    &csvchecker.RangeValidation{Min: 1, Max: 5},
}))

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
    "separator": ";", //Collumns separator
    "withHeader": true, //Boolean, if your files first line is the header
    "columns": [
        {
            "position": 1, //Position of the column to be checked
            "validation": {
                //Depending on the validator type this field will have different fields
            }
        }
    ]
}
```

### string validator
```json
{
    "type": "string",
    "allowEmpty": true, //If field in the column can be empty
    "CheckRange": {
        //if defined will be check the length of the string
        "min": 2,
        "max": 5
    }
}
```

### number validator
```json
{
    "type": "number",
    "allowEmpty": true, //If field in the column can be empty
    "CheckRange": {
        //if defined will be check the value of the number
        "min": 2,
        "max": 5
    }
}
```

### regexp validator
```json
{
    "type": "regexp",
    "Regex": "/+d/" //Regexp expression that will be applied to validate column. 
}
```

# Contribute
[Contributions](https://github.com/xabi93/csvchecker/issues) are more than welcome
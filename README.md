# Calculator
Calculator service for the Aristogram

### How to run
```shell
go run main.go --config_path="configs/prod.yaml"
```
or
```shell
CONFIG_PATH="configs/prod.yaml" go run main.go
```

### How to send request

**Port**: 12321

**Expression request rules**:
- Every operator **must be** separated by whitespace!
- "1 + 2"
- "1 - 2"
- "3 * 4"
- "4 / 5" ( 4 divided by 5 )
- "rt 2 3" ( square root of 3 )
- "rt 13 45" ( 13th root of 45 )
- "2 ^ 3" ( 2 to the power of 3 )
- "log 2 3 " ( log base 2 of 3 )

**Expression example**:
( ( 1 + 2 ) ^ 2 - rt ( 3 ^ 2 ) ( log ( 2 ^ 4 ) ( 4 * 8 ) ) ) * 8

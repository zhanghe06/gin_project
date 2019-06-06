# validator

[validator.v8](https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags)

[validator.v9](https://godoc.org/gopkg.in/go-playground/validator.v9)

这里以v8为例


## Cross-Field Validation

- eqfield
- nefield
- gtfield
- gtefield
- ltfield
- ltefield
- eqcsfield
- necsfield
- gtcsfield
- ftecsfield
- ltcsfield
- ltecsfield

说明

tag | des
--- | ---
lt | <
gt | >
eq | =
lte | <=
gte | >=

关键字处理`,` `|`

逗号"," 是validation tags的默认分隔符（与）
```
Field `validate:"excludesall=0x2C"`
```

竖线"|" 是validation tags的分隔符（或）
```
Field `validate:"excludesall=0x7C"`
```

## Validators and Tags

- structonly
- nostructlevel
- required
- omitempty
- exists
- dive
- len
- max
- min
- eq
- ne
- gt
- gte
- lt
- lte
- email
- url
- uri
- base64
- contains
- containsany
- excludes
- excludesall
- uuid
- uuid3
- uuid4
- uuid5
- ascii
- ip
- ipv4
- ipv6
- cidr
- cidrv4
- cidrv6
- mac

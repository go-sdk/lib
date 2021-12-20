# lib

![Golang](https://img.shields.io/github/workflow/status/go-sdk/lib/Golang/dev?style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/github/go-sdk/lib/dev?style=for-the-badge&token=QJ7tka53iP)
![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue?style=for-the-badge)

## Install

- Only `Go 1.17+`

```shell
go get -u github.com/go-sdk/lib
```

## Env

- `log.console.level` `string`
- `log.console.color` `bool`

- `log.file.level` `string`
- `log.file.color` `bool`
- `log.file.json` `bool`
- `log.file.path` `string`
- `log.file.max_size` `int`
- `log.file.max_age` `int`
- `log.file.max_backups` `int`
- `log.file.local_time` `bool`
- `log.file.compress` `bool`

- `seq.snowflake.epoch` `int` `13位`
- `seq.snowflake.node` `int`
- `seq.uuid.epoch` `int` `18位`

- `srv.gin.mode` `release|debug`

- `token.key` `string` `HS256`
- `token.expire` `duration`

## License

[Apache License 2.0](./LICENSE)

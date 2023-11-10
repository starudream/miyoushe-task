# Miyoushe-Task

![golang](https://img.shields.io/github/actions/workflow/status/starudream/miyoushe-task/golang.yml?style=for-the-badge&logo=github&label=golang)
![release](https://img.shields.io/github/v/release/starudream/miyoushe-task?style=for-the-badge)
![license](https://img.shields.io/github/license/starudream/miyoushe-task?style=for-the-badge)

## Config

- `global` [doc](https://github.com/starudream/go-lib/blob/v2/README.md) - [example](https://github.com/starudream/go-lib/blob/v2/app.example.yaml)

以下参数无需手动增加，可通过下方 [Account](#account) 初始化并扫码登录自动获取

```yaml
accounts:
  - phone: "手机号码，仅用作唯一标识，暂无实际作用"
    device:
      id: "设备标识，uuid，登录后建议不要修改"
      type: "手机类型，默认 2 为安卓"
      name: "手机型号，默认 Xiaomi 22011211C"
      model: "手机型号，默认 22011211C"
      version: "手机安卓版本，默认 13"
      channel: "渠道，默认 miyousheluodi"
    uid: "米游社 uid"
    gtoken: "game token"
    ctoken: "cookie token"
    mid: "米哈游 uid"
    stoken: "stoken v2"

cron:
  spec: "签到奖励执行时间，默认 0 0 8 * * * 即每天 08:00:00"
  startup: "是否启动时执行一次，默认 false"

# 打码平台配置，游戏福利签到可能需要打码，可选
dm:
    retry: "在打码失败后尝试次数，默认 1"
rrocr:
    key: ""
```

## Usage

```
> miyoushe-task -h
Usage:
  miyoushe-task [command]

Available Commands:
  account     Manage accounts
  config      Manage config
  cron        Run as cron job
  notify      Manage notify
  sign        Run sign task

Flags:
  -c, --config string   path to config file
  -h, --help            help for miyoushe-task
  -v, --version         version for miyoushe-task

Use "miyoushe-task [command] --help" for more information about a command.
```

### Account

```shell
# list accounts
miyoushe-task account list
# init account device information
miyoushe-task account init <account phone>
# login account by scan qrcode to get game token
miyoushe-task account login <account phone>
```

### SignBBS `版区签到`

```shell
# manual sign bbs for input account phone
miyoushe-task sign bbs <account phone>
```

### SignLuna `福利签到`

```shell
# manual sign luna for input account phone
miyoushe-task sign luna <account phone>
```

### Config

```shell
# write config to file
miyoushe-task config save
```

### Notify

```shell
# test notify
miyoushe-task notify send <message>
```

## Docker

```shell
mkdir miyoushe && touch miyoushe/app.yaml
docker run -it --rm -v $(pwd)/miyoushe:/miyoushe -e DEBUG=true starudream/miyoushe-task /miyoushe-task -c /miyoushe/app.yaml account init <account phone>
docker run -it --rm -v $(pwd)/miyoushe:/miyoushe -e DEBUG=true starudream/miyoushe-task /miyoushe-task -c /miyoushe/app.yaml account login <account phone>
docker run -it --rm -v $(pwd)/miyoushe:/miyoushe -e DEBUG=true starudream/miyoushe-task /miyoushe-task -c /miyoushe/app.yaml sign luna <account phone>
```

## Docker Compose

```yaml
version: "3"
services:
  miyoushe:
    image: starudream/miyoushe-task
    container_name: miyoushe
    restart: always
    command: /miyoushe-task -c /miyoushe/app.yaml cron
    volumes:
      - "./miyoushe/:/miyoushe"
    environment:
      DEBUG: "true"
      app.log.console.level: "info"
      app.log.file.enabled: "true"
      app.log.file.level: "debug"
      app.log.file.filename: "/miyoushe/app.log"
      app.cron.spec: "0 40 8 * * *"
      app.rrocr.key: "foo"
```

## Thanks

- [mihoyo-api-collect](https://github.com/UIGF-org/mihoyo-api-collect)
- [miyoushe 2.62.2 salt](https://blog.starudream.cn/2023/11/09/miyoushe-salt-2.62.2/)

## [License](./LICENSE)

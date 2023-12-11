# install act runner

## 生成config

````bash
act_runner generate-config > config.yaml
````

修改 config.yaml里面的,启用cache

````bash

cache:
  host: "宿主ip"
  port: 8080
````

````yaml
version: "3.8"

services:
  runner:
    # image: gitea/act_runner:0.2.6
    image: gitea/act_runner:nightly
    environment:
      CONFIG_FILE: /config.yaml
      GITEA_INSTANCE_URL: "https://gitea.linuxcrypt.cn"
      GITEA_RUNNER_REGISTRATION_TOKEN: "--"
      GITEA_RUNNER_NAME: "server"
      GITEA_RUNNER_LABELS: "server"
    # cache port配置
    ports:
      - 8080:8080
    volumes:
      - ./data:/data
      - ./cache:/root/.cache
      - ./config.yaml:/config.yaml
      - /var/run/docker.sock:/var/run/docker.sock
````

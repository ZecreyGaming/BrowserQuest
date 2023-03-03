# Zecrey Browser Quest Server

This is an example server project of Zecrey BrowserQuest.

## Getting Started

If you do not have docker installed, [install docker](https://dockerdocs.cn/desktop/#download-and-install).


Clone the repo and create the `config.yaml` file

```bash
  cd BrowserQuest/game/api/server/etc/ && cp config.yaml.example config.yaml
```

Modify the `config.yaml` file to configure your information, following is an example:

```yaml
  Seed: "<private_key_from_metamask>"

  NftPrefix: "Hunter Warrior"

  CollectionId: "0"

```

Then,run the development server:
```bash
  docker run -d -p 5566:5566 -v $(pwd)/config.yaml:/app/etc/config.yaml zecrey/browser-quest:0.0.1  
```

If your docker download image is stuck, please [configure Docker image source](https://mirrors.ustc.edu.cn/help/dockerhub.html#linux)

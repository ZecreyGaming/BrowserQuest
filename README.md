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

  CollectionId: "<ID of the collection you created>"

```

Each account has a collection created by default. You can query through this example curl

You can replace `amber1.zec` in example with your name for query.

Example:

```bash
   curl --location 'https://hasura.zecrey.com/v1/graphql' \
   --header 'Content-Type: application/json' \
   --data '{"query":"query MyQuery {\n  collection(where: {account: {account_name: {_eq: \"amber1.zec\"}}, l2_collection_id: {_eq: \"0\"}}) {\n    id\n  }\n}","variables":{}}'
```

Example result:

```bash
 #{"data":{"collection":[{"id":5}]}}
```

Then,run the development server:

```bash
  docker run -d -p 5566:5566 -v $(pwd)/config.yaml:/app/etc/config.yaml zecrey/browser-quest:0.0.4  
```

If your docker download image is stuck,
please [configure Docker image source](https://mirrors.ustc.edu.cn/help/dockerhub.html#linux)

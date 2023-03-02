# Zecrey_treasure_hunt

This is a sample treasure hunt game that uses zecrey's sdk in the background

## Getting Started

First, Create a database locally use this code

```bash
  docker-compose -f docker-compose.yaml up -d
```

Second,run 'Makefile' command 'make Server'

Then,run this project:
```bash
  cd ./game/api/server/
  go run ./
```
If you see the prompt "Starting server at 0.0.0.0:5566...", the operation is successful
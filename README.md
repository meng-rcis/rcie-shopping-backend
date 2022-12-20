# rcie-shopping-backend

### Command to Clear All Database Data and Docker Service

- MacOS

  ```
  rm -rf external/database/volumes && \
  docker rm db && \
  docker rm api
  ```

- Windows

  ```
  rm .\external\database\volumes; `
  docker rm db; `
  docker rm api
  ```

### Prune Docker System

    docker system prune

### Start Docker Service

    docker-compose up --build

### Software Required

- [SB Data Generator](https://soft-builder.com/sb-data-generator/)
- [Hash Password Generator Tool](https://emn178.github.io/online-tools/sha256.html)
- [Random String Generator](https://www.random.org/strings/)

### Account Password

Hash Algorithm: SHA256

Hash -> {{password}} + {{ salt }}

- Admin

```
  User: admin
  Password: AdminPassword#123
  Salt: tOvyVv6VNs
```

- Buyer (John)

```
  User: johndoe
  Password: JohnDoePassword#123
  Salt: mT8kgmRfep
```

```
  User: janecat
  Password: JaneCatPassword#123
  Salt: p5U1fGte3y
```

- Seller (Mary)

```
  User: mary001
  Password: Mary001Password#123
  Salt: xNjc22n5kY
```

```
  User: katecha
  Password: KateChaPassword#123
  Salt: ljUy3RtWer
```

- The Rest

```
  Password: TheRestPassword#123
  Salt: vA3fmg0T71
  Hash: 6a9526e49628084ffb22d6da266ec6067c6a215b5d1b916c01b473af7a4d71a8
```

to-do list
- create cicd pipeline deploying the code

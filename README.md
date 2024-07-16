# Candywatch

A youtube watchparty app

# How to use?
## Clone the repository

``` markdown
$ git clone https://github.com/itrajkov/candywatch.git
```

## Start containers

``` markdown
$ cd candywatch && docker compose up --build 
```


## Development setup
// TODO

# Roadmap
- [x] User Sessions
- [x] Rooms (create/join/leave)
- [x] Chat (within a room)
  - This sets most of the infrastructure in place for implementing the syncing protocol later.
- [x] Dockerize app
- [ ] Write tests for existing features
- [ ] Define Protocol spec
- [ ] Protocol implementation
- [ ] Youtube embed
- [ ] Wire up youtube embed with syncing mechanism.

# Candywatch

An app like [Watch2Gether](https://w2g.tv/en/) that allows for watching youtube with someone.
It uses raw WebSockets with a custom protocol for syncing within the room.

# Roadmap
- [x] User Sessions
- [x] Rooms (create/join/leave)
- [x] Chat (within a room)
  - This sets most of the infrastructure in place for implementing the syncing protocol later.
- [x] Dockerize app
- [ ] Define Protocol spec
- [ ] Protocol implementation
- [ ] Youtube embed
- [ ] Wire up youtube embed with syncing mechanism.

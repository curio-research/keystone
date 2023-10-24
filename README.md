<h1 align="center"> Keystone </h1>

Keystone is the rollup SDK for building onchain games hyper focused on performance and composability above all.

We believe that in order for onchain games as a whole to meaningfully expand, we must first and foremost deliver unparalleled experience that the average player is familiar with.

TODO: insert header image

## Current features

- High frequency tick based game logic
- State sync with support for emitting events and errors
- Atomic transactions within game logic systems

## Coming Soon

- EVM layer that composes on top of the tick-based state machine
- Better code-gen support across platforms (Unity, Typescript, etc)
- UI state explorers, simulation tools
- Parallel execution ???

## Back story

After building and launching onchain games ourselves, we realized the following problems:

- **Slow performance**: Games built using smart contract languages are slow by nature, as blockchain state machines aren’t specialized for games. We cannot build ambitious games with even hundreds of concurrent users.
- **Missing critical game features:** Since blockchains are async and transparent by nature, missing games like game tick and private information block meaningful games from thriving.
- **Isolated developer toolchains**: Smart contract languages are unable to leverage existing mature toolchains and ecosystems for games.

We pride ourselves as being unapologetic, practical idealists. We strive to combine the decades of game engineering knowledge with cutting edge blockchain research to further the space.

Keystone represents a step-function improvement in how onchain games are built. We designed keystone from grounds up to create a high-performance, data oriented game server with the composability of EVM smart contracts (coming soon).

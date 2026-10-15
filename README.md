# Zombie Army 4: Dead War - Nextendo server

Initial Go NEX authentication, secure connection, matchmaking and NAT traversal service. Gameplay remains peer-to-peer. End-to-end campaign co-op is not yet verified.

## Verified client identity

- Switch title: `01000BF0152FA000`, update **1.1.1** (title version 131072).
- Main build: `12024D08CCFD25EB068E55366BD4C544A6A6699D`.
- Game server ID: `20532500`; SNI: `g20532500-lp1.s.n.srv.nintendo.net` (confirmed in Citron log).
- Access key: `dd3825ac` at main flat offset `0x1444565`; referenced by login setup at `0x31E1BC`, which passes server ID `0x20532500`.
- Embedded SDK: NEX 4.6.6, PiaNex 5.38.0.

## Build and deployment

Uses the sibling `nextendo-nex-main` checkout, matching the existing Nextendo services. Run `go build -o zombiearmy4 .` for the target OS. The deployment script installs a new isolated container with SNI passthrough on 443, secure TCP port 60030, and token-protected dashboard on loopback 8117. Existing services are not restarted. Secrets are copied server-side from the existing Nextendo deployment, never placed in this repository.

The initial implementation uses the common modern Switch Pia Register response. Actual client traffic must confirm registration, session creation, search/join, and peer connectivity. Unimplemented game-specific calls must be implemented from their wire contract, not acknowledged with fabricated success payloads.

## First live connection

Confirmed on production: signed Nextendo login, secure Register, NAT endpoint replacement, AutoMatchmakeWithParam, CreateMatchmakeSessionWithParam and OpenParticipation. User reached the multiplayer hosting lobby. Ranking2 PutCommonData (0x7A.3) was initially unregistered; the shared Ranking2 store handler is now wired in. A second player joining and starting a campaign remains unverified.

## Disclaimer

This server ships **no** Nintendo code, keys, or copyrighted assets. It is an independent
reimplementation for use with a community-run replacement service, not affiliated with, endorsed by,
or associated with Nintendo or the game's publisher.

## License

Released under the **[PolyForm Shield License 1.0.0](LICENSE.md)**, source-available: read, use,
modify, and self-host, but do not use it to provide a product that competes with Nextendo Network.

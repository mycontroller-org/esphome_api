## Changes by Version

### v1.4.0 (2026-07-18)

Proto baseline jumped from ESPHome **2024.5.0** (library snapshot 2024-06-15) to
[**2026.7.0**](https://github.com/esphome/esphome/releases/tag/2026.7.0).
Full migration notes: [docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md](docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md).

**Contains BREAKING CHANGES** (see below and the doc above).

#### Protocol and generated API
* Sync `api.proto` / `api_options.proto` with ESPHome **2026.7.0**
* Regenerate `pkg/api` protobuf bindings; expand `helper.go` to all **148** message type IDs
  (was 111; highest ID 114 → 148)
* New entities/features on the wire: siren, update, water heater, infrared, serial proxy,
  radio frequency, Z-Wave proxy, noise key provisioning, HA action / service responses,
  Bluetooth scanner mode and connection params, voice assistant extensions

#### Client (`pkg/client`, `pkg/types`)
* Advertise Hello API version **1.14**
* Session setup: use **`Hello()`** on modern devices (ESPHome 2026.1+)
* **`Login(password)`** kept for older firmware only; password auth was **removed** in
  ESPHome 2026.1.0. Prefer Noise encryption via `encryptionKey` on `GetClient`
* `Login` treats auth-response timeout as success after Hello (modern devices ignore password)
* Answer device-originated `GetTimeRequest` (timezone-aware epoch seconds)
* Record `DisconnectReason` from device `DisconnectRequest`
* Safer shutdown (`close`-once + deadline) and register waiters before send
* `DeviceInfo` gains friendly name, manufacturer, project, BT MAC, encryption flags,
  areas/devices, Z-Wave, serial proxies, and related fields
* `GetLogEntry`: decode log `message` from `bytes`; `SendFailed` always false (field removed)

#### CLI (`esphomectl`) and examples
* Always complete session setup: `Hello()`, or legacy `Login(password)` when a password is set
* Warn on stderr when API password is used (deprecated; removed in ESPHome 2026.1.0)
* `--password` help/examples mark password auth as removed; prefer `--encryption-key`
* Persist expanded device info after login; entity list UI no longer uses removed `unique_id`
* Examples follow the same Hello / Login session pattern

#### Build and docs
* Release builds use `-trimpath` and slim ldflags (`-s -w`); README documents the same for local CLI builds
* Migration guide: `docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md`
* README / `proto/README.md` updated for 2026.7.0 and encryption-first setup

#### Breaking changes
* `api.ConnectRequest` / `api.ConnectResponse` → `api.AuthenticationRequest` / `api.AuthenticationResponse`
  (IDs 3/4; not processed by ESPHome 2026.1.0+)
* `api.HomeassistantServiceResponse` → `api.HomeassistantActionRequest`
* Entity list messages: `unique_id` removed (use `object_id` + `key`, optional `device_id`)
* `SubscribeLogsResponse.message`: `string` → `bytes`; `send_failed` removed
* `BluetoothLEAdvertisementResponse.name`: `string` → `bytes`
* Bluetooth helper constants use `*TypeID` suffix (aligned with other message IDs)
* Callers must run `Hello()` (or deprecated `Login`) after `GetClient` for a full session

### v1.3.0 (2023-01-06)
* update proto: bluetooth support included ([#7](https://github.com/mycontroller-org/esphome_api/pull/7), [@jkandasa](https://github.com/jkandasa))
* support encrypted connection ([#5](https://github.com/mycontroller-org/esphome_api/pull/5), [@jkandasa](https://github.com/jkandasa))

**Contains BREAKING CHANGES**
* get new client function name changed and encryption key argument added on the new function
```
old: func Init(clientID, address string, timeout time.Duration, handlerFunc func(proto.Message)) (*Client, error) 
new: GetClient(clientID, address, encryptionKey string, timeout time.Duration, handlerFunc func(proto.Message)) (*Client, error) 
```

### v1.2.0 (2022-07-06)
* rename pkg model to types, upgrade go version ([#2](https://github.com/mycontroller-org/esphome_api/pull/2), [@jkandasa](https://github.com/jkandasa)) **Contains BREAKING CHANGES**
### v1.1.0 (2022-06-13)
* Updated api.proto and included new message definitions ([#1](https://github.com/mycontroller-org/esphome_api/pull/1), [@mligor](https://github.com/mligor))

  
### v1.0.0 (2022-06-13)
* initial release

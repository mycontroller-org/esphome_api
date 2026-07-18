# ESPHome API Proto Update: 2024.5.0 → 2026.7.0

## Version jump

|                               | From                                                                     | To                                                                       |
| ----------------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| ESPHome release               | [**2024.5.0**](https://github.com/esphome/esphome/releases/tag/2024.5.0) | [**2026.7.0**](https://github.com/esphome/esphome/releases/tag/2026.7.0) |
| Library baseline              | Proto last synced 2024-06-15 (wire-compatible with 2024.5.0–2024.5.5)    | This update                                                              |
| Message types (`option (id)`) | 111                                                                      | 148                                                                      |
| Highest message ID            | 114                                                                      | 148                                                                      |

**From (previous library protos)**

- Snapshot committed 2024-06-15 from ESPHome `dev` at that time
- Message count and max ID match release **2024.5.0** through **2024.5.5**
- Still used `ConnectRequest` / `ConnectResponse` password auth
- Entity list messages still carried `unique_id`

**To (this update)**

- [`api.proto`](https://github.com/esphome/esphome/blob/2026.7.0/esphome/components/api/api.proto)
- [`api_options.proto`](https://github.com/esphome/esphome/blob/2026.7.0/esphome/components/api/api_options.proto)
- Pinned to ESPHome release **2026.7.0** (not `dev`)

## Summary

| Item                             | Before (2024.5.0)               | After (2026.7.0)                             |
| -------------------------------- | ------------------------------- | -------------------------------------------- |
| Proto source                     | ESPHome **2024.5.x** equivalent | ESPHome **2026.7.0**                         |
| Message types with `option (id)` | 111                             | 148                                          |
| Highest message ID               | 114                             | 148                                          |
| Password authentication          | Supported via `ConnectRequest`  | **Deprecated / removed** (ESPHome 2026.1.0+) |
| Entity `unique_id` field         | Present on list-entity messages | **Removed** (field 4 reserved)               |
| Log message payload              | `string message`                | `bytes message`                              |

Regenerated artifacts:

- `proto/api.proto`, `proto/api_options.proto` - copied from ESPHome release 2026.7.0
- `pkg/api/api.pb.go`, `pkg/api/api_options.pb.go` - `protoc-gen-go` output
- `pkg/api/helper.go` - complete type-id maps for all 148 message IDs

---

## Breaking changes

### 1. Password authentication removed (ESPHome 2026.1.0)

**What changed in the protocol**

- `ConnectRequest` / `ConnectResponse` were renamed to
  `AuthenticationRequest` / `AuthenticationResponse` (IDs **3** and **4**).
- Those messages are **deprecated** and **not processed** by ESPHome 2026.1.0+.
- Connection authentication is completed after **Hello** alone.
- Message IDs 3 and 4 are reserved and must not be reused.

**Impact on this library**

| Old API                     | New API                            |
| --------------------------- | ---------------------------------- |
| `api.ConnectRequest`        | `api.AuthenticationRequest`        |
| `api.ConnectResponse`       | `api.AuthenticationResponse`       |
| `api.ConnectRequestTypeID`  | `api.AuthenticationRequestTypeID`  |
| `api.ConnectResponseTypeID` | `api.AuthenticationResponseTypeID` |

Client API for session setup:

| Method                   | Use for                              | Behavior                                                                                                                                                                                |
| ------------------------ | ------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `client.Hello()`         | Modern devices (ESPHome 2026.1+)     | Sends `HelloRequest` (API version 1.14) and waits for `HelloResponse`.                                                                                                                  |
| `client.Login(password)` | **Deprecated** - older firmware only | Hello, then `AuthenticationRequest`; waits for `AuthenticationResponse`. Invalid password → `types.ErrPassword`. Timeout (new firmware ignores auth) is treated as success after Hello. |

**Migration**

- Call **`Hello()`** for current devices.
- Prefer **Noise encryption** (`encryptionKey` on `GetClient`) instead of API password.
- Use **`Login(password)` only** if the device still runs pre-2026.1 firmware with `api: password:`.
- Devices that only had `api: password:` should be reconfigured to
  `api: encryption: key: ...` (see [ESPHome API docs](https://esphome.io/components/api/)).

`DeviceInfo.UsesPassword` remains populated when the device sends it, but the field is
deprecated on the wire.

---

### 2. `HomeassistantServiceResponse` renamed

| Old                            | New                          | ID  |
| ------------------------------ | ---------------------------- | --- |
| `HomeassistantServiceResponse` | `HomeassistantActionRequest` | 35  |

New related message:

| Message                       | ID  | Direction                       |
| ----------------------------- | --- | ------------------------------- |
| `HomeassistantActionResponse` | 130 | Client → device (action result) |

Any type switch / assert on `*api.HomeassistantServiceResponse` must use
`*api.HomeassistantActionRequest`.

---

### 3. Entity `unique_id` removed

Field **4** (`string unique_id`) is **reserved** on nearly all
`ListEntities*Response` messages. Use:

- `object_id` + `key` for entity identity
- optional `device_id` for multi-device setups

Affected list-entity messages include binary sensor, cover, fan, light, sensor,
switch, text sensor, camera, climate, number, select, lock, button, media player,
alarm control panel, text, date, time, event, valve, datetime, and new entity types.

**Migration:** stop reading `UniqueId` / `unique_id` from list responses.

---

### 4. Log messages are bytes

`SubscribeLogsResponse`:

| Field                   | Old      | New         |
| ----------------------- | -------- | ----------- |
| `message` (field 3)     | `string` | `bytes`     |
| `send_failed` (field 4) | `bool`   | **removed** |

`types.GetLogEntry` converts `[]byte` to `string`. `LogEntry.SendFailed` is kept
as always-`false` for older callers.

---

### 5. Bluetooth LE advertisement name is bytes

`BluetoothLEAdvertisementResponse.name` (field 2): `string` → `bytes`.

---

### 6. Bluetooth type-id constant suffix normalized

Historical helper constants used an `ID` suffix (e.g. `BluetoothDeviceRequestID`).
All message type constants now use the `TypeID` suffix:

```
SubscribeBluetoothLEAdvertisementsRequestID  →  SubscribeBluetoothLEAdvertisementsRequestTypeID
BluetoothLEAdvertisementResponseID           →  BluetoothLEAdvertisementResponseTypeID
… (all former *ID bluetooth constants)
```

---

### 7. Incomplete / incorrect type maps fixed

`pkg/api/helper.go` previously:

- Stopped around Bluetooth GATT notify (ID 84) and omitted many existing IDs (85-114).
- Mapped `NumberCommandRequest` to the wrong type ID (`ClimateCommandRequestTypeID`).

It now maps **all** message IDs from the current proto.

Unknown type IDs on the wire still produce a protocol error from the connection reader.

---

## Non-breaking / additive changes

### New message types (IDs)

| ID  | Message                                | Notes                    |
| --- | -------------------------------------- | ------------------------ |
| 55  | `ListEntitiesSirenResponse`            | Siren entity             |
| 56  | `SirenStateResponse`                   |                          |
| 57  | `SirenCommandRequest`                  |                          |
| 115 | `VoiceAssistantTimerEventResponse`     | Voice assistant          |
| 116 | `ListEntitiesUpdateResponse`           | Update entity            |
| 117 | `UpdateStateResponse`                  |                          |
| 118 | `UpdateCommandRequest`                 |                          |
| 119 | `VoiceAssistantAnnounceRequest`        |                          |
| 120 | `VoiceAssistantAnnounceFinished`       |                          |
| 121 | `VoiceAssistantConfigurationRequest`   |                          |
| 122 | `VoiceAssistantConfigurationResponse`  |                          |
| 123 | `VoiceAssistantSetConfiguration`       |                          |
| 124 | `NoiseEncryptionSetKeyRequest`         | Provision encryption key |
| 125 | `NoiseEncryptionSetKeyResponse`        |                          |
| 126 | `BluetoothScannerStateResponse`        |                          |
| 127 | `BluetoothScannerSetModeRequest`       |                          |
| 128 | `ZWaveProxyFrame`                      | Z-Wave proxy             |
| 129 | `ZWaveProxyRequest`                    |                          |
| 130 | `HomeassistantActionResponse`          | HA action result         |
| 131 | `ExecuteServiceResponse`               | User service response    |
| 132 | `ListEntitiesWaterHeaterResponse`      | Water heater             |
| 133 | `WaterHeaterStateResponse`             |                          |
| 134 | `WaterHeaterCommandRequest`            |                          |
| 135 | `ListEntitiesInfraredResponse`         | Infrared                 |
| 136 | `InfraredRFTransmitRawTimingsRequest`  |                          |
| 137 | `InfraredRFReceiveEvent`               |                          |
| 138 | `SerialProxyConfigureRequest`          | Serial proxy             |
| 139 | `SerialProxyDataReceived`              |                          |
| 140 | `SerialProxyWriteRequest`              |                          |
| 141 | `SerialProxySetModemPinsRequest`       |                          |
| 142 | `SerialProxyGetModemPinsRequest`       |                          |
| 143 | `SerialProxyGetModemPinsResponse`      |                          |
| 144 | `SerialProxyRequest`                   |                          |
| 145 | `BluetoothSetConnectionParamsRequest`  |                          |
| 146 | `BluetoothSetConnectionParamsResponse` |                          |
| 147 | `SerialProxyRequestResponse`           |                          |
| 148 | `ListEntitiesRadioFrequencyResponse`   | RF entity                |

These are available as generated Go types under `pkg/api` and are registered in
`TypeID` / `NewMessageByTypeID`.

### New / extended fields on existing messages

**`DeviceInfoResponse`**

- `bluetooth_mac_address`, `api_encryption_supported`, `api_encryption_provisionable`
- `devices` (`DeviceInfo` sub-devices), `areas`, `area`
- `zwave_proxy_feature_flags`, `zwave_home_id`
- `serial_proxies`
- Existing fields such as `friendly_name`, `manufacturer`, `project_*` remain

**`DisconnectRequest`**

- Optional `reason` (`DisconnectReason` enum), e.g. provisioning window closed

**Most entity state/command/list messages**

- Optional `device_id` for multi-device configurations

**`GetTimeResponse`**

- `timezone`, `parsed_timezone`

**`ExecuteServiceRequest`**

- `call_id`, `return_response` (pairs with `ExecuteServiceResponse` ID 131)

**`ListEntitiesServicesResponse`**

- `supports_response`

**`ListEntitiesClimateResponse`**

- `feature_flags`, `temperature_unit`

**`ListEntitiesMediaPlayerResponse`**

- `supported_formats`, `feature_flags`

**`BluetoothConnectionsFreeResponse`**

- `allocated`

**`SubscribeHomeAssistantStateResponse`**

- `once`

### Client behavior improvements

- `HelloRequest` now advertises API version **1.14** (`ClientAPIVersionMajor/Minor`).
- Client answers device-originated **`GetTimeRequest`** with current epoch seconds.
- `SendAndWaitForResponse` registers the waiter **before** sending (avoids response race).
- CLI always completes session setup (`Hello()`, or deprecated `Login(password)` when configured).
- `client.NoiseEncryptionSetKey` helper for key provisioning (IDs 124/125).
- `types.DeviceInfo` exposes additional device metadata fields.

### `api_options.proto`

Extended with ESPHome code-generation options (`base_class`, field-level
`container_pointer`, `fixed_vector`, `force`, `mac_address`, etc.). These affect
ESPHome's C++ generator; they do not change wire layout for standard protobuf
fields used by this Go client.

---

## Library surface map

| Package          | Role                                                  |
| ---------------- | ----------------------------------------------------- |
| `proto/`         | Source `.proto` files (from ESPHome)                  |
| `pkg/api`        | Generated messages + type-id helpers                  |
| `pkg/connection` | Plaintext / Noise framing                             |
| `pkg/client`     | High-level client (`Hello`, `Login`, `DeviceInfo`, …) |
| `pkg/types`      | Friendly wrappers (`DeviceInfo`, `LogEntry`, errors)  |
| `cli/`           | `esphomectl`                                          |

---

## Compatibility matrix

| ESPHome device                    | Encryption | Password | Expected client usage                     |
| --------------------------------- | ---------- | -------- | ----------------------------------------- |
| < 2026.1, no encryption           | no         | optional | `GetClient(..., "")` + `Login(password)`  |
| < 2026.1, encrypted               | yes        | optional | `GetClient(..., key)` + `Login(password)` |
| ≥ 2026.1, encrypted (recommended) | yes        | n/a      | `GetClient(..., key)` + `Hello()`         |
| ≥ 2026.1, plaintext, no password  | no         | n/a      | `GetClient(..., "")` + `Hello()`          |

---

## How to regenerate after a future ESPHome proto change

```bash
# 1. Refresh protos from a specific ESPHome release tag (example: 2026.7.0)
ESPHOME_TAG=2026.7.0
curl -sL "https://raw.githubusercontent.com/esphome/esphome/${ESPHOME_TAG}/esphome/components/api/api.proto" \
  -o proto/api.proto
curl -sL "https://raw.githubusercontent.com/esphome/esphome/${ESPHOME_TAG}/esphome/components/api/api_options.proto" \
  -o proto/api_options.proto

# 2. Generate Go bindings
cd proto && ./generate.sh

# 3. Refresh pkg/api/helper.go type-id maps (script or manual) to match option (id) values
# 4. go build ./...
# 5. Add docs/API_PROTO_UPDATE_<previous>_to_<new>.md and update CHANGELOG.md
```

---

## Files touched in this update

- `proto/api.proto`, `proto/api_options.proto`, `proto/generate.sh`
- `pkg/api/api.pb.go`, `pkg/api/api_options.pb.go`, `pkg/api/helper.go`
- `pkg/client/client.go`
- `pkg/types/types.go`
- `cli/command/root/cmd.go`, `cli/command/root/login.go`, `cli/types/config.go`
- `cli/command/device/get_entities.go`
- `CHANGELOG.md`, `README.md`, `docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md`

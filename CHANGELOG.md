## Changes by Version

### 1.3.0 (2023-01-06)
* update proto: bluetooth support included ([#7](https://github.com/mycontroller-org/esphome_api/pull/7), [@jkandasa](https://github.com/jkandasa))
* support encrypted connection ([#5](https://github.com/mycontroller-org/esphome_api/pull/5), [@jkandasa](https://github.com/jkandasa))

**Contains BREAKING CHANGES**
* get new client function name changed and encryption key argument added on the new function
```
old: func Init(clientID, address string, timeout time.Duration, handlerFunc func(proto.Message)) (*Client, error) 
new: GetClient(clientID, address, encryptionKey string, timeout time.Duration, handlerFunc func(proto.Message)) (*Client, error) 
```

### 1.2.0 (2022-07-06)
* rename pkg model to types, upgrade go version ([#2](https://github.com/mycontroller-org/esphome_api/pull/2), [@jkandasa](https://github.com/jkandasa)) **Contains BREAKING CHANGES**
### 1.1.0 (2022-06-13)
* Updated api.proto and included new message definitions ([#1](https://github.com/mycontroller-org/esphome_api/pull/1), [@mligor](https://github.com/mligor))

  
### 1.0.0 (2022-06-13)
* initial release

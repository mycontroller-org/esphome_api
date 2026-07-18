package api

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

// ESPHome API message type IDs (api.proto option (id), release 2026.7.0).
const (
	UndefinedTypeID                                   = 0
	HelloRequestTypeID                                = 1
	HelloResponseTypeID                               = 2
	AuthenticationRequestTypeID                       = 3
	AuthenticationResponseTypeID                      = 4
	DisconnectRequestTypeID                           = 5
	DisconnectResponseTypeID                          = 6
	PingRequestTypeID                                 = 7
	PingResponseTypeID                                = 8
	DeviceInfoRequestTypeID                           = 9
	DeviceInfoResponseTypeID                          = 10
	ListEntitiesRequestTypeID                         = 11
	ListEntitiesBinarySensorResponseTypeID            = 12
	ListEntitiesCoverResponseTypeID                   = 13
	ListEntitiesFanResponseTypeID                     = 14
	ListEntitiesLightResponseTypeID                   = 15
	ListEntitiesSensorResponseTypeID                  = 16
	ListEntitiesSwitchResponseTypeID                  = 17
	ListEntitiesTextSensorResponseTypeID              = 18
	ListEntitiesDoneResponseTypeID                    = 19
	SubscribeStatesRequestTypeID                      = 20
	BinarySensorStateResponseTypeID                   = 21
	CoverStateResponseTypeID                          = 22
	FanStateResponseTypeID                            = 23
	LightStateResponseTypeID                          = 24
	SensorStateResponseTypeID                         = 25
	SwitchStateResponseTypeID                         = 26
	TextSensorStateResponseTypeID                     = 27
	SubscribeLogsRequestTypeID                        = 28
	SubscribeLogsResponseTypeID                       = 29
	CoverCommandRequestTypeID                         = 30
	FanCommandRequestTypeID                           = 31
	LightCommandRequestTypeID                         = 32
	SwitchCommandRequestTypeID                        = 33
	SubscribeHomeassistantServicesRequestTypeID       = 34
	HomeassistantActionRequestTypeID                  = 35
	GetTimeRequestTypeID                              = 36
	GetTimeResponseTypeID                             = 37
	SubscribeHomeAssistantStatesRequestTypeID         = 38
	SubscribeHomeAssistantStateResponseTypeID         = 39
	HomeAssistantStateResponseTypeID                  = 40
	ListEntitiesServicesResponseTypeID                = 41
	ExecuteServiceRequestTypeID                       = 42
	ListEntitiesCameraResponseTypeID                  = 43
	CameraImageResponseTypeID                         = 44
	CameraImageRequestTypeID                          = 45
	ListEntitiesClimateResponseTypeID                 = 46
	ClimateStateResponseTypeID                        = 47
	ClimateCommandRequestTypeID                       = 48
	ListEntitiesNumberResponseTypeID                  = 49
	NumberStateResponseTypeID                         = 50
	NumberCommandRequestTypeID                        = 51
	ListEntitiesSelectResponseTypeID                  = 52
	SelectStateResponseTypeID                         = 53
	SelectCommandRequestTypeID                        = 54
	ListEntitiesSirenResponseTypeID                   = 55
	SirenStateResponseTypeID                          = 56
	SirenCommandRequestTypeID                         = 57
	ListEntitiesLockResponseTypeID                    = 58
	LockStateResponseTypeID                           = 59
	LockCommandRequestTypeID                          = 60
	ListEntitiesButtonResponseTypeID                  = 61
	ButtonCommandRequestTypeID                        = 62
	ListEntitiesMediaPlayerResponseTypeID             = 63
	MediaPlayerStateResponseTypeID                    = 64
	MediaPlayerCommandRequestTypeID                   = 65
	SubscribeBluetoothLEAdvertisementsRequestTypeID   = 66
	BluetoothLEAdvertisementResponseTypeID            = 67
	BluetoothDeviceRequestTypeID                      = 68
	BluetoothDeviceConnectionResponseTypeID           = 69
	BluetoothGATTGetServicesRequestTypeID             = 70
	BluetoothGATTGetServicesResponseTypeID            = 71
	BluetoothGATTGetServicesDoneResponseTypeID        = 72
	BluetoothGATTReadRequestTypeID                    = 73
	BluetoothGATTReadResponseTypeID                   = 74
	BluetoothGATTWriteRequestTypeID                   = 75
	BluetoothGATTReadDescriptorRequestTypeID          = 76
	BluetoothGATTWriteDescriptorRequestTypeID         = 77
	BluetoothGATTNotifyRequestTypeID                  = 78
	BluetoothGATTNotifyDataResponseTypeID             = 79
	SubscribeBluetoothConnectionsFreeRequestTypeID    = 80
	BluetoothConnectionsFreeResponseTypeID            = 81
	BluetoothGATTErrorResponseTypeID                  = 82
	BluetoothGATTWriteResponseTypeID                  = 83
	BluetoothGATTNotifyResponseTypeID                 = 84
	BluetoothDevicePairingResponseTypeID              = 85
	BluetoothDeviceUnpairingResponseTypeID            = 86
	UnsubscribeBluetoothLEAdvertisementsRequestTypeID = 87
	BluetoothDeviceClearCacheResponseTypeID           = 88
	SubscribeVoiceAssistantRequestTypeID              = 89
	VoiceAssistantRequestTypeID                       = 90
	VoiceAssistantResponseTypeID                      = 91
	VoiceAssistantEventResponseTypeID                 = 92
	BluetoothLERawAdvertisementsResponseTypeID        = 93
	ListEntitiesAlarmControlPanelResponseTypeID       = 94
	AlarmControlPanelStateResponseTypeID              = 95
	AlarmControlPanelCommandRequestTypeID             = 96
	ListEntitiesTextResponseTypeID                    = 97
	TextStateResponseTypeID                           = 98
	TextCommandRequestTypeID                          = 99
	ListEntitiesDateResponseTypeID                    = 100
	DateStateResponseTypeID                           = 101
	DateCommandRequestTypeID                          = 102
	ListEntitiesTimeResponseTypeID                    = 103
	TimeStateResponseTypeID                           = 104
	TimeCommandRequestTypeID                          = 105
	VoiceAssistantAudioTypeID                         = 106
	ListEntitiesEventResponseTypeID                   = 107
	EventResponseTypeID                               = 108
	ListEntitiesValveResponseTypeID                   = 109
	ValveStateResponseTypeID                          = 110
	ValveCommandRequestTypeID                         = 111
	ListEntitiesDateTimeResponseTypeID                = 112
	DateTimeStateResponseTypeID                       = 113
	DateTimeCommandRequestTypeID                      = 114
	VoiceAssistantTimerEventResponseTypeID            = 115
	ListEntitiesUpdateResponseTypeID                  = 116
	UpdateStateResponseTypeID                         = 117
	UpdateCommandRequestTypeID                        = 118
	VoiceAssistantAnnounceRequestTypeID               = 119
	VoiceAssistantAnnounceFinishedTypeID              = 120
	VoiceAssistantConfigurationRequestTypeID          = 121
	VoiceAssistantConfigurationResponseTypeID         = 122
	VoiceAssistantSetConfigurationTypeID              = 123
	NoiseEncryptionSetKeyRequestTypeID                = 124
	NoiseEncryptionSetKeyResponseTypeID               = 125
	BluetoothScannerStateResponseTypeID               = 126
	BluetoothScannerSetModeRequestTypeID              = 127
	ZWaveProxyFrameTypeID                             = 128
	ZWaveProxyRequestTypeID                           = 129
	HomeassistantActionResponseTypeID                 = 130
	ExecuteServiceResponseTypeID                      = 131
	ListEntitiesWaterHeaterResponseTypeID             = 132
	WaterHeaterStateResponseTypeID                    = 133
	WaterHeaterCommandRequestTypeID                   = 134
	ListEntitiesInfraredResponseTypeID                = 135
	InfraredRFTransmitRawTimingsRequestTypeID         = 136
	InfraredRFReceiveEventTypeID                      = 137
	SerialProxyConfigureRequestTypeID                 = 138
	SerialProxyDataReceivedTypeID                     = 139
	SerialProxyWriteRequestTypeID                     = 140
	SerialProxySetModemPinsRequestTypeID              = 141
	SerialProxyGetModemPinsRequestTypeID              = 142
	SerialProxyGetModemPinsResponseTypeID             = 143
	SerialProxyRequestTypeID                          = 144
	BluetoothSetConnectionParamsRequestTypeID         = 145
	BluetoothSetConnectionParamsResponseTypeID        = 146
	SerialProxyRequestResponseTypeID                  = 147
	ListEntitiesRadioFrequencyResponseTypeID          = 148
)

func TypeID(message interface{}) uint64 {
	if message == nil {
		return UndefinedTypeID
	}

	// convert from pointer to normal type
	if reflect.ValueOf(message).Kind() == reflect.Pointer {
		message = reflect.ValueOf(message).Elem().Interface()
	}
	switch message.(type) {
	case HelloRequest:
		return HelloRequestTypeID

	case HelloResponse:
		return HelloResponseTypeID

	case AuthenticationRequest:
		return AuthenticationRequestTypeID

	case AuthenticationResponse:
		return AuthenticationResponseTypeID

	case DisconnectRequest:
		return DisconnectRequestTypeID

	case DisconnectResponse:
		return DisconnectResponseTypeID

	case PingRequest:
		return PingRequestTypeID

	case PingResponse:
		return PingResponseTypeID

	case DeviceInfoRequest:
		return DeviceInfoRequestTypeID

	case DeviceInfoResponse:
		return DeviceInfoResponseTypeID

	case ListEntitiesRequest:
		return ListEntitiesRequestTypeID

	case ListEntitiesBinarySensorResponse:
		return ListEntitiesBinarySensorResponseTypeID

	case ListEntitiesCoverResponse:
		return ListEntitiesCoverResponseTypeID

	case ListEntitiesFanResponse:
		return ListEntitiesFanResponseTypeID

	case ListEntitiesLightResponse:
		return ListEntitiesLightResponseTypeID

	case ListEntitiesSensorResponse:
		return ListEntitiesSensorResponseTypeID

	case ListEntitiesSwitchResponse:
		return ListEntitiesSwitchResponseTypeID

	case ListEntitiesTextSensorResponse:
		return ListEntitiesTextSensorResponseTypeID

	case ListEntitiesDoneResponse:
		return ListEntitiesDoneResponseTypeID

	case SubscribeStatesRequest:
		return SubscribeStatesRequestTypeID

	case BinarySensorStateResponse:
		return BinarySensorStateResponseTypeID

	case CoverStateResponse:
		return CoverStateResponseTypeID

	case FanStateResponse:
		return FanStateResponseTypeID

	case LightStateResponse:
		return LightStateResponseTypeID

	case SensorStateResponse:
		return SensorStateResponseTypeID

	case SwitchStateResponse:
		return SwitchStateResponseTypeID

	case TextSensorStateResponse:
		return TextSensorStateResponseTypeID

	case SubscribeLogsRequest:
		return SubscribeLogsRequestTypeID

	case SubscribeLogsResponse:
		return SubscribeLogsResponseTypeID

	case CoverCommandRequest:
		return CoverCommandRequestTypeID

	case FanCommandRequest:
		return FanCommandRequestTypeID

	case LightCommandRequest:
		return LightCommandRequestTypeID

	case SwitchCommandRequest:
		return SwitchCommandRequestTypeID

	case SubscribeHomeassistantServicesRequest:
		return SubscribeHomeassistantServicesRequestTypeID

	case HomeassistantActionRequest:
		return HomeassistantActionRequestTypeID

	case GetTimeRequest:
		return GetTimeRequestTypeID

	case GetTimeResponse:
		return GetTimeResponseTypeID

	case SubscribeHomeAssistantStatesRequest:
		return SubscribeHomeAssistantStatesRequestTypeID

	case SubscribeHomeAssistantStateResponse:
		return SubscribeHomeAssistantStateResponseTypeID

	case HomeAssistantStateResponse:
		return HomeAssistantStateResponseTypeID

	case ListEntitiesServicesResponse:
		return ListEntitiesServicesResponseTypeID

	case ExecuteServiceRequest:
		return ExecuteServiceRequestTypeID

	case ListEntitiesCameraResponse:
		return ListEntitiesCameraResponseTypeID

	case CameraImageResponse:
		return CameraImageResponseTypeID

	case CameraImageRequest:
		return CameraImageRequestTypeID

	case ListEntitiesClimateResponse:
		return ListEntitiesClimateResponseTypeID

	case ClimateStateResponse:
		return ClimateStateResponseTypeID

	case ClimateCommandRequest:
		return ClimateCommandRequestTypeID

	case ListEntitiesNumberResponse:
		return ListEntitiesNumberResponseTypeID

	case NumberStateResponse:
		return NumberStateResponseTypeID

	case NumberCommandRequest:
		return NumberCommandRequestTypeID

	case ListEntitiesSelectResponse:
		return ListEntitiesSelectResponseTypeID

	case SelectStateResponse:
		return SelectStateResponseTypeID

	case SelectCommandRequest:
		return SelectCommandRequestTypeID

	case ListEntitiesSirenResponse:
		return ListEntitiesSirenResponseTypeID

	case SirenStateResponse:
		return SirenStateResponseTypeID

	case SirenCommandRequest:
		return SirenCommandRequestTypeID

	case ListEntitiesLockResponse:
		return ListEntitiesLockResponseTypeID

	case LockStateResponse:
		return LockStateResponseTypeID

	case LockCommandRequest:
		return LockCommandRequestTypeID

	case ListEntitiesButtonResponse:
		return ListEntitiesButtonResponseTypeID

	case ButtonCommandRequest:
		return ButtonCommandRequestTypeID

	case ListEntitiesMediaPlayerResponse:
		return ListEntitiesMediaPlayerResponseTypeID

	case MediaPlayerStateResponse:
		return MediaPlayerStateResponseTypeID

	case MediaPlayerCommandRequest:
		return MediaPlayerCommandRequestTypeID

	case SubscribeBluetoothLEAdvertisementsRequest:
		return SubscribeBluetoothLEAdvertisementsRequestTypeID

	case BluetoothLEAdvertisementResponse:
		return BluetoothLEAdvertisementResponseTypeID

	case BluetoothDeviceRequest:
		return BluetoothDeviceRequestTypeID

	case BluetoothDeviceConnectionResponse:
		return BluetoothDeviceConnectionResponseTypeID

	case BluetoothGATTGetServicesRequest:
		return BluetoothGATTGetServicesRequestTypeID

	case BluetoothGATTGetServicesResponse:
		return BluetoothGATTGetServicesResponseTypeID

	case BluetoothGATTGetServicesDoneResponse:
		return BluetoothGATTGetServicesDoneResponseTypeID

	case BluetoothGATTReadRequest:
		return BluetoothGATTReadRequestTypeID

	case BluetoothGATTReadResponse:
		return BluetoothGATTReadResponseTypeID

	case BluetoothGATTWriteRequest:
		return BluetoothGATTWriteRequestTypeID

	case BluetoothGATTReadDescriptorRequest:
		return BluetoothGATTReadDescriptorRequestTypeID

	case BluetoothGATTWriteDescriptorRequest:
		return BluetoothGATTWriteDescriptorRequestTypeID

	case BluetoothGATTNotifyRequest:
		return BluetoothGATTNotifyRequestTypeID

	case BluetoothGATTNotifyDataResponse:
		return BluetoothGATTNotifyDataResponseTypeID

	case SubscribeBluetoothConnectionsFreeRequest:
		return SubscribeBluetoothConnectionsFreeRequestTypeID

	case BluetoothConnectionsFreeResponse:
		return BluetoothConnectionsFreeResponseTypeID

	case BluetoothGATTErrorResponse:
		return BluetoothGATTErrorResponseTypeID

	case BluetoothGATTWriteResponse:
		return BluetoothGATTWriteResponseTypeID

	case BluetoothGATTNotifyResponse:
		return BluetoothGATTNotifyResponseTypeID

	case BluetoothDevicePairingResponse:
		return BluetoothDevicePairingResponseTypeID

	case BluetoothDeviceUnpairingResponse:
		return BluetoothDeviceUnpairingResponseTypeID

	case UnsubscribeBluetoothLEAdvertisementsRequest:
		return UnsubscribeBluetoothLEAdvertisementsRequestTypeID

	case BluetoothDeviceClearCacheResponse:
		return BluetoothDeviceClearCacheResponseTypeID

	case SubscribeVoiceAssistantRequest:
		return SubscribeVoiceAssistantRequestTypeID

	case VoiceAssistantRequest:
		return VoiceAssistantRequestTypeID

	case VoiceAssistantResponse:
		return VoiceAssistantResponseTypeID

	case VoiceAssistantEventResponse:
		return VoiceAssistantEventResponseTypeID

	case BluetoothLERawAdvertisementsResponse:
		return BluetoothLERawAdvertisementsResponseTypeID

	case ListEntitiesAlarmControlPanelResponse:
		return ListEntitiesAlarmControlPanelResponseTypeID

	case AlarmControlPanelStateResponse:
		return AlarmControlPanelStateResponseTypeID

	case AlarmControlPanelCommandRequest:
		return AlarmControlPanelCommandRequestTypeID

	case ListEntitiesTextResponse:
		return ListEntitiesTextResponseTypeID

	case TextStateResponse:
		return TextStateResponseTypeID

	case TextCommandRequest:
		return TextCommandRequestTypeID

	case ListEntitiesDateResponse:
		return ListEntitiesDateResponseTypeID

	case DateStateResponse:
		return DateStateResponseTypeID

	case DateCommandRequest:
		return DateCommandRequestTypeID

	case ListEntitiesTimeResponse:
		return ListEntitiesTimeResponseTypeID

	case TimeStateResponse:
		return TimeStateResponseTypeID

	case TimeCommandRequest:
		return TimeCommandRequestTypeID

	case VoiceAssistantAudio:
		return VoiceAssistantAudioTypeID

	case ListEntitiesEventResponse:
		return ListEntitiesEventResponseTypeID

	case EventResponse:
		return EventResponseTypeID

	case ListEntitiesValveResponse:
		return ListEntitiesValveResponseTypeID

	case ValveStateResponse:
		return ValveStateResponseTypeID

	case ValveCommandRequest:
		return ValveCommandRequestTypeID

	case ListEntitiesDateTimeResponse:
		return ListEntitiesDateTimeResponseTypeID

	case DateTimeStateResponse:
		return DateTimeStateResponseTypeID

	case DateTimeCommandRequest:
		return DateTimeCommandRequestTypeID

	case VoiceAssistantTimerEventResponse:
		return VoiceAssistantTimerEventResponseTypeID

	case ListEntitiesUpdateResponse:
		return ListEntitiesUpdateResponseTypeID

	case UpdateStateResponse:
		return UpdateStateResponseTypeID

	case UpdateCommandRequest:
		return UpdateCommandRequestTypeID

	case VoiceAssistantAnnounceRequest:
		return VoiceAssistantAnnounceRequestTypeID

	case VoiceAssistantAnnounceFinished:
		return VoiceAssistantAnnounceFinishedTypeID

	case VoiceAssistantConfigurationRequest:
		return VoiceAssistantConfigurationRequestTypeID

	case VoiceAssistantConfigurationResponse:
		return VoiceAssistantConfigurationResponseTypeID

	case VoiceAssistantSetConfiguration:
		return VoiceAssistantSetConfigurationTypeID

	case NoiseEncryptionSetKeyRequest:
		return NoiseEncryptionSetKeyRequestTypeID

	case NoiseEncryptionSetKeyResponse:
		return NoiseEncryptionSetKeyResponseTypeID

	case BluetoothScannerStateResponse:
		return BluetoothScannerStateResponseTypeID

	case BluetoothScannerSetModeRequest:
		return BluetoothScannerSetModeRequestTypeID

	case ZWaveProxyFrame:
		return ZWaveProxyFrameTypeID

	case ZWaveProxyRequest:
		return ZWaveProxyRequestTypeID

	case HomeassistantActionResponse:
		return HomeassistantActionResponseTypeID

	case ExecuteServiceResponse:
		return ExecuteServiceResponseTypeID

	case ListEntitiesWaterHeaterResponse:
		return ListEntitiesWaterHeaterResponseTypeID

	case WaterHeaterStateResponse:
		return WaterHeaterStateResponseTypeID

	case WaterHeaterCommandRequest:
		return WaterHeaterCommandRequestTypeID

	case ListEntitiesInfraredResponse:
		return ListEntitiesInfraredResponseTypeID

	case InfraredRFTransmitRawTimingsRequest:
		return InfraredRFTransmitRawTimingsRequestTypeID

	case InfraredRFReceiveEvent:
		return InfraredRFReceiveEventTypeID

	case SerialProxyConfigureRequest:
		return SerialProxyConfigureRequestTypeID

	case SerialProxyDataReceived:
		return SerialProxyDataReceivedTypeID

	case SerialProxyWriteRequest:
		return SerialProxyWriteRequestTypeID

	case SerialProxySetModemPinsRequest:
		return SerialProxySetModemPinsRequestTypeID

	case SerialProxyGetModemPinsRequest:
		return SerialProxyGetModemPinsRequestTypeID

	case SerialProxyGetModemPinsResponse:
		return SerialProxyGetModemPinsResponseTypeID

	case SerialProxyRequest:
		return SerialProxyRequestTypeID

	case BluetoothSetConnectionParamsRequest:
		return BluetoothSetConnectionParamsRequestTypeID

	case BluetoothSetConnectionParamsResponse:
		return BluetoothSetConnectionParamsResponseTypeID

	case SerialProxyRequestResponse:
		return SerialProxyRequestResponseTypeID

	case ListEntitiesRadioFrequencyResponse:
		return ListEntitiesRadioFrequencyResponseTypeID

	default:
		return UndefinedTypeID
	}
}

func NewMessageByTypeID(typeID uint64) proto.Message {
	switch typeID {
	case 1:
		return new(HelloRequest)

	case 2:
		return new(HelloResponse)

	case 3:
		return new(AuthenticationRequest)

	case 4:
		return new(AuthenticationResponse)

	case 5:
		return new(DisconnectRequest)

	case 6:
		return new(DisconnectResponse)

	case 7:
		return new(PingRequest)

	case 8:
		return new(PingResponse)

	case 9:
		return new(DeviceInfoRequest)

	case 10:
		return new(DeviceInfoResponse)

	case 11:
		return new(ListEntitiesRequest)

	case 12:
		return new(ListEntitiesBinarySensorResponse)

	case 13:
		return new(ListEntitiesCoverResponse)

	case 14:
		return new(ListEntitiesFanResponse)

	case 15:
		return new(ListEntitiesLightResponse)

	case 16:
		return new(ListEntitiesSensorResponse)

	case 17:
		return new(ListEntitiesSwitchResponse)

	case 18:
		return new(ListEntitiesTextSensorResponse)

	case 19:
		return new(ListEntitiesDoneResponse)

	case 20:
		return new(SubscribeStatesRequest)

	case 21:
		return new(BinarySensorStateResponse)

	case 22:
		return new(CoverStateResponse)

	case 23:
		return new(FanStateResponse)

	case 24:
		return new(LightStateResponse)

	case 25:
		return new(SensorStateResponse)

	case 26:
		return new(SwitchStateResponse)

	case 27:
		return new(TextSensorStateResponse)

	case 28:
		return new(SubscribeLogsRequest)

	case 29:
		return new(SubscribeLogsResponse)

	case 30:
		return new(CoverCommandRequest)

	case 31:
		return new(FanCommandRequest)

	case 32:
		return new(LightCommandRequest)

	case 33:
		return new(SwitchCommandRequest)

	case 34:
		return new(SubscribeHomeassistantServicesRequest)

	case 35:
		return new(HomeassistantActionRequest)

	case 36:
		return new(GetTimeRequest)

	case 37:
		return new(GetTimeResponse)

	case 38:
		return new(SubscribeHomeAssistantStatesRequest)

	case 39:
		return new(SubscribeHomeAssistantStateResponse)

	case 40:
		return new(HomeAssistantStateResponse)

	case 41:
		return new(ListEntitiesServicesResponse)

	case 42:
		return new(ExecuteServiceRequest)

	case 43:
		return new(ListEntitiesCameraResponse)

	case 44:
		return new(CameraImageResponse)

	case 45:
		return new(CameraImageRequest)

	case 46:
		return new(ListEntitiesClimateResponse)

	case 47:
		return new(ClimateStateResponse)

	case 48:
		return new(ClimateCommandRequest)

	case 49:
		return new(ListEntitiesNumberResponse)

	case 50:
		return new(NumberStateResponse)

	case 51:
		return new(NumberCommandRequest)

	case 52:
		return new(ListEntitiesSelectResponse)

	case 53:
		return new(SelectStateResponse)

	case 54:
		return new(SelectCommandRequest)

	case 55:
		return new(ListEntitiesSirenResponse)

	case 56:
		return new(SirenStateResponse)

	case 57:
		return new(SirenCommandRequest)

	case 58:
		return new(ListEntitiesLockResponse)

	case 59:
		return new(LockStateResponse)

	case 60:
		return new(LockCommandRequest)

	case 61:
		return new(ListEntitiesButtonResponse)

	case 62:
		return new(ButtonCommandRequest)

	case 63:
		return new(ListEntitiesMediaPlayerResponse)

	case 64:
		return new(MediaPlayerStateResponse)

	case 65:
		return new(MediaPlayerCommandRequest)

	case 66:
		return new(SubscribeBluetoothLEAdvertisementsRequest)

	case 67:
		return new(BluetoothLEAdvertisementResponse)

	case 68:
		return new(BluetoothDeviceRequest)

	case 69:
		return new(BluetoothDeviceConnectionResponse)

	case 70:
		return new(BluetoothGATTGetServicesRequest)

	case 71:
		return new(BluetoothGATTGetServicesResponse)

	case 72:
		return new(BluetoothGATTGetServicesDoneResponse)

	case 73:
		return new(BluetoothGATTReadRequest)

	case 74:
		return new(BluetoothGATTReadResponse)

	case 75:
		return new(BluetoothGATTWriteRequest)

	case 76:
		return new(BluetoothGATTReadDescriptorRequest)

	case 77:
		return new(BluetoothGATTWriteDescriptorRequest)

	case 78:
		return new(BluetoothGATTNotifyRequest)

	case 79:
		return new(BluetoothGATTNotifyDataResponse)

	case 80:
		return new(SubscribeBluetoothConnectionsFreeRequest)

	case 81:
		return new(BluetoothConnectionsFreeResponse)

	case 82:
		return new(BluetoothGATTErrorResponse)

	case 83:
		return new(BluetoothGATTWriteResponse)

	case 84:
		return new(BluetoothGATTNotifyResponse)

	case 85:
		return new(BluetoothDevicePairingResponse)

	case 86:
		return new(BluetoothDeviceUnpairingResponse)

	case 87:
		return new(UnsubscribeBluetoothLEAdvertisementsRequest)

	case 88:
		return new(BluetoothDeviceClearCacheResponse)

	case 89:
		return new(SubscribeVoiceAssistantRequest)

	case 90:
		return new(VoiceAssistantRequest)

	case 91:
		return new(VoiceAssistantResponse)

	case 92:
		return new(VoiceAssistantEventResponse)

	case 93:
		return new(BluetoothLERawAdvertisementsResponse)

	case 94:
		return new(ListEntitiesAlarmControlPanelResponse)

	case 95:
		return new(AlarmControlPanelStateResponse)

	case 96:
		return new(AlarmControlPanelCommandRequest)

	case 97:
		return new(ListEntitiesTextResponse)

	case 98:
		return new(TextStateResponse)

	case 99:
		return new(TextCommandRequest)

	case 100:
		return new(ListEntitiesDateResponse)

	case 101:
		return new(DateStateResponse)

	case 102:
		return new(DateCommandRequest)

	case 103:
		return new(ListEntitiesTimeResponse)

	case 104:
		return new(TimeStateResponse)

	case 105:
		return new(TimeCommandRequest)

	case 106:
		return new(VoiceAssistantAudio)

	case 107:
		return new(ListEntitiesEventResponse)

	case 108:
		return new(EventResponse)

	case 109:
		return new(ListEntitiesValveResponse)

	case 110:
		return new(ValveStateResponse)

	case 111:
		return new(ValveCommandRequest)

	case 112:
		return new(ListEntitiesDateTimeResponse)

	case 113:
		return new(DateTimeStateResponse)

	case 114:
		return new(DateTimeCommandRequest)

	case 115:
		return new(VoiceAssistantTimerEventResponse)

	case 116:
		return new(ListEntitiesUpdateResponse)

	case 117:
		return new(UpdateStateResponse)

	case 118:
		return new(UpdateCommandRequest)

	case 119:
		return new(VoiceAssistantAnnounceRequest)

	case 120:
		return new(VoiceAssistantAnnounceFinished)

	case 121:
		return new(VoiceAssistantConfigurationRequest)

	case 122:
		return new(VoiceAssistantConfigurationResponse)

	case 123:
		return new(VoiceAssistantSetConfiguration)

	case 124:
		return new(NoiseEncryptionSetKeyRequest)

	case 125:
		return new(NoiseEncryptionSetKeyResponse)

	case 126:
		return new(BluetoothScannerStateResponse)

	case 127:
		return new(BluetoothScannerSetModeRequest)

	case 128:
		return new(ZWaveProxyFrame)

	case 129:
		return new(ZWaveProxyRequest)

	case 130:
		return new(HomeassistantActionResponse)

	case 131:
		return new(ExecuteServiceResponse)

	case 132:
		return new(ListEntitiesWaterHeaterResponse)

	case 133:
		return new(WaterHeaterStateResponse)

	case 134:
		return new(WaterHeaterCommandRequest)

	case 135:
		return new(ListEntitiesInfraredResponse)

	case 136:
		return new(InfraredRFTransmitRawTimingsRequest)

	case 137:
		return new(InfraredRFReceiveEvent)

	case 138:
		return new(SerialProxyConfigureRequest)

	case 139:
		return new(SerialProxyDataReceived)

	case 140:
		return new(SerialProxyWriteRequest)

	case 141:
		return new(SerialProxySetModemPinsRequest)

	case 142:
		return new(SerialProxyGetModemPinsRequest)

	case 143:
		return new(SerialProxyGetModemPinsResponse)

	case 144:
		return new(SerialProxyRequest)

	case 145:
		return new(BluetoothSetConnectionParamsRequest)

	case 146:
		return new(BluetoothSetConnectionParamsResponse)

	case 147:
		return new(SerialProxyRequestResponse)

	case 148:
		return new(ListEntitiesRadioFrequencyResponse)

	default:
		return nil
	}
}

package api

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

// Request and response types from/to esphome
const (
	UndefinedTypeID = iota
	HelloRequestTypeID
	HelloResponseTypeID
	ConnectRequestTypeID
	ConnectResponseTypeID
	DisconnectRequestTypeID
	DisconnectResponseTypeID
	PingRequestTypeID
	PingResponseTypeID
	DeviceInfoRequestTypeID
	DeviceInfoResponseTypeID
	ListEntitiesRequestTypeID
	ListEntitiesBinarySensorResponseTypeID
	ListEntitiesCoverResponseTypeID
	ListEntitiesFanResponseTypeID
	ListEntitiesLightResponseTypeID
	ListEntitiesSensorResponseTypeID
	ListEntitiesSwitchResponseTypeID
	ListEntitiesTextSensorResponseTypeID
	ListEntitiesDoneResponseTypeID
	SubscribeStatesRequestTypeID
	BinarySensorStateResponseTypeID
	CoverStateResponseTypeID
	FanStateResponseTypeID
	LightStateResponseTypeID
	SensorStateResponseTypeID
	SwitchStateResponseTypeID
	TextSensorStateResponseTypeID
	SubscribeLogsRequestTypeID
	SubscribeLogsResponseTypeID
	CoverCommandRequestTypeID
	FanCommandRequestTypeID
	LightCommandRequestTypeID
	SwitchCommandRequestTypeID
	SubscribeHomeAssistantServicesRequestTypeID
	HomeAssistantServiceResponseTypeID
	GetTimeRequestTypeID
	GetTimeResponseTypeID
	SubscribeHomeAssistantStatesRequestTypeID
	SubscribeHomeAssistantStateResponseTypeID
	HomeAssistantStateResponseTypeID
	ListEntitiesServicesResponseTypeID
	ExecuteServiceRequestTypeID
	ListEntitiesCameraResponseTypeID
	CameraImageResponseTypeID
	CameraImageRequestTypeID
	ListEntitiesClimateResponseTypeID
	ClimateStateResponseTypeID
	ClimateCommandRequestTypeID
	ListEntitiesNumberResponseTypeID
	NumberStateResponseTypeID
	NumberCommandRequestTypeID
	ListEntitiesSelectResponseTypeID
	SelectStateResponseTypeID
	SelectCommandRequestTypeID
	UnknownTypeID55
	UnknownTypeID56
	UnknownTypeID57
	ListEntitiesLockResponseTypeID
	LockStateResponseTypeID
	LockCommandRequestTypeID
	ListEntitiesButtonResponseTypeID
	ButtonCommandRequestTypeID
	ListEntitiesMediaPlayerResponseTypeID
	MediaPlayerStateResponseTypeID
	MediaPlayerCommandRequestTypeID
	SubscribeBluetoothLEAdvertisementsRequestID
	BluetoothLEAdvertisementResponseID
	BluetoothDeviceRequestID
	BluetoothDeviceConnectionResponseID
	BluetoothGATTGetServicesRequestID
	BluetoothGATTGetServicesResponseID
	BluetoothGATTGetServicesDoneResponseID
	BluetoothGATTReadRequestID
	BluetoothGATTReadResponseID
	BluetoothGATTWriteRequestID
	BluetoothGATTReadDescriptorRequestID
	BluetoothGATTWriteDescriptorRequestID
	BluetoothGATTNotifyRequestID
	BluetoothGATTNotifyDataResponseID
	SubscribeBluetoothConnectionsFreeRequestID
	BluetoothConnectionsFreeResponseID
	BluetoothGATTErrorResponseID
	BluetoothGATTWriteResponseID
	BluetoothGATTNotifyResponseID
)

func TypeID(message interface{}) uint64 {
	if message == nil {
		return UndefinedTypeID
	}

	// convert from pointer to normal type
	if reflect.ValueOf(message).Kind() == reflect.Ptr {
		message = reflect.ValueOf(message).Elem().Interface()
	}
	switch message.(type) {

	case HelloRequest:
		return HelloRequestTypeID

	case HelloResponse:
		return HelloResponseTypeID

	case ConnectRequest:
		return ConnectRequestTypeID

	case ConnectResponse:
		return ConnectResponseTypeID

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
		return SubscribeHomeAssistantServicesRequestTypeID

	case HomeassistantServiceResponse:
		return HomeAssistantServiceResponseTypeID

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

	case NumberCommandRequest:
		return ClimateCommandRequestTypeID

	case SelectCommandRequest:
		return SelectCommandRequestTypeID

	case ButtonCommandRequest:
		return ButtonCommandRequestTypeID

	case LockCommandRequest:
		return LockCommandRequestTypeID

	case MediaPlayerCommandRequest:
		return MediaPlayerCommandRequestTypeID

	case ListEntitiesNumberResponse:
		return ListEntitiesNumberResponseTypeID

	case NumberStateResponse:
		return NumberStateResponseTypeID

	case ListEntitiesSelectResponse:
		return ListEntitiesSelectResponseTypeID

	case SelectStateResponse:
		return SelectStateResponseTypeID

	case ListEntitiesLockResponse:
		return ListEntitiesLockResponseTypeID

	case LockStateResponse:
		return LockStateResponseTypeID

	case ListEntitiesButtonResponse:
		return ListEntitiesButtonResponseTypeID

	case ListEntitiesMediaPlayerResponse:
		return ListEntitiesMediaPlayerResponseTypeID

	case MediaPlayerStateResponse:
		return MediaPlayerStateResponseTypeID

	case SubscribeBluetoothLEAdvertisementsRequest:
		return SubscribeBluetoothLEAdvertisementsRequestID

	case BluetoothLEAdvertisementResponse:
		return BluetoothLEAdvertisementResponseID

	case BluetoothDeviceRequest:
		return BluetoothDeviceRequestID

	case BluetoothDeviceConnectionResponse:
		return BluetoothDeviceConnectionResponseID

	case BluetoothGATTGetServicesRequest:
		return BluetoothGATTGetServicesRequestID

	case BluetoothGATTGetServicesResponse:
		return BluetoothGATTGetServicesResponseID

	case BluetoothGATTGetServicesDoneResponse:
		return BluetoothGATTGetServicesDoneResponseID

	case BluetoothGATTReadRequest:
		return BluetoothGATTReadRequestID

	case BluetoothGATTReadResponse:
		return BluetoothGATTReadResponseID

	case BluetoothGATTWriteRequest:
		return BluetoothGATTWriteRequestID

	case BluetoothGATTReadDescriptorRequest:
		return BluetoothGATTReadDescriptorRequestID

	case BluetoothGATTWriteDescriptorRequest:
		return BluetoothGATTWriteDescriptorRequestID

	case BluetoothGATTNotifyRequest:
		return BluetoothGATTNotifyRequestID

	case BluetoothGATTNotifyDataResponse:
		return BluetoothGATTNotifyDataResponseID

	case SubscribeBluetoothConnectionsFreeRequest:
		return SubscribeBluetoothConnectionsFreeRequestID

	case BluetoothConnectionsFreeResponse:
		return BluetoothConnectionsFreeResponseID

	case BluetoothGATTErrorResponse:
		return BluetoothGATTErrorResponseID

	case BluetoothGATTWriteResponse:
		return BluetoothGATTWriteResponseID

	case BluetoothGATTNotifyResponse:
		return BluetoothGATTNotifyResponseID

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
		return new(ConnectRequest)

	case 4:
		return new(ConnectResponse)

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
		return new(HomeassistantServiceResponse)

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

	// case 55:
	// 	return new(UnknownTypeID55)
	//
	// case 56:
	// 	return new(UnknownTypeID56)
	//
	// case 57:
	// 	return new(UnknownTypeID57)

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

	default:
		return nil
	}
}

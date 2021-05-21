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

	default:
		return nil
	}
}

package starling

type CommonProperties struct {
	Type          string `json:"type"`
	ID            string `json:"id"`
	Where         string `json:"where"`
	Name          string `json:"name"`
	SerialNumber  string `json:"serialNumber"`
	StructureName string `json:"structureName"`
}

type ThermostatProperties struct {
	CommonProperties
	DisplayTemperatureUnits           string  `json:"displayTemperatureUnits"`
	BackplateTemperature              float64 `json:"backplateTemperature"`
	CurrentTemperature                float64 `json:"currentTemperature"`
	TargetTemperature                 float64 `json:"targetTemperature"`
	TargetHeatingThresholdTemperature float64 `json:"targetHeatingThresholdTemperature"`
	TargetCoolingThresholdTemperature float64 `json:"targetCoolingThresholdTemperature"`
	HumidityPercent                   float64 `json:"humidityPercent"`
	HVACMode                          string  `json:"hvacMode"`
	HVACState                         string  `json:"hvacState"`
	CanHeat                           bool    `json:"canHeat"`
	CanCool                           bool    `json:"canCool"`
	EcoMode                           bool    `json:"ecoMode"`
}

type TemperatureSensorProperties struct {
	BatteryStatus      string `json:"batteryStatus"`
	CurrentTemperature int    `json:"currentTemperature"`
}

type ProtectProperties struct {
	CommonProperties
	BatteryStatus     string `json:"batteryStatus"`
	CODetected        bool   `json:"coDetected"`
	ManualTestActive  bool   `json:"manualTestActive"`
	OccupancyDetected bool   `json:"occupancyDetected"`
	SmokeDetected     bool   `json:"smokeDetected"`
}

type CameraProperties struct {
	AnimalDetected    bool   `json:"animalDetected"`
	BatteryIsCharging bool   `json:"batteryIsCharging"`
	BatteryLevel      int    `json:"batteryLevel"`
	BatteryStatus     string `json:"batteryStatus"`
	CameraEnabled     bool   `json:"cameraEnabled"`
	CameraModel       string `json:"cameraModel"`
	ChimeEnabled      bool   `json:"chimeEnabled"`
	DoorbellPushed    bool   `json:"doorbellPushed"`
	FloodlightOn      bool   `json:"floodlightOn"`
	MotionDetected    bool   `json:"motionDetected"`
	PackageDelivered  bool   `json:"packageDelivered"`
	PackageRetrieved  bool   `json:"packageRetrieved"`
	PersonDetected    bool   `json:"personDetected"`
	RunningOnBattery  bool   `json:"runningOnBattery"`
	SoundDetected     bool   `json:"soundDetected"`
	SupportsStreaming bool   `json:"supportsStreaming"`
	TrickleCharging   bool   `json:"trickleCharging"`
	VehicleDetected   bool   `json:"vehicleDetected"`
}

type GuardProperties struct {
	BatteryStatus string `json:"batteryStatus"`
	CurrentState  string `json:"currentState"`
	TargetState   string `json:"targetState"`
}

type DetectProperties struct {
	BatteryStatus string `json:"batteryStatus"`
	ButtonPushed  bool   `json:"buttonPushed"`
	ContactState  string `json:"contactState"`
	FixtureType   string `json:"fixtureType"`
	IsTampered    bool   `json:"isTampered"`
}

type YaleLockProperties struct {
	AutoRelockEnabled   bool   `json:"autoRelockEnabled"`
	BatteryStatus       string `json:"batteryStatus"`
	CurrentState        string `json:"currentState"`
	IsTampered          bool   `json:"isTampered"`
	OneTouchLockEnabled bool   `json:"oneTouchLockEnabled"`
	PrivacyModeEnabled  bool   `json:"privacyModeEnabled"`
	TargetState         string `json:"targetState"`
}

type HomeAwayControlProperties struct {
	HomeState string `json:"homeState"`
}

type WeatherServiceProperties struct {
	CurrentTemperature int    `json:"currentTemperature"`
	HumidityPercent    int    `json:"humidityPercent"`
	SerialNumber       string `json:"serialNumber"`
	StructureName      string `json:"structureName"`
}

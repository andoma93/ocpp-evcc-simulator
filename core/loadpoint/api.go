package loadpoint

import (
	"time"

	"github.com/evcc-io/evcc/api"
)

// Controller gives access to loadpoint
type Controller interface {
	LoadpointControl(API)
}

// API is the external loadpoint API
type API interface {
	// Name returns the defined loadpoint name
	Name() string

	//
	// status
	//

	// GetStatus returns the charging status
	GetStatus() api.ChargeStatus

	//
	// settings
	//

	// GetMode returns the charge mode
	GetMode() api.ChargeMode
	// SetMode sets the charge mode
	SetMode(api.ChargeMode)
	// GetTargetEnergy returns the charge target energy
	GetTargetEnergy() float64
	// SetTargetEnergy sets the charge target energy
	SetTargetEnergy(float64)
	// GetTargetSoc returns the charge target soc
	GetTargetSoc() int
	// SetTargetSoc sets the charge target soc
	SetTargetSoc(int)
	// GetMinSoc returns the charge minimum soc
	GetMinSoc() int
	// SetMinSoc sets the charge minimum soc
	SetMinSoc(int)
	// GetPhases returns the enabled phases
	GetPhases() int
	// SetPhases sets the enabled phases
	SetPhases(int) error

	// SetTargetCharge sets the charge targetSoc
	SetTargetCharge(time.Time, int) error
	// RemoteControl sets remote status demand
	RemoteControl(string, RemoteDemand)

	//
	// power and energy
	//

	// HasChargeMeter determines if a physical charge meter is attached
	HasChargeMeter() bool
	// GetChargePower returns the current charging power
	GetChargePower() float64
	// GetMinCurrent returns the min charging current
	GetMinCurrent() float64
	// SetMinCurrent sets the min charging current
	SetMinCurrent(float64)
	// GetMaxCurrent returns the max charging current
	GetMaxCurrent() float64
	// SetMaxCurrent sets the max charging current
	SetMaxCurrent(float64)
	// GetMinPower returns the min charging power for a single phase
	GetMinPower() float64
	// GetMaxPower returns the max charging power taking active phases into account
	GetMaxPower() float64

	//
	// charge progress
	//

	// GetRemainingDuration is the estimated remaining charging duration
	GetRemainingDuration() time.Duration
	// GetRemainingEnergy is the remaining charge energy in Wh
	GetRemainingEnergy() float64

	//
	// vehicles
	//

	// GetVehicle gets the active vehicle
	GetVehicle() api.Vehicle
	// SetVehicle sets the active vehicle
	SetVehicle(vehicle api.Vehicle)
	// StartVehicleDetection allows triggering vehicle detection for debugging purposes
	StartVehicleDetection()
}

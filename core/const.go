package core

const (
	phasesConfigured = "phasesConfigured" // configured phases (1/3, 0 for auto on 1p3p chargers, nil for plain chargers)
	phasesEnabled    = "phasesEnabled"    // enabled phases (1/3)
	phasesActive     = "phasesActive"     // active phases as used by vehicle (1/2/3)

	vehicleDetectionActive = "vehicleDetectionActive" // vehicle detection is active (bool)

	vehicleRange     = "vehicleRange"     // vehicle range
	vehicleOdometer  = "vehicleOdometer"  // vehicle odometer
	vehicleSoc       = "vehicleSoc"       // vehicle soc
	vehicleTargetSoc = "vehicleTargetSoc" // vehicle soc limit

	minSoc                   = "minSoc"                   // min soc goal
	targetSoc                = "targetSoc"                // target charging soc goal
	targetTime               = "targetTime"               // target charging finish time goal
	targetTimeActive         = "targetTimeActive"         // target charging plan has determined current slot to be an active slot
	targetTimeProjectedStart = "targetTimeProjectedStart" // target charging plan start time (earliest slot)
)

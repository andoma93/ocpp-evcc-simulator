package ocpp

import (
	"errors"
	"fmt"
	"sync"

	"github.com/evcc-io/evcc/util"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
)

type CS struct {
	mu  sync.Mutex
	log *util.Logger
	ocpp16.CentralSystem
	cps map[string]*CP
}

func (cs *CS) Register(id string, cp *CP) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, ok := cs.cps[id]; ok && id == "" {
		return errors.New("cannot have >1 chargepoint with empty station id")
	}

	cs.cps[id] = cp

	return nil
}

// errorHandler logs error channel
func (cs *CS) errorHandler(errC <-chan error) {
	for err := range errC {
		cs.log.ERROR.Println(err)
	}
}

func (cs *CS) chargepointByID(id string) (*CP, error) {
	cp, ok := cs.cps[id]
	if !ok {
		return nil, fmt.Errorf("unknown charge point: %s", id)
	}
	return cp, nil
}

func (cs *CS) NewChargePoint(chargePoint ocpp16.ChargePointConnection) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if cp, err := cs.chargepointByID(chargePoint.ID()); err != nil {
		if cp, ok := cs.cps[""]; ok {
			cs.log.INFO.Printf("chargepoint connected, registering: %s", chargePoint.ID())

			// update id
			cp.RegisterID(chargePoint.ID())
			cs.cps[chargePoint.ID()] = cp
			delete(cs.cps, "")

			cp.Connect()

			return
		}

		cs.log.WARN.Printf("chargepoint connected, ignoring: %s", chargePoint.ID())
	} else {
		cs.log.DEBUG.Printf("chargepoint connected: %s", chargePoint.ID())
		cp.Connect()
	}
}

func (cs *CS) ChargePointDisconnected(chargePoint ocpp16.ChargePointConnection) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, err := cs.chargepointByID(chargePoint.ID()); err != nil {
		cs.log.ERROR.Printf("chargepoint disconnected: %v", err)
	} else {
		cs.log.DEBUG.Printf("chargepoint disconnected: %s", chargePoint.ID())
	}
}

func (cs *CS) Debug(args ...interface{}) {
	cs.log.TRACE.Println(args...)
}

func (cs *CS) Debugf(fmt string, args ...interface{}) {
	cs.log.TRACE.Printf(fmt, args...)
}

func (cs *CS) Info(args ...interface{}) {
	cs.log.DEBUG.Println(args...)
}

func (cs *CS) Infof(fmt string, args ...interface{}) {
	cs.log.DEBUG.Printf(fmt, args...)
}

func (cs *CS) Error(args ...interface{}) {
	cs.log.ERROR.Println(args...)
}

func (cs *CS) Errorf(fmt string, args ...interface{}) {
	cs.log.ERROR.Printf(fmt, args...)
}

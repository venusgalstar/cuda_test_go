package cudnn

/* WAS Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// AlgorithmPerformance is a representation of cudnnAlgorithmPerformance_t.
type AlgorithmPerformance struct {
	internal C.cudnnAlgorithmPerformance_t

	n int

	algoDesc *AlgorithmDescriptor
	status   Status
	time     float32
	memory   uintptr
}

// NewAlgorithmPerformance creates `n` new cudnnAlgorithmPerformance objects, wrapped in AlgorithmPerformance.
func NewAlgorithmPerformance(algoDesc *AlgorithmDescriptor, n int, status Status, time float32, memory uintptr) (retVal *AlgorithmPerformance, err error) {
	var internal C.cudnnAlgorithmPerformance_t
	if err := result(C.cudnnCreateAlgorithmPerformance(&internal, C.int(n))); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetAlgorithmPerformance(internal, algoDesc.internal, status.C(), C.float(time), C.size_t(memory))); err != nil {
		return nil, err
	}

	retVal = &AlgorithmPerformance{
		internal: internal,
		algoDesc: algoDesc,
		status:   status,
		time:     time,
		memory:   memory,
	}
	runtime.SetFinalizer(retVal, destroyAlgorithmPerformance)
	return retVal, nil
}

// C returns the cgo representation.
func (a *AlgorithmPerformance) C() C.cudnnAlgorithmPerformance_t { return a.internal }

// AlgoDesc returns the internal algoDesc.
func (a *AlgorithmPerformance) AlgoDesc() *AlgorithmDescriptor { return a.algoDesc }

// Status returns the internal status.
func (a *AlgorithmPerformance) Status() Status { return a.status }

// Time returns the internal time.
func (a *AlgorithmPerformance) Time() float32 { return a.time }

// Memory returns the internal memory.
func (a *AlgorithmPerformance) Memory() uintptr { return a.memory }

// N returns how many algorithm performance objects were created in the graphics card.
func (a *AlgorithmPerformance) N() int { return a.n }

func destroyAlgorithmPerformance(obj *AlgorithmPerformance) {
	C.cudnnDestroyAlgorithmPerformance(&obj.internal, C.int(obj.n))
}

package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
)

// Dropout is a representation of cudnnDropoutDescriptor_t.
//
// The usecase of Dropout is quite different from the rest of the APIs in this library. There is a two stage API:
//		drop := NewDropout(...)
//		drop.Use(ctx, states....)
//
// This is because the Dropout is largely tied to run-time. An additional `.IsReady` method is added to indicate if the dropout state is ready to be used
//
// However, if your runtime is known ahead of time, the `NewDropoutWithContext` creation function can be used.
type Dropout struct {
	internal C.cudnnDropoutDescriptor_t

	handle           *Context
	dropout          float32
	states           Memory
	stateSizeInBytes uintptr
	seed             uint64
	reqStateSize     uintptr
}

// NewDropout creates a Dropout descriptor. It is not usable by default because some additional stateful information needs to be passed in
func NewDropout(dropout float64) (retVal *Dropout, err error) {
	var internal C.cudnnDropoutDescriptor_t
	if err := result(C.cudnnCreateDropoutDescriptor(&internal)); err != nil {
		return nil, err
	}
	retVal = &Dropout{
		internal: internal,
		dropout:  float32(dropout),
	}
	runtime.SetFinalizer(retVal, destroyDropout)
	return retVal, nil
}

// NewDropout creates a new Dropout with the given context (handle, states, etc)
func NewDropoutWithContext(dropout float64, handle *Context, states Memory, stateSizeInBytes uintptr, seed uint64) (retVal *Dropout, err error) {
	if retVal, err = NewDropout(dropout); err != nil {
		return
	}
	err = retVal.Use(handle, states, stateSizeInBytes, seed)
	return
}

// Use is the second stage of the two-stage API.
func (dr *Dropout) Use(ctx *Context, states Memory, stateSizeInBytes uintptr, seed uint64) error {
	dr.handle = ctx
	dr.states = states
	dr.stateSizeInBytes = stateSizeInBytes
	dr.seed = seed

	return result(C.cudnnSetDropoutDescriptor(dr.internal, dr.handle.internal, C.float(dr.dropout), unsafe.Pointer(dr.states.Uintptr()), C.size_t(dr.stateSizeInBytes), C.ulonglong(dr.seed)))
}

// IsReady indicates if the dropout operator is ready to be used
func (dr *Dropout) IsReady() bool {
	return dr.handle != nil && dr.states != nil && dr.stateSizeInBytes != 0
}

// Reset resets the state to be not ready. It does NOT reset the dropout ratio.
func (dr *Dropout) Reset() {
	dr.handle = nil
	dr.states = nil
	dr.stateSizeInBytes = 0
	dr.seed = 0
}

// Handle returns the internal handle.
func (dr *Dropout) Handle() *Context { return dr.handle }

// Dropout returns the internal dropout ratio.
func (dr *Dropout) Dropout() float32 { return dr.dropout }

// StateSizeInBytes returns the internal stateSizeInBytes.
func (dr *Dropout) StateSizeInBytes() uintptr { return dr.stateSizeInBytes }

// Seed returns the internal seed.
func (dr *Dropout) Seed() uint64 { return dr.seed }

func (dr *Dropout) States() Memory { return dr.states }

func (dr *Dropout) RequiredStateSize(ctx *Context) (uintptr, error) {
	if dr.reqStateSize > 0 {
		return dr.reqStateSize, nil
	}

	var minSize C.size_t
	if err := result(C.cudnnDropoutGetStatesSize(ctx.internal, &minSize)); err != nil {
		return 0, errors.Wrapf(err, "Unable to get minimum state size")
	}

	dr.reqStateSize = uintptr(minSize)
	return dr.reqStateSize, nil
}

// BUG(anyone): the memory for the scratch space isn't freed. This could potentially lead to some issues
func destroyDropout(obj *Dropout) { obj.Reset(); C.cudnnDestroyDropoutDescriptor(obj.internal) }

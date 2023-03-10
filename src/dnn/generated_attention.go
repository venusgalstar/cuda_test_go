package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// Attention is a representation of cudnnAttnDescriptor_t.
type Attention struct {
	internal C.cudnnAttnDescriptor_t

	attnMode        uint
	nHeads          int
	smScaler        float64
	dataType        DataType
	computePrec     DataType
	mathType        MathType
	attnDropoutDesc *Dropout
	postDropoutDesc *Dropout
	qSize           int
	kSize           int
	vSize           int
	qProjSize       int
	kProjSize       int
	vProjSize       int
	oProjSize       int
	qoMaxSeqLength  int
	kvMaxSeqLength  int
	maxBatchSize    int
	maxBeamSize     int
}

// NewAttention creates a new Attention.
func NewAttention(attnMode uint, nHeads int, smScaler float64, dataType DataType, computePrec DataType, mathType MathType, attnDropoutDesc *Dropout, postDropoutDesc *Dropout, qSize int, kSize int, vSize int, qProjSize int, kProjSize int, vProjSize int, oProjSize int, qoMaxSeqLength int, kvMaxSeqLength int, maxBatchSize int, maxBeamSize int) (retVal *Attention, err error) {
	var internal C.cudnnAttnDescriptor_t
	if err := result(C.cudnnCreateAttnDescriptor(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetAttnDescriptor(internal, C.uint(attnMode), C.int(nHeads), C.double(smScaler), dataType.C(), computePrec.C(), mathType.C(), attnDropoutDesc.internal, postDropoutDesc.internal, C.int(qSize), C.int(kSize), C.int(vSize), C.int(qProjSize), C.int(kProjSize), C.int(vProjSize), C.int(oProjSize), C.int(qoMaxSeqLength), C.int(kvMaxSeqLength), C.int(maxBatchSize), C.int(maxBeamSize))); err != nil {
		return nil, err
	}

	retVal = &Attention{
		internal:        internal,
		attnMode:        attnMode,
		nHeads:          nHeads,
		smScaler:        smScaler,
		dataType:        dataType,
		computePrec:     computePrec,
		mathType:        mathType,
		attnDropoutDesc: attnDropoutDesc,
		postDropoutDesc: postDropoutDesc,
		qSize:           qSize,
		kSize:           kSize,
		vSize:           vSize,
		qProjSize:       qProjSize,
		kProjSize:       kProjSize,
		vProjSize:       vProjSize,
		oProjSize:       oProjSize,
		qoMaxSeqLength:  qoMaxSeqLength,
		kvMaxSeqLength:  kvMaxSeqLength,
		maxBatchSize:    maxBatchSize,
		maxBeamSize:     maxBeamSize,
	}
	runtime.SetFinalizer(retVal, destroyAttention)
	return retVal, nil
}

// C returns the internal cgo representation.
func (a *Attention) C() C.cudnnAttnDescriptor_t { return a.internal }

// AttnMode returns the internal attnMode.
func (a *Attention) AttnMode() uint { return a.attnMode }

// NHeads returns the internal nHeads.
func (a *Attention) NHeads() int { return a.nHeads }

// SmScaler returns the internal smScaler.
func (a *Attention) SmScaler() float64 { return a.smScaler }

// DataType returns the internal dataType.
func (a *Attention) DataType() DataType { return a.dataType }

// ComputePrec returns the internal computePrec.
func (a *Attention) ComputePrec() DataType { return a.computePrec }

// MathType returns the internal mathType.
func (a *Attention) MathType() MathType { return a.mathType }

// AttnDropoutDesc returns the internal attnDropoutDesc.
func (a *Attention) AttnDropoutDesc() *Dropout { return a.attnDropoutDesc }

// PostDropoutDesc returns the internal postDropoutDesc.
func (a *Attention) PostDropoutDesc() *Dropout { return a.postDropoutDesc }

// QSize returns the internal qSize.
func (a *Attention) QSize() int { return a.qSize }

// KSize returns the internal kSize.
func (a *Attention) KSize() int { return a.kSize }

// VSize returns the internal vSize.
func (a *Attention) VSize() int { return a.vSize }

// QProjSize returns the internal qProjSize.
func (a *Attention) QProjSize() int { return a.qProjSize }

// KProjSize returns the internal kProjSize.
func (a *Attention) KProjSize() int { return a.kProjSize }

// VProjSize returns the internal vProjSize.
func (a *Attention) VProjSize() int { return a.vProjSize }

// OProjSize returns the internal oProjSize.
func (a *Attention) OProjSize() int { return a.oProjSize }

// QoMaxSeqLength returns the internal qoMaxSeqLength.
func (a *Attention) QoMaxSeqLength() int { return a.qoMaxSeqLength }

// KvMaxSeqLength returns the internal kvMaxSeqLength.
func (a *Attention) KvMaxSeqLength() int { return a.kvMaxSeqLength }

// MaxBatchSize returns the internal maxBatchSize.
func (a *Attention) MaxBatchSize() int { return a.maxBatchSize }

// MaxBeamSize returns the internal maxBeamSize.
func (a *Attention) MaxBeamSize() int { return a.maxBeamSize }

func destroyAttention(obj *Attention) { C.cudnnDestroyAttnDescriptor(obj.internal) }

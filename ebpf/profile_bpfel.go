// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64

package ebpfspy

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type profileBssArg struct {
	TgidFilter    uint32
	CollectUser   uint8
	CollectKernel uint8
	_             [2]byte
}

type profilePidData struct {
	CurrentStateAddr  uint64
	TlsKeyAddr        uint64
	GilLockedAddr     uint64
	GilLastHolderAddr uint64
	Offsets           struct {
		PyObjectType            int64
		PyTypeObjectName        int64
		PyThreadStateFrame      int64
		PyThreadStateThread     int64
		PyFrameObjectBack       int64
		PyFrameObjectCode       int64
		PyFrameObjectLineno     int64
		PyFrameObjectLocalsplus int64
		PyCodeObjectFilename    int64
		PyCodeObjectName        int64
		PyCodeObjectVarnames    int64
		PyTupleObjectItem       int64
		StringData              int64
		StringSize              int64
	}
}

type profilePyEvent struct {
	Pid              uint32
	Tid              uint32
	Comm             [16]int8
	ThreadStateMatch uint8
	GilState         uint8
	PthreadIdMatch   uint8
	StackStatus      uint8
	_                [4]byte
	StackLen         int64
	Stack            [75]uint32
	_                [4]byte
}

type profileSampleKey struct {
	Pid       uint32
	_         [4]byte
	KernStack int64
	UserStack int64
	Comm      [16]int8
}

type profileSampleStateT struct {
	Offsets struct {
		PyObjectType            int64
		PyTypeObjectName        int64
		PyThreadStateFrame      int64
		PyThreadStateThread     int64
		PyFrameObjectBack       int64
		PyFrameObjectCode       int64
		PyFrameObjectLineno     int64
		PyFrameObjectLocalsplus int64
		PyCodeObjectFilename    int64
		PyCodeObjectName        int64
		PyCodeObjectVarnames    int64
		PyTupleObjectItem       int64
		StringData              int64
		StringSize              int64
	}
	CurCpu                 uint64
	SymbolCounter          int64
	FramePtr               uint64
	PythonStackProgCallCnt int64
	Event                  profilePyEvent
}

type profileSymbol struct {
	Classname [32]int8
	Name      [64]int8
	File      [128]int8
}

// loadProfile returns the embedded CollectionSpec for profile.
func loadProfile() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ProfileBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load profile: %w", err)
	}

	return spec, err
}

// loadProfileObjects loads profile and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*profileObjects
//	*profilePrograms
//	*profileMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadProfileObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadProfile()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// profileSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type profileSpecs struct {
	profileProgramSpecs
	profileMapSpecs
}

// profileSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type profileProgramSpecs struct {
	DoPerfEvent     *ebpf.ProgramSpec `ebpf:"do_perf_event"`
	OnEvent         *ebpf.ProgramSpec `ebpf:"on_event"`
	ReadPythonStack *ebpf.ProgramSpec `ebpf:"read_python_stack"`
}

// profileMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type profileMapSpecs struct {
	Args          *ebpf.MapSpec `ebpf:"args"`
	Counts        *ebpf.MapSpec `ebpf:"counts"`
	PyEvents      *ebpf.MapSpec `ebpf:"py_events"`
	PyPidConfig   *ebpf.MapSpec `ebpf:"py_pid_config"`
	PyProgs       *ebpf.MapSpec `ebpf:"py_progs"`
	PyStateHeap   *ebpf.MapSpec `ebpf:"py_state_heap"`
	PyStubsEvents *ebpf.MapSpec `ebpf:"py_stubs_events"`
	PySymbols     *ebpf.MapSpec `ebpf:"py_symbols"`
	Stacks        *ebpf.MapSpec `ebpf:"stacks"`
}

// profileObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadProfileObjects or ebpf.CollectionSpec.LoadAndAssign.
type profileObjects struct {
	profilePrograms
	profileMaps
}

func (o *profileObjects) Close() error {
	return _ProfileClose(
		&o.profilePrograms,
		&o.profileMaps,
	)
}

// profileMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadProfileObjects or ebpf.CollectionSpec.LoadAndAssign.
type profileMaps struct {
	Args          *ebpf.Map `ebpf:"args"`
	Counts        *ebpf.Map `ebpf:"counts"`
	PyEvents      *ebpf.Map `ebpf:"py_events"`
	PyPidConfig   *ebpf.Map `ebpf:"py_pid_config"`
	PyProgs       *ebpf.Map `ebpf:"py_progs"`
	PyStateHeap   *ebpf.Map `ebpf:"py_state_heap"`
	PyStubsEvents *ebpf.Map `ebpf:"py_stubs_events"`
	PySymbols     *ebpf.Map `ebpf:"py_symbols"`
	Stacks        *ebpf.Map `ebpf:"stacks"`
}

func (m *profileMaps) Close() error {
	return _ProfileClose(
		m.Args,
		m.Counts,
		m.PyEvents,
		m.PyPidConfig,
		m.PyProgs,
		m.PyStateHeap,
		m.PyStubsEvents,
		m.PySymbols,
		m.Stacks,
	)
}

// profilePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadProfileObjects or ebpf.CollectionSpec.LoadAndAssign.
type profilePrograms struct {
	DoPerfEvent     *ebpf.Program `ebpf:"do_perf_event"`
	OnEvent         *ebpf.Program `ebpf:"on_event"`
	ReadPythonStack *ebpf.Program `ebpf:"read_python_stack"`
}

func (p *profilePrograms) Close() error {
	return _ProfileClose(
		p.DoPerfEvent,
		p.OnEvent,
		p.ReadPythonStack,
	)
}

func _ProfileClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed profile_bpfel.o
var _ProfileBytes []byte

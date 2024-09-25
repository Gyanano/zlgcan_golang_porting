package zlgcan

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

const (
	ZCAN_PCI5121              = 0x1
	ZCAN_PCI9810              = 0x2
	ZCAN_USBCAN1              = 0x3
	ZCAN_USBCAN2              = 0x4
	ZCAN_PCI9820              = 0x5
	ZCAN_CAN232               = 0x6
	ZCAN_PCI5110              = 0x7
	ZCAN_CANLITE              = 0x8
	ZCAN_ISA9620              = 0x9
	ZCAN_ISA5420              = 0xa
	ZCAN_PC104CAN             = 0xb
	ZCAN_CANETUDP             = 0xc
	ZCAN_CANETE               = 0xc
	ZCAN_DNP9810              = 0xd
	ZCAN_PCI9840              = 0xe
	ZCAN_PC104CAN2            = 0xf
	ZCAN_PCI9820I             = 0x10
	ZCAN_CANETTCP             = 0x11
	ZCAN_PCIE_9220            = 0x12
	ZCAN_PCI5010U             = 0x13
	ZCAN_USBCAN_E_U           = 0x14
	ZCAN_USBCAN_2E_U          = 0x15
	ZCAN_PCI5020U             = 0x16
	ZCAN_EG20T_CAN            = 0x17
	ZCAN_PCIE9221             = 0x18
	ZCAN_WIFICAN_TCP          = 0x19
	ZCAN_WIFICAN_UDP          = 0x1a
	ZCAN_PCIe9120             = 0x1b
	ZCAN_PCIe9110             = 0x1c
	ZCAN_PCIe9140             = 0x1d
	ZCAN_USBCAN_4E_U          = 0x1f
	ZCAN_CANDTU_200UR         = 0x20
	ZCAN_CANDTU_MINI          = 0x21
	ZCAN_USBCAN_8E_U          = 0x22
	ZCAN_CANREPLAY            = 0x23
	ZCAN_CANDTU_NET           = 0x24
	ZCAN_CANDTU_100UR         = 0x25
	ZCAN_PCIE_CANFD_100U      = 0x26
	ZCAN_PCIE_CANFD_200U      = 0x27
	ZCAN_PCIE_CANFD_400U      = 0x28
	ZCAN_USBCANFD_200U        = 0x29
	ZCAN_USBCANFD_100U        = 0x2a
	ZCAN_USBCANFD_MINI        = 0x2b
	ZCAN_CANFDCOM_100IE       = 0x2c
	ZCAN_CANSCOPE             = 0x2d
	ZCAN_CLOUD                = 0x2e
	ZCAN_CANDTU_NET_400       = 0x2f
	ZCAN_CANFDNET_TCP         = 0x30
	ZCAN_CANFDNET_200U_TCP    = 0x30
	ZCAN_CANFDNET_UDP         = 0x31
	ZCAN_CANFDNET_200U_UDP    = 0x31
	ZCAN_CANFDWIFI_TCP        = 0x32
	ZCAN_CANFDWIFI_100U_TCP   = 0x32
	ZCAN_CANFDWIFI_UDP        = 0x33
	ZCAN_CANFDWIFI_100U_UDP   = 0x33
	ZCAN_CANFDNET_400U_TCP    = 0x34
	ZCAN_CANFDNET_400U_UDP    = 0x35
	ZCAN_CANFDBLUE_200U       = 0x36
	ZCAN_CANFDNET_100U_TCP    = 0x37
	ZCAN_CANFDNET_100U_UDP    = 0x38
	ZCAN_CANFDNET_800U_TCP    = 0x39
	ZCAN_CANFDNET_800U_UDP    = 0x3a
	ZCAN_USBCANFD_800U        = 0x3b
	ZCAN_PCIE_CANFD_100U_EX   = 0x3c
	ZCAN_PCIE_CANFD_400U_EX   = 0x3d
	ZCAN_PCIE_CANFD_200U_MINI = 0x3e
	ZCAN_PCIE_CANFD_200U_M2   = 0x3f
	ZCAN_CANFDDTU_400_TCP     = 0x40
	ZCAN_CANFDDTU_400_UDP     = 0x41
	ZCAN_CANFDWIFI_200U_TCP   = 0x42
	ZCAN_CANFDWIFI_200U_UDP   = 0x43
	ZCAN_OFFLINE_DEVICE       = 0x62
	ZCAN_VIRTUAL_DEVICE       = 0x63
)

const (
	ZCAN_STATUS_ERR         = 0
	ZCAN_STATUS_OK          = 1
	ZCAN_STATUS_ONLINE      = 2
	ZCAN_STATUS_OFFLINE     = 3
	ZCAN_STATUS_UNSUPPORTED = 4
)

const (
	ZCAN_TYPE_CAN   = 0x0
	ZCAN_TYPE_CANFD = 0x1
)

const (
	INVALID_DEVICE_HANDLE  = 0
	INVALID_CHANNEL_HANDLE = 0
)

type ZCAN_DEVICE_INFO struct {
	hw_Version     uint16
	fw_Version     uint16
	dr_Version     uint16
	in_Version     uint16
	irq_Num        uint16
	can_Num        uint8
	str_Serial_Num [20]uint8
	str_hw_Type    [40]uint8
	reserved       [8]uint16
}

func (info *ZCAN_DEVICE_INFO) _version(version uint16) string {
	if version/0xFF >= 9 {
		return fmt.Sprintf("V%02x.%02x", version>>8, version&0xff)
	}
	return fmt.Sprintf("V%d.%02x", version>>8, version&0xff)
}

func (info *ZCAN_DEVICE_INFO) HwVersion() string {
	return info._version(info.hw_Version)
}

func (info *ZCAN_DEVICE_INFO) FwVersion() string {
	return info._version(info.fw_Version)
}

func (info *ZCAN_DEVICE_INFO) DrVersion() string {
	return info._version(info.dr_Version)
}

func (info *ZCAN_DEVICE_INFO) InVersion() string {
	return info._version(info.in_Version)
}

func (info *ZCAN_DEVICE_INFO) IrqNum() uint16 {
	return info.irq_Num
}

func (info *ZCAN_DEVICE_INFO) CanNum() uint8 {
	return info.can_Num
}

func (info *ZCAN_DEVICE_INFO) Serial() string {
	serial := ""
	for c := range info.str_Serial_Num {
		if info.str_Serial_Num[c] > 0 {
			serial += string(info.str_Serial_Num[c])
		} else {
			break
		}
	}
	return serial
}

func (info *ZCAN_DEVICE_INFO) HwType() string {
	hwType := ""
	for c := range info.str_hw_Type {
		if info.str_hw_Type[c] > 0 {
			hwType += string(info.str_hw_Type[c])
		} else {
			break
		}
	}
	return hwType
}

type _ZCAN_CHANNEL_CAN_INIT_CONFIG struct {
	AccCode  uint32
	AccMask  uint32
	Reserved uint32
	Filter   uint8
	Timing0  uint8
	Timing1  uint8
	Mode     uint8
}
type _ZCAN_CHANNEL_CANFD_INIT_CONFIG struct {
	AccCode    uint32
	AccMask    uint32
	AbitTiming uint32
	DbitTiming uint32
	Brp        uint32
	Filter     uint8
	Mode       uint8
	Pad        uint16
	Reserved   uint32
}

// type _ZCAN_CHANNEL_INIT_CONFIG struct {
//     Can   _ZCAN_CHANNEL_CAN_INIT_CONFIG
//     Canfd _ZCAN_CHANNEL_CANFD_INIT_CONFIG
// }

type ZCAN_NORMAL_CHANNEL_INIT_CONFIG struct {
	CanType uint32
	Config  _ZCAN_CHANNEL_CAN_INIT_CONFIG
}
type ZCAN_CANFD_CHANNEL_INIT_CONFIG struct {
	CanType uint32
	Config  _ZCAN_CHANNEL_CANFD_INIT_CONFIG
}
type ZCAN_CHANNEL_ERR_INFO struct {
	ErrorCode      uint32
	PassiveErrData [3]uint8
	ArLostErrData  uint8
}
type ZCAN_CHANNEL_STATUS struct {
	ErrInterrupt uint8
	RegMode      uint8
	RegStatus    uint8
	RegALCapture uint8
	RegECCapture uint8
	RegEWLimit   uint8
	RegRECounter uint8
	RegTECounter uint8
	Reserved     uint32
}
type ZCAN_CAN_FRAME struct {
	Id      uint32
	Dlc     uint8
	X__pad  uint8
	X__res0 uint8
	X__res1 uint8
	Data    [8]uint8
}

func (f *ZCAN_CAN_FRAME) GenerateID(can_id uint32, err, rtr, eff uint8) {
	id := ((can_id << 3) | (uint32(err) << 2) | (uint32(rtr) << 1) | uint32(eff))
	f.Id = id
}

func (f *ZCAN_CAN_FRAME) GetFrameID() uint32 {
	return f.Id >> 3
}

func (f *ZCAN_CAN_FRAME) GetFrameERR() uint8 {
	return uint8(f.Id&0x04) >> 2
}

func (f *ZCAN_CAN_FRAME) GetFrameRTR() uint8 {
	return uint8(f.Id&0x02) >> 1
}

func (f *ZCAN_CAN_FRAME) GetFrameEFF() uint8 {
	return uint8(f.Id & 0x01)
}

type ZCAN_CANFD_FRAME struct {
	Id      uint32
	Len     uint8
	Flags   uint8
	X__res0 uint8
	X__res1 uint8
	Data    [64]uint8
}

func (f *ZCAN_CANFD_FRAME) GenerateID(can_id uint32, err, rtr, eff uint8) {
	id := ((can_id << 3) | (uint32(err) << 2) | (uint32(rtr) << 1) | uint32(eff))
	f.Id = id
}

func (f *ZCAN_CANFD_FRAME) GetFrameID() uint32 {
	return f.Id >> 3
}

func (f *ZCAN_CANFD_FRAME) GetFrameERR() uint8 {
	return uint8(f.Id&0x04) >> 2
}

func (f *ZCAN_CANFD_FRAME) GetFrameRTR() uint8 {
	return uint8(f.Id&0x02) >> 1
}

func (f *ZCAN_CANFD_FRAME) GetFrameEFF() uint8 {
	return uint8(f.Id & 0x01)
}

func (f *ZCAN_CANFD_FRAME) GenerateFlags(brs, esi, res uint8) {
	flags := ((brs << 7) | (esi << 6) | res)
	f.Flags = flags
}

func (f *ZCAN_CANFD_FRAME) GetFrameBRS() uint8 {
	return uint8(f.Flags&0x80) >> 7
}

func (f *ZCAN_CANFD_FRAME) GetFrameESI() uint8 {
	return uint8(f.Flags&0x40) >> 6
}

func (f *ZCAN_CANFD_FRAME) GetFrameRES() uint8 {
	return uint8(f.Flags & 0x3F)
}

type ZCAN_Transmit_Data struct {
	Frame ZCAN_CAN_FRAME
	Type  uint32
}
type ZCAN_Receive_Data struct {
	Frame     ZCAN_CAN_FRAME
	Timestamp uint64
}
type ZCAN_TransmitFD_Data struct {
	Frame ZCAN_CANFD_FRAME
	Type  uint32
}
type ZCAN_ReceiveFD_Data struct {
	Frame     ZCAN_CANFD_FRAME
	Timestamp uint64
}
type ZCAN_AUTO_TRANSMIT_OBJ struct {
	Enable   uint16
	Index    uint16
	Interval uint32
	Obj      ZCAN_Transmit_Data
}
type ZCANFD_AUTO_TRANSMIT_OBJ struct {
	Enable   uint16
	Index    uint16
	Interval uint32
	Obj      ZCAN_TransmitFD_Data
}
type ZCAN_IProperty struct {
	SetValue     *[0]byte
	GetValue     *[0]byte
	GetPropertys *[0]byte
}

type ZCAN struct {
	dll syscall.Handle
}

func NewZCAN(dllPath string) (*ZCAN, error) {
	var dll syscall.Handle
	if runtime.GOOS == "windows" {
		dll, _ = syscall.LoadLibrary(".\\zlgcan_x64\\zlgcan.dll")
	} else {
		fmt.Println("No support now!")
		return nil, fmt.Errorf("unsupported OS")
	}
	return &ZCAN{dll: dll}, nil
}

func (zc *ZCAN) OpenDevice(deviceType int, deviceIndex int, reserved int) int {
	if zc.dll == 0 {
		return -1
	}
	openDevice, _ := syscall.GetProcAddress(zc.dll, "ZCAN_OpenDevice")
	ret, _, _ := syscall.SyscallN(
		openDevice,
		uintptr(deviceType),
		uintptr(deviceIndex),
		uintptr(reserved))
	return int(ret)
}

func (zc *ZCAN) CloseDevice(deviceHandle int) int {
	if zc.dll == 0 {
		return -1
	}
	closeDevice, _ := syscall.GetProcAddress(zc.dll, "ZCAN_CloseDevice")
	ret, _, _ := syscall.SyscallN(closeDevice, uintptr(deviceHandle))
	return int(ret)
}

func (zc *ZCAN) GetDeviceInf(deviceHandle int) *ZCAN_DEVICE_INFO {
	info := ZCAN_DEVICE_INFO{}
	getDeviceInf, _ := syscall.GetProcAddress(zc.dll, "ZCAN_GetDeviceInf")
	ret, _, _ := syscall.SyscallN(getDeviceInf, uintptr(deviceHandle), uintptr(unsafe.Pointer(&info)))
	if ret == ZCAN_STATUS_OK {
		return (*ZCAN_DEVICE_INFO)(&info)
	}
	fmt.Println("error calling ZCAN_GetDeviceInf:", ret)
	return nil
}

func (zc *ZCAN) IsDeviceOnLine(deviceHandle int) int {
	isDeviceOnline, _ := syscall.GetProcAddress(zc.dll, "ZCAN_IsDeviceOnLine")
	ret, _, _ := syscall.SyscallN(isDeviceOnline, uintptr(deviceHandle))
	return int(ret)
}

func (zc *ZCAN) InitCAN(deviceHandle int, canIndex uint, initConfig *ZCAN_NORMAL_CHANNEL_INIT_CONFIG) int {
	initCAN, _ := syscall.GetProcAddress(zc.dll, "ZCAN_InitCAN")
	ret, _, _ := syscall.SyscallN(initCAN, uintptr(deviceHandle), uintptr(canIndex), uintptr(unsafe.Pointer(initConfig)))
	return int(ret)
}

func (zc *ZCAN) InitCANFD(deviceHandle int, canIndex uint, initConfig *ZCAN_CANFD_CHANNEL_INIT_CONFIG) int {
	initCAN, _ := syscall.GetProcAddress(zc.dll, "ZCAN_InitCAN")
	ret, _, _ := syscall.SyscallN(initCAN, uintptr(deviceHandle), uintptr(canIndex), uintptr(unsafe.Pointer(initConfig)))
	return int(ret)
}

func (zc *ZCAN) StartCAN(channelHandle int) uint {
	startCAN, _ := syscall.GetProcAddress(zc.dll, "ZCAN_StartCAN")
	ret, _, _ := syscall.SyscallN(startCAN, uintptr(channelHandle))
	return uint(ret)
}

func (zc *ZCAN) ResetCAN(channelHandle int) uint {
	resetCAN, _ := syscall.GetProcAddress(zc.dll, "ZCAN_ResetCAN")
	ret, _, _ := syscall.SyscallN(resetCAN, uintptr(channelHandle))
	return uint(ret)
}

func (zc *ZCAN) ClearBuffer(channelHandle int) uint {
	clearBuffer, _ := syscall.GetProcAddress(zc.dll, "ZCAN_ClearBuffer")
	ret, _, _ := syscall.SyscallN(clearBuffer, uintptr(channelHandle))
	return uint(ret)
}

func (zc *ZCAN) ReadChannelErrInfo(channelHandle int) (*ZCAN_CHANNEL_ERR_INFO, error) {
	errInfo := ZCAN_CHANNEL_ERR_INFO{}
	readChannelErrInfo, _ := syscall.GetProcAddress(zc.dll, "ZCAN_ReadChannelErrInfo")
	ret, _, _ := syscall.SyscallN(readChannelErrInfo, uintptr(channelHandle), uintptr(unsafe.Pointer(&errInfo)))
	if ret == ZCAN_STATUS_OK {
		return (*ZCAN_CHANNEL_ERR_INFO)(&errInfo), nil
	}
	fmt.Println("error calling ZCAN_ReadChannelErrInfo:", ret)
	return nil, fmt.Errorf("error calling ZCAN_ReadChannelErrInfo")
}

func (zc *ZCAN) ReadChannelStatus(channelHandle int) (*ZCAN_CHANNEL_STATUS, error) {
	status := ZCAN_CHANNEL_STATUS{}
	readChannelStatus, _ := syscall.GetProcAddress(zc.dll, "ZCAN_ReadChannelStatus")
	ret, _, _ := syscall.SyscallN(readChannelStatus, uintptr(channelHandle), uintptr(unsafe.Pointer(&status)))
	if ret == ZCAN_STATUS_OK {
		return (*ZCAN_CHANNEL_STATUS)(&status), nil
	}
	fmt.Println("error calling ZCAN_ReadChannelStatus:", ret)
	return nil, fmt.Errorf("error calling ZCAN_ReadChannelStatus")
}

func (zc *ZCAN) GetReceiveNum(channelHandle int, canType uint) uint {
	getReceiveNum, _ := syscall.GetProcAddress(zc.dll, "ZCAN_GetReceiveNum")
	ret, _, _ := syscall.SyscallN(getReceiveNum, uintptr(channelHandle), uintptr(canType))
	return uint(ret)
}

func (zc *ZCAN) Transmit(channelHandle int, stdMsg []ZCAN_Transmit_Data, len uint) uint {
	transmit, _ := syscall.GetProcAddress(zc.dll, "ZCAN_Transmit")
	ret, _, _ := syscall.SyscallN(transmit, uintptr(channelHandle), uintptr(unsafe.Pointer(&stdMsg[0])), uintptr(len))
	return uint(ret)
}

func (zc *ZCAN) Receive(channelHandle int, rcvNum uint, waitTime int) ([]ZCAN_Receive_Data, uint) {
	msgs := make([]ZCAN_Receive_Data, rcvNum)
	receive, _ := syscall.GetProcAddress(zc.dll, "ZCAN_Receive")
	ret, _, _ := syscall.SyscallN(receive, uintptr(channelHandle), uintptr(unsafe.Pointer(&msgs[0])), uintptr(rcvNum), uintptr(waitTime))
	return msgs, uint(ret)
}

func (zc *ZCAN) TransmitFD(channelHandle int, fdMsg []ZCAN_TransmitFD_Data, len uint) uint {
	transmitFD, _ := syscall.GetProcAddress(zc.dll, "ZCAN_TransmitFD")
	ret, _, _ := syscall.SyscallN(transmitFD, uintptr(channelHandle), uintptr(unsafe.Pointer(&fdMsg[0])), uintptr(len))
	return uint(ret)
}

func (zc *ZCAN) ReceiveFD(channelHandle int, rcvNum uint, waitTime int) ([]ZCAN_ReceiveFD_Data, uint) {
	if waitTime == 0 {
		waitTime = -1
	}
	msgs := make([]ZCAN_ReceiveFD_Data, rcvNum)
	receiveFD, _ := syscall.GetProcAddress(zc.dll, "ZCAN_ReceiveFD")
	ret, _, _ := syscall.SyscallN(receiveFD, uintptr(channelHandle), uintptr(unsafe.Pointer(&msgs[0])), uintptr(rcvNum), uintptr(waitTime))
	return msgs, uint(ret)
}

func (zc *ZCAN) GetIProperty(deviceHandle int) (*ZCAN_IProperty, error) {
	getIProperty, _ := syscall.GetProcAddress(zc.dll, "GetIProperty")
	ret, _, callErr := syscall.SyscallN(getIProperty, uintptr(deviceHandle))
	if callErr != 0 {
		return nil, fmt.Errorf("error calling GetIProperty: %w", callErr)
	}
	// transform the ret to a pointer for ZCAN_IProperty
	iproperty := (*ZCAN_IProperty)(unsafe.Pointer(ret))
	return iproperty, nil
}

func (zc *ZCAN) SetValue(iproperty *ZCAN_IProperty, path, value string) uint {
	setValue := iproperty.SetValue
	cPath := C.CString(path)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cPath))
	defer C.free(unsafe.Pointer(cValue))
	ret, _, _ := syscall.SyscallN(uintptr(unsafe.Pointer(setValue)), uintptr(unsafe.Pointer(cPath)), uintptr(unsafe.Pointer(cValue)))
	return uint(ret)
}

func (zc *ZCAN) GetValue(iproperty *ZCAN_IProperty, path string) string {
	getValue := iproperty.GetValue
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	ret, _, _ := syscall.SyscallN(uintptr(unsafe.Pointer(getValue)), uintptr(unsafe.Pointer(cPath)))
	return C.GoString((*C.char)(unsafe.Pointer(ret)))
}

func (zc *ZCAN) ReleaseIProperty(iproperty *ZCAN_IProperty) uint {
	proc, _ := syscall.GetProcAddress(zc.dll, "ReleaseIProperty")
	ret, _, _ := syscall.SyscallN(proc, uintptr(unsafe.Pointer(iproperty)))
	return uint(ret)
}

func can_start(zcanlib *ZCAN, handle int, channel int) int {
	ip, _ := zcanlib.GetIProperty(handle)
	ret := zcanlib.SetValue(ip, "/initenal_resistance", "1")
	if ret != ZCAN_STATUS_OK {
		fmt.Println("Set resistance failed")
	}
	ret = zcanlib.SetValue(ip, fmt.Sprintf("%d/clock", channel), "60000000")
	if ret != ZCAN_STATUS_OK {
		fmt.Println("Set clock failed")
	}
	zcanlib.ReleaseIProperty(ip)

	initCfg := ZCAN_CANFD_CHANNEL_INIT_CONFIG{
		CanType: ZCAN_TYPE_CANFD,
	}
	initCfg.Config.Mode = 0
	initCfg.Config.AbitTiming = 101166  // 1Mbps
	initCfg.Config.DbitTiming = 8487694 // 1Mbps

	can_handle := zcanlib.InitCANFD(handle, 0, &initCfg)
	if can_handle == INVALID_CHANNEL_HANDLE {
		fmt.Println("Init CANFD failed")
	}

	_ = zcanlib.StartCAN(can_handle)
	return can_handle
}

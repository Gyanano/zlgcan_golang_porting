package zlgcan

import (
	"fmt"
	"syscall"
	"testing"
)

// Test for Open&Close
func TestOpenAndCloseDevice(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	ret := zcanlib.CloseDevice(handle)
	if ret != 1 {
		t.Fatalf("Close Device failed! ret: %v", ret)
	}
}

// Test for GetInfo
func TestGetDeviceInfo(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	deviceInfo := zcanlib.GetDeviceInf(handle)
	t.Logf("Device Information:\n%+v\n", deviceInfo)
	t.Logf("Device Information:\n%+v\n", deviceInfo)
	t.Logf("Device hw_Version: %s\n", deviceInfo.HwVersion())
	t.Logf("Device fw_Version: %s\n", deviceInfo.FwVersion())
	t.Logf("Device dr_Version: %s\n", deviceInfo.DrVersion())
	t.Logf("Device in_Version: %s\n", deviceInfo.InVersion())
	t.Logf("Device irq_Num: %d\n", deviceInfo.IrqNum())
	t.Logf("Device can_Num: %d\n", deviceInfo.CanNum())
	t.Logf("Device str_Serial_Num: %s\n", deviceInfo.Serial())
	t.Logf("Device str_hw_Type: %s\n", deviceInfo.HwType())

	ret := zcanlib.CloseDevice(handle)
	if ret != 1 {
		t.Fatalf("Close Device failed! ret: %v", ret)
	}
}

// Test for DeviceOnLine
func TestDeviceOnLine(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	ret := zcanlib.IsDeviceOnLine(handle)
	// If the device is not online, the value will not be 2.
	if ret != 2 {
		t.Fatalf("Device OnLine failed! ret: %v", ret)
		return
	}
	// t.Logf("Device OnLine: %d\n", ret)

	ret = zcanlib.CloseDevice(handle)
	if ret != 1 {
		t.Fatalf("Close Device failed! ret: %v", ret)
	}
}

// Test for InitCAN(Channel Initialize)
func TestInitCAN(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	initCfg := ZCAN_CANFD_CHANNEL_INIT_CONFIG{
		CanType: ZCAN_TYPE_CAN, // use normal mode because the fd mode need to set the clock
	}
	initCfg.Config.Mode = 0

	ret := zcanlib.InitCANFD(handle, 0, &initCfg)
	if ret == INVALID_CHANNEL_HANDLE {
		t.Fatalf("The result of initializing the Channel 0 is %d", ret)
		return
	}
}

// Test for GetIProperty and ReleaseIProperty
func TestGetAndReleaseIProperty(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	ip, err := zcanlib.GetIProperty(handle)
	if err != nil {
		t.Fatalf("GetIProperty failed: %v", err)
		return
	}
	defer zcanlib.ReleaseIProperty(ip)
	t.Logf("GetIProperty: %+v\n", ip)
}

// Test for Open&Close CAN Channel
func TestOpenAndCloseCAN(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)

	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		t.Fatalf("Open Device failed: %v", err)
		return
	}

	ip, err := zcanlib.GetIProperty(handle)
	if err != nil {
		t.Fatalf("GetIProperty failed: %v", err)
		return
	}
	ret := zcanlib.SetValue(ip, "/initenal_resistance", "1")
	if ret != ZCAN_STATUS_OK {
		t.Fatalf("SetValue failed! ret: %v", ret)
		return
	}
	ret = zcanlib.SetValue(ip, "0/clock", "60000000")
	if ret != ZCAN_STATUS_OK {
		t.Fatalf("SetValue failed! ret: %v", ret)
		return
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
		t.Fatalf("The result of initializing the Channel 0 is %d", can_handle)
		return
	}

	ret = zcanlib.StartCAN(can_handle)
	if ret != ZCAN_STATUS_OK {
		t.Fatalf("StartCAN failed! ret: %v", ret)
		return
	}
	// t.Logf("StartCAN success! ret: %v", ret)
	// time.Sleep(time.Second * 1)
	ret = zcanlib.ResetCAN(can_handle)
	if ret != ZCAN_STATUS_OK {
		t.Fatalf("ResetCAN failed! ret: %v", ret)
		return
	}
	// t.Logf("ResetCAN success! ret: %v", ret)
}

// Test for Transmit&Receive
func TestTransmitAndReceive(t *testing.T) {
	zcanlib, err := NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
	if err != nil {
		t.Fatalf("Failed to load ZCAN DLL: %v", err)
		return
	}
	defer syscall.FreeLibrary(zcanlib.dll)
	handle := zcanlib.OpenDevice(ZCAN_USBCANFD_200U, 0, 0)
	if handle == INVALID_DEVICE_HANDLE {
		fmt.Println("Open Device failed!")
		return
	}
	t.Logf("device handle:%d.\n", handle)

	info := zcanlib.GetDeviceInf(handle)
	if info == nil {
		t.Fatalf("Get device info failed!")
		return
	}
	t.Logf("Device Information:\n%+v\n", info)

	chanHandle := can_start(zcanlib, handle, 0)
	print("channel handle:", chanHandle, "\n")

	// Send CAN Messages
	transmitNum := 10
	msgs := make([]ZCAN_Transmit_Data, transmitNum)
	for msg_id := range msgs {
		msgs[msg_id].Type = 2
		// 生成 ID
		msgs[msg_id].Frame.GenerateID(uint32(msg_id), 0, 0, 0)
		msgs[msg_id].Frame.Dlc = 8
		for j := 0; j < 8; j++ {
			msgs[msg_id].Frame.Data[j] = uint8(j)
		}
	}
	ret := zcanlib.Transmit(chanHandle, msgs, uint(transmitNum))
	t.Logf("Transmit Num: %d.\n", ret)

	// Send CANFD Messages
	transmitNum = 10
	msgs_fd := make([]ZCAN_TransmitFD_Data, transmitNum)
	for msg_id := range msgs_fd {
		msgs_fd[msg_id].Type = 2
		// 生成 ID
		msgs_fd[msg_id].Frame.GenerateID(uint32(msg_id), 0, 0, 0)
		msgs_fd[msg_id].Frame.GenerateFlags(1, 0, 0)
		msgs_fd[msg_id].Frame.Len = 8
		for j := 0; j < 8; j++ {
			msgs_fd[msg_id].Frame.Data[j] = uint8(j)
		}
	}
	ret = zcanlib.TransmitFD(chanHandle, msgs_fd, uint(transmitNum))
	t.Logf("Transmit FD Num: %d.\n", ret)

	// Receive CAN Messages
	for {
		rcv_num := zcanlib.GetReceiveNum(chanHandle, ZCAN_TYPE_CAN)
		rcv_num_fd := zcanlib.GetReceiveNum(chanHandle, ZCAN_TYPE_CANFD)
		if rcv_num > 0 {
			var rcv_msg []ZCAN_Receive_Data
			t.Logf("Receive CAN Num: %d.\n", rcv_num)
			rcv_msg, rcv_num = zcanlib.Receive(chanHandle, rcv_num, 0)
			for msg_id := range rcv_msg {
				var hex_str string = ""
				for b := range rcv_msg[msg_id].Frame.Data {
					hex_str += fmt.Sprintf("%02x ", rcv_msg[msg_id].Frame.Data[b])
				}
				t.Logf(
					"[%d]:ts:%d, id:%d, dlc:%d, eff:%d, rtr:%d, err:%d, data:%s\n",
					msg_id,
					rcv_msg[msg_id].Timestamp,
					rcv_msg[msg_id].Frame.GetFrameID(), // because the Frame.Id has three bits for the eff, rtr and err
					rcv_msg[msg_id].Frame.Dlc,
					rcv_msg[msg_id].Frame.GetFrameEFF(),
					rcv_msg[msg_id].Frame.GetFrameRTR(),
					rcv_msg[msg_id].Frame.GetFrameERR(),
					hex_str,
				)
			}
		} else if rcv_num_fd > 0 {
			var rcv_msg []ZCAN_ReceiveFD_Data
			t.Logf("Receive FD Num: %d.\n", rcv_num_fd)
			rcv_msg, _ = zcanlib.ReceiveFD(chanHandle, rcv_num_fd, 1000)
			for msg_id := range rcv_msg {
				var hex_str string = ""
				for b := range rcv_msg[msg_id].Frame.Data {
					hex_str += fmt.Sprintf("%02x ", rcv_msg[msg_id].Frame.Data[b])
				}
				t.Logf(
					"[%d]:ts:%d, id:%d, len:%d, eff:%d, rtr:%d, err:%d, esi:%d, brs:%d, data:%s\n",
					msg_id,
					rcv_msg[msg_id].Timestamp,
					rcv_msg[msg_id].Frame.GetFrameID(),
					rcv_msg[msg_id].Frame.Len,
					rcv_msg[msg_id].Frame.GetFrameEFF(),
					rcv_msg[msg_id].Frame.GetFrameRTR(),
					rcv_msg[msg_id].Frame.GetFrameERR(),
					rcv_msg[msg_id].Frame.GetFrameESI(),
					rcv_msg[msg_id].Frame.GetFrameBRS(),
					hex_str,
				)
			}
		} else {
			break
		}

	}

	// Close CAN
	zcanlib.ResetCAN(chanHandle)
	// Close Device
	zcanlib.CloseDevice(handle)
	t.Log("Close Device success!")
}

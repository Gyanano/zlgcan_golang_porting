# ZLGCAN Go Wrapper

This project is a Go language wrapper for ZLG's CAN box driver. It provides a Go interface for interacting with ZLG CAN devices.

## Features

- Support for opening and closing CAN devices
- Retrieving device information
- Checking device online status
- Initializing CAN and CANFD channels
- Sending and receiving CAN/CANFD messages
- Getting and setting device properties

## Installation

Ensure you have Go installed on your system. Then, clone this repository:

```
git clone https://github.com/Gyanano/zlgcan-go.git
```

## Usage

1. Import the package:

```go
import "github.com/Gyanano/zlgcan-go"
```

2. Create a ZCAN instance:

```go
zcanlib, err := zlgcan.NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
if err != nil {
    // Handle error
}
```

3. Open a device:

```go
handle := zcanlib.OpenDevice(zlgcan.ZCAN_USBCANFD_200U, 0, 0)
if handle == zlgcan.INVALID_DEVICE_HANDLE {
    // Handle error
}
```

4. Use other functions, such as sending and receiving messages:

```go
// Send CAN message
msgs := make([]zlgcan.ZCAN_Transmit_Data, 1)
// Set message content
ret := zcanlib.Transmit(chanHandle, msgs, 1)

// Receive CAN message
rcvNum := zcanlib.GetReceiveNum(chanHandle, zlgcan.ZCAN_TYPE_CAN)
if rcvNum > 0 {
    rcvMsg, _ := zcanlib.Receive(chanHandle, rcvNum, 0)
    // Process received messages
}
```

5. Close the device:

```go
zcanlib.CloseDevice(handle)
```

## Testing

The project includes a series of unit tests covering the main functionalities. To run the tests:

```
go test
```

## Notes

- This project has only been tested in a Windows environment.
- Ensure that ZLG's CAN device drivers are properly installed.
- Please carefully read ZLG's original documentation to understand the specific uses and parameter meanings of each function before use.

## Contributing

Issues and pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
# ZLGCAN Go 封装

这个项目是ZLG公司CAN盒驱动程序的Go语言封装。它提供了一个Go语言接口,用于与ZLG的CAN设备进行交互。

## 功能特性

- 支持打开和关闭CAN设备
- 获取设备信息
- 检查设备在线状态
- 初始化CAN和CANFD通道
- 发送和接收CAN/CANFD消息
- 获取和设置设备属性

## 安装

确保你的系统中已安装Go语言环境。然后,克隆此仓库:

```
git clone https://github.com/your-username/zlgcan-go.git
```

## 使用方法

1. 导入包:

```go
import "github.com/your-username/zlgcan-go"
```

2. 创建ZCAN实例:

```go
zcanlib, err := zlgcan.NewZCAN(".\\zlgcan_x64\\zlgcan.dll")
if err != nil {
    // 处理错误
}
```

3. 打开设备:

```go
handle := zcanlib.OpenDevice(zlgcan.ZCAN_USBCANFD_200U, 0, 0)
if handle == zlgcan.INVALID_DEVICE_HANDLE {
    // 处理错误
}
```

4. 使用其他功能,如发送和接收消息:

```go
// 发送CAN消息
msgs := make([]zlgcan.ZCAN_Transmit_Data, 1)
// 设置消息内容
ret := zcanlib.Transmit(chanHandle, msgs, 1)

// 接收CAN消息
rcvNum := zcanlib.GetReceiveNum(chanHandle, zlgcan.ZCAN_TYPE_CAN)
if rcvNum > 0 {
    rcvMsg, _ := zcanlib.Receive(chanHandle, rcvNum, 0)
    // 处理接收到的消息
}
```

5. 关闭设备:

```go
zcanlib.CloseDevice(handle)
```

## 测试

项目包含了一系列单元测试,涵盖了主要功能。运行测试:

```
go test
```

## 注意事项

- 此项目仅在Windows环境下测试过。
- 确保ZLG的CAN设备驱动程序已正确安装。
- 使用前请仔细阅读ZLG原始文档,了解各函数的具体用途和参数含义。

## 贡献
欢迎提交问题和拉取请求。对于重大更改,请先开issue讨论您想要改变的内容。

## 许可证
[MIT](https://choosealicense.com/licenses/mit/)
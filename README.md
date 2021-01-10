# raspberry
raspberry相关开发
## 目录
1. fanAccordCPU 根据温度激活树莓派的风扇，需要使用GPIO引脚进行供电，如果直插，考虑使用电压3V3的风扇
2. LED 01Studio PiHAT 灯组测试
2. BUTTON 01Studio PiHAT 按钮针脚状态变更测试
***
## 注意
1. 引脚有三种编码方式：
    > - BCM编码——GPIO2
    > - WiringPi编码——WiringPi8
    > - Board编码——3或者P1_3
    > - 注意程序库使用的哪一种，需要一一对应。一般使用BCM编码即可。
2. 第二需要注意的是GPIO操作**需要root权限**，否则有可能在不抛出错误情况下照常运行。
***
## 参考
[GPIO引脚定义](https://pinout.xyz/)  
![RPI.GPIO](https://raw.githubusercontent.com/kintansky/raspberry/main/PinOut.png)  


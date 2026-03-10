package pack

import "sync"

var once sync.Once

// Init 初始化并注册所有服务
func Init() {
	once.Do(func() {
		// 注册 Jio
		Register(NewJioH5PackService())
		// Register(NewJioPCPackService())   // 未来扩展
		// Register(NewJioIOSPackService())  // 未来扩展

		// 注册 Airtel
		Register(NewAirtelH5PackService())
		// Register(NewAirtelPCPackService())  // 未来扩展

		Register(NewVIPackService())
	})
}

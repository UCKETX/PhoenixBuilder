package Happy2018new_depends

// 描述玩家的游戏名称
type PlayerName string

// 描述多个玩家的坐标和朝向信息
type PlayersPosInfo map[PlayerName]PosInfo

// 描述单个玩家的坐标和朝向信息
type PosInfo struct {
	Dimension byte       // 玩家所处的维度
	Position  [3]float32 // 玩家的位置
	YRot      float32    // 玩家的偏航角
}

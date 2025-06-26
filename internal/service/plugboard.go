package service

// プラグボードの設定
type PlugboardConfig struct {
	PlugboardConfig []string
}

type Plugboard struct {
	plugboard map[string]string
}

// プラグボードの設定をmap形式に変換
func PlugboardJsonToMap(plugboardConfig PlugboardConfig) *Plugboard {
	pb := &Plugboard{
		plugboard: make(map[string]string),
	}
	for _, pairs := range plugboardConfig.PlugboardConfig {
		pair0 := string(pairs[0])
		pair1 := string(pairs[1])
		pb.plugboard[pair0] = pair1
		pb.plugboard[pair1] = pair0
	}
	return pb
}

// プラグボードで文字を変換
func (pb *Plugboard) PlugboardTransform(input string) string {
	if mapped, ok := pb.plugboard[input]; ok {
		return mapped
	}
	return input
}
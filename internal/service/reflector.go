package service

// リフレクターの設定
type ReflectorConfig struct {
	reflectorConfig []string
}

type Reflector struct {
	reflector map[string]string
}

// リフレクターの設定をmap形式に変換
func ReflectorJsonToMap(config ReflectorConfig) *Reflector {
	ref := &Reflector{
		reflector: make(map[string]string),
	}
	for _, pairs := range config.reflectorConfig {
		pair0 := string(pairs[0])
		pair1 := string(pairs[1])
		ref.reflector[pair0] = pair1
		ref.reflector[pair1] = pair0
	}
	return ref
}

// リフレクターで文字を変換
func (ref *Reflector) ReflectorTransform(input string) string {
	return ref.reflector[input]
}
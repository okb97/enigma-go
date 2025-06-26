package service

// ローターの各タイプの配線
var rotorWiringPatterns = map[string]string{
	"I":   "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"II":  "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"III": "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"IV":  "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"V":   "VZBRGITYUPSDNHLXAWMJQOFECK",
}

// ローターのノッチ位置
var rotorNotchPositions = map[string]string{
	"I":   "Q",
	"II":  "E",
	"III": "V",
	"IV":  "J",
	"V":   "Z",
}

// ローターの設定
type RotorConfig struct {
	RotorType     string
	RotorPosition string
	RingSetting   int
}

type Rotor struct {
	rotorType       string
	forwardWiring   map[string]string
	backwardWiring  map[string]string
	notchPosition   string
	currentPosition int
	ringSetting     int
}

// ローターを初期化
func InitialRotor(config RotorConfig) *Rotor {
	notchPosition := rotorNotchPositions[config.RotorType]
	forwardWiring := buildForwardWiring(config.RotorType)
	backwardWiring := buildBackwardWiring(config.RotorType)
	currentPosition := int(config.RotorPosition[0] - 'A')
	return &Rotor{
		rotorType:       config.RotorType,
		forwardWiring:   forwardWiring,
		backwardWiring:  backwardWiring,
		notchPosition:   notchPosition,
		currentPosition: currentPosition,
		ringSetting:     config.RingSetting,
	}

}

// 順方向の配線をMap形式に変換
func buildForwardWiring(rotorType string) map[string]string {
	wiringPosition := rotorWiringPatterns[rotorType]
	forwardWiringMap := make(map[string]string)
	for idx, _ := range wiringPosition {
		forwardWiringMap[string(rune('A'+idx))] = string(wiringPosition[idx])
	}
	return forwardWiringMap
}

// 逆方向の配線をMap形式に変換
func buildBackwardWiring(rotorType string) map[string]string {
	wiringPosition := rotorWiringPatterns[rotorType]
	backwardWiringMap := make(map[string]string)
	for idx, _ := range wiringPosition {
		backwardWiringMap[string(wiringPosition[idx])] = string(rune('A' + idx))
	}
	return backwardWiringMap
}

// 順方向のローターの変換
func (r *Rotor) ForwardTransform(input string) string {
	currentPos := r.currentPosition
	inputPos := int(input[0] - 'A')
	transformPos := (inputPos + currentPos - r.ringSetting + 26) % 26
	transformString := string(rune('A') + rune(transformPos))
	transformdString := r.forwardWiring[transformString]
	transformdPos := int(transformdString[0] - 'A')
	mappedPos := (transformdPos - currentPos + r.ringSetting + 26) % 26
	return string(rune('A') + rune(mappedPos))
}

// 逆方向のローターの変換
func (r *Rotor) BackwardTransform(input string) string {
	currentPos := r.currentPosition
	inputPos := int(rune(input[0]) - 'A')
	transformPos := (inputPos + currentPos - r.ringSetting + 26) % 26
	transformString := string(rune('A') + rune(transformPos))
	transformdString := r.backwardWiring[transformString]
	transformdPos := int(transformdString[0] - 'A')
	mappedPos := (transformdPos - currentPos + r.ringSetting + 26) % 26
	return string(rune('A') + rune(mappedPos))
}

// ローターを回転
func (r *Rotor) Rotate() {
	r.currentPosition = (r.currentPosition + 1) % 26
}

func (r *Rotor) GetPosition() string {
	return string(rune('A') + rune(r.currentPosition))
}

func (r *Rotor) IsAtNotch() bool {
	currenChar := string(rune('A') + rune(r.currentPosition))
	return currenChar == r.notchPosition
}

// 位置を設定
func (r *Rotor) SetPosition(position string) {
	r.currentPosition = int(position[0] - 'A')
}

// 初期位置にリセット
func (r *Rotor) Reset() {
	r.currentPosition = 0
}

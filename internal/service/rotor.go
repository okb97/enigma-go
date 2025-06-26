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
	rotorType     string
	rotorPosition string
	ringSetting   int
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
	notchPosition := rotorNotchPositions[config.rotorType]
	forwardWiring := buildForwardWiring(config.rotorType)
	backwardWiring := buildBackwardWiring(config.rotorType)
	currentPosition := 0
	return &Rotor{
		rotorType:       config.rotorType,
		forwardWiring:   forwardWiring,
		backwardWiring:  backwardWiring,
		notchPosition:   notchPosition,
		currentPosition: currentPosition,
		ringSetting:     config.ringSetting,
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
	return string('A' + mappedPos)
}

// 逆方向のローターの変換
func (r *Rotor) BackwardTransform(input string) string {
	currentPos := r.currentPosition
	inputPos := int(rune(input[0]) - 'A')
	transformPos := (inputPos - currentPos + r.ringSetting + 26) % 26
	transformString := string(rune('A') + rune(transformPos))
	transformdString := r.backwardWiring[transformString]
	transformdPos := int(transformdString[0] - 'A')
	mappedPos := (transformdPos + currentPos - r.ringSetting + 26) % 26
	return string('A' + mappedPos)
}

// ローターを回転
func (r *Rotor) Rotate() {
	r.currentPosition = (r.currentPosition + 1) % 26
}

func (r *Rotor) GetPosition() string {
	return string('A' + r.currentPosition)
}

func (r *Rotor) IsAtNotch() bool {
	currenChar := string('A' + r.currentPosition)
	return currenChar == r.notchPosition
}

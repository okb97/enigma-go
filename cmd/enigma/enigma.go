package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"enigma-go/internal/service"
)

// エニグママシンの設定
type EnigmaConfig struct {
	Name      string            `json:"name"`
	Rotors    []RotorConfigJSON `json:"rotors"`
	Reflector []string          `json:"reflector"`
	Plugboard []string          `json:"plugboard"`
}

// ローター設定（JSON用）
type RotorConfigJSON struct {
	Type        string `json:"type"`
	Position    string `json:"position"`
	RingSetting int    `json:"ring_setting"`
}

// エニグママシン本体
type EnigmaMachine struct {
	rotors    []*service.Rotor
	reflector *service.Reflector
	plugboard *service.Plugboard
}

func main() {
	// コマンドライン引数の定義
	var input = flag.String("input", "", "暗号化するメッセージ")
	var configFile = flag.String("config", "", "設定ファイルのパス")

	flag.Parse()

	enigmaConfig, err := LoadConfig(*configFile)
	if err != nil {
		fmt.Printf("設定ファイルの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	enigma := NewEnigmaMachine(*enigmaConfig)

	encrypted := enigma.encrypt(*input)
	fmt.Print(encrypted)
}

func (em *EnigmaMachine) encrypt(input string) string {
	// 前処理追加
	processedInput := preprocessMessage(input)

	encrypted := ""
	for _, char := range processedInput {
		if char >= 'A' && char <= 'Z' {
			encryptedChar := em.encryptChar(string(char))
			encrypted += encryptedChar
		}
	}
	return encrypted
}

func (em *EnigmaMachine) encryptChar(input string) string {
	// 1. ローター回転（暗号化前に実行）
	em.handleRotorStepping()

	// 2. プラグボード入力変換
	step1 := em.plugboard.PlugboardTransform(input)

	// 3. 順方向ローター変換（右→左）
	step2 := em.rotors[2].ForwardTransform(step1) // 右端ローター
	step3 := em.rotors[1].ForwardTransform(step2) // 中央ローター
	step4 := em.rotors[0].ForwardTransform(step3) // 左端ローター

	// 4. リフレクター変換
	step5 := em.reflector.ReflectorTransform(step4)

	// 5. 逆方向ローター変換（左→右）
	step6 := em.rotors[0].BackwardTransform(step5) // 左端ローター
	step7 := em.rotors[1].BackwardTransform(step6) // 中央ローター
	step8 := em.rotors[2].BackwardTransform(step7) // 右端ローター

	// 6. プラグボード出力変換
	result := em.plugboard.PlugboardTransform(step8)

	return result
}

// エニグママシンを初期化
func NewEnigmaMachine(enigmaConfig EnigmaConfig) *EnigmaMachine {
	plugboardConfig := service.PlugboardConfig{PlugboardConfig: enigmaConfig.Plugboard}
	plugboard := service.PlugboardJsonToMap(plugboardConfig)

	reflectorConfig := service.ReflectorConfig{ReflectorConfig: enigmaConfig.Reflector}
	reflector := service.ReflectorJsonToMap(reflectorConfig)

	rotors := make([]*service.Rotor, 0, 3)
	for _, rotorConfig := range enigmaConfig.Rotors {
		internalConfig := service.RotorConfig{
			RotorType:     rotorConfig.Type,
			RotorPosition: rotorConfig.Position,
			RingSetting:   rotorConfig.RingSetting,
		}
		rotor := service.InitialRotor(internalConfig)
		rotors = append(rotors, rotor)
	}
	return &EnigmaMachine{
		rotors:    rotors,
		plugboard: plugboard,
		reflector: reflector,
	}
}

func (em *EnigmaMachine) handleRotorStepping() {
	// 最もシンプルな実装：右端ローターのみ回転（ノッチ無視）
	em.rotors[2].Rotate()
}

func preprocessMessage(message string) string {
	message = strings.ToUpper(message)
	result := ""
	for _, char := range message {
		if char >= 'A' && char <= 'Z' {
			result += string(char)
		}
	}
	return result
}

// 設定ファイルを読み込む
func LoadConfig(filename string) (*EnigmaConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config EnigmaConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

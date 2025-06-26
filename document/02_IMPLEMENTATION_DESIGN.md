# 実装設計書

## 2.1 アーキテクチャ概要

### パッケージ構成
```
enigma-go/
├── cmd/enigma/           # CLIアプリケーション
│   └── Enigma.go         # メインエントリーポイント
├── internal/service/     # 内部サービス層
│   ├── plugboard.go      # プラグボード処理
│   ├── reflector.go      # リフレクター処理
│   └── rotor.go          # ローター処理
└── config/               # 設定ファイル
    └── test_config.json  # テスト用設定
```

### 依存関係
- CLI層 → サービス層
- サービス層は独立（相互依存なし）

## 2.2 データ構造設計

### CLI層の構造体

```go
// エニグママシンの設定（JSON用）
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
```

### サービス層の構造体

```go
// プラグボード
type PlugboardConfig struct {
    PlugboardConfig []string
}

type Plugboard struct {
    plugboard map[string]string
}

// リフレクター
type ReflectorConfig struct {
    ReflectorConfig []string
}

type Reflector struct {
    reflector map[string]string
}

// ローター
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
```

## 2.3 エニグマ暗号化フロー

### 処理順序
1. **ローター回転** - 暗号化前に実行
2. **プラグボード入力変換**
3. **順方向ローター変換** (右→左)
4. **リフレクター変換**
5. **逆方向ローター変換** (左→右)
6. **プラグボード出力変換**

### 実装コード例
```go
func (e *EnigmaMachine) EncryptChar(input string) string {
    // 1. ローター回転
    e.handleRotorStepping()
    
    // 2. プラグボード入力変換
    step1 := e.plugboard.PlugboardTransform(input)
    
    // 3. 順方向ローター変換（右→左）
    step2 := e.rotors[2].ForwardTransform(step1)  // 右端
    step3 := e.rotors[1].ForwardTransform(step2)  // 中央
    step4 := e.rotors[0].ForwardTransform(step3)  // 左端
    
    // 4. リフレクター変換
    step5 := e.reflector.ReflectorTransform(step4)
    
    // 5. 逆方向ローター変換（左→右）
    step6 := e.rotors[0].BackwardTransform(step5) // 左端
    step7 := e.rotors[1].BackwardTransform(step6) // 中央
    step8 := e.rotors[2].BackwardTransform(step7) // 右端
    
    // 6. プラグボード出力変換
    result := e.plugboard.PlugboardTransform(step8)
    
    return result
}
```

## 2.4 ローター回転メカニズム

### 基本原理
- 右端ローターは毎回1つ回転
- 現在は最もシンプルな実装（右端ローターのみ）
- 将来的にダブルステッピング機能を追加予定

### 実装
```go
func (em *EnigmaMachine) handleRotorStepping() {
    // 最もシンプルな実装：右端ローターのみ回転
    em.rotors[2].Rotate()
}
```

## 2.5 自己逆元特性

### エニグマの重要な特性
- 同じ設定で2回処理すると元に戻る
- A→Bなら必ずB→A
- この特性により暗号化と復号化が同一処理

### テスト例
```bash
# 暗号化
go run cmd/enigma/Enigma.go -input "HELLO" -config "config/test_config.json"
# 出力例: UVJPX

# 復号化（同じコマンド）
go run cmd/enigma/Enigma.go -input "UVJPX" -config "config/test_config.json"
# 出力: HELLO
```
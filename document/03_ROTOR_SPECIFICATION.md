# ローター仕様書

## 3.1 ローター基本仕様

### 対応ローター種別
- **I**: 配線パターン "EKMFLGDQVZNTOWYHXUSPAIBRCJ", ノッチ位置 "Q"
- **II**: 配線パターン "AJDKSIRUXBLHWTMCQGZNPYFVOE", ノッチ位置 "E"
- **III**: 配線パターン "BDFHJLCPRTXVZNYEIWGAKMUSQO", ノッチ位置 "V"
- **IV**: 配線パターン "ESOVPZJAYQUIRHXLNFTGKDCMWB", ノッチ位置 "J"
- **V**: 配線パターン "VZBRGITYUPSDNHLXAWMJQOFECK", ノッチ位置 "Z"

### ローター構成要素
- **配線パターン**: 文字変換のマッピング
- **ノッチ位置**: 次ローターを回転させる位置
- **現在位置**: 0-25の数値（A=0, B=1, ..., Z=25）
- **リング設定**: 内部オフセット調整値

## 3.2 変換処理アルゴリズム

### 順方向変換（ForwardTransform）
```go
func (r *Rotor) ForwardTransform(input string) string {
    currentPos := r.currentPosition
    inputPos := int(input[0] - 'A')
    
    // 入力位置調整
    transformPos := (inputPos + currentPos - r.ringSetting + 26) % 26
    transformString := string(rune('A') + rune(transformPos))
    
    // 配線による変換
    transformdString := r.forwardWiring[transformString]
    transformdPos := int(transformdString[0] - 'A')
    
    // 出力位置調整
    mappedPos := (transformdPos - currentPos + r.ringSetting + 26) % 26
    return string(rune('A') + rune(mappedPos))
}
```

### 逆方向変換（BackwardTransform）
```go
func (r *Rotor) BackwardTransform(input string) string {
    currentPos := r.currentPosition
    inputPos := int(rune(input[0]) - 'A')
    
    // 入力位置調整（順方向と符号が逆）
    transformPos := (inputPos + currentPos - r.ringSetting + 26) % 26
    transformString := string(rune('A') + rune(transformPos))
    
    // 逆配線による変換
    transformdString := r.backwardWiring[transformString]
    transformdPos := int(transformdString[0] - 'A')
    
    // 出力位置調整（順方向と符号が逆）
    mappedPos := (transformdPos - currentPos + r.ringSetting + 26) % 26
    return string(rune('A') + rune(mappedPos))
}
```

## 3.3 配線マップ生成

### 順方向配線
```go
func buildForwardWiring(rotorType string) map[string]string {
    wiringPosition := rotorWiringPatterns[rotorType]
    forwardWiringMap := make(map[string]string)
    for idx, _ := range wiringPosition {
        forwardWiringMap[string(rune('A'+idx))] = string(wiringPosition[idx])
    }
    return forwardWiringMap
}
```

### 逆方向配線
```go
func buildBackwardWiring(rotorType string) map[string]string {
    wiringPosition := rotorWiringPatterns[rotorType]
    backwardWiringMap := make(map[string]string)
    for idx, _ := range wiringPosition {
        backwardWiringMap[string(wiringPosition[idx])] = string(rune('A' + idx))
    }
    return backwardWiringMap
}
```

## 3.4 ローター回転

### 基本回転
```go
func (r *Rotor) Rotate() {
    r.currentPosition = (r.currentPosition + 1) % 26
}
```

### 位置管理
```go
// 現在位置取得
func (r *Rotor) GetPosition() string {
    return string(rune('A') + rune(r.currentPosition))
}

// ノッチ位置判定
func (r *Rotor) IsAtNotch() bool {
    currenChar := string(rune('A') + rune(r.currentPosition))
    return currenChar == r.notchPosition
}

// 位置設定
func (r *Rotor) SetPosition(position string) {
    r.currentPosition = int(position[0] - 'A')
}
```

## 3.5 初期化プロセス

### ローター作成
```go
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
```
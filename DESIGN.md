# エニグママシン Go実装 全体設計書

## 概要
エニグママシンの暗号化・復号化機能をGoで実装するプロジェクトです。
入力された文字列をエニグママシンのアルゴリズムで暗号化し、同じ設定で復号化できるシステムを構築します。

本書は全体設計を記載しており、各機能の詳細設計は別途作成します。

## システム構成

### 1. プロジェクト構造
```
enigma-go/
├── cmd/
│   └── main.go           # エントリーポイント
├── internal/
│   ├── enigma/
│   │   ├── machine.go    # エニグママシン本体
│   │   ├── rotor.go      # ローター実装
│   │   ├── reflector.go  # リフレクター実装
│   │   └── plugboard.go  # プラグボード実装
│   ├── config/
│   │   └── config.go     # 設定管理
│   └── utils/
│       └── validator.go  # 入力検証
├── pkg/
│   └── enigma/
│       └── api.go        # 外部APIインターフェース
├── go.mod
├── go.sum
└── README.md
```

### 2. 主要コンポーネント

#### 2.1 エニグママシン (EnigmaMachine)
- **責務**: 全体の暗号化・復号化処理を統括
- **構成要素**:
  - ローター（複数）
  - リフレクター
  - プラグボード
- **主要メソッド**:
  - `NewEnigmaMachine(config Config) *EnigmaMachine`
  - `Encrypt(input string) (string, error)`
  - `Decrypt(input string) (string, error)`
  - `Reset()`

#### 2.2 ローター (Rotor)
- **責務**: 文字の置換と回転機能
- **属性**:
  - 配線パターン (wiring)
  - 現在位置 (position)
  - ノッチ位置 (notch)
  - リング設定 (ringSetting)
- **主要メソッド**:
  - `NewRotor(rotorType string, position rune, ringSetting int) *Rotor`
  - `Forward(input int) int`
  - `Backward(input int) int`
  - `Rotate() bool`
  - `AtNotch() bool`

#### 2.3 リフレクター (Reflector)
- **責務**: 信号の反射処理
- **属性**:
  - 配線パターン (wiring)
- **主要メソッド**:
  - `NewReflector(reflectorType string) *Reflector`
  - `Reflect(input int) int`

#### 2.4 プラグボード (Plugboard)
- **責務**: 文字のペア交換
- **属性**:
  - プラグ設定 (plugs)
- **主要メソッド**:
  - `NewPlugboard(plugs []string) *Plugboard`
  - `Transform(input int) int`

### 3. データ構造

#### 3.1 設定構造体
```go
type Config struct {
    Rotors      []RotorConfig
    Reflector   string
    Plugboard   []string
    RingSettings []int
}

type RotorConfig struct {
    Type     string
    Position rune
}
```

#### 3.2 エニグママシン構造体
```go
type EnigmaMachine struct {
    rotors    []*Rotor
    reflector *Reflector
    plugboard *Plugboard
    config    Config
}
```

### 4. 暗号化処理フロー

1. **入力検証**: 文字列の妥当性チェック（A-Z、空白文字等）
2. **前処理**: 大文字変換、空白・記号の処理
3. **暗号化処理**:
   - プラグボード変換（前）
   - ローター回転処理
   - ローター順方向処理（右→左）
   - リフレクター処理
   - ローター逆方向処理（左→右）
   - プラグボード変換（後）
4. **後処理**: 出力形式の整形

### 5. 復号化処理
- 暗号化と同じ処理（エニグママシンの特性により、同じ設定で暗号化すると復号化される）

### 6. 設定管理

#### 6.1 サポートするローター種類
- I, II, III, IV, V（標準的なエニグマローター）
- 各ローターの配線パターンとノッチ位置を定義

#### 6.2 サポートするリフレクター種類
- B, C（標準的なリフレクター）

#### 6.3 設定例
```go
config := Config{
    Rotors: []RotorConfig{
        {Type: "I", Position: 'A'},
        {Type: "II", Position: 'B'},
        {Type: "III", Position: 'C'},
    },
    Reflector: "B",
    Plugboard: []string{"AB", "CD", "EF"},
    RingSettings: []int{1, 1, 1},
}
```

### 7. エラーハンドリング

#### 7.1 カスタムエラー型
```go
type EnigmaError struct {
    Type    string
    Message string
}
```

#### 7.2 エラーパターン
- 無効な入力文字
- 不正な設定値
- ローター/リフレクター種類の不一致
- プラグボード設定の重複

### 8. 外部インターフェース

#### 8.1 CLI インターフェース
```bash
# 暗号化
./enigma-go encrypt -input "HELLO WORLD" -config config.json

# 復号化
./enigma-go decrypt -input "ENCRYPTED_TEXT" -config config.json

# 対話モード
./enigma-go interactive
```

#### 8.2 設定ファイル形式（JSON）
```json
{
  "rotors": [
    {"type": "I", "position": "A"},
    {"type": "II", "position": "B"},
    {"type": "III", "position": "C"}
  ],
  "reflector": "B",
  "plugboard": ["AB", "CD", "EF"],
  "ring_settings": [1, 1, 1]
}
```

### 9. 実装優先度

#### Phase 1: 基本機能
1. ローター実装
2. リフレクター実装
3. 基本的な暗号化・復号化

#### Phase 2: 拡張機能
1. プラグボード実装
2. 設定管理機能
3. CLI インターフェース

## 参考資料

- エニグママシンの歴史的仕様
- 暗号学の基礎理論
- Go言語のベストプラクティス
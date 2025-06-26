# CLI使用方法

## 4.1 基本使用方法

### コマンド形式
```bash
go run cmd/enigma/Enigma.go -input "メッセージ" -config "設定ファイルパス"
```

### 必須パラメータ
- **-input**: 暗号化・復号化するメッセージ
- **-config**: JSON設定ファイルのパス

## 4.2 使用例

### 暗号化
```bash
go run cmd/enigma/Enigma.go -input "HELLO WORLD" -config "config/test_config.json"
```

### 復号化
```bash
# 上記の出力結果を使用
go run cmd/enigma/Enigma.go -input "暗号化結果" -config "config/test_config.json"
```

### 1文字暗号化
```bash
go run cmd/enigma/Enigma.go -input "A" -config "config/test_config.json"
```

## 4.3 入力文字列の処理

### 前処理ルール
1. **大文字変換**: 小文字は自動的に大文字に変換
2. **文字フィルタリング**: A-Z以外の文字は除去
3. **スペース除去**: 空白文字は自動的に除去

### 処理例
```go
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
```

## 4.4 設定ファイル仕様

### 基本設定ファイル例
```json
{
  "name": "test_setting_001",
  "rotors": [
    {
      "type": "I",
      "position": "A", 
      "ring_setting": 0
    },
    {
      "type": "II",
      "position": "B",
      "ring_setting": 3
    },
    {
      "type": "III",
      "position": "V",
      "ring_setting": 6
    }
  ],
  "reflector": ["AB", "CD", "EF", "GH", "IJ", "KL", "MN", "OP", "QR", "ST", "UV", "WY", "XZ"],
  "plugboard": ["AB", "CD", "EF", "GH", "IJ"]
}
```

### 設定項目説明
- **name**: 設定の識別名（省略可能）
- **rotors**: 3つのローター設定（左端、中央、右端の順）
- **reflector**: 13個のペア（26文字すべてをペア化）
- **plugboard**: プラグボード設定（0-13個のペア）

## 4.5 エラーハンドリング

### 引数不足エラー
```bash
$ go run cmd/enigma/Enigma.go
使用方法: enigma -input "HELLO WORLD" -config config.json
  -config string
        設定ファイルのパス
  -input string
        暗号化するメッセージ
```

### 設定ファイルエラー
```bash
$ go run cmd/enigma/Enigma.go -input "HELLO" -config "存在しないファイル.json"
設定ファイルの読み込みエラー: open 存在しないファイル.json: no such file or directory
```
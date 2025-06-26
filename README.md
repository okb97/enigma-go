# Enigma Go

エニグママシンをGoで実装したプロジェクト

## 概要

第二次世界大戦中にドイツ軍が使用した暗号機「エニグマ」の動作をGoで再現したプロジェクトです。
エニグマの自己逆元特性により、同じ設定で暗号化と復号化を行うことができます。

## 特徴

- ✅ エニグマの暗号化アルゴリズムを実装
- ✅ ローター、リフレクター、プラグボードの再現
- ✅ CLIによる操作
- ✅ JSON形式での設定管理
- ✅ 自己逆元特性（暗号化=復号化）

## 技術仕様

- **言語**: Go
- **設定**: JSON
- **CLI**: flag パッケージ
- **対象**: エニグマ（3ローター構成）
- **ローター**: I, II, III, IV, V対応
- **文字処理**: A-Z（26文字）

## 使用方法

### 基本コマンド
```bash
go run cmd/enigma/Enigma.go -input "メッセージ" -config "設定ファイル"
```

### 暗号化例
```bash
go run cmd/enigma/Enigma.go -input "HELLO WORLD" -config "config/test_config.json"
# 出力例: UVJPX
```

### 復号化例
```bash
go run cmd/enigma/Enigma.go -input "UVJPX" -config "config/test_config.json"
# 出力: HELLO
```

## 設定ファイル例

```json
{
  "name": "test_setting_001",
  "rotors": [
    {"type": "I", "position": "A", "ring_setting": 0},
    {"type": "II", "position": "B", "ring_setting": 3},
    {"type": "III", "position": "V", "ring_setting": 6}
  ],
  "reflector": ["AB", "CD", "EF", "GH", "IJ", "KL", "MN", "OP", "QR", "ST", "UV", "WY", "XZ"],
  "plugboard": ["AB", "CD", "EF", "GH", "IJ"]
}
```

## プロジェクト構成

```
enigma-go/
├── cmd/enigma/           # CLIアプリケーション
│   └── Enigma.go         # メインエントリーポイント
├── internal/service/     # 内部サービス層
│   ├── plugboard.go      # プラグボード処理
│   ├── reflector.go      # リフレクター処理
│   └── rotor.go          # ローター処理
├── config/               # 設定ファイル
│   └── test_config.json  # テスト用設定
└── document/             # 設計文書
    ├── 01_JSON_CONFIG_DESIGN.md
    ├── 02_IMPLEMENTATION_DESIGN.md
    ├── 03_ROTOR_SPECIFICATION.md
    ├── 04_CLI_USAGE.md
    └── 05_TESTING_GUIDE.md
```
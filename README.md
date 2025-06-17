# Enigma Go

エニグママシンをGoで実装するプロジェクト

## 概要

第二次世界大戦中にドイツ軍が使用した暗号機「エニグママシン」の動作をGoで再現したプロジェクトです。

## 特徴

- 🔐 エニグママシンの暗号化アルゴリズムを実装
- ⚙️ ローター、リフレクター、プラグボードの再現
- 🎯 CLI インターフェースによる簡単操作
- 📝 JSON形式での設定管理

## エニグママシンについて

エニグママシンの詳しい仕組みについては、[ENIGMA_MECHANISM.md](ENIGMA_MECHANISM.md) を参照。

## 使用予定の技術

- **言語**: Go
- **設定**: JSON
- **CLI**: コマンドライン引数とフラグ

## 使用例（予定）

```bash
# 暗号化
./enigma-go encrypt -input "HELLO WORLD" -config config.json

# 復号化  
./enigma-go decrypt -input "ILBDA AMTAZ" -config config.json

# 対話モード
./enigma-go interactive
```

# Enigma Go

エニグママシンをGoで実装するプロジェクト

## 概要

第二次世界大戦中にドイツ軍が使用した暗号機「エニグマ」の動作をGoで再現したプロジェクトです。

## 特徴

- エニグマの暗号化アルゴリズムを実装
- ローター、リフレクター、プラグボードの完全再現
- CLI インターフェースによる簡単操作
- JSON形式での設定管理

## 設計ドキュメント

- [JSON設定データ構造設計](document/01_JSON_CONFIG_DESIGN.md)
- [エニグママシン本体設計](document/02_ENIGMA_MACHINE_DESIGN.md)
- [プラグボード設計](document/03_PLUGBOARD_DESIGN.md)
- [ローター設計](document/04_ROTOR_DESIGN.md)
- [リフレクター設計](document/05_REFLECTOR_DESIGN.md)
- [関数処理方法設計](document/06_FUNCTION_DESIGN.md)
- [エラーハンドリング設計](document/07_ERROR_HANDLING_DESIGN.md)

## 使用予定の技術

- **言語**: Go
- **設定**: JSON
- **CLI**: コマンドライン引数とフラグ
- **対象**: 標準的なエニグマ（3ローター構成）

## 使用例（予定）

```bash
# 暗号化
./enigma-go encrypt -input "HELLO WORLD" -config config.json

# 復号化  
./enigma-go decrypt -input "ILBDA AMTAZ" -config config.json

# 対話モード
./enigma-go interactive
```

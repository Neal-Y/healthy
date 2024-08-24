# Healthy AI 營養分析器

**版本**: 1.0.0  
**作者**: [Neal_Yang]  
**日期**: [8/01]

## 項目概述

Healthy AI 營養分析器是一個基於 Web 的應用程式，允許用戶上傳食物圖片，應用程式使用 OpenAI 的 GPT API 來分析圖片並提供估算的營養信息，包括`熱量`、`蛋白質`和`纖維`。該項目展示了我在 Golang、Gin 框架、圖片處理、整合外部 API（OpenAI）以及雲端部署方面的技術能力。

## 功能

- **圖片上傳**: 用戶可以上傳菜品或食物的圖片。
- **AI 分析**: 應用程式調用 OpenAI 的 GPT API 分析圖片並返回 JSON 格式的熱量、蛋白質和纖維的總量。
- **響應式前端**: 簡單而優雅的用戶界面，用於上傳圖片並顯示結果。
- **雲端部署**: 部署在 AWS EC2 上並且結合 RDS 支持實時使用。
- **可擴展架構**: 採用 Clean Architecture 架構，便於擴展和維護。

## 技術棧

- **後端**: Golang, Gin, GORM
- **前端**: HTML, CSS, JavaScript 
- **資料庫**: MySQL（通過 GORM）
- **API 整合**: OpenAI API
- **雲端部署**: AWS EC2, RDS
- **任務調度**: 使用 Cron 定期清理任務
- **版本控制**: Git

## 架構設計

該項目遵循 Clean Architecture 設計模式，將業務邏輯與外部服務和接口解耦。

```plaintext

├── cmd                                    # 主程序入口點和單獨運行的清理程序
│   ├── cleanup                            # 用於定期清理任務的程式
│   └── service                            # 主應用程序的啟動邏輯
├── config                                 # 配置文件，包含應用程式和數據庫設置
│   ├── config.go                          # 應用程序的配置邏輯
│   └── mysql                              # MySQL 配置與初始化
├── constant                               # 定義應用中的常量，如 GPT API 配置
├── cron                                   # 定期任務邏輯，調度清理任務
├── database                               # 數據庫相關，包含遷移文件
│   └── migrations                         # 數據庫遷移文件，用於管理表結構
├── docker-compose.yaml                    # Docker 容器編排文件
├── go.mod                                 # Go 模組的依賴管理
├── go.sum                                 # Go 模組依賴的校驗和版本信息
├── handler                                # HTTP 請求處理層，負責文件上傳和下載邏輯
│   └── file                               # 文件處理的具體邏輯
├── infrastructure                         # 基礎設施層，主要處理數據庫連接
├── model                                  # 數據模型，用於數據庫和 DTO（數據傳輸對象）
├── repository                             # 數據庫訪問層，處理數據的 CRUD 操作
├── route                                  # Gin 框架的路由定義
├── service                                # 服務層，處理業務邏輯
├── template                               # 前端模板文件夾
├── uploads                                # 用戶上傳的文件存儲位置
└── util                                   # 通用工具類，包含 GPT 調用等功能
```

# 未來改進
- 用戶認證: 添加用戶登錄功能，以便跟蹤每個用戶的營養分析歷史記錄。
- 改進 UI: 優化前端設計，提升用戶體驗。
- 支持多圖片上傳: 支持單次上傳多個圖片進行分析。
- 詳細報告生成: 為用戶生成 PDF 格式的營養分析報告，供下載。


   

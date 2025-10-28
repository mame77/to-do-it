## githubの使い方について

### branch
branch名は基本feature/OOにする
OOの部分は機能名をつける

example
```
git checkout -b feature/task
```

### commit
commitメッセージはOO: なにをしたか #issueナンバー

example
```
git feat:構造体の追加 #2
```

OOの部分に関しては

- feat: 機能追加
- update: 機能変更
- fix: バグ修正
- docs: ドキュメント修正
- refactor: リファクタリング
- test: テスト

### pull reqest
reviewerにTL(まめ)を入れる
assigneeに自分で入る
developmentにissueを入れる


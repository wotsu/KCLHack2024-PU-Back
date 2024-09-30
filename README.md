# KCLHack2024-PU-Back

# ローカル環境を最新にする
$ git switch main

$ git pull

# ブランチを切り替える
# ブランチ名はそのブランチの内容が分かるように
$ git switch -c ブランチ名

# 作業内容をcommit&pushする

# そのブランチで作業が終わればmainブランチにプルリクエストを出す
# 原則として一つのブランチにつき一つの機能
# 例: loginブランチでログイン機能が完成したらプルリクを出し、ログアウト機能はlogoutブランチに切り替えてブランチを別にする

今いるブランチの確認
`git branch`

ブランチの変更
`git switch (ブランチ名)`

ブランチの作成と移動
`git switch -c (ブランチ名)`

コミットするファイルを追加(ファイル名に . で一括追加)
`git add (ファイル名)`

コミット
`git commit -m "(コミットメッセージ)"`

コミットしたファイルをpush
`git push origin (ブランチ名)`

ローカルのコードを最新にする(mainブランチで)
`git pull`
# github.comのAPI調査など

Githubで自分がつけたスターの情報を取る。

- 言語の種類
- 日付

これを元に分布グラフっぽいの作る。

## OAuth認証なしでスターの一覧取れる

`curl https://api.github.com/users/:username/starred`

### 例)

`https://api.github.com/users/kyokomi/starred`

## OAuth認証なしでユーザー情報もとれるみたい

`https://api.github.com/users/:username`


### 例)

`https://api.github.com/users/kyokomi`




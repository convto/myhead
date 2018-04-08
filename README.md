# myhead
UNIXの `head` コマンドを再現したコマンドを作成した。
リモートファイルにもアクセスできるように機能追加した

## Installation
`$ go get github.com/saroteki/myhead`

## Usage
remoteのファイルは実際に存在するのでそのまま動作確認できます
### local file
```
$ myhead -n=2 FILE_PATH

==> FILE_PATH <==
FILE_CONTENT
FILE_CONTENT
==> FILE_PATH end <==

$
```
### remote file
```
$ myhead -n=2 -remote=true https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt

==> https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt <==
Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
==> https://s3-ap-northeast-1.amazonaws.com/myhead/test.txt end <==

$
```

## Author
Yuya Okumur (saroteki)
# my_gsa

My Go Static Analyzers.

| Name         | Description                                 |
|--------------|---------------------------------------------|
| argcount     | 関数定義で、多すぎる引数を検知              |
| trashcomment | 無意味な関数コメントを検知                  |
| gophersample | 識別子gopherを検知                          |
| structctx    | 構造体のフィールドに定義されたcontextを検知 |
| privatetag   | 非公開フィールドにつけられたtagを検知       |

check

```shell
make run
make test
```

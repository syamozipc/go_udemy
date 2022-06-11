# goについて基本

- scope chainはjsと一緒
- アクセス修飾子が無いらしい（以下はfunc nameでもpropertyでも同様）
    - 小文字始まりはprivate（そのpackage内でのみ使用可能）
    - 大文字始まりはpublic（package外で使用可能）

# mapについて

- PHPやJSでは、mapでは順番通りに要素が格納されるが、Goではランダムに並び替えられる
    - そのため、必ずkeyを指定し、取得時はそのkeyを指定する必要がある
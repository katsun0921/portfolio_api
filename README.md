# portfolio_api

Portfolio API created with Golang.

## DATABASE

今回データベース構築はmigrateでcliで実行しています。

<https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>

### CLI

Linux

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

MacOS(Homebrewを使用)

```bash
brew install golang-migrate
```

Moduleをinstall。今回はmysqlを使用

```bash
go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate
```

### MySQL

<https://github.com/golang-migrate/migrate/tree/master/database/mysql>

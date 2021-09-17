# GORM v2 Compatable with Unique Index

Fork from and fully compatible with: [jinzhu/gorm v1.21.14](https://github.com/go-gorm/gorm)


```sh
go get -u github.com/Laisky/gorm/v2
```

## New Features

1. support unique-index with soft-delete
2. support filter specified stack by regexp in log
3. log SQL result

### Support unique-index with soft-delete

1. Add `DeletedFlag` in model struct

(or embeded `gorm.ModelSupportUnique`)

```go
type Model struct {
    ID          uint `gorm:"primary_key"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   *time.Time `sql:"index"`
    DeletedFlag uint `gorm:"type:INT UNSIGNED DEFAULT 0 NOT NULL"`

    Foo string `gorm:"unique_index"`
}
```

2. Run Migrate

will auto create unique index on `foo,deleted_flag`

3. Soft Delete

```go
mdl.ID = 1
Model.Delete(mdl)
```

will execute SQL like:

```sql
update model set deleted_at = NOW(), deleted_flag = id where id = 1
```

4. Undelete

```go
mdl.Undelete()
db.Save(mdl)
```

### Support filter specified stack by regexp in log

before create gorm db:

```go
// stacks that match these regexps will be ignored
gorm.AddLogFileIgnoreStackPattern(
    regexp.MustCompile("core/model/base/pagination_sql.go"),
    regexp.MustCompile("utils/helpers.go"),
)
```

supported tags: eq, lt, lte, gt, gte, in, ints, strs, like, like-bin

more details in `tag.go`

### Log SQL Result

```go
db = db.LogSQLResult(true)
```

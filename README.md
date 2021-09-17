# GORM Compatable with Unique Index

Fork from and fully compatible with: [jinzhu/gorm v1.9.16](https://github.com/jinzhu/gorm)


```sh
go get github.com/Laisky/gorm v1.9.18
```

## New Features

1. support unique-index with soft-delete
2. support filter specified stack by regexp in log
3. support ignore specified SQL log by regexp
4. generate SQL condition by tag


###  support unique-index with soft-delete

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

### Support ignore specified SQL log by regexp

```go
type Request struct {
	HostID          uint         `form:"host_id" sql:"column:h.id;op:eq"`
	HostIPs         string       `form:"host_ips" sql:"column:h.ip;op:ips"`
	HostName        string       `form:"host_name" sql:"column:h.name;op:like"`
}

req := new(Request)
db = gorm.ApplySQLCondition(db, req)
// equal to:
//   db = db.Where("h.id = ?", req.HostID).
//       Where("h.ip IN (?)", strings.Split(req.HostIPs, ",")).
//       Where("h.id LIKE %?%", req.HostName)

```

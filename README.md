# GORM Compatable with Unique Index

Fork from and fully compatible with: [jinzhu/gorm v1.9.16](https://github.com/jinzhu/gorm)


```sh
go get github.com/Laisky/gorm v1.9.18
```

## New Features


### Compatible with soft-delete and unique-index

#### 1. Add `DeletedFlag` in model struct

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

#### 2. Run Migrate

will auto create unique index on `foo,deleted_flag`

#### 3. Soft Delete

```go
mdl.ID = 1
Model.Delete(mdl)
```

will execute SQL like:

```sql
update model set deleted_at = NOW(), deleted_flag = id where id = 1
```

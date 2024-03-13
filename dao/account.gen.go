// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gin-admin-api/model"
)

func newAccountEntity(db *gorm.DB, opts ...gen.DOOption) accountEntity {
	_accountEntity := accountEntity{}

	_accountEntity.accountEntityDo.UseDB(db, opts...)
	_accountEntity.accountEntityDo.UseModel(&model.AccountEntity{})

	tableName := _accountEntity.accountEntityDo.TableName()
	_accountEntity.ALL = field.NewAsterisk(tableName)
	_accountEntity.ID = field.NewInt64(tableName, "id")
	_accountEntity.Username = field.NewString(tableName, "username")
	_accountEntity.Password = field.NewString(tableName, "password")
	_accountEntity.Name = field.NewString(tableName, "name")
	_accountEntity.Mobile = field.NewString(tableName, "mobile")
	_accountEntity.Email = field.NewString(tableName, "email")
	_accountEntity.Avatar = field.NewString(tableName, "avatar")
	_accountEntity.IsAdmin = field.NewInt64(tableName, "is_admin")
	_accountEntity.Status = field.NewInt64(tableName, "status")
	_accountEntity.LastLoginIP = field.NewString(tableName, "last_login_ip")
	_accountEntity.LastLoginDate = field.NewField(tableName, "last_login_date")
	_accountEntity.Salt = field.NewString(tableName, "salt")
	_accountEntity.Token = field.NewString(tableName, "token")
	_accountEntity.ExpireTime = field.NewField(tableName, "expire_time")
	_accountEntity.CreatedAt = field.NewField(tableName, "created_at")
	_accountEntity.UpdatedAt = field.NewField(tableName, "updated_at")
	_accountEntity.DeletedAt = field.NewField(tableName, "deleted_at")

	_accountEntity.fillFieldMap()

	return _accountEntity
}

type accountEntity struct {
	accountEntityDo

	ALL           field.Asterisk
	ID            field.Int64  // 主键id
	Username      field.String // 用户名
	Password      field.String // 密码
	Name          field.String // 真实姓名
	Mobile        field.String // 手机号码
	Email         field.String // 邮箱地址
	Avatar        field.String // 用户头像
	IsAdmin       field.Int64  // 是否为超级管理员:0否,1是
	Status        field.Int64  // 状态1是正常,0是禁用
	LastLoginIP   field.String // 最后登录ip地址
	LastLoginDate field.Field  // 最后登录时间
	Salt          field.String // 密码盐
	Token         field.String // token
	ExpireTime    field.Field  // token过期时间
	CreatedAt     field.Field  // 创建时间
	UpdatedAt     field.Field  // 更新时间
	DeletedAt     field.Field  // 软删除时间

	fieldMap map[string]field.Expr
}

func (a accountEntity) Table(newTableName string) *accountEntity {
	a.accountEntityDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountEntity) As(alias string) *accountEntity {
	a.accountEntityDo.DO = *(a.accountEntityDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountEntity) updateTableName(table string) *accountEntity {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.Name = field.NewString(table, "name")
	a.Mobile = field.NewString(table, "mobile")
	a.Email = field.NewString(table, "email")
	a.Avatar = field.NewString(table, "avatar")
	a.IsAdmin = field.NewInt64(table, "is_admin")
	a.Status = field.NewInt64(table, "status")
	a.LastLoginIP = field.NewString(table, "last_login_ip")
	a.LastLoginDate = field.NewField(table, "last_login_date")
	a.Salt = field.NewString(table, "salt")
	a.Token = field.NewString(table, "token")
	a.ExpireTime = field.NewField(table, "expire_time")
	a.CreatedAt = field.NewField(table, "created_at")
	a.UpdatedAt = field.NewField(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *accountEntity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountEntity) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 17)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["name"] = a.Name
	a.fieldMap["mobile"] = a.Mobile
	a.fieldMap["email"] = a.Email
	a.fieldMap["avatar"] = a.Avatar
	a.fieldMap["is_admin"] = a.IsAdmin
	a.fieldMap["status"] = a.Status
	a.fieldMap["last_login_ip"] = a.LastLoginIP
	a.fieldMap["last_login_date"] = a.LastLoginDate
	a.fieldMap["salt"] = a.Salt
	a.fieldMap["token"] = a.Token
	a.fieldMap["expire_time"] = a.ExpireTime
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a accountEntity) clone(db *gorm.DB) accountEntity {
	a.accountEntityDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountEntity) replaceDB(db *gorm.DB) accountEntity {
	a.accountEntityDo.ReplaceDB(db)
	return a
}

type accountEntityDo struct{ gen.DO }

type IAccountEntityDo interface {
	gen.SubQuery
	Debug() IAccountEntityDo
	WithContext(ctx context.Context) IAccountEntityDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAccountEntityDo
	WriteDB() IAccountEntityDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAccountEntityDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAccountEntityDo
	Not(conds ...gen.Condition) IAccountEntityDo
	Or(conds ...gen.Condition) IAccountEntityDo
	Select(conds ...field.Expr) IAccountEntityDo
	Where(conds ...gen.Condition) IAccountEntityDo
	Order(conds ...field.Expr) IAccountEntityDo
	Distinct(cols ...field.Expr) IAccountEntityDo
	Omit(cols ...field.Expr) IAccountEntityDo
	Join(table schema.Tabler, on ...field.Expr) IAccountEntityDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAccountEntityDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAccountEntityDo
	Group(cols ...field.Expr) IAccountEntityDo
	Having(conds ...gen.Condition) IAccountEntityDo
	Limit(limit int) IAccountEntityDo
	Offset(offset int) IAccountEntityDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountEntityDo
	Unscoped() IAccountEntityDo
	Create(values ...*model.AccountEntity) error
	CreateInBatches(values []*model.AccountEntity, batchSize int) error
	Save(values ...*model.AccountEntity) error
	First() (*model.AccountEntity, error)
	Take() (*model.AccountEntity, error)
	Last() (*model.AccountEntity, error)
	Find() ([]*model.AccountEntity, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountEntity, err error)
	FindInBatches(result *[]*model.AccountEntity, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AccountEntity) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAccountEntityDo
	Assign(attrs ...field.AssignExpr) IAccountEntityDo
	Joins(fields ...field.RelationField) IAccountEntityDo
	Preload(fields ...field.RelationField) IAccountEntityDo
	FirstOrInit() (*model.AccountEntity, error)
	FirstOrCreate() (*model.AccountEntity, error)
	FindByPage(offset int, limit int) (result []*model.AccountEntity, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAccountEntityDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a accountEntityDo) Debug() IAccountEntityDo {
	return a.withDO(a.DO.Debug())
}

func (a accountEntityDo) WithContext(ctx context.Context) IAccountEntityDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountEntityDo) ReadDB() IAccountEntityDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountEntityDo) WriteDB() IAccountEntityDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountEntityDo) Session(config *gorm.Session) IAccountEntityDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountEntityDo) Clauses(conds ...clause.Expression) IAccountEntityDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountEntityDo) Returning(value interface{}, columns ...string) IAccountEntityDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountEntityDo) Not(conds ...gen.Condition) IAccountEntityDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountEntityDo) Or(conds ...gen.Condition) IAccountEntityDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountEntityDo) Select(conds ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountEntityDo) Where(conds ...gen.Condition) IAccountEntityDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountEntityDo) Order(conds ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountEntityDo) Distinct(cols ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountEntityDo) Omit(cols ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountEntityDo) Join(table schema.Tabler, on ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountEntityDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountEntityDo) RightJoin(table schema.Tabler, on ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountEntityDo) Group(cols ...field.Expr) IAccountEntityDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountEntityDo) Having(conds ...gen.Condition) IAccountEntityDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountEntityDo) Limit(limit int) IAccountEntityDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountEntityDo) Offset(offset int) IAccountEntityDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountEntityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountEntityDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountEntityDo) Unscoped() IAccountEntityDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountEntityDo) Create(values ...*model.AccountEntity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountEntityDo) CreateInBatches(values []*model.AccountEntity, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountEntityDo) Save(values ...*model.AccountEntity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountEntityDo) First() (*model.AccountEntity, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountEntity), nil
	}
}

func (a accountEntityDo) Take() (*model.AccountEntity, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountEntity), nil
	}
}

func (a accountEntityDo) Last() (*model.AccountEntity, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountEntity), nil
	}
}

func (a accountEntityDo) Find() ([]*model.AccountEntity, error) {
	result, err := a.DO.Find()
	return result.([]*model.AccountEntity), err
}

func (a accountEntityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountEntity, err error) {
	buf := make([]*model.AccountEntity, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountEntityDo) FindInBatches(result *[]*model.AccountEntity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountEntityDo) Attrs(attrs ...field.AssignExpr) IAccountEntityDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountEntityDo) Assign(attrs ...field.AssignExpr) IAccountEntityDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountEntityDo) Joins(fields ...field.RelationField) IAccountEntityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountEntityDo) Preload(fields ...field.RelationField) IAccountEntityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountEntityDo) FirstOrInit() (*model.AccountEntity, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountEntity), nil
	}
}

func (a accountEntityDo) FirstOrCreate() (*model.AccountEntity, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountEntity), nil
	}
}

func (a accountEntityDo) FindByPage(offset int, limit int) (result []*model.AccountEntity, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accountEntityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountEntityDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountEntityDo) Delete(models ...*model.AccountEntity) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountEntityDo) withDO(do gen.Dao) *accountEntityDo {
	a.DO = *do.(*gen.DO)
	return a
}

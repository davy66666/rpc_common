package dbx

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

// Options 用于传递分表信息
type Options struct {
	BaseTableName string
	HashField     string
}

var (
	suffixFn = getTableSuffix
)

func getTableSuffix(id int64) string {
	return fmt.Sprintf("%d", id%10)
}

func SetSuffixFn(fn func(id int64) string) {
	suffixFn = fn
}

func RegisterHooks(db *gorm.DB, opts ...Options) error {
	for _, opt := range opts {
		// 创建操作回调
		err := db.Callback().Create().Before("gorm:create").Register("sharding:before_create", func(db *gorm.DB) {
			setShardedTable(db, opt)
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// 更新操作回调
		err = db.Callback().Update().Before("gorm:update").Register("sharding:before_update", func(db *gorm.DB) {
			setShardedTable(db, opt)
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// 删除操作回调
		err = db.Callback().Delete().Before("gorm:delete").Register("sharding:before_delete", func(db *gorm.DB) {
			setShardedTable(db, opt)
		})
		if err != nil {
			return errors.WithStack(err)
		}

		// 查询操作回调
		err = db.Callback().Query().Before("gorm:query").Register("sharding:before_query", func(db *gorm.DB) {
			setShardedTable(db, opt)
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func setShardedTable(db *gorm.DB, opts Options) {
	if db.Statement.Schema != nil && db.Statement.Schema.Table == opts.BaseTableName {
		var id int64
		// 获取哈希字段的值
		if db.Statement.ReflectValue.Kind() == reflect.Ptr {
			idField := db.Statement.ReflectValue.Elem().FieldByName(opts.HashField)
			if idField.IsValid() {
				id = idField.Interface().(int64)
			}
		} else {
			idField := db.Statement.ReflectValue.FieldByName(opts.HashField)
			if idField.IsValid() {
				id = idField.Interface().(int64)
			}
		}
		db.Statement.Table = opts.BaseTableName + "_" + suffixFn(id)
	}
}

// ReplaceSelectStatementPlugin GORM 替换掉 * 的Hook方法
type ReplaceSelectStatementPlugin struct {
	gorm.Plugin
}

// Name 替换 Select 语句的回调方法
func (p *ReplaceSelectStatementPlugin) Name() string {
	return "ReplaceSelectStatementPlugin"
}

// Initialize 在执行查询之前调用的回调方法
func (p *ReplaceSelectStatementPlugin) Initialize(db *gorm.DB) (err error) {
	// 在查询前注册一个钩子
	return db.Callback().Query().Before("gorm:query").Register("replaceSelectStatement", p.replaceSelectStatement)
}

// 替换 Select 语句的具体逻辑
func (p *ReplaceSelectStatementPlugin) replaceSelectStatement(db *gorm.DB) {
	// 获取表名称
	var b bool
	for _, s := range db.Statement.Clauses {
		if s.Name == "SELECT" && s.Expression != nil {
			b = true
		}
	}
	if len(db.Statement.Selects) == 0 && !b {
		db.Statement.Select(db.Statement.Schema.DBNames)
	}
}

// 注册一个插件来移除RETURNING子句
type noReturningPlugin struct{}

func (p *noReturningPlugin) Name() string {
	return "no_returning_plugin"
}

func (p *noReturningPlugin) Initialize(db *gorm.DB) error {
	err := db.Callback().Create().Before("gorm:create").Register(
		"disable_returning",
		func(d *gorm.DB) {
			d.Statement.Clauses["RETURNING"] = clause.Clause{}
		})
	if err != nil {
		return err
	}
	return nil
}

type CheckTenantPlugin struct {
	RequiredField string
}

func (p *CheckTenantPlugin) Name() string {
	return "check_tenant_plugin"
}

func (p *CheckTenantPlugin) Initialize(db *gorm.DB) error {
	check := func(tx *gorm.DB) {
		// 检查 WHERE 语句中是否包含指定字段
		hasField := false

		// 如果是软删除场景，WHERE 可能在 Clauses["WHERE"].Expression 中
		if whereClause, ok := tx.Statement.Clauses["WHERE"]; ok {
			if expr, ok := whereClause.Expression.(clause.Where); ok {
				for _, v := range expr.Exprs {
					if hasField {
						break
					}

					switch e := v.(type) {
					case clause.Eq:
						if col, ok := e.Column.(clause.Column); ok {
							if col.Name == p.RequiredField {
								hasField = true
								break
							}
						}
					case clause.IN:
						if col, ok := e.Column.(clause.Column); ok {
							if col.Name == p.RequiredField {
								hasField = true
								break
							}
						}
					case clause.Expr:
						if strings.Contains(e.SQL, p.RequiredField) {
							hasField = true
							break
						}
					}
				}
			}
		}

		if !hasField {
			err := tx.AddError(fmt.Errorf("missing required condition in WHERE: %s", p.RequiredField))
			if err != nil {
				return
			}
		}
	}

	// 在不同生命周期中注册检查逻辑
	if err := db.Callback().Query().Before("gorm:query").Register("check_tenant_before_query", check); err != nil {
		return errors.WithStack(err)
	}

	if err := db.Callback().Delete().Before("gorm:delete").Register("check_tenant_before_delete", check); err != nil {
		return errors.WithStack(err)
	}

	if err := db.Callback().Update().Before("gorm:update").Register("check_tenant_before_update", check); err != nil {
		return errors.WithStack(err)
	}

	if err := db.Callback().Row().Before("gorm:row").Register("check_tenant_before_row", check); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

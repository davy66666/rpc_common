package dbx

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type BatchUpdateOption struct {
	DB        *gorm.DB                 // GORM DB（可带事务）
	Table     string                   // 表名
	WhereKeys []string                 // WHERE 的字段名（支持单个或多个）
	Data      []map[string]interface{} // 每条记录：必须包含 WhereKeys 和要更新的字段
}

// BatchUpdateByCase 每次最多500条更新,更新的字段控制在10个以内,
func BatchUpdateByCase(ctx context.Context, opt BatchUpdateOption) error {
	// 1. 参数校验
	if err := validateBatchUpdateParams(opt); err != nil {
		return err
	}

	// 2. 获取需要更新的字段列表
	fields, err := getUpdateFields(opt)
	if err != nil {
		return err
	}

	// 3. 构建 CASE 表达式
	caseSQL, args, err := buildCaseExpressions(opt, fields)
	if err != nil {
		return err
	}

	// 4. 构建 WHERE IN 子句
	whereClause, whereValues, err := buildWhereClause(opt)
	if err != nil {
		return err
	}

	// 5. 拼接并执行 SQL
	sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s",
		opt.Table,
		strings.Join(caseSQL, ", "),
		whereClause,
	)

	// 安全检查
	if len(sql) > 800_000 {
		return fmt.Errorf("SQL too long (>800KB); consider reducing batch size")
	}

	logx.WithContext(ctx).Debugw("批量更新 SQL",
		logx.Field("sql", sql),
		logx.Field("args", append(args, whereValues...)),
	)

	return opt.DB.WithContext(ctx).Exec(sql, append(args, whereValues...)...).Error
}

// 参数校验
func validateBatchUpdateParams(opt BatchUpdateOption) error {
	if len(opt.Data) == 0 {
		return fmt.Errorf("empty data")
	}
	if len(opt.WhereKeys) == 0 {
		return fmt.Errorf("empty where keys")
	}
	if opt.DB == nil {
		return fmt.Errorf("nil db")
	}
	if opt.Table == "" {
		return fmt.Errorf("empty table name")
	}
	if len(opt.Data) > 500 {
		return fmt.Errorf("batch size exceeds limit (500)")
	}
	return nil
}

// 获取需要更新的字段
func getUpdateFields(opt BatchUpdateOption) ([]string, error) {
	fieldSet := make(map[string]struct{})
	for _, row := range opt.Data {
		for k := range row {
			skip := false
			for _, wk := range opt.WhereKeys {
				if wk == k {
					skip = true
					break
				}
			}
			if !skip {
				fieldSet[k] = struct{}{}
			}
		}
	}

	if len(fieldSet) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	fields := make([]string, 0, len(fieldSet))
	for f := range fieldSet {
		fields = append(fields, f)
	}

	return fields, nil
}

// 构建 CASE 表达式
func buildCaseExpressions(opt BatchUpdateOption, fields []string) ([]string, []interface{}, error) {
	caseSQL := make([]string, 0, len(fields))
	args := make([]interface{}, 0, len(opt.Data)*len(fields)*(len(opt.WhereKeys)+1))

	for _, field := range fields {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("`%s` = CASE", field))

		for _, row := range opt.Data {
			// 检查是否包含所有 where keys
			for _, k := range opt.WhereKeys {
				if _, ok := row[k]; !ok {
					return nil, nil, fmt.Errorf("missing where key '%s' in data", k)
				}
			}
			// 检查是否包含当前字段
			if _, ok := row[field]; !ok {
				return nil, nil, fmt.Errorf("missing field '%s' in data", field)
			}

			sb.WriteString(" WHEN ")
			if len(opt.WhereKeys) == 1 {
				sb.WriteString(fmt.Sprintf("`%s` = ?", opt.WhereKeys[0]))
				args = append(args, row[opt.WhereKeys[0]])
			} else {
				sb.WriteString("(")
				for i, k := range opt.WhereKeys {
					if i > 0 {
						sb.WriteString(" AND ")
					}
					sb.WriteString(fmt.Sprintf("`%s` = ?", k))
					args = append(args, row[k])
				}
				sb.WriteString(")")
			}
			sb.WriteString(" THEN ?")
			args = append(args, row[field])
		}
		sb.WriteString(" END")
		caseSQL = append(caseSQL, sb.String())
	}

	return caseSQL, args, nil
}

// 构建 WHERE 子句
func buildWhereClause(opt BatchUpdateOption) (string, []interface{}, error) {
	var whereClause strings.Builder
	whereValues := make([]interface{}, 0, len(opt.Data)*len(opt.WhereKeys))

	if len(opt.WhereKeys) == 1 {
		// 单字段条件
		key := opt.WhereKeys[0]
		whereClause.WriteString(fmt.Sprintf("`%s` IN (", key))
		placeholders := strings.Repeat("?,", len(opt.Data))
		whereClause.WriteString(strings.TrimRight(placeholders, ",") + ")")

		for _, row := range opt.Data {
			if val, ok := row[key]; ok {
				whereValues = append(whereValues, val)
			} else {
				return "", nil, fmt.Errorf("missing where key '%s' in data", key)
			}
		}
	} else {
		// 多字段条件
		whereClause.WriteString("(")
		for i, k := range opt.WhereKeys {
			if i > 0 {
				whereClause.WriteString(",")
			}
			whereClause.WriteString(fmt.Sprintf("`%s`", k))
		}
		whereClause.WriteString(") IN (")

		// 构建占位符
		rowPlaceholder := "(" + strings.Repeat("?,", len(opt.WhereKeys)-1) + "?),"
		allPlaceholders := strings.Repeat(rowPlaceholder, len(opt.Data))
		whereClause.WriteString(strings.TrimRight(allPlaceholders, ",") + ")")

		// 收集值
		for _, row := range opt.Data {
			for _, k := range opt.WhereKeys {
				if val, ok := row[k]; ok {
					whereValues = append(whereValues, val)
				} else {
					return "", nil, fmt.Errorf("missing where key '%s' in data", k)
				}
			}
		}
	}

	return whereClause.String(), whereValues, nil
}

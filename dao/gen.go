// generate.go
package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	//common = flag.String("c",
	//	"dev:zZblBgi6XMjIfH8h@tcp(34.96.209.243:20001)/ks_dbm2?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
	//	"数据库连接地址")
	//log = flag.String("c",
	//	"dev:zZblBgi6XMjIfH8h@tcp(34.96.209.243:20001)/ks_logs?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
	//	"数据库连接地址")
	activity = flag.String("c",
		"dev:zZblBgi6XMjIfH8h@tcp(34.96.209.243:20001)/ks_activity?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		"数据库连接地址")
	//common = flag.String("c",
	//	"joey:6NkzMqtJFsYE4L3x@tcp(20.189.74.163:23001)/ks_dbm2?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
	//	"数据库连接地址")
	//postgre = flag.String("s",
	//	"host=34.96.209.243 user=postgres password=4esg2O5qmIGY*Zyf8duEJrSTK3jzoZKB dbname=ks_dbm1 port=10005 sslmode=disable TimeZone=Asia/Shanghai",
	//	"数据库连接地址")

	outPath   string
	modelPath string
)

// main.go
//
//go:generate go run gen.go
func main() {
	flag.Args()
	//outPath = "./common/query"
	//modelPath = "./model"
	//generate(*common)

	outPath = "./common/query"
	modelPath = "./model"
	generate(*activity)
}

func generate(sqlConf string, tables ...string) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(sqlConf))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	//db, err := gorm.Open(postgres.Open(sqlConf))
	//if err != nil {
	//	panic(fmt.Errorf("cannot establish db connection: %w", err))
	//}

	// 生成实例
	config := gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      outPath,
		ModelPkgPath: modelPath,

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	}
	config.WithImportPkgPath([]string{"gorm.io/plugin/soft_delete"}...)
	g := gen.NewGenerator(config)
	// 设置目标 db
	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"timestamp": func(detailType gorm.ColumnType) (dataType string) { return "sql.NullTime" },
		"datetime":  func(detailType gorm.ColumnType) (dataType string) { return "sql.NullTime" },
		"decimal":   func(detailType gorm.ColumnType) (dataType string) { return "float64" },
		"double":    func(detailType gorm.ColumnType) (dataType string) { return "float64" },
		"float":     func(detailType gorm.ColumnType) (dataType string) { return "float64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	// jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `balance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	// })
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "updated_at")
		tag.Set("type", "bigint")
		tag.Set("autoUpdateTime", "milli")
		return tag
	})
	autoCreateTimeField := gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "created_at")
		tag.Set("type", "bigint")
		tag.Set("autoCreateTime", "milli")
		return tag
	})
	softDeleteTimeField := gen.FieldGORMTag("deleted_at", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "deleted_at")
		tag.Set("type", "bigint")
		tag.Set("softDelete", "milli")
		return tag
	})
	softDeleteField := gen.FieldType("deleted_at", "soft_delete.DeletedAt")
	tagGen := gen.FieldJSONTagWithNS(func(columnName string) string {
		return fmt.Sprintf(`%s" db:"%s`, columnName, columnName)
	})

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteTimeField, softDeleteField, tagGen}

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖

	if tables == nil {
		tables, err = db.Migrator().GetTables()
		if err != nil {
			panic(err)
		}
	}

	// 自动生成每个表的 Service 接口和 DAO
	for _, table := range tables {
		m := g.GenerateModel(table, fieldOpts...)
		for i, f := range m.Fields {
			if i == 0 {
				continue
			}
			if f.ColumnName == "created_at" || f.ColumnName == "updated_at" || f.ColumnName == "deleted_at" {
				continue
			}
			var hasDefault bool
			for k, v := range f.GORMTag {
				if k == "default" {
					hasDefault = true
					break
				}
				if k == "type" {
					for _, v1 := range v {
						// 文本类型不能设置默认值
						if strings.Contains(v1, "text") {
							hasDefault = true
							break
						}
					}
				}
			}

			if !hasDefault {
				switch f.Type {
				case "int64", "int32", "int8", "int", "uint64", "uint32", "uint", "uint8", "float64", "float32":
					m.Fields[i].GORMTag.Append("default", "0")
				case "string":
					m.Fields[i].GORMTag.Append("default", `''`)
				}
			}
		}
		g.ApplyBasic(m)
	}

	g.Execute()
}

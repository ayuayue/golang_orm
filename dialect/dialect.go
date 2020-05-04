package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

//Dialect 定义数据库连接
type Dialect interface {
	DataTypeOf(typ reflect.Value) string //将go中的数据类型转为数据库的数据类型
	TableExistSQL(tableName string) (string,[]interface{}) //判断表是否存在
}

//RegisterDialect 注册一个数据库连接方式
func RegisterDialect(name string,dialect Dialect) {
	dialectsMap[name] = dialect
}
//GetDialect 获取连接方式
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect,ok =dialectsMap[name]
	return 
}
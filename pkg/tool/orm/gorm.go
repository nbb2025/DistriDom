package orm

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

//	func QueryToMap(query string, args ...interface{}) (map[string]interface{}, error) {
//		res, err := QueryToMapArr(query, args...)
//		if res == nil || len(res) == 0 || err != nil {
//			return map[string]interface{}{}, err
//		}
//		return res[0], err
//	}
func SetDB(initDB *gorm.DB) {
	if db != nil {
		return
	}
	db = initDB
}

func DB() *gorm.DB {
	return db
}

// 避免调用QueryToMapArr导致log层级不准确
func QueryToMap(query string, args ...interface{}) (map[string]interface{}, error) {
	//ldb :=
	//logger.AutoLogQueryToAny(db.Statement.SQL.String())
	// 执行原生 SQL 查询
	rows, err := db.Raw(query, args...).Rows()
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return map[string]interface{}{}, err
	}

	var result map[string]interface{}

	// 遍历查询结果
	if rows.Next() {
		// 创建一个包含 interface{} 的 slice 来存储行数据
		columnPointers := make([]interface{}, len(columns))
		for i := range columnPointers {
			columnPointers[i] = new(interface{})
		}

		// 将行数据扫描到 columnPointers
		if err := rows.Scan(columnPointers...); err != nil {
			return map[string]interface{}{}, err
		}

		// 创建一个 map 并将列名和对应的值存储到 map 中
		result = make(map[string]interface{})
		for i, colName := range columns {
			val := columnPointers[i].(*interface{})

			// 将值转换为 string
			result[colName] = toString(*val)
		}
	} else {
		result = make(map[string]interface{})
	}

	return result, nil
}

func QueryToMapArr(query string, args ...interface{}) ([]map[string]interface{}, error) {
	//ldb :=
	//logger.AutoLogQueryToAny(db.Statement.SQL.String())
	// 执行原生 SQL 查询
	rows, err := db.Raw(query, args...).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	// 遍历查询结果
	for rows.Next() {
		// 创建一个包含 interface{} 的 slice 来存储行数据
		columnPointers := make([]interface{}, len(columns))
		for i := range columnPointers {
			columnPointers[i] = new(interface{})
		}

		// 将行数据扫描到 columnPointers
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// 创建一个 map 并将列名和对应的值存储到 map 中
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnPointers[i].(*interface{})

			// 将值转换为 string
			rowMap[colName] = toString(*val)
		}

		results = append(results, rowMap)
	}
	if results == nil {
		results = make([]map[string]interface{}, 0)
	}
	return results, nil
}

// 将 interface{} 转换为 string 的辅助函数
func toString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case []byte:
		return string(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}

func GetLangStr(lang string) string {
	langStr := ""
	switch lang {
	case "0":
		langStr = "ch"
	case "1":
		langStr = "en"
	default:
		langStr = "ch"
	}
	return langStr
}

func ParseParamIntToString(m map[string]interface{}) {
	for k, v := range m {
		switch value := v.(type) {
		case int:
			m[k] = strconv.Itoa(value)
		case int8:
			m[k] = strconv.FormatInt(int64(value), 10)
		case int16:
			m[k] = strconv.FormatInt(int64(value), 10)
		case int32:
			m[k] = strconv.FormatInt(int64(value), 10)
		case int64:
			m[k] = strconv.FormatInt(value, 10)
		case uint:
			m[k] = strconv.FormatUint(uint64(value), 10)
		case uint8:
			m[k] = strconv.FormatUint(uint64(value), 10)
		case uint16:
			m[k] = strconv.FormatUint(uint64(value), 10)
		case uint32:
			m[k] = strconv.FormatUint(uint64(value), 10)
		case uint64:
			m[k] = strconv.FormatUint(value, 10)
		case float32:
			m[k] = strconv.FormatFloat(float64(value), 'f', -1, 32)
		case float64:
			m[k] = strconv.FormatFloat(value, 'f', -1, 64)
		}
	}
}

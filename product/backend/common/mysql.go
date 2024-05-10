package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//创建mysql 连接
func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(198.19.249.180:3306)/go_mx?charset=utf8")
	return
}

//获取返回值，获取一条
func GetResultRow(rows *sql.Rows) map[string]string {
	//这行代码从结果集中获取列的名称，并将其存储在 columns 变量中
	columns, _ := rows.Columns()
	//创建了一个长度为列数的切片 scanArgs，用于在 rows.Scan 函数中存储每一列的数据
	scanArgs := make([]interface{}, len(columns))
	//创建了另一个与 scanArgs 长度相同的切片 values，用于存储每一列扫描后的具体数值
	values := make([]interface{}, len(columns))
	//这个循环将 values 切片中每个元素的地址赋给 scanArgs，以便 rows.Scan 函数可以正确地将每一列的值扫描到相应的位置。
	for j := range values {
		scanArgs[j] = &values[j]
	}
	//创建一个空的 map，用于存储结果数据
	record := make(map[string]string)
	//循环遍历结果集中的每一行
	for rows.Next() {
		//将当前行的数据扫描到values切片中
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				//如果当前值不为空，则将其转换为 string 类型并存储在 record map 中，键为该列的名称
				record[columns[i]] = string(v.([]byte))
			}
		}
	}
	return record
}

//获取所有
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	//返回所有列
	columns, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(columns))
	//这里表示一行填充数据
	scans := make([]interface{}, len(columns))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := columns[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return result
}

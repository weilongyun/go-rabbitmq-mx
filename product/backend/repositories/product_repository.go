package repositories

import (
	"backend/common"
	"backend/datamodels"
	"database/sql"
	"log"
	"strconv"
)

const PRODUCT_TABLE = "product"

//定义商品接口
type IProductRepository interface {
	Conn() error //数据库链接
	Insert(*datamodels.Product) (int64, error)
	Delete(string2 string) (bool, error)
	Update(*datamodels.Product) error
	SelectByProductId(string2 string) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}
type ProductRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

//初始化构造函数
func NewProductRepositoryManager(tableName string, mysqlConn *sql.DB) IProductRepository {
	return &ProductRepositoryManager{
		table:     tableName,
		mysqlConn: mysqlConn,
	}
}

//db链接
func (p *ProductRepositoryManager) Conn() (err error) {
	if p.mysqlConn == nil {
		// 连接mysql
		conn, err := common.NewMysqlConn()
		if err != nil {
			log.Fatalln("mysql ProductRepositoryManager Conn error", err)
			return err
		}
		p.mysqlConn = conn
	}
	if len(p.table) == 0 {
		p.table = PRODUCT_TABLE
	}
	//这里有一个小知识点，如果定义了返回名称(err error)，直接return会默认返回，当前例子return 代表 return err
	//如果返回值定义error没有定义名称就需要return nil
	return
}

//插入数据
func (p *ProductRepositoryManager) Insert(product *datamodels.Product) (id int64, err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager Insert error", err)
		return 0, err
	}
	//预编译采用占位符，防止sql注入
	sql := "insert into" + p.table + "set product_id = ？，product_name=?,product_num=?,image=?,product_price=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Prepare error", err)
		return
	}
	resp, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductPrice)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Exec error", err)
		return
	}
	return resp.LastInsertId()
}

//通过商品id删除
func (p *ProductRepositoryManager) Delete(product_id string) (isSuccess bool, err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager Delete error", err)
		return false, err
	}
	sql := "delete from" + p.table + "where product_id=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Delete Prepare error", err)
		return
	}
	_, err = stmt.Exec(product_id)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Delete Exec error", err)
		return
	}
	return true, nil
}

func (p *ProductRepositoryManager) Update(product *datamodels.Product) (err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager Update error", err)
		return err
	}
	sql := "update" + p.table + "set product_id=? and product_name=? and product_num=? and image=? and product_price=? where id=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Update Prepare error", err)
		return
	}
	_, err = stmt.Exec(product.ProductID, product.ProductName, product.ProductNum, product.ProductImage, product.ProductPrice)
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager Update Exec error", err)
		return
	}
	return
}

//通过商品id获取商品信息
func (p ProductRepositoryManager) SelectByProductId(product_id string) (product *datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager Update error", err)
		return nil, err
	}
	sql := "select * from" + p.table + "where product_id=" + product_id
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager SelectById Query error", err)
		return
	}
	//获取单条记录
	resp := common.GetResultRow(rows)
	if len(resp) == 0 {
		log.Println("mysql ProductRepositoryManager SelectById GetResultRow empty")
		return
	}
	//map转结构体 ,利用反射
	common.DataToStructByTagSql(resp, product)
	return
}

//查询所有数据
func (p *ProductRepositoryManager) SelectAll() (arrProductInfo []*datamodels.Product, err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager SelectAll error", err)
		return nil, err
	}
	sql := "select * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Fatalln("mysql ProductRepositoryManager SelectAll Query error", err)
		return
	}
	resp := common.GetResultRows(rows)
	if len(resp) == 0 {
		log.Println("mysql ProductRepositoryManager SelectAll GetResultRows empty")
		return
	}
	for _, v := range resp {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		arrProductInfo = append(arrProductInfo, product)
	}
	return
}

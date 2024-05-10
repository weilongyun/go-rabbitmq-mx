package repositories

import (
	"backend/common"
	"backend/datamodels"
	"database/sql"
	"log"
)

const PRODUCT_TABLE = "product"

//定义商品接口
type IProductRepository interface {
	Conn() error //数据库链接
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) (bool, error)
	Update(*datamodels.Product) error
	SelectById(int64) (*datamodels.Product, error)
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
func (p *ProductRepositoryManager) Insert(product *datamodels.Product) (product_id int64, err error) {
	if err := p.Conn(); err != nil {
		log.Fatalln("mysql ProductRepositoryManager Insert error", err)
		return
	}
	//预编译采用占位符，防止sql注入
	sql := "insert into" + p.table + "set product_name=?,product_num=?,image=?,product_price=?"
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

func (p *ProductRepositoryManager) Delete(i int64) (bool, error) {
	panic("implement me")
}

func (p *ProductRepositoryManager) Update(product *datamodels.Product) error {
	panic("implement me")
}

func (p ProductRepositoryManager) SelectById(i int64) (*datamodels.Product, error) {
	panic("implement me")
}

func (p *ProductRepositoryManager) SelectAll() ([]*datamodels.Product, error) {
	panic("implement me")
}

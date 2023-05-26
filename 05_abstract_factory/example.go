package abstractfactory

import "fmt"

// OrderMainDAO 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 订单详情
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory 抽象工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 关系型数据库 OrderMainDAO 实现
type RDBMainDAO struct{}

func (R *RDBMainDAO) SaveOrderMain() {
	fmt.Println("Order Main saved by RDB.")
}

// RDBDetailDAO 关系型数据库 OrderDetailDAO 实现
type RDBDetailDAO struct{}

func (RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("Order Detail saved by RDB.")
}

// RDBDAOFactory DAOFactory 的 RDB 实现
type RDBDAOFactory struct{}

func (rdb *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (rdb *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML OrderMainDAO 实现
type XMLMainDAO struct{}

func (x *XMLMainDAO) SaveOrderMain() {
	fmt.Println("Order Main saved by XML.")
}

// XMLDetailDAO XML OrderDetailDAO 实现
type XMLDetailDAO struct{}

func (XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("Order Detail saved by XML.")
}

// XMLDAOFactory DAOFactory 的 XML 实现
type XMLDAOFactory struct{}

func (xml *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (xml *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

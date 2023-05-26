package abstractfactory

import (
	"fmt"
	"testing"
)

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRDBDAOFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXMLDAOFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}

func Test(t *testing.T) {
	ExampleRDBDAOFactory()
	fmt.Println("--------------")
	ExampleXMLDAOFactory()
}

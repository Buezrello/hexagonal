package domain

type CustomRepositoryStub struct {
	customers []Customer
}

func (s CustomRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Igor", City: "Haifa", Zipcode: "110011", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Bob", City: "Jerusalem", Zipcode: "220022", DateOfBirth: "2000-02-02", Status: "1"},
	}

	return CustomRepositoryStub{customers}
}

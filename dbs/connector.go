package dbs

type DBConnector interface {
	GetInstance() DBConnector
	Connect(uri string)
	Disconnect()
}

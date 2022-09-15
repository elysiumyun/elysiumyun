package datasource

type datasource struct {
	Mysql mysqlClient
}

func (ds *datasource) InitDataSource() {
	ds.Mysql.openMysql()
}

func getDatasource() *datasource {
	return &datasource{}
}

var Get *datasource = getDatasource()

package datasource

type datasource struct{}

func (ds *datasource) InitDataSource() {

}

func getDatasource() *datasource {
	return &datasource{}
}

var Get *datasource = getDatasource()

package datasource

type datasource struct{}

func (ds *datasource) InitDataSource() {

}

func Get() *datasource {
	return &datasource{}
}

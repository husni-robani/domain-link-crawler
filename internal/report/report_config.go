package report

type Report interface{
	prepareData() interface{}
	generate(data interface{}, fileName string)
}
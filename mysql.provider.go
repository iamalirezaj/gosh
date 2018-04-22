package gosh

type MysqlProvider Provider

func (MysqlProvider) Boot(application Application) Application {

	// boot code here
	return application
}

func (provider MysqlProvider) Register(app Application) Application {

	return app
}
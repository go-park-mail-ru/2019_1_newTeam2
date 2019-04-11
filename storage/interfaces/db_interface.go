package interfaces

type DBInterface interface {
	UserManager
	LangManager
	CardManager
	DictManager
	// interface for other managers
	// method to connect to db
	// method to query sql
	// method to exec sql
	// ..
}

package interfaces

type DBInterface interface {
	UserManager
	LanguageManager
	CardManager
	DictionaryManager
	// interface for other managers
	// method to connect to db
	// method to query sql
	// method to exec sql
	// ..
}

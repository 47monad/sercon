// Code generated from Pkl module `o47monad.sercon.MongodbConfig`. DO NOT EDIT.
package mongodbconf

type URIBuilder struct {
	Username *string `pkl:"username"`

	Password *string `pkl:"password"`

	DbName string `pkl:"dbName"`

	Hosts []string `pkl:"hosts"`

	Options *Options `pkl:"options"`
}

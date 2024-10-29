// Code generated from Pkl module `o47monad.sercon.MongodbConfig`. DO NOT EDIT.
package mongodbconf

type URIBuilder struct {
	Username *string `pkl:"username" json:"username,omitempty"`

	Password *string `pkl:"password" json:"password,omitempty"`

	DbName string `pkl:"dbName" json:"dbName,omitempty"`

	Hosts []string `pkl:"hosts" json:"hosts,omitempty"`

	Options *Options `pkl:"options" json:"options,omitempty"`
}

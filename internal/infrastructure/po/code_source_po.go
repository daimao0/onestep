package po

type CodeSourcePO struct {

	// Model common fields
	Model

	// Name the local name of code source, such as GitHub project name
	Name string `gorm:"column:name;type:varchar(255) not null;default:'';comment:the local name of code source, such as GitHub project name;uniqueIndex:uk_t_code_source_name"`

	// RemoteURL the remote url of code source,such as https://github.com/onestep
	RemoteURL string `gorm:"column:remote_url;type:varchar(255) not null;default:'';comment:the code source remote url;uniqueIndex:uk_t_code_source_remote_url"`

	// RemoteToken is usually a GitHub personal access token
	RemoteToken string `gorm:"column:remote_token;type:varchar(255) not null;default:'';comment:the code source remote token;uniqueIndex:uk_t_code_source_remote_url"`

	// Desc the code source description
	Desc string `gorm:"column:desc;type:varchar(255) not null;default:'';comment:the code source description;uniqueIndex:uk_t_code_source_desc"`
}

package po

type CodeSourcePO struct {

	// Model common fields
	Model

	// Username the user of code source, such as GitHub username
	Username string `gorm:"column:username;type:varchar(255) not null;default:'';comment:such as GitHub username;uniqueIndex:uk_t_code_source_username"`

	// RemoteURL the remote url of code source,such as https://github.com/onestep
	RemoteURL string `gorm:"column:remote_url;type:varchar(255) not null;default:'';comment:the code source remote url;uniqueIndex:uk_t_code_source_remote_url"`

	// RemoteToken is usually a GitHub personal access token
	PersonalAccessToken string `gorm:"column:personal_access_token;type:varchar(255) not null;default:'';comment:github personal access token;uniqueIndex:uk_t_code_source_personal_access_token"`

	// Desc the code source description
	Desc string `gorm:"column:desc;type:varchar(255) not null;default:'';comment:the code source description;uniqueIndex:uk_t_code_source_desc"`
}

func (p *CodeSourcePO) TableName() string {
	return "t_code_source"
}

package po

type CodeSource struct {

	// Model common fields
	Model

	// Name: the local name of code source, such as GitHub project name
	Name string `gorm:"column:name;type:varchar(255) not null;default:'';comment:the local name of code source, such as GitHub project name;uniqueIndex:uk_t_code_source_name"`

	// GitURL: the git url of code source
	GitURL string `gorm:"column:git_url;type:varchar(255) not null;default:'';comment:workspace name;uniqueIndex:uk_t_code_source_git_url"`

	// Desc: the code source description
	Desc string `gorm:"column:desc;type:varchar(255) not null;default:'';comment:the code source description;uniqueIndex:uk_t_code_source_desc"`
}

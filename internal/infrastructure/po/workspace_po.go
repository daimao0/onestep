package po

// WorkspacePO maps to the database table 'workspace'
type WorkspacePO struct {
	// Model common fields
	Model

	// Name of workspace
	Name string `gorm:"column:name;type:varchar(255) not null;default:'';comment:workspace name;uniqueIndex:uk_t_workspace_name"`
}

func (p *WorkspacePO) TableName() string {
	return "t_workspace"
}

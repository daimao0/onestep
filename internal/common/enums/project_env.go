package enums

// @Author CY Yan
// @Date 2025/5/28 23:38

// Env is code source project env enum
type Env string

const (
	FAT Env = "fat"
	UAT Env = "uat"
	PRE Env = "pre"
	PRO Env = "pro"
)

func (p *Env) String() string {
	return string(*p)
}

package id_util

var snowflake = NewSnowflake(1)

// GenID use snowflake to generate global unique id
func GenID() int {
	return snowflake.GenerateID()
}

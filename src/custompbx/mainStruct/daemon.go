package mainStruct

type DaemonState struct {
	DatabaseConnection    bool  `json:"database_connection"`
	ESLConnection         bool  `json:"esl_connection"`
	DataBaseError         error `json:"data_base_error"`
	ESLError              error `json:"esl_error"`
	CdrDatabaseConnection bool  `json:"cdr_database_connection"`
	CdrDataBaseError      error `json:"cdr_data_base_error"`
	StunServerStatus      bool  `json:"stun_server_status"`
}

package atopdb

//AuditLog struct of audit log
type AuditLog struct {
	Message string `bson:"message"`
	LastTS  string `bson:"lastTS"`
}

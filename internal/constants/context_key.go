package constants

type ContextKey string

const (
	/*
		buat constant bernama ContextTenantID
		bertipe ContextKey
		dengan isi "tenant_id"
	*/
	ContextTenantID ContextKey = "tenant_id"
	ContextUserID   ContextKey = "user_id"
	ContextUsername ContextKey = "username"
	ContextRole     ContextKey = "role"
)

package domain

import (
	"encoding/json"
	"time"
)

// AuditEntry represents a single audit log entry
type AuditEntry struct {
	ID          string          `json:"id"`
	SessionID   string          `json:"sessionId"`
	UserID      string          `json:"userId"`
	Action      string          `json:"action"`
	Timestamp   time.Time       `json:"timestamp"`
	Details     json.RawMessage `json:"details,omitempty"`
	IPAddress   string          `json:"ipAddress,omitempty"`
	UserAgent   string          `json:"userAgent,omitempty"`
}

// AuditResponse represents the paginated audit log response
type AuditResponse struct {
	TotalCount int           `json:"totalCount"`
	Items      []AuditEntry  `json:"items"`
}

// AuditAction represents the type of action performed
type AuditAction string

// Common audit actions
const (
	ActionCreate    AuditAction = "create"
	ActionEdit      AuditAction = "edit"
	ActionMerge     AuditAction = "merge"
	ActionReorder   AuditAction = "reorder"
	ActionComment   AuditAction = "comment"
	ActionExport    AuditAction = "export"
	ActionShare     AuditAction = "share"
	ActionUnshare   AuditAction = "unshare"
	ActionView      AuditAction = "view"
)

// Pagination parameters
type PaginationParams struct {
	Limit  int
	Offset int
}

// Validate ensures pagination parameters are within acceptable bounds
func (p *PaginationParams) Validate() {
	if p.Limit <= 0 {
		p.Limit = 50 // default
	}
	if p.Limit > 100 {
		p.Limit = 100 // max
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
} 
package schema

// Role enum
const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleAsisten    = "asisten"
	RolePraktikan  = "praktikan"
)

// Status enum
const (
	StatusUnscheduled = "unscheduled"
	StatusScheduled   = "scheduled"
	StatusFinished    = "finished"
	StatusCompleted   = "completed"
	StatusCancelled   = "cancelled"
)

// AttendanceStatus enum
const (
	AttendanceStatusHadir      = "hadir"
	AttendanceStatusSakit      = "sakit"
	AttendanceStatusIzin       = "izin"
	AttendanceStatusTidakHadir = "tidak_hadir"
)

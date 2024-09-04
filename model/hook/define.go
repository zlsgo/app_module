package hook

type Event string

const (
	EventMigrationIndexDone Event = "MigrationIndexDone"
	EventMigrationStart     Event = "MigrationStart"
	EventMigrationDone      Event = "MigrationDone"
)

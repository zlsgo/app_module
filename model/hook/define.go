package hook

type Event string

const (
	// Migration events
	EventMigrationIndexDone Event = "MigrationIndexDone"
	EventMigrationStart     Event = "MigrationStart"
	EventMigrationDone      Event = "MigrationDone"

	// Insert events
	EventBeforeInsert Event = "BeforeInsert"
	EventAfterInsert  Event = "AfterInsert"

	// Update events
	EventBeforeUpdate Event = "BeforeUpdate"
	EventAfterUpdate  Event = "AfterUpdate"

	// Delete events
	EventBeforeDelete Event = "BeforeDelete"
	EventAfterDelete  Event = "AfterDelete"
)

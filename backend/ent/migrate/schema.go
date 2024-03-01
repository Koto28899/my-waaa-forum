// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "to_post_id", Type: field.TypeInt64},
		{Name: "likes_count", Type: field.TypeInt64, Default: 0},
		{Name: "is_top", Type: field.TypeBool, Default: false},
		{Name: "body", Type: field.TypeString, Size: 1023},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// FollowsColumns holds the columns for the "follows" table.
	FollowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "to_role_id", Type: field.TypeInt64},
	}
	// FollowsTable holds the schema information for the "follows" table.
	FollowsTable = &schema.Table{
		Name:       "follows",
		Columns:    FollowsColumns,
		PrimaryKey: []*schema.Column{FollowsColumns[0]},
	}
	// HelpsColumns holds the columns for the "helps" table.
	HelpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_post_id", Type: field.TypeInt64},
		{Name: "adopt_comment_id", Type: field.TypeInt64, Nullable: true},
	}
	// HelpsTable holds the schema information for the "helps" table.
	HelpsTable = &schema.Table{
		Name:       "helps",
		Columns:    HelpsColumns,
		PrimaryKey: []*schema.Column{HelpsColumns[0]},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "notification_type", Type: field.TypeEnum, Enums: []string{"adopt", "reply", "like", "at", "from_admin"}},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "secne_type", Type: field.TypeEnum, Enums: []string{"post", "comment", "reply"}},
		{Name: "sence_id", Type: field.TypeInt64},
		{Name: "to_role_id", Type: field.TypeInt64},
		{Name: "is_checked", Type: field.TypeBool, Default: false},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "to_section_id", Type: field.TypeInt64},
		{Name: "likes_count", Type: field.TypeInt64, Default: 0},
		{Name: "is_top", Type: field.TypeBool, Default: false},
		{Name: "is_highlight", Type: field.TypeBool, Default: false},
		{Name: "title", Type: field.TypeString, Size: 127},
		{Name: "body", Type: field.TypeString, Size: 1023},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// RepliesColumns holds the columns for the "replies" table.
	RepliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "to_post_id", Type: field.TypeInt64},
		{Name: "likes_count", Type: field.TypeInt64, Default: 0},
		{Name: "is_top", Type: field.TypeBool, Default: false},
		{Name: "body", Type: field.TypeString, Size: 1023},
	}
	// RepliesTable holds the schema information for the "replies" table.
	RepliesTable = &schema.Table{
		Name:       "replies",
		Columns:    RepliesColumns,
		PrimaryKey: []*schema.Column{RepliesColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 321},
		{Name: "hashed_pwd", Type: field.TypeString},
		{Name: "password_changed_at", Type: field.TypeTime},
		{Name: "role_name", Type: field.TypeString, Size: 21},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"user", "admin"}, Default: "user"},
		{Name: "statement", Type: field.TypeString, Nullable: true, Size: 127},
		{Name: "posts_count", Type: field.TypeInt64, Default: 0},
		{Name: "comments_count", Type: field.TypeInt64, Default: 0},
		{Name: "replies_count", Type: field.TypeInt64, Default: 0},
		{Name: "likes_count", Type: field.TypeInt64, Default: 0},
		{Name: "helps_count", Type: field.TypeInt64, Default: 0},
		{Name: "fans_count", Type: field.TypeInt64, Default: 0},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SectionColumns holds the columns for the "section" table.
	SectionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "section_name", Type: field.TypeString, Unique: true, Size: 21},
		{Name: "statement", Type: field.TypeString, Nullable: true, Size: 127},
		{Name: "manager_id", Type: field.TypeInt64},
	}
	// SectionTable holds the schema information for the "section" table.
	SectionTable = &schema.Table{
		Name:       "section",
		Columns:    SectionColumns,
		PrimaryKey: []*schema.Column{SectionColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "client_ip", Type: field.TypeString},
		{Name: "is_blocked", Type: field.TypeBool, Default: false},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// StarsColumns holds the columns for the "stars" table.
	StarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "secne_type", Type: field.TypeEnum, Enums: []string{"post", "comment", "reply"}},
		{Name: "sence_id", Type: field.TypeInt64},
	}
	// StarsTable holds the schema information for the "stars" table.
	StarsTable = &schema.Table{
		Name:       "stars",
		Columns:    StarsColumns,
		PrimaryKey: []*schema.Column{StarsColumns[0]},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_post_id", Type: field.TypeInt64},
		{Name: "register", Type: field.TypeBool, Default: false},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
	}
	// VoteEventsColumns holds the columns for the "vote_events" table.
	VoteEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "from_role_id", Type: field.TypeInt64},
		{Name: "to_vote_id", Type: field.TypeInt64},
		{Name: "to_vote_option", Type: field.TypeInt64},
	}
	// VoteEventsTable holds the schema information for the "vote_events" table.
	VoteEventsTable = &schema.Table{
		Name:       "vote_events",
		Columns:    VoteEventsColumns,
		PrimaryKey: []*schema.Column{VoteEventsColumns[0]},
	}
	// VoteOptionsColumns holds the columns for the "vote_options" table.
	VoteOptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "vote_id", Type: field.TypeInt64},
		{Name: "info", Type: field.TypeString, Size: 127},
		{Name: "count", Type: field.TypeInt64, Default: 0},
	}
	// VoteOptionsTable holds the schema information for the "vote_options" table.
	VoteOptionsTable = &schema.Table{
		Name:       "vote_options",
		Columns:    VoteOptionsColumns,
		PrimaryKey: []*schema.Column{VoteOptionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		FollowsTable,
		HelpsTable,
		NotificationsTable,
		PostsTable,
		RepliesTable,
		RolesTable,
		SectionTable,
		SessionsTable,
		StarsTable,
		VotesTable,
		VoteEventsTable,
		VoteOptionsTable,
	}
)

func init() {
	RepliesTable.Annotation = &entsql.Annotation{
		Table: "replies",
	}
	SectionTable.Annotation = &entsql.Annotation{
		Table: "section",
	}
	VoteEventsTable.Annotation = &entsql.Annotation{
		Table: "vote_events",
	}
	VoteOptionsTable.Annotation = &entsql.Annotation{
		Table: "vote_options",
	}
}
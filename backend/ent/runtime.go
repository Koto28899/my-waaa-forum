// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/comment"
	"backend/ent/follow"
	"backend/ent/help"
	"backend/ent/notification"
	"backend/ent/post"
	"backend/ent/reply"
	"backend/ent/role"
	"backend/ent/schema"
	"backend/ent/section"
	"backend/ent/session"
	"backend/ent/star"
	"backend/ent/vote"
	"backend/ent/voteevent"
	"backend/ent/voteoption"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[1].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[2].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescDeletedAt is the schema descriptor for deleted_at field.
	commentDescDeletedAt := commentFields[3].Descriptor()
	// comment.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	comment.UpdateDefaultDeletedAt = commentDescDeletedAt.UpdateDefault.(func() time.Time)
	// commentDescFromRoleID is the schema descriptor for from_role_id field.
	commentDescFromRoleID := commentFields[4].Descriptor()
	// comment.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	comment.FromRoleIDValidator = commentDescFromRoleID.Validators[0].(func(int64) error)
	// commentDescToPostID is the schema descriptor for to_post_id field.
	commentDescToPostID := commentFields[5].Descriptor()
	// comment.ToPostIDValidator is a validator for the "to_post_id" field. It is called by the builders before save.
	comment.ToPostIDValidator = commentDescToPostID.Validators[0].(func(int64) error)
	// commentDescLikesCount is the schema descriptor for likes_count field.
	commentDescLikesCount := commentFields[6].Descriptor()
	// comment.DefaultLikesCount holds the default value on creation for the likes_count field.
	comment.DefaultLikesCount = commentDescLikesCount.Default.(int64)
	// comment.LikesCountValidator is a validator for the "likes_count" field. It is called by the builders before save.
	comment.LikesCountValidator = commentDescLikesCount.Validators[0].(func(int64) error)
	// commentDescIsTop is the schema descriptor for is_top field.
	commentDescIsTop := commentFields[7].Descriptor()
	// comment.DefaultIsTop holds the default value on creation for the is_top field.
	comment.DefaultIsTop = commentDescIsTop.Default.(bool)
	// commentDescBody is the schema descriptor for body field.
	commentDescBody := commentFields[8].Descriptor()
	// comment.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	comment.BodyValidator = func() func(string) error {
		validators := commentDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	comment.IDValidator = commentDescID.Validators[0].(func(int64) error)
	followFields := schema.Follow{}.Fields()
	_ = followFields
	// followDescCreatedAt is the schema descriptor for created_at field.
	followDescCreatedAt := followFields[1].Descriptor()
	// follow.DefaultCreatedAt holds the default value on creation for the created_at field.
	follow.DefaultCreatedAt = followDescCreatedAt.Default.(func() time.Time)
	// followDescUpdatedAt is the schema descriptor for updated_at field.
	followDescUpdatedAt := followFields[2].Descriptor()
	// follow.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	follow.DefaultUpdatedAt = followDescUpdatedAt.Default.(func() time.Time)
	// follow.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	follow.UpdateDefaultUpdatedAt = followDescUpdatedAt.UpdateDefault.(func() time.Time)
	// followDescDeletedAt is the schema descriptor for deleted_at field.
	followDescDeletedAt := followFields[3].Descriptor()
	// follow.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	follow.UpdateDefaultDeletedAt = followDescDeletedAt.UpdateDefault.(func() time.Time)
	// followDescFromRoleID is the schema descriptor for from_role_id field.
	followDescFromRoleID := followFields[4].Descriptor()
	// follow.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	follow.FromRoleIDValidator = followDescFromRoleID.Validators[0].(func(int64) error)
	// followDescToRoleID is the schema descriptor for to_role_id field.
	followDescToRoleID := followFields[5].Descriptor()
	// follow.ToRoleIDValidator is a validator for the "to_role_id" field. It is called by the builders before save.
	follow.ToRoleIDValidator = followDescToRoleID.Validators[0].(func(int64) error)
	// followDescID is the schema descriptor for id field.
	followDescID := followFields[0].Descriptor()
	// follow.IDValidator is a validator for the "id" field. It is called by the builders before save.
	follow.IDValidator = followDescID.Validators[0].(func(int64) error)
	helpFields := schema.Help{}.Fields()
	_ = helpFields
	// helpDescCreatedAt is the schema descriptor for created_at field.
	helpDescCreatedAt := helpFields[1].Descriptor()
	// help.DefaultCreatedAt holds the default value on creation for the created_at field.
	help.DefaultCreatedAt = helpDescCreatedAt.Default.(func() time.Time)
	// helpDescUpdatedAt is the schema descriptor for updated_at field.
	helpDescUpdatedAt := helpFields[2].Descriptor()
	// help.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	help.DefaultUpdatedAt = helpDescUpdatedAt.Default.(func() time.Time)
	// help.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	help.UpdateDefaultUpdatedAt = helpDescUpdatedAt.UpdateDefault.(func() time.Time)
	// helpDescDeletedAt is the schema descriptor for deleted_at field.
	helpDescDeletedAt := helpFields[3].Descriptor()
	// help.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	help.UpdateDefaultDeletedAt = helpDescDeletedAt.UpdateDefault.(func() time.Time)
	// helpDescFromPostID is the schema descriptor for from_post_id field.
	helpDescFromPostID := helpFields[4].Descriptor()
	// help.FromPostIDValidator is a validator for the "from_post_id" field. It is called by the builders before save.
	help.FromPostIDValidator = helpDescFromPostID.Validators[0].(func(int64) error)
	// helpDescAdoptCommentID is the schema descriptor for adopt_comment_id field.
	helpDescAdoptCommentID := helpFields[5].Descriptor()
	// help.AdoptCommentIDValidator is a validator for the "adopt_comment_id" field. It is called by the builders before save.
	help.AdoptCommentIDValidator = helpDescAdoptCommentID.Validators[0].(func(int64) error)
	// helpDescID is the schema descriptor for id field.
	helpDescID := helpFields[0].Descriptor()
	// help.IDValidator is a validator for the "id" field. It is called by the builders before save.
	help.IDValidator = helpDescID.Validators[0].(func(int64) error)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[1].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationFields[2].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescDeletedAt is the schema descriptor for deleted_at field.
	notificationDescDeletedAt := notificationFields[3].Descriptor()
	// notification.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	notification.UpdateDefaultDeletedAt = notificationDescDeletedAt.UpdateDefault.(func() time.Time)
	// notificationDescFromRoleID is the schema descriptor for from_role_id field.
	notificationDescFromRoleID := notificationFields[5].Descriptor()
	// notification.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	notification.FromRoleIDValidator = notificationDescFromRoleID.Validators[0].(func(int64) error)
	// notificationDescSenceID is the schema descriptor for sence_id field.
	notificationDescSenceID := notificationFields[7].Descriptor()
	// notification.SenceIDValidator is a validator for the "sence_id" field. It is called by the builders before save.
	notification.SenceIDValidator = notificationDescSenceID.Validators[0].(func(int64) error)
	// notificationDescToRoleID is the schema descriptor for to_role_id field.
	notificationDescToRoleID := notificationFields[8].Descriptor()
	// notification.ToRoleIDValidator is a validator for the "to_role_id" field. It is called by the builders before save.
	notification.ToRoleIDValidator = notificationDescToRoleID.Validators[0].(func(int64) error)
	// notificationDescIsChecked is the schema descriptor for is_checked field.
	notificationDescIsChecked := notificationFields[9].Descriptor()
	// notification.DefaultIsChecked holds the default value on creation for the is_checked field.
	notification.DefaultIsChecked = notificationDescIsChecked.Default.(bool)
	// notificationDescID is the schema descriptor for id field.
	notificationDescID := notificationFields[0].Descriptor()
	// notification.IDValidator is a validator for the "id" field. It is called by the builders before save.
	notification.IDValidator = notificationDescID.Validators[0].(func(int64) error)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[1].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[2].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescDeletedAt is the schema descriptor for deleted_at field.
	postDescDeletedAt := postFields[3].Descriptor()
	// post.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	post.UpdateDefaultDeletedAt = postDescDeletedAt.UpdateDefault.(func() time.Time)
	// postDescFromRoleID is the schema descriptor for from_role_id field.
	postDescFromRoleID := postFields[4].Descriptor()
	// post.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	post.FromRoleIDValidator = postDescFromRoleID.Validators[0].(func(int64) error)
	// postDescToSectionID is the schema descriptor for to_section_id field.
	postDescToSectionID := postFields[5].Descriptor()
	// post.ToSectionIDValidator is a validator for the "to_section_id" field. It is called by the builders before save.
	post.ToSectionIDValidator = postDescToSectionID.Validators[0].(func(int64) error)
	// postDescLikesCount is the schema descriptor for likes_count field.
	postDescLikesCount := postFields[6].Descriptor()
	// post.DefaultLikesCount holds the default value on creation for the likes_count field.
	post.DefaultLikesCount = postDescLikesCount.Default.(int64)
	// post.LikesCountValidator is a validator for the "likes_count" field. It is called by the builders before save.
	post.LikesCountValidator = postDescLikesCount.Validators[0].(func(int64) error)
	// postDescIsTop is the schema descriptor for is_top field.
	postDescIsTop := postFields[7].Descriptor()
	// post.DefaultIsTop holds the default value on creation for the is_top field.
	post.DefaultIsTop = postDescIsTop.Default.(bool)
	// postDescIsHighlight is the schema descriptor for is_highlight field.
	postDescIsHighlight := postFields[8].Descriptor()
	// post.DefaultIsHighlight holds the default value on creation for the is_highlight field.
	post.DefaultIsHighlight = postDescIsHighlight.Default.(bool)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[9].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = func() func(string) error {
		validators := postDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescBody is the schema descriptor for body field.
	postDescBody := postFields[10].Descriptor()
	// post.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	post.BodyValidator = func() func(string) error {
		validators := postDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.IDValidator is a validator for the "id" field. It is called by the builders before save.
	post.IDValidator = postDescID.Validators[0].(func(int64) error)
	replyFields := schema.Reply{}.Fields()
	_ = replyFields
	// replyDescCreatedAt is the schema descriptor for created_at field.
	replyDescCreatedAt := replyFields[1].Descriptor()
	// reply.DefaultCreatedAt holds the default value on creation for the created_at field.
	reply.DefaultCreatedAt = replyDescCreatedAt.Default.(func() time.Time)
	// replyDescUpdatedAt is the schema descriptor for updated_at field.
	replyDescUpdatedAt := replyFields[2].Descriptor()
	// reply.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reply.DefaultUpdatedAt = replyDescUpdatedAt.Default.(func() time.Time)
	// reply.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reply.UpdateDefaultUpdatedAt = replyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// replyDescDeletedAt is the schema descriptor for deleted_at field.
	replyDescDeletedAt := replyFields[3].Descriptor()
	// reply.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	reply.UpdateDefaultDeletedAt = replyDescDeletedAt.UpdateDefault.(func() time.Time)
	// replyDescFromRoleID is the schema descriptor for from_role_id field.
	replyDescFromRoleID := replyFields[4].Descriptor()
	// reply.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	reply.FromRoleIDValidator = replyDescFromRoleID.Validators[0].(func(int64) error)
	// replyDescToPostID is the schema descriptor for to_post_id field.
	replyDescToPostID := replyFields[5].Descriptor()
	// reply.ToPostIDValidator is a validator for the "to_post_id" field. It is called by the builders before save.
	reply.ToPostIDValidator = replyDescToPostID.Validators[0].(func(int64) error)
	// replyDescLikesCount is the schema descriptor for likes_count field.
	replyDescLikesCount := replyFields[6].Descriptor()
	// reply.DefaultLikesCount holds the default value on creation for the likes_count field.
	reply.DefaultLikesCount = replyDescLikesCount.Default.(int64)
	// reply.LikesCountValidator is a validator for the "likes_count" field. It is called by the builders before save.
	reply.LikesCountValidator = replyDescLikesCount.Validators[0].(func(int64) error)
	// replyDescIsTop is the schema descriptor for is_top field.
	replyDescIsTop := replyFields[7].Descriptor()
	// reply.DefaultIsTop holds the default value on creation for the is_top field.
	reply.DefaultIsTop = replyDescIsTop.Default.(bool)
	// replyDescBody is the schema descriptor for body field.
	replyDescBody := replyFields[8].Descriptor()
	// reply.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	reply.BodyValidator = func() func(string) error {
		validators := replyDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// replyDescID is the schema descriptor for id field.
	replyDescID := replyFields[0].Descriptor()
	// reply.IDValidator is a validator for the "id" field. It is called by the builders before save.
	reply.IDValidator = replyDescID.Validators[0].(func(int64) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[1].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[2].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescDeletedAt is the schema descriptor for deleted_at field.
	roleDescDeletedAt := roleFields[3].Descriptor()
	// role.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	role.UpdateDefaultDeletedAt = roleDescDeletedAt.UpdateDefault.(func() time.Time)
	// roleDescEmail is the schema descriptor for email field.
	roleDescEmail := roleFields[4].Descriptor()
	// role.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	role.EmailValidator = func() func(string) error {
		validators := roleDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roleDescHashedPwd is the schema descriptor for hashed_pwd field.
	roleDescHashedPwd := roleFields[5].Descriptor()
	// role.HashedPwdValidator is a validator for the "hashed_pwd" field. It is called by the builders before save.
	role.HashedPwdValidator = roleDescHashedPwd.Validators[0].(func(string) error)
	// roleDescPasswordChangedAt is the schema descriptor for password_changed_at field.
	roleDescPasswordChangedAt := roleFields[6].Descriptor()
	// role.DefaultPasswordChangedAt holds the default value on creation for the password_changed_at field.
	role.DefaultPasswordChangedAt = roleDescPasswordChangedAt.Default.(func() time.Time)
	// roleDescRoleName is the schema descriptor for role_name field.
	roleDescRoleName := roleFields[7].Descriptor()
	// role.RoleNameValidator is a validator for the "role_name" field. It is called by the builders before save.
	role.RoleNameValidator = func() func(string) error {
		validators := roleDescRoleName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(role_name string) error {
			for _, fn := range fns {
				if err := fn(role_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roleDescStatement is the schema descriptor for statement field.
	roleDescStatement := roleFields[9].Descriptor()
	// role.StatementValidator is a validator for the "statement" field. It is called by the builders before save.
	role.StatementValidator = roleDescStatement.Validators[0].(func(string) error)
	// roleDescPostsCount is the schema descriptor for posts_count field.
	roleDescPostsCount := roleFields[10].Descriptor()
	// role.DefaultPostsCount holds the default value on creation for the posts_count field.
	role.DefaultPostsCount = roleDescPostsCount.Default.(int64)
	// role.PostsCountValidator is a validator for the "posts_count" field. It is called by the builders before save.
	role.PostsCountValidator = roleDescPostsCount.Validators[0].(func(int64) error)
	// roleDescCommentsCount is the schema descriptor for comments_count field.
	roleDescCommentsCount := roleFields[11].Descriptor()
	// role.DefaultCommentsCount holds the default value on creation for the comments_count field.
	role.DefaultCommentsCount = roleDescCommentsCount.Default.(int64)
	// role.CommentsCountValidator is a validator for the "comments_count" field. It is called by the builders before save.
	role.CommentsCountValidator = roleDescCommentsCount.Validators[0].(func(int64) error)
	// roleDescRepliesCount is the schema descriptor for replies_count field.
	roleDescRepliesCount := roleFields[12].Descriptor()
	// role.DefaultRepliesCount holds the default value on creation for the replies_count field.
	role.DefaultRepliesCount = roleDescRepliesCount.Default.(int64)
	// role.RepliesCountValidator is a validator for the "replies_count" field. It is called by the builders before save.
	role.RepliesCountValidator = roleDescRepliesCount.Validators[0].(func(int64) error)
	// roleDescLikesCount is the schema descriptor for likes_count field.
	roleDescLikesCount := roleFields[13].Descriptor()
	// role.DefaultLikesCount holds the default value on creation for the likes_count field.
	role.DefaultLikesCount = roleDescLikesCount.Default.(int64)
	// role.LikesCountValidator is a validator for the "likes_count" field. It is called by the builders before save.
	role.LikesCountValidator = roleDescLikesCount.Validators[0].(func(int64) error)
	// roleDescHelpsCount is the schema descriptor for helps_count field.
	roleDescHelpsCount := roleFields[14].Descriptor()
	// role.DefaultHelpsCount holds the default value on creation for the helps_count field.
	role.DefaultHelpsCount = roleDescHelpsCount.Default.(int64)
	// role.HelpsCountValidator is a validator for the "helps_count" field. It is called by the builders before save.
	role.HelpsCountValidator = roleDescHelpsCount.Validators[0].(func(int64) error)
	// roleDescFansCount is the schema descriptor for fans_count field.
	roleDescFansCount := roleFields[15].Descriptor()
	// role.DefaultFansCount holds the default value on creation for the fans_count field.
	role.DefaultFansCount = roleDescFansCount.Default.(int64)
	// role.FansCountValidator is a validator for the "fans_count" field. It is called by the builders before save.
	role.FansCountValidator = roleDescFansCount.Validators[0].(func(int64) error)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[0].Descriptor()
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(int64) error)
	sectionFields := schema.Section{}.Fields()
	_ = sectionFields
	// sectionDescCreatedAt is the schema descriptor for created_at field.
	sectionDescCreatedAt := sectionFields[1].Descriptor()
	// section.DefaultCreatedAt holds the default value on creation for the created_at field.
	section.DefaultCreatedAt = sectionDescCreatedAt.Default.(func() time.Time)
	// sectionDescUpdatedAt is the schema descriptor for updated_at field.
	sectionDescUpdatedAt := sectionFields[2].Descriptor()
	// section.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	section.DefaultUpdatedAt = sectionDescUpdatedAt.Default.(func() time.Time)
	// section.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	section.UpdateDefaultUpdatedAt = sectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sectionDescDeletedAt is the schema descriptor for deleted_at field.
	sectionDescDeletedAt := sectionFields[3].Descriptor()
	// section.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	section.UpdateDefaultDeletedAt = sectionDescDeletedAt.UpdateDefault.(func() time.Time)
	// sectionDescSectionName is the schema descriptor for section_name field.
	sectionDescSectionName := sectionFields[4].Descriptor()
	// section.SectionNameValidator is a validator for the "section_name" field. It is called by the builders before save.
	section.SectionNameValidator = func() func(string) error {
		validators := sectionDescSectionName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(section_name string) error {
			for _, fn := range fns {
				if err := fn(section_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sectionDescStatement is the schema descriptor for statement field.
	sectionDescStatement := sectionFields[5].Descriptor()
	// section.StatementValidator is a validator for the "statement" field. It is called by the builders before save.
	section.StatementValidator = sectionDescStatement.Validators[0].(func(string) error)
	// sectionDescManagerID is the schema descriptor for manager_id field.
	sectionDescManagerID := sectionFields[6].Descriptor()
	// section.ManagerIDValidator is a validator for the "manager_id" field. It is called by the builders before save.
	section.ManagerIDValidator = sectionDescManagerID.Validators[0].(func(int64) error)
	// sectionDescID is the schema descriptor for id field.
	sectionDescID := sectionFields[0].Descriptor()
	// section.IDValidator is a validator for the "id" field. It is called by the builders before save.
	section.IDValidator = sectionDescID.Validators[0].(func(int64) error)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescRoleID is the schema descriptor for role_id field.
	sessionDescRoleID := sessionFields[1].Descriptor()
	// session.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	session.RoleIDValidator = sessionDescRoleID.Validators[0].(func(int64) error)
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[2].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// session.UpdateDefaultCreatedAt holds the default value on update for the created_at field.
	session.UpdateDefaultCreatedAt = sessionDescCreatedAt.UpdateDefault.(func() time.Time)
	// sessionDescExpiresAt is the schema descriptor for expires_at field.
	sessionDescExpiresAt := sessionFields[3].Descriptor()
	// session.DefaultExpiresAt holds the default value on creation for the expires_at field.
	session.DefaultExpiresAt = sessionDescExpiresAt.Default.(func() time.Time)
	// sessionDescUserAgent is the schema descriptor for user_agent field.
	sessionDescUserAgent := sessionFields[4].Descriptor()
	// session.UserAgentValidator is a validator for the "user_agent" field. It is called by the builders before save.
	session.UserAgentValidator = sessionDescUserAgent.Validators[0].(func(string) error)
	// sessionDescClientIP is the schema descriptor for client_ip field.
	sessionDescClientIP := sessionFields[5].Descriptor()
	// session.ClientIPValidator is a validator for the "client_ip" field. It is called by the builders before save.
	session.ClientIPValidator = sessionDescClientIP.Validators[0].(func(string) error)
	// sessionDescIsBlocked is the schema descriptor for is_blocked field.
	sessionDescIsBlocked := sessionFields[6].Descriptor()
	// session.DefaultIsBlocked holds the default value on creation for the is_blocked field.
	session.DefaultIsBlocked = sessionDescIsBlocked.Default.(bool)
	starFields := schema.Star{}.Fields()
	_ = starFields
	// starDescCreatedAt is the schema descriptor for created_at field.
	starDescCreatedAt := starFields[1].Descriptor()
	// star.DefaultCreatedAt holds the default value on creation for the created_at field.
	star.DefaultCreatedAt = starDescCreatedAt.Default.(func() time.Time)
	// starDescUpdatedAt is the schema descriptor for updated_at field.
	starDescUpdatedAt := starFields[2].Descriptor()
	// star.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	star.DefaultUpdatedAt = starDescUpdatedAt.Default.(func() time.Time)
	// star.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	star.UpdateDefaultUpdatedAt = starDescUpdatedAt.UpdateDefault.(func() time.Time)
	// starDescDeletedAt is the schema descriptor for deleted_at field.
	starDescDeletedAt := starFields[3].Descriptor()
	// star.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	star.UpdateDefaultDeletedAt = starDescDeletedAt.UpdateDefault.(func() time.Time)
	// starDescFromRoleID is the schema descriptor for from_role_id field.
	starDescFromRoleID := starFields[4].Descriptor()
	// star.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	star.FromRoleIDValidator = starDescFromRoleID.Validators[0].(func(int64) error)
	// starDescSenceID is the schema descriptor for sence_id field.
	starDescSenceID := starFields[6].Descriptor()
	// star.SenceIDValidator is a validator for the "sence_id" field. It is called by the builders before save.
	star.SenceIDValidator = starDescSenceID.Validators[0].(func(int64) error)
	// starDescID is the schema descriptor for id field.
	starDescID := starFields[0].Descriptor()
	// star.IDValidator is a validator for the "id" field. It is called by the builders before save.
	star.IDValidator = starDescID.Validators[0].(func(int64) error)
	voteFields := schema.Vote{}.Fields()
	_ = voteFields
	// voteDescCreatedAt is the schema descriptor for created_at field.
	voteDescCreatedAt := voteFields[1].Descriptor()
	// vote.DefaultCreatedAt holds the default value on creation for the created_at field.
	vote.DefaultCreatedAt = voteDescCreatedAt.Default.(func() time.Time)
	// voteDescUpdatedAt is the schema descriptor for updated_at field.
	voteDescUpdatedAt := voteFields[2].Descriptor()
	// vote.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vote.DefaultUpdatedAt = voteDescUpdatedAt.Default.(func() time.Time)
	// vote.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vote.UpdateDefaultUpdatedAt = voteDescUpdatedAt.UpdateDefault.(func() time.Time)
	// voteDescDeletedAt is the schema descriptor for deleted_at field.
	voteDescDeletedAt := voteFields[3].Descriptor()
	// vote.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	vote.UpdateDefaultDeletedAt = voteDescDeletedAt.UpdateDefault.(func() time.Time)
	// voteDescFromPostID is the schema descriptor for from_post_id field.
	voteDescFromPostID := voteFields[4].Descriptor()
	// vote.FromPostIDValidator is a validator for the "from_post_id" field. It is called by the builders before save.
	vote.FromPostIDValidator = voteDescFromPostID.Validators[0].(func(int64) error)
	// voteDescRegister is the schema descriptor for register field.
	voteDescRegister := voteFields[5].Descriptor()
	// vote.DefaultRegister holds the default value on creation for the register field.
	vote.DefaultRegister = voteDescRegister.Default.(bool)
	// voteDescID is the schema descriptor for id field.
	voteDescID := voteFields[0].Descriptor()
	// vote.IDValidator is a validator for the "id" field. It is called by the builders before save.
	vote.IDValidator = voteDescID.Validators[0].(func(int64) error)
	voteeventFields := schema.VoteEvent{}.Fields()
	_ = voteeventFields
	// voteeventDescCreatedAt is the schema descriptor for created_at field.
	voteeventDescCreatedAt := voteeventFields[1].Descriptor()
	// voteevent.DefaultCreatedAt holds the default value on creation for the created_at field.
	voteevent.DefaultCreatedAt = voteeventDescCreatedAt.Default.(func() time.Time)
	// voteeventDescUpdatedAt is the schema descriptor for updated_at field.
	voteeventDescUpdatedAt := voteeventFields[2].Descriptor()
	// voteevent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	voteevent.DefaultUpdatedAt = voteeventDescUpdatedAt.Default.(func() time.Time)
	// voteevent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	voteevent.UpdateDefaultUpdatedAt = voteeventDescUpdatedAt.UpdateDefault.(func() time.Time)
	// voteeventDescDeletedAt is the schema descriptor for deleted_at field.
	voteeventDescDeletedAt := voteeventFields[3].Descriptor()
	// voteevent.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	voteevent.UpdateDefaultDeletedAt = voteeventDescDeletedAt.UpdateDefault.(func() time.Time)
	// voteeventDescFromRoleID is the schema descriptor for from_role_id field.
	voteeventDescFromRoleID := voteeventFields[4].Descriptor()
	// voteevent.FromRoleIDValidator is a validator for the "from_role_id" field. It is called by the builders before save.
	voteevent.FromRoleIDValidator = voteeventDescFromRoleID.Validators[0].(func(int64) error)
	// voteeventDescToVoteID is the schema descriptor for to_vote_id field.
	voteeventDescToVoteID := voteeventFields[5].Descriptor()
	// voteevent.ToVoteIDValidator is a validator for the "to_vote_id" field. It is called by the builders before save.
	voteevent.ToVoteIDValidator = voteeventDescToVoteID.Validators[0].(func(int64) error)
	// voteeventDescToVoteOption is the schema descriptor for to_vote_option field.
	voteeventDescToVoteOption := voteeventFields[6].Descriptor()
	// voteevent.ToVoteOptionValidator is a validator for the "to_vote_option" field. It is called by the builders before save.
	voteevent.ToVoteOptionValidator = voteeventDescToVoteOption.Validators[0].(func(int64) error)
	// voteeventDescID is the schema descriptor for id field.
	voteeventDescID := voteeventFields[0].Descriptor()
	// voteevent.IDValidator is a validator for the "id" field. It is called by the builders before save.
	voteevent.IDValidator = voteeventDescID.Validators[0].(func(int64) error)
	voteoptionFields := schema.VoteOption{}.Fields()
	_ = voteoptionFields
	// voteoptionDescCreatedAt is the schema descriptor for created_at field.
	voteoptionDescCreatedAt := voteoptionFields[1].Descriptor()
	// voteoption.DefaultCreatedAt holds the default value on creation for the created_at field.
	voteoption.DefaultCreatedAt = voteoptionDescCreatedAt.Default.(func() time.Time)
	// voteoptionDescUpdatedAt is the schema descriptor for updated_at field.
	voteoptionDescUpdatedAt := voteoptionFields[2].Descriptor()
	// voteoption.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	voteoption.DefaultUpdatedAt = voteoptionDescUpdatedAt.Default.(func() time.Time)
	// voteoption.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	voteoption.UpdateDefaultUpdatedAt = voteoptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// voteoptionDescDeletedAt is the schema descriptor for deleted_at field.
	voteoptionDescDeletedAt := voteoptionFields[3].Descriptor()
	// voteoption.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	voteoption.UpdateDefaultDeletedAt = voteoptionDescDeletedAt.UpdateDefault.(func() time.Time)
	// voteoptionDescVoteID is the schema descriptor for vote_id field.
	voteoptionDescVoteID := voteoptionFields[4].Descriptor()
	// voteoption.VoteIDValidator is a validator for the "vote_id" field. It is called by the builders before save.
	voteoption.VoteIDValidator = voteoptionDescVoteID.Validators[0].(func(int64) error)
	// voteoptionDescInfo is the schema descriptor for info field.
	voteoptionDescInfo := voteoptionFields[5].Descriptor()
	// voteoption.InfoValidator is a validator for the "info" field. It is called by the builders before save.
	voteoption.InfoValidator = func() func(string) error {
		validators := voteoptionDescInfo.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(info string) error {
			for _, fn := range fns {
				if err := fn(info); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// voteoptionDescCount is the schema descriptor for count field.
	voteoptionDescCount := voteoptionFields[6].Descriptor()
	// voteoption.DefaultCount holds the default value on creation for the count field.
	voteoption.DefaultCount = voteoptionDescCount.Default.(int64)
	// voteoption.CountValidator is a validator for the "count" field. It is called by the builders before save.
	voteoption.CountValidator = voteoptionDescCount.Validators[0].(func(int64) error)
	// voteoptionDescID is the schema descriptor for id field.
	voteoptionDescID := voteoptionFields[0].Descriptor()
	// voteoption.IDValidator is a validator for the "id" field. It is called by the builders before save.
	voteoption.IDValidator = voteoptionDescID.Validators[0].(func(int64) error)
}

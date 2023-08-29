package dao

import (
	"fmt"
	"strings"

	"github.com/ArtalkJS/Artalk/internal/cache"
	"github.com/ArtalkJS/Artalk/internal/entity"
)

// TODO modify strange function parameter table
func (dao *Dao) FindAndStoreCache(name string, dest interface{}, queryDBResult func() interface{}) error {
	if dao.cache == nil {
		// directly call queryDB while cache is disabled
		_ = queryDBResult()
		return nil
	}

	return cache.FindAndStoreCache(name, dest, queryDBResult)
}

func (dao *Dao) FindComment(id uint, checkers ...func(*entity.Comment) bool) entity.Comment {
	var comment entity.Comment

	dao.FindAndStoreCache(fmt.Sprintf("comment#id=%d", id), &comment, func() interface{} {
		dao.DB().Where("id = ?", id).First(&comment)
		return &comment
	})

	// the case with checkers
	for _, c := range checkers {
		if !c(&comment) {
			return entity.Comment{}
		}
	}

	return comment
}

// (Cached：parent-comments)
func (dao *Dao) FindCommentChildrenShallow(parentID uint, checkers ...func(*entity.Comment) bool) []entity.Comment {
	var children []entity.Comment
	var childIDs []uint

	dao.FindAndStoreCache(fmt.Sprintf("parent-comments#pid=%d", parentID), &childIDs, func() interface{} {
		dao.DB().Model(&entity.Comment{}).Where(&entity.Comment{Rid: parentID}).Select("id").Find(&childIDs)
		return &childIDs
	})

	for _, childID := range childIDs {
		child := dao.FindComment(childID, checkers...)
		if !child.IsEmpty() {
			children = append(children, child)
		}
	}

	return children
}

func (dao *Dao) FindCommentChildren(parentID uint, checkers ...func(*entity.Comment) bool) []entity.Comment {
	allChildren := []entity.Comment{}
	dao._findCommentChildrenOnce(&allChildren, parentID, checkers...) // TODO: children 数量限制
	return allChildren
}

func (dao *Dao) _findCommentChildrenOnce(source *[]entity.Comment, parentID uint, checkers ...func(*entity.Comment) bool) {
	// TODO 子评论排序问题
	children := dao.FindCommentChildrenShallow(parentID, checkers...)

	for _, child := range children {
		*source = append(*source, child)
		dao._findCommentChildrenOnce(source, child.ID, checkers...) // recurse
	}
}

// 查找用户 (精确查找 name & email)
func (dao *Dao) FindUser(name string, email string) entity.User {
	var user entity.User

	// 查询缓存
	dao.FindAndStoreCache(fmt.Sprintf("user#name=%s;email=%s", strings.ToLower(name), strings.ToLower(email)), &user, func() interface{} {
		// 不区分大小写
		dao.DB().Where("LOWER(name) = LOWER(?) AND LOWER(email) = LOWER(?)", name, email).First(&user)
		return &user
	})

	return user
}

// 查找用户 ID (仅根据 email)
func (dao *Dao) FindUserIdsByEmail(email string) []uint {
	var userIds = []uint{}

	// 查询缓存
	dao.FindAndStoreCache(fmt.Sprintf("user_id#email=%s", strings.ToLower(email)), &userIds, func() interface{} {
		dao.DB().Model(&entity.User{}).Where("LOWER(email) = LOWER(?)", email).Pluck("id", &userIds)

		return &userIds
	})

	return userIds
}

// 查找用户 (仅根据 email)
func (dao *Dao) FindUsersByEmail(email string) []entity.User {
	userIds := dao.FindUserIdsByEmail(email)

	users := []entity.User{}
	for _, id := range userIds {
		users = append(users, dao.FindUserByID(id))
	}

	return users
}

// 查找用户 (通过 ID)
func (dao *Dao) FindUserByID(id uint) entity.User {
	var user entity.User

	// 查询缓存
	dao.FindAndStoreCache(fmt.Sprintf("user#id=%d", id), &user, func() interface{} {
		dao.DB().Where("id = ?", id).First(&user)
		return &user
	})

	return user
}

func (dao *Dao) FindPage(key string, siteName string) entity.Page {
	var page entity.Page

	dao.FindAndStoreCache(fmt.Sprintf("page#key=%s;site_name=%s", key, siteName), &page, func() interface{} {
		dao.DB().Where(&entity.Page{Key: key, SiteName: siteName}).First(&page)
		return &page
	})

	return page
}

func (dao *Dao) FindPageByID(id uint) entity.Page {
	var page entity.Page

	dao.FindAndStoreCache(fmt.Sprintf("page#id=%d", id), &page, func() interface{} {
		dao.DB().Where("id = ?", id).First(&page)
		return &page
	})

	return page
}

func (dao *Dao) FindSite(name string) entity.Site {
	var site entity.Site

	// 查询缓存
	dao.FindAndStoreCache(fmt.Sprintf("site#name=%s", name), &site, func() interface{} {
		dao.DB().Where("name = ?", name).First(&site)
		return &site
	})

	return site
}

func (dao *Dao) FindSiteByID(id uint) entity.Site {
	var site entity.Site

	dao.FindAndStoreCache(fmt.Sprintf("site#id=%d", id), &site, func() interface{} {
		dao.DB().Where("id = ?", id).First(&site)
		return &site
	})

	return site
}

func (dao *Dao) FindAllSites() []entity.Site {
	var sites []entity.Site
	dao.DB().Model(&entity.Site{}).Find(&sites)

	return sites
}

// #region Notify
func (dao *Dao) FindNotify(userID uint, commentID uint) entity.Notify {
	var notify entity.Notify
	dao.DB().Where("user_id = ? AND comment_id = ?", userID, commentID).First(&notify)
	return notify
}

func (dao *Dao) FindNotifyForComment(commentID uint, key string) entity.Notify {
	var notify entity.Notify
	dao.DB().Where(entity.Notify{CommentID: commentID, Key: key}).First(&notify)
	return notify
}

func (dao *Dao) FindUnreadNotifies(userID uint) []entity.Notify {
	if userID == 0 {
		return []entity.Notify{}
	}

	var notifies []entity.Notify
	dao.DB().Where("user_id = ? AND is_read = ?", userID, false).Find(&notifies)

	return notifies
}

func (dao *Dao) FindNotifyParentComment(n *entity.Notify) entity.Comment {
	comment := dao.FetchCommentForNotify(n)
	if comment.Rid == 0 {
		return entity.Comment{}
	}

	return dao.FindComment(comment.Rid)
}

//#endregion

// #region Vote
func (dao *Dao) GetVoteNum(targetID uint, voteType string) int {
	var num int64
	dao.DB().Model(&entity.Vote{}).Where("target_id = ? AND type = ?", targetID, voteType).Count(&num)
	return int(num)
}

func (dao *Dao) GetVoteNumUpDown(targetID uint, voteTo string) (int, int) {
	var up int64
	var down int64
	dao.DB().Model(&entity.Vote{}).Where("target_id = ? AND type = ?", targetID, voteTo+"_up").Count(&up)
	dao.DB().Model(&entity.Vote{}).Where("target_id = ? AND type = ?", targetID, voteTo+"_down").Count(&down)
	return int(up), int(down)
}

//#endregion

// #region 管理员账号检测
var allAdmins *[]entity.User = nil

func (dao *Dao) GetAllAdmins() []entity.User {
	if allAdmins == nil {
		var admins []entity.User
		dao.DB().Where(&entity.User{IsAdmin: true}).Find(&admins)
		allAdmins = &admins
	}

	return *allAdmins
}

func (dao *Dao) GetAllAdminIDs() []uint {
	admins := dao.GetAllAdmins()
	ids := []uint{}
	for _, a := range admins {
		ids = append(ids, a.ID)
	}
	return ids
}

func (dao *Dao) IsAdminUser(userID uint) bool {
	admins := dao.GetAllAdmins()
	for _, admin := range admins {
		if admin.ID == userID {
			return true
		}
	}

	return false
}

func (dao *Dao) IsAdminUserByNameEmail(name string, email string) bool {
	admins := dao.GetAllAdmins()
	for _, admin := range admins {
		// Name 和 Email 都匹配才是管理员
		if strings.EqualFold(admin.Name, name) &&
			strings.EqualFold(admin.Email, email) {
			return true
		}
	}

	return false
}

//#endregion

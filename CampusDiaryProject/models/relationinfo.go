package models

const (
	NoRelationship   = iota //没关系
	Follow                  //是本人关注的人
	Fans                    //是本人的粉丝
	FocusOnEachOther        //互相关注
)

//RelationInfo 用户关系信息结构
type RelationInfo struct {
	Status   int       `json:"status"`
	UserInfo *UserInfo `json:"user_info"`		//所查看用户的个人信息
}

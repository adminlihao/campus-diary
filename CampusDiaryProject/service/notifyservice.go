package service

import (
	"CampusDiaryProject/models"
	"CampusDiaryProject/utils/redisutil"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"
)

//GetNotifyHistory 获取通知历史消息
func GetNotifyHistory(userid int, notifyType string, pageIndex int64, pageSize int64) ([]*models.NotifyHistory, error) {
	userId := strconv.Itoa(userid)
	notifyHistorys, err := redisutil.RedisZRevRangeWithScores(models.NOTIFY_HISTORY+notifyType+":"+userId, (pageIndex-1)*pageSize, pageIndex*pageSize-1)
	if err != nil {
		return nil, err
	}
	var notifyHistoryList []*models.NotifyHistory
	for _, v := range notifyHistorys {
		notify := make(map[string]interface{})
		json.Unmarshal([]byte(v.Member.(string)), &notify)
		notifyHistory := &models.NotifyHistory{
			SendTime: time.Unix(0, int64(v.Score)).Format("2006-01-02 15:04:05"),
			Notify:   notify,
		}
		notifyHistoryList = append(notifyHistoryList, notifyHistory)
	}
	return notifyHistoryList, nil
}

//ResetUnreadNotify    重置未读通知消息数量
func ResetUnreadNotify(userid int, notifyType string) error {
	userId := strconv.Itoa(userid)
	err := redisutil.RedisHSet(models.UNREAD_NOTIFY+notifyType, userId, 0)
	if err != nil {
		return err
	}
	return nil
}

//GetUnreadNotify  获取未读通知消息
func GetUnreadNotify(userid int) (map[string]int, error) {
	userId := strconv.Itoa(userid)
	var unreadNotify = make(map[string]int)
	follow, err := redisutil.RedisHGet(models.UNREAD_NOTIFY+"follow", userId)
	if err != nil {
		return nil, err
	}
	unreadFollow, _ := strconv.Atoi(follow)
	unreadNotify["follow"] = unreadFollow
	giveLike, err := redisutil.RedisHGet(models.UNREAD_NOTIFY+"giveLike", userId)
	if err != nil {
		return nil, err
	}
	unreadGiveLike, _ := strconv.Atoi(giveLike)
	unreadNotify["giveLike"] = unreadGiveLike
	comment, err := redisutil.RedisHGet(models.UNREAD_NOTIFY+"comment", userId)
	if err != nil {
		return nil, err
	}
	unreadComment, _ := strconv.Atoi(comment)
	unreadNotify["comment"] = unreadComment
	return unreadNotify, nil
}

//GetUnreadChat 获取未读聊天信息通知
func GetUnreadChat(userid int) ([]*models.UnreadChat, error) {
	userId := strconv.Itoa(userid)
	results, err := redisutil.RedisHGetAll(models.UNREAD_CHAT + userId)
	if err != nil {
		return nil, err
	}
	var unreadChatList []*models.UnreadChat
	for k, v := range results {
		res := strings.Split(v, "+")
		unreadChat := &models.UnreadChat{
			SendTime:   res[0],
			UnreadNum:  res[1],
			ChatUserID: k,
		}
		unreadChatList = append(unreadChatList, unreadChat)
	}
	sort.Slice(unreadChatList, func(i, j int) bool {
		return unreadChatList[i].SendTime > unreadChatList[j].SendTime
	})
	return unreadChatList, nil
}

//ResetUnreadChat 重置未读聊天信息数目
func ResetUnreadChat(userid string, chatUserid string, sendTime string) error {
	res, _ := redisutil.RedisHGet(models.UNREAD_CHAT+userid, chatUserid)
	if sendTime == "" {
		sendTime = strings.Split(res, "+")[0]
	}
	err := redisutil.RedisHSet(models.UNREAD_CHAT+userid, chatUserid, sendTime+"+0")
	if err != nil {
		return err
	}
	return nil
}

//IncrUnreadChat 加1未读聊天信息数目
func IncrUnreadChat(userid string, chatUserid string, sendTime string, unreadNum int) error {
	err := redisutil.RedisHSet(models.UNREAD_CHAT+userid, chatUserid, sendTime+"+"+strconv.Itoa(unreadNum))
	if err != nil {
		return err
	}
	return nil
}

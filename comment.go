// Copyright (c) 2018, 0xrgb. All right reserved.
// Use of this source code is governed by a license that can be found in the LICENSE file.

package rcdc

import (
	"encoding/json"
	"time"
)

// CommentKind 타입은 댓글의 종류를 나타낸다.
type CommentKind int

// CommentKind에 들어가는 상수
const (
	CommentNormal  CommentKind = iota // 댓글
	CommentDCCon                      // 디시콘
	CommentVoice                      // 보이스리플
	CommentSpecial                    // 댓글돌이
)

// CommentDate 상수는 디시인시아드에서 사용되는 댓글 시간 포맷을 나타낸다.
const CommentDate = "06.01.02 15:04:05"

// CommentTZ 상수는 디시인사이드에서 사용되는 댓글 시간의 UTC를 나타낸다.
var CommentTZ = time.FixedZone("Asia/Seoul", +9*60*60) // UTC +9

// Comment 구조체는 디시인사이드 댓글 하나를 나타낸다.
type Comment struct {
	ID   string
	Nick NickName
	IP   string
	Memo string
	Date time.Time

	// Experimental support
	VoteUp   int
	VoteDown int

	// TODO: 디시콘, 보이스리플, 댓글돌이 감지
	Kind CommentKind
}

// commentResponse 구조체는 디시인사이드에서 받은 JSON을 컨버팅하기 위해 내부적으로 사용한다.
type commentResponse struct {
	TotalCnt   int           `json:"total_cnt,string"`
	CommentCnt int           `json:"comment_cnt"`
	Comments   []commentJSON `json:"comments"`
}

type commentJSON struct {
	ID       string `json:"user_id"`
	Nickname string `json:"name"`
	IP       string `json:"ip"`
	Date     string `json:"reg_date"`
	VoteUp   int    `json:"recommend_cnt"`
	VoteDown int    `json:"nonrecommend_cnt"`
	Memo     string `json:"memo"`
}

// parseCommentResponse 함수는 디시인사이드에서 반환된 JSON을 받아 댓글 목록을 만든다.
func parseCommentResponse(data []byte) ([]Comment, error) {
	var resp commentResponse

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	ret := make([]Comment, len(resp.Comments))
	for idx, val := range resp.Comments {
		// TODO: 반고닉일때 NickHalfFix로 주기 (현재는 NickFix)
		var nick NickName
		if val.ID != "" {
			nick = NickName{
				Name: val.Nickname,
				Kind: NickFix,
			}
		} else {
			nick = NickName{
				Name: val.Nickname,
				Kind: NickNoLogin,
			}
		}

		tx, err := time.ParseInLocation(CommentDate, val.Date, CommentTZ)
		if err != nil {
			return nil, err
		}

		ret[idx] = Comment{
			ID:       val.ID,
			Nick:     nick,
			IP:       val.IP,
			Memo:     val.Memo,
			Date:     tx,
			VoteUp:   val.VoteUp,
			VoteDown: val.VoteDown,
		}
	}

	return ret, nil
}

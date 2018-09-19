// Copyright (c) 2018, 0xrgb. All right reserved.
// Use of this source code is governed by a license that can be found in the LICENSE file.

package rcdc

import "testing"

func TestParseCommentResponse01(t *testing.T) {
	// From http://gall.dcinside.com/ps/1132
	var PAYLOAD = []byte(`{"total_cnt":"4","comment_cnt":4,"comments":[{"no":"3996","headnum":"-2730","parent":"1132","ismember":"3556042","islevel":"9","user_id":"lchbest10","name":"\uc5d0\ub974\uc528","u_key":"","ip":"","reg_date":"18.09.16 20:08:28","nicktype":"20","t_ch1":"","t_ch2":"0","recommend_cnt":0,"nonrecommend_cnt":0,"vr_type":"","voice":null,"rcnt":"0","c_no":0,"depth":0,"del_yn":"N","is_delete":"0","memo":"\uc5ec\uae30\uc11c \uc724\uc131\uc6b0  \ucd94\ucc9c\ud558\uba74 \ud63c\ub0a0\uae4c\uc694? ","my_cmt":"N","del_btn":"Y","mod_btn":"N","gallog_icon":"<span class='nickname in' title='\uc5d0\ub974\uc528' >\uc5d0\ub974\uc528<\/span><a class='writer_nikcon'><img src='http:\/\/nstatic.dcinside.com\/dc\/w\/images\/fix_nik.gif' border=0 title='lchbest** : \uac24\ub85c\uadf8\ub85c \uc774\ub3d9\ud569\ub2c8\ub2e4.'  width='12'  height='11'  style='margin-left:2px;cursor:pointer;' onClick=\"window.open('\/\/gallog.dcinside.com\/lchbest10');\" alt='\uac24\ub85c\uadf8\ub85c \uc774\ub3d9\ud569\ub2c8\ub2e4.'><\/a>","vr_player":false,"vr_player_tag":"","next_type":0},{"no":"3997","headnum":"-2731","parent":"1132","ismember":"0","islevel":"10","user_id":"","name":"\u3141","u_key":"","ip":"220.127.*.*","reg_date":"18.09.16 21:31:49","nicktype":"00","t_ch1":"","t_ch2":"0","recommend_cnt":0,"nonrecommend_cnt":0,"vr_type":"","voice":null,"rcnt":"0","c_no":0,"depth":0,"del_yn":"N","is_delete":"0","memo":"k&r","my_cmt":"N","del_btn":"Y","mod_btn":"N","gallog_icon":"<span class=\"nickname\">\u3141<span class=\"ip\">(220.127.*.*)<\/span><\/span>","vr_player":false,"vr_player_tag":"","next_type":0},{"no":"3998","headnum":"-2732","parent":"1132","ismember":"0","islevel":"10","user_id":"","name":"\u3147\u3147","u_key":"","ip":"223.38.*.*","reg_date":"18.09.16 21:49:44","nicktype":"00","t_ch1":"0","t_ch2":"0","recommend_cnt":0,"nonrecommend_cnt":0,"vr_type":"","voice":null,"rcnt":"0","c_no":0,"depth":0,"del_yn":"N","is_delete":"0","memo":"\uc720\ud29c\ube0c\uc5d0 \ub530\ubc30\uc528 \u3131\u3131","my_cmt":"N","del_btn":"Y","mod_btn":"N","gallog_icon":"<span class=\"nickname\">\u3147\u3147<span class=\"ip\">(223.38.*.*)<\/span><\/span>","vr_player":false,"vr_player_tag":"","next_type":0},{"no":"3999","headnum":"-2733","parent":"1132","ismember":"0","islevel":"10","user_id":"","name":"\u3147\u3147","u_key":"","ip":"223.38.*.*","reg_date":"18.09.16 21:50:10","nicktype":"00","t_ch1":"0","t_ch2":"0","recommend_cnt":0,"nonrecommend_cnt":0,"vr_type":"","voice":null,"rcnt":"0","c_no":0,"depth":0,"del_yn":"N","is_delete":"0","memo":"\ucc45\uc73c\ub85c \ud558\uace0\uc2f6\uc73c\uba74 c++\uae30\ucd08\ud50c\ub7ec\uc2a4","my_cmt":"N","del_btn":"Y","mod_btn":"N","gallog_icon":"<span class=\"nickname\">\u3147\u3147<span class=\"ip\">(223.38.*.*)<\/span><\/span>","vr_player":false,"vr_player_tag":"","next_type":0}],"pagination":"<em>1<\/em>"}`)
	var answer = []struct {
		name string
		kind NickKind
		memo string
	}{
		{"에르씨", NickFix, "여기서 윤성우  추천하면 혼날까요? "},
		{"ㅁ", NickNoLogin, "k&r"},
		{"ㅇㅇ", NickNoLogin, "유튜브에 따배씨 ㄱㄱ"},
		{"ㅇㅇ", NickNoLogin, "책으로 하고싶으면 c++기초플러스"},
	}

	val, err := parseCommentResponse(PAYLOAD)
	if err != nil {
		t.Error("cannot parse JSON:", err)
		return
	}

	if len(val) != 4 {
		t.Error("number of comment is wrong")
		return
	}

	for i := range val {
		if answer[i].name != val[i].Nick.Name {
			t.Errorf("expected name %q, got %q", answer[i].name, val[i].Nick.Name)
		}
		if answer[i].kind != val[i].Nick.Kind {
			t.Errorf("expected kind %v, got %v", answer[i].kind, val[i].Nick.Kind)
		}
		if answer[i].memo != val[i].Memo {
			t.Errorf("expected memo %q, got %q", answer[i].memo, val[i].Memo)
		}
	}
}

// Copyright (c) 2018, 0xrgb. All right reserved.
// Use of this source code is governed by a license that can be found in the LICENSE file.

package rcdc

// NickKind 타입은 디시인사이드 닉네임의 종류를 표현할 때 사용한다.
// 현재는 유동, 반고정닉, 고정닉이 있다.
type NickKind int

// Nickkind에 들어가는 상수
const (
	NickInvalid NickKind = iota // 초기화가 되지 않은 값
	NickNoLogin                 // 유동
	NickFix                     // 고정닉
	NickHalfFix                 // 반고정닉
)

// NickName 구조체는 디시인사이드 닉네임을 표현할 때 사용한다.
type NickName struct {
	Name string
	Kind NickKind
}

// 고정닉에는 "●", 반고정닉에는 "◐", 유동에는 "○"를 뒤에 붙여 출력한다.
// 값이 잘못되어 있을 경우 "(nil)"을 대신 출력한다.
func (nick NickName) String() string {
	switch nick.Kind {
	case NickNoLogin:
		return nick.Name + "○"
	case NickHalfFix:
		return nick.Name + "◐"
	case NickFix:
		return nick.Name + "●"
	}

	return "(nil)"
}

// IsZero 메소드는 Nickname이 초기화 되어있지 않다면 true를 반환한다.
func (nick NickName) IsZero() bool {
	return nick.Kind == NickInvalid
}

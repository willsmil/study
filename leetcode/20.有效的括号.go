package leetcode

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串 👍 2607 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := NewStack1()
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		} else if c == ')' || c == '}' || c == ']' {
			if !matched(stack.Pop(), c) {
				return false
			}
		} else {
			return false
		}
	}
	if stack.Size() > 0 {
		return false
	}
	return true
}

func matched(a, b int32) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	return false
}

type Stack1 struct {
	data []int32
}

func NewStack1() *Stack1 {
	return &Stack1{data: []int32{}}
}

func (s *Stack1) Peek() int32 {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack1) Push(i int32) {
	s.data = append(s.data, i)
}

func (s *Stack1) Pop() int32 {
	if len(s.data) == 0 {
		return 0
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}

func (s *Stack1) Size() int {
	return len(s.data)
}

func (s *Stack1) IsEmpty() bool {
	return len(s.data) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

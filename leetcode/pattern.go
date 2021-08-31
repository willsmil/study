package leetcode

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

type GrayRuleService struct {
}

func (s *GrayRuleService) validRule(ctx context.Context, visitInfo int, ruleId string) bool {
	// fmt.Println("valid:", ruleId)
	return cast.ToInt(ruleId)%2 == 1
}

func (s *GrayRuleService) parseRule2(ctx context.Context, visitInfo int, rule string) bool {
	// fmt.Println("parse rule:", rule)
	if rule == "" {
		return false
	}
	var res bool

	for len(rule) > 0 {
		fmt.Println("parse rule:", rule)
		switch rule[0] {
		case '!':
			fmt.Println("!", rule[1:])
			//return !s.parseRule2(ctx, visitInfo, rule[1:])
			idx, pattern := findFirstPattern(rule[1:])
			res = !s.parseRule2(ctx, visitInfo, pattern)
			rule = rule[idx+1:]
			// fmt.Println(idx, pattern, res, rule)
		case '&':
			fmt.Println(res, "&&", rule[1:])
			return res && s.parseRule2(ctx, visitInfo, rule[1:])
		case '|':
			fmt.Println(res, "||", rule[1:])
			return res || s.parseRule2(ctx, visitInfo, rule[1:])
		case '(':
			idx, pattern := findFirstParenthesisPattern(rule)
			res = s.parseRule2(ctx, visitInfo, pattern)
			rule = rule[idx:]
		case ')':
			rule = rule[1:]
		default:
			idx, pattern := findFirstIdPattern(rule)
			res = s.validRule(ctx, visitInfo, pattern)
			rule = rule[idx:]
			// fmt.Println("left:", rule)
		}
		// fmt.Println("pre res:", res, "left rule:", rule)
	}
	return res
}

func (s *GrayRuleService) parseRule(ctx context.Context, visitInfo int, rule string) bool {
	if len(rule) == 0 {
		return false
	}
	// rule = strings.Replace(rule, "!!", "", -1)
	res := true
	for i := 0; i < len(rule); {
		// fmt.Println("start parse: i =", i, "left rule:", rule[i:])
		switch rule[0] {
		case '!':
			return !s.parseRule(ctx, visitInfo, rule[1:])
			//var preOp byte = '&'
			//if i > 0 {
			//	if rule[i-1] == '&' || rule[i-1] == '|' {
			//		preOp = rule[i-1]
			//	}
			//}
			//
			//i++
			////ruleId := parseRuleId(rule, &i)
			////if ruleId == "" {
			////	break
			////}
			//if preOp == '&' {
			//	res = res && !s.parseRule(ctx, visitInfo, rule[i:])
			//} else {
			//	res = res || !s.parseRule(ctx, visitInfo, rule[i:])
			//}
			//fmt.Println(string(preOp), "!", ruleId, "-->", res)
		case '&':
			// i++
			res = res && s.parseRule(ctx, visitInfo, rule[1:])
			//ruleId := parseRuleId(rule, &i)
			//if ruleId == "" {
			//	break
			//}
			//res = res && s.validRule(ctx, visitInfo, ruleId)
			//fmt.Println("&", ruleId, "-->", res)
		case '|':
			// i++
			res = res || s.parseRule(ctx, visitInfo, rule[i:])
			//i++
			//ruleId := parseRuleId(rule, &i)
			//if ruleId == "" {
			//	break
			//}
			//res = res || s.validRule(ctx, visitInfo, ruleId)
			//fmt.Println("|", ruleId, "-->", res)
		case '(':
			//var preOp byte = '&'
			//preNot := false
			//if i > 0 {
			//	if rule[i-1] == '&' || rule[i-1] == '|' {
			//		preOp = rule[i-1]
			//	} else if rule[i-1] == '!' {
			//		preNot = true
			//		if i-1 == 0 {
			//			preOp = '!'
			//		} else {
			//			if rule[i-2] == '&' || rule[i-2] == '|' {
			//				preOp = rule[i-2]
			//			}
			//		}
			//	}
			//}
			//i++
			// var i = 0
			subRule := parseParentheses(rule, &i)
			fmt.Println("-------sub rule start:", subRule)
			subRes := s.parseRule(ctx, visitInfo, subRule)
			fmt.Println("-------sub rule end:", subRule, "res-->", subRes)
			//
			//if preNot {
			//	subRes = !subRes
			//}
			//if preOp == '&' {
			//	res = res && subRes
			//} else {
			//	res = res || subRes
			//}
			//fmt.Println("sub rule: (", subRule, ")", "preOp:", string(preOp), preNot, "-->", res)
		// case ')':
		// return
		//i++
		//fmt.Println("skip ')'")
		default:
			// var i = 0
			ruleId := parseRuleId(rule, &i)
			res = s.validRule(ctx, visitInfo, ruleId)
			// fmt.Println(ruleId, "-->", res)
		}
	}
	return res
}
func parseRuleId(rule string, i *int) string {
	j := *i
	if *i == len(rule)-1 {
		*i++
		return rule[len(rule)-1:]
	}
	if rule[*i] == '(' || rule[*i] == '!' {
		return ""
	}
	for ; *i < len(rule); *i++ {
		if rule[*i] == '!' || rule[*i] == '&' || rule[*i] == '|' || rule[*i] == '(' || rule[*i] == ')' {
			break
		}
	}
	return rule[j:*i]
}
func findFirstPattern(rule string) (int, string) {
	if rule[0] == '(' {
		return findFirstParenthesisPattern(rule)
	}
	return findFirstIdPattern(rule)
}
func findFirstIdPattern(rule string) (int, string) {
	var i = 0
	for ; i < len(rule); i++ {
		if rule[i] == '!' || rule[i] == '&' || rule[i] == '|' || rule[i] == '(' || rule[i] == ')' {
			break
		}
	}
	return i, rule[:i]
}
func findFirstParenthesisPattern(rule string) (int, string) {
	num := 1
	var i = 1
	for ; i < len(rule); i++ {
		if num == 0 {
			break
		}
		if rule[i] == '(' {
			num++
		} else if rule[i] == ')' {
			num--
		}
	}
	return i, rule[1:i]
}
func parseParentheses(rule string, i *int) string {
	j := *i
	num := 1
	for ; *i < len(rule); *i++ {
		if num == 0 {
			break
		}
		if rule[*i] == ')' {
			num--
		} else if rule[*i] == '(' {
			num++
		}
	}
	return rule[j : *i-1]
}
func ruleLegal(rule string) bool {
	match, err := regexp.MatchString("[^0-9&|!()]", rule)
	if err != nil || match {
		return false
	}
	if strings.Count(rule, "(") != strings.Count(rule, ")") {
		return false
	}
	rule = strings.ReplaceAll(rule, "&", "?")
	rule = strings.ReplaceAll(rule, "|", "?")
	if strings.Contains(rule, "??") || strings.Contains(rule, "!!") || strings.Contains(rule, "!?") {
		return false
	}
	return true
}

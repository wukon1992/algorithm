package main

import(
	"math"
)

//有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

/**
 * 首先确认都要闭合则字符串长度必须为偶数，因此判断是否为偶数，不是则false
 * 遇到右括号则需要判断是否闭合。因此先将右括号为key，左括号为value放在map中。
 * 定义一个栈存放遍历的符号，当符号为左括号放入栈中，当右括号则从栈顶取出一个左括号。用右括号的map值来匹配。匹配失败则不满足
 * 字符串遍历
 */
func isValid(s string) bool {
    l := len(s)
    if l % 2 != 0 {
        return false
    }
    stack := make([]byte,0)
    hs := map[byte]byte{ // 右括号则需要闭合
        '}':'{',
        ')':'(',
        ']':'[',
    }

    for i := 0; i < l; i++ {
        if hs[s[i]] > 0 {
            if len(stack) > 0 && stack[len(stack)-1] ==  hs[s[i]]{
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }else{
            stack = append(stack,s[i])
        }
    }
    return len(stack) == 0
}

// //最小栈 力扣155
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
// 定一个辅助栈来存储，栈中的对应元素的最小值,
/**
 * 定义一个最小栈
 */
type MinStack struct{
	stack: []int
	minStack: []int
}

/**
 * 初始化最小栈
 * @return {[type]} [description]
 */
func initMinStack(){
	return &MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64}	
	}
}
func min(x,y int) {
	if x > y {
		return y
	}
	return x
}
func(this *MinStack)Push(x int){
	this.stack = append(this.stack,x)
	this.minStack = append(this.minStack,min(x, this.minStack(len(this.minStack)-1)))
}
func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}
func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}

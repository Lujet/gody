package main

import "fmt"

/*
	基本数据(基础类型)

	计算机底层都是位, 而实际操作则是基于大小固定的单元中的数值, 称为字(word);
	go 的数据类型宽泛, 并有多种组织方式, 向下匹配硬件特性, 向上满足程序员所需,
	从而可以方便的表示复杂数据结构;

	数据类型分为:
		- 基础类型(basic type)
		- 聚合类型(aggregate type)
		- 引用类型(reference type)
		- 接口类型(interface type)

	引用类型全都间接指向程序变量或状态, 于是操作所引用数据的效果就会遍及该数据
	的全部引用.

	int 和 uint 在特定的平台上, 其大小与原生的有符号整数\无符号整数相同, 或等于
	该平台上运算效率最高的值, int 是使用最广泛的数值类型; 这两种类型大小相等,
	都是32位或64位, 但不能认为它们一定是32位或一定就是64位, 即使在同样的硬件平台
	上, 不同的编译器可能选用不同的大小;

	rune 是int32类型的同义词, 用于指明一个值是 Unicode 码点(code point); 这两个
	名称可以互换; 同样, byte 类型是 uint8 类型的同义词, 强调一个值是原始数据而非
	量值;(TODO: 量值怎么理解?)

	uintptr 无符号整数, 大小不明确, 但足以完整存放指针; uintptr 类型仅仅用于底层
	编程, 例如在 go 程序和 c 程序库或操作系统的接口界面;

	int, uint, uintptr 都有别于其大小明确的相似类型的类型; int 和 int32 是不同的
	类型, 尽管 int 天然的大小就是32位, 并且 int 值若要当做 int32 使用, 必须显式
	转换, 反之亦然;

	有符号整数以补码表示(TODO:补码), 保留最高位作为符号位; 无符号整数由全部位构
	成其非负值;

	算术运算符+, -, *, / 可应用于整数, 浮点数和复数, 而取模运算符%仅能用于整数,
	取模运算符%的行为因编程语言而异; 对go而言, 取模余数的正负号总是与被除数一致,
	于是 -5%3 和 -5%-3 的都 -2; 除法运算(/)的行为取决于操作数是否都为整型, 整数
	相除, 商会舍弃小数部分, 于是5.0/4.0得到1.25, 而 5/4 结果是1.

	不论是有符号数还是无符号数, 若表示算法运算结果所需的位超出该类型的范围, 就
	称为溢出, 溢出的高位部分会无提示的丢弃;

	全部基本类型的值都可以比较, 两个相同类型的值可以用 == 和!= 运算符比较;
	整数, 浮点数和字符串还能根据比较运算符排序;

	go 具备下列位运算符:(TODO:位运算)
	对操作数的运算逐位独立进行, 不涉及算术进位或正负号:
	- &  位运算 AND
	- |  位运算 OR
	- ^  位运算 XOR
	- &^ 位清空 AND NOT

	移位运算:
	- << 左移
	- >> 右移

	运算符 ^ 作为二元运算符表示按位 "异或"(XOR); 若作为一元前缀运算符, 则它表示
	按位取反或按位取补, 运算结果就是操作数逐位取反;
	运算符 &^ 是按位清除(AND NOT), 对于 z=x&^y, 若 y 的某位为1, 则 z 的对应位
	为0; 否则, 它就等于x的对应位;(TODO: 应用场景)

	在移位运算 x<<n, x>>n 中, 操作数n决定位移量, 而且n必须为无符号型; 操作数x
	可以是有符号型也可是无符号型; 算术上, 左移运算 x<<n 等于x乘以2^n; 而右移
	运算 x>>n 等价于x除以2^n, 向下取整;
	左移以0填补右边空位, 无符号整数右移同样以0填补左边空位; 但有符号数的右移
	操作是按符号位的值填补空位; 因此, 如果将整数以位模式处理, 必须使用无符号
	整型;(TODO)

	对于每种基础类型T, 若允许转换, 操作T(x)会将x的值转换为类型T, 很多整型-整型
	转换不会引起值的变化, 仅告知编译器应如何解读该值, 不过, 缩减大小的整型转换,
	以及整型和浮点型的相互转换, 可能改变值或损失精度;

	浮点型转成整型, 会舍弃小数部分, 趋零截尾(正值向下取整, 负值向上取整), 如果
	有些转换的操作数的值超出了目标类型的取值范围, 就应当避免这种转换, 因为其
	行为依赖具体实现;


	不论有无大小和符号限制, 源码中的整数都能写成常见的十进制数; 也能写成八进制
	数, 以0开头(如: 0666); 还能写成十六进制数(以0x或0X开头); 八进制目前的应用
	场景很少(表示POSIX文件系统的权限), 而十六进制广泛用于强调其位模式, 而非数值
	大小;
*/

// 使用位运算将一个 uint8 值作为位集(bitset)处理, 其含有8个独立的位,
// 高效且紧凑; TODO: 集合相关概念
func HandleBit() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Println(x, y)
	// 谓词 08 表示在这个输出结果前补零, 补够8位
	fmt.Printf("%08b\n", x)    // 00100010 集合 {1, 5}
	fmt.Printf("%08b\n", y)    // 00000110 集合 {1, 2}
	fmt.Printf("%08b\n", x&y)  // 00000010 交集 {1}
	fmt.Printf("%08b\n", x|y)  // 00100110 并集 {1,2,5})
	fmt.Printf("%08b\n", x^y)  // 00100100 对称差 {2, 5} (TODO:)
	fmt.Printf("%08b\n", x&^y) // 00100000 差集 {5}

}

/*
	尽管go具备无符号整型数和相关算术运算, 也尽管某些量值不可能为负, 但是往往
	采用有符号整型数, 如数组的长度(即使直观上应该选uint);
	无符号整数往往只用于位运算和特定算术运算符, 如实现位集时, 解析二进制格式
	的文件, 或散列和加密; 一般情况下无符号整数极少用于表示非负值;
*/
func UseUint() {
	medals := []string{"gold", "silver", "bronze"}
	length := uint(len(medals) - 1)
	for i := length; i >= 0; i-- {
		fmt.Println(medals[i])
		// panic: runtime error: index out of range [18446744073709551615] with length 3
		// i 为 uint, 条件 i >= 0 恒成立; 第3轮迭代后, 有i==0,
		// 语句 i-- 使得i变为uint型的最大值(如: 2^64-1), 而非 -1, 导致切片元素
		// 访问越界, 引发宕机
	}
}

/*
	如果使用 fmt 包输出数字, 可以用谓词 %d, %o, %x 指定进位制基数
	两个 fmt 技巧:
		- [1] 告知 Printf 重复使用第一个操作数
		- # 告知 Printf 输出相应的前缀0, 0x 或 0X
*/
func Print() {
	o := 0666                                 // 8进制
	fmt.Printf("%d %[1]o %#[1]o %#[1]x\n", o) // 438 666 0666 0x1b6
}

func main() {
	// HandleBit()
	// UseUint()
	// Print()
	Literal()
}

/*
	中文字符(rune literal) 的形式是字符写在一对单引号内
	用 %c 输出文字符号, 如果想让输出带单引号则用 %q
*/
func Literal() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
	fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'
}

package main

import "fmt"

/*
	字符串是不可变的字节序列(内存分配在只读data段), 可以包含任何数据,
	包括0值字节, 主要是人可读的文本, 习惯上, 文本字符串被解读成 utf-8 编
	码的 unicode 码点(文字符号)序列;

	内置的 len 函数返回的字符串的字节数(而非文字符号的数目), 下标访问操作
	s[i] 则取得第 i 个字节, 第 i 个字节不一定就是第 i 个字符, 因为非 ASCII
	字符的 utf-8 码点需要两个字节或多个字节

	字符串可以通过比较运算符做比较, 如 == 和 <, 比较运算按字节进行, 结果服从
	本身的字典排序(TODO: 何为字典排序)

	不可变意味着两个字符串能安全的共用同一段底层内存, 使得复制任何长度字符串
	的开销都低廉;


	字符串包含字节数组, 创建后就无法改变, 相反, 字节slice的元素允许随意修改,
	所以 []byte 不能和 string 共用同一段底层内存;
	字符串可以和字节slice相互转换:
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	[]byte(s) 转换操作会分配新的字节数组, 拷贝填入s含有的字节, 并生成一个slice
	引用指向整个数组; 复制有必要确保 s 的字节维持不变, 反之, 用string(b) 将
	字节slice转换成字符串也会产生一份副本, 保证 b 也不可变;

	/* 字符串和字节slice相互转换都要分配新的字节数组, 产生内存消耗

	字符串和数字的相互转换(常用)
	整数转换为字符串:
		- fmt.Sprinf  谓词: %b, %d, %o(8进制), %x
		- strconv.Itoa  ('Integer to ASCII')
	字符串转换为整数:
		- strconv.ParseInt
		- strconv.Atoi  ('ASCII to Integer')
*/

// 尽管可以将新值赋予字符串变量, 但是字符串值无法改变: 字符串值本身所包含的
// 字节序列用不可变(TODO:理解)
func Assign() {
	s := "left foot"
	t := s
	s += ", right foot"
	// 并不改变 s 原有的字符串值, 只是将 += 语句生成的新字符串赋予 s, 同时
	// t 仍然持有旧的字符串值
	fmt.Println(s)
	fmt.Println(t)
	// s[0] = 't' // 因为字符串不可改变, 所以字符串内部的数据不运行修改
}

/*
	Unicode

	早期计算机只处理一个字符集 ASCII 码, ASCII 码使用7位表示128个"字符": 大小
	写英文字母, 数字, 各种标点和设备控制符; 但是要包含世界上各种语言的所有
	字符是不够用的; 所以使用 Unicode 表示, 所有语言的文书体系的全部字符,
	还有重音字符以及许多其他的字符都可以各自赋予一个叫 Unicode 码点的标准数字,
	在 go 中, 这些字符记号称为文字符号(rune);

	适合保存单个文字符号的数据类型就是 int32, 所以 rune 类型作为 int32 类型的
	别名;

	可以将文字符号的序列表示成  int32 值序列, 这种表示方式称为 utf-32 或 ucs-4,
	每个 unicode 码点的编码长度相同, 都是32位, 这种编码简单化一, 但是因为大多数
	面向计算机的可读文本是 ASCII 码, 每个字符只需8位(1字节), 导致了不必要的存储
	空间消耗, 而使用广泛的字符数也少于 65536 个, 用16位就可容纳, 使用 utf-8
	可进行兼容;


	UTF-8

	UTF-8 以字节为单位对 Unicode 码点作变长编码, UTF-8 是现行的一种 Unicode
	标准; 每个字符用1-4个字节表示, ASCII 字符的编码仅占1个字节, 而其他常用
	的文书字符的编码只是2或3个字节;
	一个字符符号编码的首字母的高位指明了后面还有多少字节; 若最高位是0, 则标示
	着它是7位的 ASCII 码, 其文字符号的编码仅占1字节, 这样就兼容传统的 ASCII 码,
	若最高几位是 110, 则文字符号的编码占用2个字节, 第二个字节以 10 开始, 更长
	的编码依次类推:
	0xxxxxxx								 文字符号				0-127 (ASCII)
	110xxxxx 10xxxxxx                        128-2047               少于128个未使用的值
	1110xxxx 10xxxxxx 10xxxxxx               2048-65535             少于2048个未使用的值
	11110xxx 10xxxxxx 10xxxxxx 10xxxxxx      65536 - 0x10ffff       其他

	变长编码的字符串无法按下标直接访问第 n 个字符, 但也获得了其他的优势:
	- UTF-8 编码紧凑, 兼容 ASCII, 并且自同步: 最多追溯3字节就能定位一个字符的
		起始位置;
	- UTF-8 还是前缀编码, 因此能从左到右解码而不产生歧义, 也无须超前
		预读(TODO:), 于是查找文字符号仅需搜索它自身的字节, 不必考虑前文内容;
	- 文字符号的字典字节顺序与 Unicode 码点顺序一致(Unicode设计), 因此按 UTF-8
		编码排序自然就是对文字符号排序; UTF-8 本身不会嵌入 NUL 字节(0值), 这
		便于某些程序语言用 NUL 标记字符串结尾;

	unicode 包有判别文字符号值特性的函数, 如 IsDigit, IsLetter, IsUpper 和
	IsLower, 每个函数以单个文字符号值作为参数, 并返回布尔值;
*/

/*
	go 的源文件总是以 UTF-8 编码, 同时, go 程序操作的文本字符串也优先采用
  	UTF-8 编码,  unicode 包针对单个文字符号的函数(例如区分字母和数字,
  	转换大小写), 而 unicode/utf8 包则提供了按 utf8 编码和解码文字符号的函数
  	go 语言中, 字符串字面量的转义使得可以用码点的值来指明 Unicode 字符,
  	\uhhhh 表示16位码点值, \Uhhhhhhhh 表示32位码点值, 其中h代表一个16进制数字,
  	32位形式的码点值几乎不会用到
	下面几个字符串字面量都表示长度为6字节的系统串
	"世界"
	"\xe4\xb8\x96\xe7\x95\x8c"
	"\u4e16\u754c"
	"\U00004e16\U0000754c"

	如果字符串含有任意二进制数, 对字符串的读取会产生一个专门的 Unicode 字符
	'\uFFFD' 替换它, 其输出通常是个黑色六角形或类似钻石的形状, 里面有个白色
	问号.
*/

func UseUnicode() {
	s := "世界"
	fmt.Printf("% x\n", s) // e4 b8 96 e7 95 8c
	// 谓词 %x(% 和 x 之间有空格) 表示以十六进制输出, 并在每两个位数间插入空格
}
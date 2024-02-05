一个自动化生成结构体默认函数的工具.
主要逻辑来自 Kubernetes 的代码生成工具 [gengo](https://github.com/kubernetes/gengo/blob/master/examples/defaulter-gen/generators/defaulter.go).

## 使用
对于给定的目录, godefault 会自动扫描所有顶层结构体并生成默认函数.

生成行为完全受评论中的 tag 控制.

任何需要生成默认函数的包都应该在包注释中包括 `// +k8s:defaulter-gen=`.
这个 tag 有两种用法.
如果其值是 true|false,
则代表应该/不应该为紧随其后的结构体生成默认函数.
```go
// +k8s:defaulter-gen=true

type Defaulted struct {
```
注意此时改行注释后必须紧跟空行.

如果 tag 的值是其他内容, 则代表某个字段名, 工具会为所有包含这个字段的结构体生成默认函数.
```go
// +k8s:defaulter-gen=TypeMeta
```

字段的默认值由注释 `// +default=` 控制, 支持通过 JSON 字符串为复杂结构体赋值.
```go
type Defaulted struct {
	// +default="bar"
	StringDefault string

	// Default is forced to empty string
	// Specifying the default is a no-op
	// +default=""
	StringEmptyDefault string

	// Not specifying a default still defaults for non-omitempty
	StringEmpty string

	// +default="default"
	StringPointer *string

	// +default=64
	Int64 *int64

	// +default=32
	Int32 *int32

	// +default=1
	IntDefault int

	// +default=0
	IntEmptyDefault int

	// Default is forced to 0
	IntEmpty int

	// +default=0.5
	FloatDefault float64

	// +default=0.0
	FloatEmptyDefault float64

	FloatEmpty float64

	// +default=["foo", "bar"]
	List []Item
	// +default={"s": "foo", "i": 5}
	Sub *SubStruct

	//+default=[{"s": "foo1", "i": 1}, {"s": "foo2"}]
	StructList []SubStruct

	//+default=[{"s": "foo1", "i": 1}, {"s": "foo2"}]
	PtrStructList []*SubStruct

	//+default=["foo"]
	StringList []string

	// Default is forced to empty struct
	OtherSub SubStruct

	// +default={"foo": "bar"}
	Map map[string]Item

	// +default={"foo": {"S": "string", "I": 1}}
	StructMap map[string]SubStruct

	// +default={"foo": {"S": "string", "I": 1}}
	PtrStructMap map[string]*SubStruct

	// A default specified here overrides the default for the Item type
	// +default="banana"
	AliasPtr Item
}
```
更多用例请参考 [examples](./examples/_output_tests), 其中
`type.go` 为结构体定义,`zz_generated.go` 为生成的默认函数.

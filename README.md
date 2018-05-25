### English Introduction:
Imagine if there has a program that can helps me automatically generate `Error Codes Doc` in the `API` interface.

#### For example:

some codes like below:

```golang

func HandleCreateUser(w http.ResponseWriter,r *http.Request) map[string]interface{} {
	if r.Body == nil {
		return util.GetCommonErr(23,"create user req body is fucking null?")
	}
	return util.GetCommonSuccess("success")
}

```

Assuming that you are a API developer at this time, in order to make the front end application call, you need to write the API document, and the part of the error message is included.

If your entire server program has hundreds or even thousands of rows of erroneous information output code, do you write hundreds of articles when you write documents? What a trouble it is!

##### And my open source project can solve this problem.

#### Step 1

Setting up your configuration file.

```json
{
  "TargetFileSuffix":[".go"],
  "TargetErrorFuncName":["util.GetCommonErr","util.GetErrWithTips"],
  "FilterFileName":["core"],
  "ParamsColumnNames":[" Code "," Meaning ","Tips"],
  "ParamsSplitChar":","
}
```

#### Step 2

Enter your code directory and run the program.

```golang
func TestDocPrinter(t *testing.T) {
	p := NewDefaultErrorDocPrinter(NewDefaultMarkDownErrorDocPrinter())
	if p == nil {
		return
	}
	fmt.Println(p.printErrorDoc("../../errorDocPrinter"))
}
```

#### Step 3

Done

| Code | Meaning |Tips|
| - | - | - |
|-9|invalid create user|--空缺--|
|-4|invalid create user|--空缺--|
|-1|create user failed|创建用户失败|
|88|创建评论失败|--空缺--|
|3110|error params|--空缺--|
|3111|update failed|--空缺--|
|3112|yellow 内容涉黄|--空缺--|
|3113|forbid 禁止访问|--空缺--|
|3114|empty id|--空缺--|
|3115|服务端开启事务失败|--空缺--|
|3116|服务端事务提交失败|--空缺--|
|3117|update effect row <= 0|--空缺--|
|3118|RowsAffected 失败|--空缺--|
|3119|更新只有部分成功|--空缺--|
|3120|empty userId|--空缺--|
|3121|too lager|--空缺--|
|3122|user not exits|--空缺--|
|3123|非法更新|--空缺--|
|3124|参数个数长度限制|--空缺--|
|3126|服务端事务提交失败|--空缺--|
|3127|invalid money|--空缺--|
|3128|money not enough|--空缺--|
|3129|创建消费记录失败|--空缺--|






































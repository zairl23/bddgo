BDD in Go Using godogs
==========

在这个工程例子里面，我们使用[godog](https://github.com/DATA-DOG/godog)库来实践一下go语言的BDD开发。


## 准备

首先我们需要安装godog

```go
    go get github.com/DATA-DOG/godog/cmd/godog
```

然后验证一下gogod是否可以使用，执行：

```go
    godog 
```

终端打印出：

    No scenarios
    No steps
    35.327µs

表示我们已经安装好了godog。

## 实践

我们参照官方库中的[example](https://github.com/DATA-DOG/godog/tree/master/examples)来学习实践一把吧！


我们创建一个工程bddgo, cd $GOPATH/src/bddgo

现在想象我们有一个狗粮带，里面装满是狗粮，首先我们来描述他的功能:

$GOPATH/src/bddgo/features/godogs.feature

```
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining
```

此时我们执行一下: godog, 会输出：

```
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12          # features/godogs.feature:6
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

1 scenarios (1 undefined)
3 steps (3 undefined)
150.851µs

You can implement step definitions for undefined steps with these snippets:

func thereAreGodogs(arg1 int) error {
    return godog.ErrPending
}

func iEat(arg1 int) error {
    return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
    s.Step(`^I eat (\d+)$`, iEat)
    s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
```

godog给了我们足够的信息，指导我们下一步该如何推进功能的开发。我们创建测试文件godog_test.go

我们现在仅仅需要一个存储狗粮数量的变量：$GOPATH/src/bddgo/godogs.go

```go
package main

// Godogs available to eat
var Godogs int

func main() { 
	/* usual main func */ 
}
```

测试文件： $GOPATH/src/bddgo/godog_test.go

```go
package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

func thereAreGodogs(available int) error {
	Godogs = available
	return nil
}

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
```

我们再次执行godog，会输出:

```
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12          # features/godogs.feature:6
    Given there are 12 godogs        # godogs_test.go:10 -> thereAreGodogs
    When I eat 5                     # godogs_test.go:14 -> iEat
    Then there should be 7 remaining # godogs_test.go:22 -> thereShouldBeRemaining

1 scenarios (1 passed)
3 steps (3 passed)
343.997µs
```





BDD in go using Ginkgo
=============================

在这里我们使用[ginkgo](http://onsi.github.io/ginkgo/)这个库来学习在go工程里面实践bdd开发。

## Install

```shell
    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/onsi/gomega
```

## Start

### bootstrap

create a project in $GOPATH/src/book

```
ginkgo bootstrap
```

will create a file book_suite_test.go

```go
package book_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Suite")
}

```

run ginkgo or go test:

will output:

    Running Suite: Book Suite
    =========================
    Random Seed: 1541744323
    Will run 0 of 0 specs


    Ran 0 of 0 Specs in 0.001 seconds
    SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped
    PASS

    Ginkgo ran 1 suite in 5.675110497s
    Test Suite Passed

### Adding specs

run
```shell
ginkgo generate book
```

created a new file book_test.go

```go
package book_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "bddgo/ginkgo/book"
)

var _ = Describe("Book", func() {

})

```

The function in the Describe will contain our specs. Let’s add a few now to test loading books from JSON:

```go
var _ = Describe("Book", func() {
    var (
        longBook  Book
        shortBook Book
    )

    BeforeEach(func() {
        longBook = Book{
            Title:  "Les Miserables",
            Author: "Victor Hugo",
            Pages:  1488,
        }

        shortBook = Book{
            Title:  "Fox In Socks",
            Author: "Dr. Seuss",
            Pages:  24,
        }
    })

    Describe("Categorizing book length", func() {
        Context("With more than 300 pages", func() {
            It("should be a novel", func() {
                Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
            })
        })

        Context("With fewer than 300 pages", func() {
            It("should be a short story", func() {
                Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
            })
        })
    })
})
```

Then run the test:

```shell
ginkgo
```
the test failed, let's write the CategoryByLength function :

```go
package book


type Book struct {
	Title string
	Author string 
	Pages int
}

func (b Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

```

Again, type ginkgo, will output below:

    Running Suite: Book Suite
    =========================
    Random Seed: 1541745249
    Will run 2 of 2 specs

    ••
    Ran 2 of 2 Specs in 0.000 seconds
    SUCCESS! -- 2 Passed | 0 Failed | 0 Pending | 0 Skipped
    PASS

    Ginkgo ran 1 suite in 1.904428835s
    Test Suite Passed


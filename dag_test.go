package trp

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("DAGBuilder", func() {

  var builder *DAGBuilder

  BeforeEach(func() {
    builder = NewDAGBuilder()
  })

  It("reports new cycles", func() {
    Expect(builder.AddEdge("A", "B")).To(Succeed())
    Expect(builder.AddEdge("B", "C")).To(Succeed())
    Expect(builder.AddEdge("C", "A")).NotTo(Succeed())

    Expect(builder.TSort()).To(Equal([]string{"A", "B", "C"}))
  })

})
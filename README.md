## Ascii Folding Library [![Build Status](https://snap-ci.com/hkulekci/ascii-folding/branch/master/build_image)](https://snap-ci.com/hkulekci/ascii-folding/branch/master) [![codecov](https://codecov.io/gh/hkulekci/ascii-folding/branch/master/graph/badge.svg)](https://codecov.io/gh/hkulekci/ascii-folding)

A Golang port of the Apache Lucene ASCII Folding Filter. This library converts 
alphabetic, numeric, and symbolic Unicode characters into their ASCII 
equivalents, if one exists.

Example Usage in a program:

```go
package main

import "github.com/hkulekci/ascii-folding"

func main() {
    print(asciiFolding.Fold("L' heure n' est plus à vanter cette franche camaraderie qui avait marqué les relations entre"))
}
```

Source : [http://svn.apache.org/repos/asf/lucene/java/tags/lucene_solr_4_5_1/lucene/analysis/common/src/java/org/apache/lucene/analysis/miscellaneous/ASCIIFoldingFilter.java](http://svn.apache.org/repos/asf/lucene/java/tags/lucene_solr_4_5_1/lucene/analysis/common/src/java/org/apache/lucene/analysis/miscellaneous/ASCIIFoldingFilter.java)

# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=termcolor
GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
	termcolor.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w termcolor.go
	${GOFMT} -w test.go


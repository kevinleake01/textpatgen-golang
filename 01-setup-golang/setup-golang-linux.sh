#!/bin/sh
#
# This file will download various documents, tutorials and instructions
# about the Go Programming Language, plus how to set up using GCCGO and
# LLVM/LLGO.
#
# This program runs in Linux, and will require Subversion, Mercurial (Hg),
# and Git.
#
# ----- GOLANG.ORG -----
hg clone https://code.google.com/p/go/
hg clone https://code.google.com/p/go.wiki/
hg clone https://code.google.com/p/go.text/
hg clone https://code.google.com/p/go.example/
hg clone https://code.google.com/p/go.empty/
hg clone https://code.google.com/p/go.crypto/
hg clone https://code.google.com/p/go.net/
hg clone https://code.google.com/p/go.codereview/
hg clone https://code.google.com/p/go.image/
hg clone https://code.google.com/p/go.talks/
hg clone https://code.google.com/p/go.exp/
hg clone https://code.google.com/p/go.tools/
hg clone https://code.google.com/p/go.benchmarks/

# ----- GCCGO -----
svn co http://gcc.gnu.org/svn/gcc/branches/gccgo/

# ----- LLVM/LLGO -----
git clone https://github.com/go-llvm/llvm.git
git clone https://github.com/go-llvm/llgo.git


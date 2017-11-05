+++
author = ["Florin Pățan"]
date = "2016-12-15T00:00:00+00:00"
title = "Finding good packages in the Sea of Open-Source"
series = ["Advent 2016"]
+++

So far we've seen different ways to use Go to build a variety of applications,
from [versioning your data pipelines](https://blog.gopheracademy.com/advent-2016/pachyderm/) to building your own [BBQ grill controller](https://blog.gopheracademy.com/advent-2016/qpid/).

But often times you'll be looking you'll be looking for reusing code from
others in order to allow you to focus on what your code and not reinvent the
wheel.

But how do you know which package to use? And when should you use a package
versus write your own?

Let that sink in for a bit and now think of when you are learning a language
and you don't really know what is considered idiomatic or not.

Those are hard questions. They become even harder if you have to use these
third-party packages to build applications that are going to be used into
production code.

Usually you have a checklist of items that you'd like a package or library
to meat before you can use it. My list looks like this (not in the order of 
importance):

- is it used by others?
- how does the documentation look like?
- does it do what I need it to do (and preferably nothing more)?
- is the author maintaining the package? How are issues / PRs addressed? Are
there any contributors to the project?
- license
- what's the code quality like? Does it have tests?

If you wonder why code quality is the last one on the list the answer is
simple: because I can assert the first five in just a few minutes while
reading the code can take a bit longer.

So how do you actually find packages?

For Go we have [godoc.org](https://godoc.org/) and [go-search.org](http://go-search.org/) which allow you
to search for Go packages. But with over 150.000 projects and packages indexed
by them, finding a "good" package could quickly become a problem.

Fortunately, both sort results by popularity (do not confuse with Github stars,
we'll come back to this a bit later). My default search tool is GoDoc but you
can also use GoSearch as GoDoc references it when you search for a package that
is not found in its index.

Popularity for Go packages means how often they are imported which will already
answer the first question.

Then I can quickly have a look at the documentation as well. And here is were
one of the Go's lesser known features shines. The documentation of a package
can be [tested](https://blog.golang.org/examples) so if I see that then I can
have the guarantee that the project is already in a good shape.

Having briefly checked the documentation, I can quickly go to the place where
it's hosted and see what's the situation looking like for the other factors.

Remember I've said that you shouldn't confuse popularity with number of Github
stars? Well this is the time to ignore the Github stars as again. I'll stick
with Github for now since it's more or less the de facto place for source code
hosting.

Back to our package, lets say that we are looking at a package which has a few
contributors (that's a nice touch) and a few open issues and maybe a few PRs.
All those things could indicate that in fact people are contributing to the
package.

The first thing I do is to look at the readme and see what's in there.
Important information that I can quickly see is:
- does the package have any badge for a CI server?
- does the package have instructions for how to install / use?
- are there any examples there that might be relevant? Or maybe an FAQ / link
to a wiki?

Having checked those out, I move to the issues / PRs. What I'm interested there
is how is the maintainer is working with potential contributors. That is very
important since I might end up myself opening an issue / PR in the future and
I would like to know what I'm going to expect.

By now I have a rough idea if I want to invest time in reading the code or I
need to move over to alternatives or not.

For me, the number of contributors also depends on the project type / where
it's going to be used. If I'm looking at replacing my template engine, I would
expect at least a few. It also depends on how long was the project around, if
it's a brand new one, then chances are it will have less contributors than a
more established one. However, if the contributors are actively discouraged by
the author, then I'm sure I'm going to stay away from it and only use it if I
really can't find  an alternative and writing one myself would be too time
consuming.

And now, for the fun part. Say I like everything that I see so far and the
license allows me to use it in the current project I work on.

The final step, me doing a code review of this.

If you are new to the language this is the hardest part probably because while
everything above is interchangeable with other languages, reviewing someone
else's code and making sure it's idiomatic it's going to be next to impossible.

However, don't be scared by this, as this is where Go shines, readability. And
if you are unsure of that, we have tools that can help out checking common for
common problems.

The first argument for this claim is that Go is a simple language, with just
25 keywords. This means that you'll be able to review all packages out there
very quickly after learning the language without any hurdles because you may
have forgotten what's a symbol or another or what does a keyword do versus the
other or some other language construct.

There are a few guidelines for [code review](https://github.com/golang/go/wiki/CodeReviewComments) that you can read to
help out with the review. Furthermore, [Effective Go](https://golang.org/doc/effective_go.html) and articles like
[Idiomatic Go](https://dmitri.shuralyov.com/idiomatic-go), [Avoid gotchas](https://divan.github.io/posts/avoid_gotchas) or 
[Gotchas and common mistakes in Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) can help you
quickly understand what are the common patters you should look out for.

A quick way to check some of these is to have a look at the package using a
more automated way. First and the simplest is to run [Go Report Card](https://goreportcard.com/) and see
what problems does the package have. Things like `gofmt` should be a must for
any package out there so that quickly can help identify packages which are not
correctly formatted.

If Go Report Card checks out ok, it's time to download the package locally and
run [Go Meta Linter](https://github.com/alecthomas/gometalinter) against it.
That will quickly highlight any problems with the package. Hopefully this will
not discover anything but if that happens then I can at least open an issue so
that the author is informed about this. If I like how the code looks like and
I decide to use the package, now I know I can also quickly contribute to it.

Placing your full trust in automated tools however might also lead you down the
wrong path as perfect valid code, that can pass checks, can also be not
idiomatic. That's why it's good to use these tools but always have a look at
what's happening under the hood for yourself.

An important property of Go packages is how well they work with other packages
so if you want to have a single takeaway from this article then this would be:
Go as a language encourages composition and this can be observed throughout the
standard library. Many of the popular packages in the open-source community
also respect and embrace this property and allow you to quickly combine them
in order to achieve the powerful functionality you need.

Finally, there's a shortcut you can take, tho I have to admit, it involves a
lot of trust on people on the Internet. Ask about the package you are looking
at in places like the [Gophers Slack](https://gophers.slack.com/), you can get
an [invite here](https://invite.slack.golangbridge.org/), the [Go Forum](https://forum.golangbridge.org/), 
or [golang-nuts mailing list](https://groups.google.com/forum/#!forum/golang-nuts).

But what should you do if you cannot find a good package or any package at all?

If you find a package that maybe doesn't cover everything that you need or
the quality it's not up on par, I highly suggest contributing to it. 

You can read in [this article](http://blog.sgmansfield.com/2016/06/working-with-forks-in-go/) how to work with forks for
Go packages and you can hopefully quickly bring the package in a state that
will be on par with your needs. If the author is busy maybe you can ask to
have commit rights or even become the maintainer of the package (mind you that
this might not happen if you are a first time contributor so don't be upset if
your request gets denied).

Everything else failing, you can start building your own package and make sure
that everything else you'd like to see in a package is also present in it.

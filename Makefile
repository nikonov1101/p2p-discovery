GO=go
OS=$(shell uname |awk '{print tolower($0)}')

default: build/seeker build/caster

build/seeker:
	@echo "+ $@"
	${GO} build -tags nocgo -o seeker_${OS} ./seeker

build/caster:
	@echo "+ $@"
	${GO} build -tags nocgo -o caster_${OS} ./caster

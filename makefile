APP       := SanXiangBank
TARGET    := Config
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
-include scripts/makefile.tars.gomod

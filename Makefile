CC = "g++"
FILENAME = "main"
BUILDDIR = "build"
SRCDIR = "src"
CFLAGS? = 

all:
	if [[ -d bin ]];
	then
	rm -rf build 
	fi;
		g++ -o $(BUILDDIR)/$(FILENAME) $(SRCDIR)/$(FILENAME).cpp

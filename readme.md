__INSTALLATION__
             
             + setup golang binaries
             + clone or download this repository
             + install dependencies using `dep` utility
             + in directory with project run following command:
               > `$ dep init`
             + build project using the following command:
               If you're usging GO gompiler > `$ go build -ldflags "-w -s" redirect.go`
               If you're usging GCCGO gompiler > `$ go build -compuler gccgo -gccgoflags "-w -s" redirect.go`
               
             __USING__  
             
             + perform settings in `settings.ini` file  
             + run program:
               > `$ ./redirect run`
             + take profit
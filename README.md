
╭━━━┳━━━┳━╮╭━┳━━━┳╮╱╱╭━━━┳━━━╮
┃╭━╮┃╭━╮┃┃╰╯┃┃╭━╮┃┃╱╱┃╭━╮┃╭━╮┃
┃╰━━┫┃╱┃┃╭╮╭╮┃┃╱┃┃┃╱╱┃┃╱┃┃╰━━╮
╰━━╮┃╰━╯┃┃┃┃┃┃╰━╯┃┃╱╭┫╰━╯┣━━╮┃
┃╰━╯┃╭━╮┃┃┃┃┃┃╭━╮┃╰━╯┃╭━╮┃╰━╯┃
╰━━━┻╯╱╰┻╯╰╯╰┻╯╱╰┻━━━┻╯╱╰┻━━━╯

TODO :
- [  ] Add More Source For Subdomain Finder
- [  ] Create Docker Container
- [  ] ...

* hot to user list -list domain.txt
``sh
$ assetfinder --subs-only google.com | httprobe >> domain.txt
$ samalas -list domain.txt
``
# saint-hubert              ![alt text](http://www.clubdesbrunosetdesccs.org/wp-content/uploads/2020/09/bruno-saint-hubert-francais-photo-1024x1024-1-553x400.jpg)
Sniff broken links on web pages

```
http -h GET https://github.com/bananananaaefaefZGTAHZHnanananana | grep 'status:' | sed 's/.*://' 
```

Algorythm big steps:
- [x] GET request on the page
- [x] Check if any link on the page
- [ ] retrieve all href content (http, https)
- [ ] Create slice
- For each href create struct with link as key
- [ ] For each key element of the data structure
  - [ ] GET request
  - [ ] Retrieve HTTP status code
  - [ ] add HTTP code status to value of each key
- [ ] Display struct in CLI

Future features to have:
- [ ] Nice CLI UI (with colors !),
- [ ] Arg to display only broken links,
- [ ] Give location of file to make HTTP GET calls
- [ ] Parse a website given a root hook (example.com => retrieve all routes based on hrefs)

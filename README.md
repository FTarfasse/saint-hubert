![alt text](http://www.clubdesbrunosetdesccs.org/wp-content/uploads/2020/09/bruno-saint-hubert-francais-photo-1024x1024-1-553x400.jpg)
<br>#SAINT-HUBERT: Sniff the dead urls of your webpage from the terminal !
 
```go build -o sniff```

Project roadmap:
- [x] GET request on the page
- [x] Check if any link on the page
- [x] retrieve all href content (http, https)
- [x] Create slice
- For each href create struct with link as key
- [x] For each key element of the data structure
  - [x] GET request
  - [x] Retrieve HTTP status code
  - [x] add HTTP code status to value of each key
  - [x] Remove duplicates from array
- [x] Display struct in CLI
- [x] Save output as JSON to file

Future features to have:
- [x] Nice CLI UI (with colors !),
- [x] Arg to display only issues
 links,
- [ ] Add argument options
- [x] Give Json back
- [ ] Give location of file to make parsing
- [ ] Parse a website given a root hook (example.com => retrieve all routes based on hrefs)

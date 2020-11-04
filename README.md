# saint-hubert              ![alt text](http://www.clubdesbrunosetdesccs.org/wp-content/uploads/2020/09/bruno-saint-hubert-francais-photo-1024x1024-1-553x400.jpg)
Sniff broken links on web pages

```
http -h GET https://github.com/bananananaaefaefZGTAHZHnanananana | grep 'status:' | sed 's/.*://' 
```

Algorythm big steps:
- [ ] GET request on the page
- [ ] retrieve all "http" / "https" / 'http' / 'https'
- [ ] add them on data structure
- [ ] for each key element of the data structure
  - [ ] GET request
  - [ ] add value to key the HTTP code status
- [ ] Display as table all adresses and the corresponding HTTP status

Future features to have:
- [ ] CLI UI (with colors !),
- [ ] Arg to display only errors,
- [ ] Give location of file to make HTTP GET calls

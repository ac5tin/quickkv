# 📂 QuickKV


### Roadmap
- [ ] sort first
- [ ] quick search content
- [x] encrypt
- [ ] cluster
- [x] web rest api


### Encrypt
```
quickkv -p 8000 -enc
```

### Start Server
```
quickkv -p 8310
```

### API
Set Key
```
POST /set/:key
```

Get Key
```
GET /get/:key
```

Get All
```
GET /
```

Delete Key
```
DEL /del/:key
```

Push (array)
```
POST /push/:key
```

Unshift (array)
```
POST /unshift/:key
```

Delete element from array
```
POST /rm/:key
```

Pop array
```
GET /pop/:key
```

Shift array
```
GET /shift/:key
```

Search by key prefix
```
GET /prefix?q=foo
```

Search key
```
GET /search?q=foo
```

RESET
```
DELETE /reset
```
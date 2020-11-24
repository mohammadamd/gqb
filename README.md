# gqb
It's a minimal query builder for golang to unified your queries and make them more readable

## example:
```go
	q := Select("id", "title").
		Table("TEST").
		Where("id", "=", "123").
		Where("title", "like", "annad").
		OrWhere("id", "=", "222").
		OrderBy("id", "desc").
		Offset(10).
		Limit(12).
		Generate()
```
### Output: 
``` SELECT id,title FROM TEST  WHERE id = 123 AND WHERE title like "annad" OR WHERE id = 222 ORDER BY id desc LIMIT 12 OFFSET 10 ```

## Todo:
- [X] Support Select Query
- [ ] Support Union
- [ ] Support Insert
- [ ] Support Update
- [ ] Support Delete
- [ ] Support Scoped Where
- [ ] ... (Fill this by yourself)

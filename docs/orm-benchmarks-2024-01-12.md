# Results

- orm-benchmark (with no flags)
```
Reports:
Insert
dbr:	4296	300735 ns/op	2688 B/op	65 allocs/op
sqlc:	4434	304944 ns/op	2788 B/op	62 allocs/op
bun:	3924	340627 ns/op	5014 B/op	13 allocs/op
goqu:	3232	364189 ns/op	13087 B/op	378 allocs/op
dbx:	2804	473715 ns/op	5157 B/op	105 allocs/op
rel:	2691	493346 ns/op	2623 B/op	45 allocs/op
pop:	1723	684888 ns/op	9591 B/op	238 allocs/op

InsertMulti
bun:	871	1457195 ns/op	42536 B/op	219 allocs/op
goqu:	756	1633260 ns/op	228345 B/op	11572 allocs/op
rel:	754	1653186 ns/op	306914 B/op	3265 allocs/op
dbx:	bulk-insert is not supported
pop:	bulk-insert is not supported
dbr:	bulk-insert is not supported
sqlc:	bulk-insert is not supported

Update
sqlc:	8761	145378 ns/op	877 B/op	14 allocs/op
dbr:	4074	313935 ns/op	2651 B/op	57 allocs/op
pop:	4072	321233 ns/op	6047 B/op	186 allocs/op
bun:	3805	334563 ns/op	4762 B/op	5 allocs/op
goqu:	3487	358742 ns/op	11122 B/op	339 allocs/op
rel:	2677	482435 ns/op	3048 B/op	45 allocs/op
dbx:	2594	493044 ns/op	4514 B/op	103 allocs/op

Read
sqlc:	8342	155186 ns/op	2092 B/op	51 allocs/op
pop:	7713	164745 ns/op	3197 B/op	67 allocs/op
bun:	7665	172808 ns/op	5829 B/op	39 allocs/op
dbr:	7376	178646 ns/op	2168 B/op	36 allocs/op
rel:	7406	179053 ns/op	2320 B/op	47 allocs/op
goqu:	5412	212999 ns/op	12509 B/op	275 allocs/op
dbx:	3972	326289 ns/op	3431 B/op	72 allocs/op

ReadSlice
sqlc:	3501	340362 ns/op	62678 B/op	1150 allocs/op
pop:	3004	388250 ns/op	69156 B/op	1307 allocs/op
dbx:	2942	410831 ns/op	43971 B/op	1355 allocs/op
bun:	2835	418216 ns/op	34062 B/op	1124 allocs/op
dbr:	2763	418354 ns/op	30800 B/op	1253 allocs/op
goqu:	2288	533020 ns/op	63400 B/op	2386 allocs/op
rel:	1736	683787 ns/op	141473 B/op	2553 allocs/op
```
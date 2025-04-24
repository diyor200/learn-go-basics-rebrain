Ready to GO!

komandalar:
go test -bench=. -benchtime=5s -cpuprofile=cpu.profile - memprofile=mem.profile   - cpu va mem profillarini saqlash bilan benchmark testni ishga tushirish
go tool pprof -http=:9090 [profile_file]               - web interfeysda saqlangan profillarni ko'rish (bizning holatda [profile_file] cpu.profile ga teng)
 
Ready to GO!

komandalar:
go test -bench=.                                        - benchmark testlarni ishga tushiradi
go test -bench=BenchmarkMySlowFunction                  - bunda aynan qaysi test funksiyani ishga tushirimiz ko'rsatiladi
go test -bench=BenchmarkMySlowFunction -benchmem        - bunda qo'shimcha ma'lumotlar chiqaradi(1ta operatsiyaga qancha xotiradan joy ajratilganini)
go test -bench=BenchmarkMySlowFunction -benchtime 30s   - bu funksiyani qancha vaqt yurishini ko'rsatadi
go test -bench=BenchmarkMySlowFunction -count 2         - bu testni necha marta ishga tushirishni ko'rsatadi
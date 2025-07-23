## 2023-10-16 İlerleme
✅ **gRPC servis temeli**
- [x] Protobuf tanımları (`proto/core.proto`)
- [x] Temel çağrı yönlendirme (`call_router.go`)
- [x] gRPC sunucusu ayağa kaldırıldı ve test edildi (`main.go`)

---

## 🔄 Sonraki Adımlar
- [ ] **signal**: Temel SIP `INVITE` handler'ını oluşturmak.
- [ ] **core**: `signal` servisinden gelen isteklere yanıt verecek adaptörü yazmak.
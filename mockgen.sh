#mock infra
mockgen -package=mock_infra -source=infra/db.go -destination=infra/mock/db.go

#mock repo
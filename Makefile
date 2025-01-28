.PHONY: gen-frontend
gen-frontend: ## gen frontend page code of {page}. example: make gen-frontend page=category_page
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --idl ../../idl/frontend/${page}.proto --service frontend -module github.com/SGNYYYY/gomall/app/frontend

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/SGNYYYY/gomall/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/SGNYYYY/gomall/app/${svc} --pass "-use github.com/SGNYYYY/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto